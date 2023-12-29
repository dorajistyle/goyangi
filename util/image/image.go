package image

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"strconv"

	"github.com/disintegration/gift"
	"github.com/dorajistyle/goyangi/util/stringHelper"
	"github.com/spf13/viper"
)

func GenerateURL(imageURLPrefix string, types string, id int64, name string) string {
	imageURLBuffer := new(bytes.Buffer)
	stringHelper.Concat(imageURLBuffer, imageURLPrefix)
	stringHelper.Concat(imageURLBuffer, types)
	stringHelper.Concat(imageURLBuffer, "/")
	stringHelper.Concat(imageURLBuffer, strconv.FormatInt(id, 10))
	stringHelper.Concat(imageURLBuffer, "/")
	stringHelper.Concat(imageURLBuffer, name)
	return imageURLBuffer.String()
}

func LargeFilter() *gift.GIFT {
	return ResizeFilter(viper.GetInt("image.large.width"), 0)
}

func MediumFilter() *gift.GIFT {
	return ResizeFilter(viper.GetInt("image.medium.width"), 0)
}

func ThumbnailFilter() *gift.GIFT {
	return ResizeFilter(viper.GetInt("image.thumbnail.width"), 0)
}

func ResizeFilter(width int, height int) *gift.GIFT {
	g := gift.New(gift.Resize(width, height, gift.LanczosResampling))
	return g
}

func ApplyFilter(src image.Image, g *gift.GIFT) image.Image {
	dst := image.NewRGBA(g.Bounds(src.Bounds()))
	g.Draw(dst, src)
	return dst
}

func ParseJpeg(wb *bytes.Buffer, r io.Reader, g *gift.GIFT) error {
	src, err := jpeg.Decode(r)
	if err != nil {
		return err
	}
	dst := ApplyFilter(src, g)
	err = jpeg.Encode(wb, dst, nil)
	return err
}

func ParsePng(wb *bytes.Buffer, r io.Reader, g *gift.GIFT) error {
	src, err := png.Decode(r)
	dst := ApplyFilter(src, g)
	err = png.Encode(wb, dst)
	return err
}

func ParseGif(wb *bytes.Buffer, r io.Reader, g *gift.GIFT) error {
	src, err := gif.Decode(r)
	dst := ApplyFilter(src, g)
	err = gif.Encode(wb, dst, nil)
	return err
}

func ResizeJpeg(wb *bytes.Buffer, r io.Reader, g *gift.GIFT, width int, height int) error {
	src, err := jpeg.Decode(r)
	bounds := src.Bounds()
	dst := src
	if width < bounds.Dx() || height < bounds.Dy() {
		dst = ApplyFilter(src, g)
	}
	err = jpeg.Encode(wb, dst, nil)
	return err
}

func ResizePng(wb *bytes.Buffer, r io.Reader, g *gift.GIFT, width int, height int) error {
	src, err := png.Decode(r)
	bounds := src.Bounds()
	dst := src
	if width < bounds.Dx() || height < bounds.Dy() {
		dst = ApplyFilter(src, g)
	}
	err = png.Encode(wb, dst)
	return err
}

func ResizeGif(wb *bytes.Buffer, r io.Reader, g *gift.GIFT, width int, height int) error {
	src, err := gif.Decode(r)
	bounds := src.Bounds()
	dst := src
	if width < bounds.Dx() || height < bounds.Dy() {
		dst = ApplyFilter(src, g)
	}
	err = gif.Encode(wb, dst, nil)
	return err
}

func ParseImage(imageFormat string, r io.Reader, g *gift.GIFT) (*bytes.Buffer, error) {
	wb := new(bytes.Buffer)
	var err error
	switch imageFormat {
	case "image/jpeg":
		err = ParseJpeg(wb, r, g)
	case "image/png":
		err = ParsePng(wb, r, g)
	case "image/gif":
		err = ParseGif(wb, r, g)
	default:
		err = fmt.Errorf("unsupported image type. %s\n", imageFormat)
	}
	return wb, err
}

func ResizeLarge(imageFormat string, r io.Reader) (*bytes.Buffer, error) {
	wb := new(bytes.Buffer)
	g := MediumFilter()
	var err error
	w, h := viper.GetInt("image.large.width"), viper.GetInt("image.large.height")
	switch imageFormat {
	case "image/jpeg":
		err = ResizeJpeg(wb, r, g, w, h)
	case "image/png":
		err = ResizePng(wb, r, g, w, h)
	case "image/gif":
		err = ResizeGif(wb, r, g, w, h)
	default:
		err = fmt.Errorf("unsupported image type. %s\n", imageFormat)
	}
	return wb, err
}

func ResizeMedium(imageFormat string, r io.Reader) (*bytes.Buffer, error) {
	wb := new(bytes.Buffer)
	g := MediumFilter()
	var err error
	w, h := viper.GetInt("image.medium.width"), viper.GetInt("image.medium.height")
	switch imageFormat {
	case "image/jpeg":
		err = ResizeJpeg(wb, r, g, w, h)
	case "image/png":
		err = ResizePng(wb, r, g, w, h)
	case "image/gif":
		err = ResizeGif(wb, r, g, w, h)
	default:
		err = fmt.Errorf("unsupported image type. %s\n", imageFormat)
	}
	return wb, err
}

func ResizeThumbnail(imageFormat string, r io.Reader) (*bytes.Buffer, error) {
	wb := new(bytes.Buffer)
	g := ThumbnailFilter()
	var err error
	w, h := viper.GetInt("image.thumbnail.width"), viper.GetInt("image.thumbnail.height")
	switch imageFormat {
	case "image/jpeg":
		err = ResizeJpeg(wb, r, g, w, h)
	case "image/png":
		err = ResizePng(wb, r, g, w, h)
	case "image/gif":
		err = ResizeGif(wb, r, g, w, h)
	default:
		err = fmt.Errorf("unsupported image type. %s\n", imageFormat)
	}
	return wb, err
}
