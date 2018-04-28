package main

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

const (
	imgQuality = 95
)

type transformFunc func(img image.Image, width int, height int, filter imaging.ResampleFilter) *image.NRGBA

func noTransform(img image.Image, width int, height int, filter imaging.ResampleFilter) *image.NRGBA {
	return imaging.Clone(img)
}

type size struct {
	width  int
	height int
}

func parseSize(s string) (size, error) {
	nums := strings.Split(s, "x")
	if len(nums) != 2 {
		return size{}, fmt.Errorf("Invalid size %q", s)
	}

	w, err := strconv.ParseInt(nums[0], 10, 64)
	if err != nil {
		return size{}, fmt.Errorf("Invalid size %q: %s", s, err)
	}

	h, err := strconv.ParseInt(nums[1], 10, 64)
	if err != nil {
		return size{}, fmt.Errorf("Invalid size %q: %s", s, err)
	}

	return size{
		width:  int(w),
		height: int(h),
	}, nil
}

func scaleImage(path string) (thumbnail io.Reader, fullsize io.Reader, err error) {
	var img image.Image
	img, err = loadImage(path)
	if err != nil {
		return
	}

	thumbnail, err = generateThumbnail(img)
	if err != nil {
		return
	}

	fullsize, err = fitImage(img)
	return
}

func fitImage(img image.Image) (io.Reader, error) {
	s, err := parseSize(cfg.FullSize)
	if err != nil {
		return nil, err
	}

	origW, origH := img.Bounds().Max.X, img.Bounds().Max.Y
	scaleFactor := math.Max(float64(s.width)/float64(origW), float64(s.height)/float64(origH))

	var t transformFunc = noTransform
	if scaleFactor < 1 {
		t = imaging.Fit
	}

	return renderImage(img, s.width, s.height, t)
}

func generateThumbnail(img image.Image) (io.Reader, error) {
	s, err := parseSize(cfg.ThumbnailSize)
	if err != nil {
		return nil, err
	}

	return renderImage(img, s.width, s.height, imaging.Thumbnail)
}

func renderImage(img image.Image, w, h int, trans transformFunc) (io.Reader, error) {
	outImg := trans(img, w, h, imaging.Lanczos)

	buf := new(bytes.Buffer)
	if err := imaging.Encode(buf, outImg, imaging.JPEG, imaging.JPEGQuality(imgQuality)); err != nil {
		return nil, err
	}

	return buf, nil
}

func loadImage(path string) (image.Image, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	return imaging.Decode(fp)
}
