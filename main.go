package main

import (
	"flag"
	"fmt"
	"image"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/unixpickle/essentials"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: brightnesses [flags] <path>")
		flag.PrintDefaults()
	}

	var rename bool
	flag.BoolVar(&rename, "rename", false, "rename images by brightness")
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	paths, err := imagePaths(flag.Args()[0])
	if err != nil {
		essentials.Die(err)
	}

	for _, path := range paths {
		r, err := os.Open(path)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		img, _, err := image.Decode(r)
		r.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, "decode "+path+":", err)
			continue
		}
		b := brightness(img)
		fmt.Println(b, path)
		if rename {
			dir, base := filepath.Split(path)
			base = strconv.Itoa(b) + " " + base
			os.Rename(path, filepath.Join(dir, base))
		}
	}
}

func imagePaths(path string) ([]string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !info.IsDir() {
		return []string{path}, nil
	}
	listing, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, x := range listing {
		res = append(res, filepath.Join(path, x.Name()))
	}
	return res, nil
}

func brightness(img image.Image) int {
	var sum int64
	bounds := img.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			sum += int64(r) >> 8
			sum += int64(g) >> 8
			sum += int64(b) >> 8
		}
	}
	return int(sum / int64(3*bounds.Dx()*bounds.Dy()))
}
