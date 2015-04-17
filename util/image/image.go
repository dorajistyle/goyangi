package image

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/daddye/vips"
	"github.com/disintegration/gift"
	"github.com/dorajistyle/goyangi/config"
)

func MediumOption() vips.Options {
	options := vips.Options{
		Width:        config.ImageWidth,
		Height:       0,
		Crop:         false,
		Extend:       vips.EXTEND_WHITE,
		Interpolator: vips.BILINEAR,
		Gravity:      vips.CENTRE,
		Quality:      95,
	}
	return options
}

func ThumbnailOption() vips.Options {
	options := vips.Options{
		Width:        config.ThumbnailWidth,
		Height:       0,
		Crop:         false,
		Extend:       vips.EXTEND_WHITE,
		Interpolator: vips.BILINEAR,
		Gravity:      vips.CENTRE,
		Quality:      95,
	}
	return options
}

func ResizeMediumVips(inBuf []byte) ([]byte, error) {
	return vips.Resize(inBuf, MediumOption())
}

func ResizeThumbnailVips(inBuf []byte) ([]byte, error) {
	return vips.Resize(inBuf, ThumbnailOption())
}

func MediumFilter() *gift.GIFT {
	return ResizeFilter(config.ImageWidth, 0)
}

func ThumbnailFilter() *gift.GIFT {
	return ResizeFilter(config.ThumbnailWidth, 0)
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
		//		log.Debugf("image type : %s %d %d\n", imageFormat, config.Width, config.Height)
	case "image/png":
		err = ParsePng(wb, r, g)
	case "image/gif":
		err = ParseGif(wb, r, g)
	default:
		err = fmt.Errorf("unsupported image type. %s\n", imageFormat)
	}
	return wb, err
}

func ResizeMedium(imageFormat string, r io.Reader) (*bytes.Buffer, error) {
	wb := new(bytes.Buffer)
	g := MediumFilter()
	var err error
	switch imageFormat {
	case "image/jpeg":
		err = ResizeJpeg(wb, r, g, config.ImageWidth, config.ImageHeight)
	case "image/png":
		err = ResizePng(wb, r, g, config.ImageWidth, config.ImageHeight)
	case "image/gif":
		err = ResizeGif(wb, r, g, config.ImageWidth, config.ImageHeight)
	default:
		err = fmt.Errorf("unsupported image type. %s\n", imageFormat)
	}
	return wb, err
}

func ResizeThumbnail(imageFormat string, r io.Reader) (*bytes.Buffer, error) {
	wb := new(bytes.Buffer)
	g := ThumbnailFilter()
	var err error
	switch imageFormat {
	case "image/jpeg":
		err = ResizeJpeg(wb, r, g, config.ImageWidth, config.ImageHeight)
	case "image/png":
		err = ResizePng(wb, r, g, config.ImageWidth, config.ImageHeight)
	case "image/gif":
		err = ResizeGif(wb, r, g, config.ImageWidth, config.ImageHeight)
	default:
		err = fmt.Errorf("unsupported image type. %s\n", imageFormat)
	}
	return wb, err
}
