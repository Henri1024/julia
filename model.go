package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"sync"
)

type Julia struct {
	filename     string
	maxIteration int
	width        int
	height       int
	CX           float64
	CY           float64
	Scale        float64
	Img          *image.RGBA
}

func NewJulia(cX, cY, scale float64) *Julia {
	return &Julia{
		CX:           cX,
		CY:           cY,
		Scale:        scale,
		filename:     fmt.Sprintf("juliaset%f,%f,%f.png", cX, cY, scale),
		width:        1080,
		height:       1080,
		maxIteration: 126,
	}
}

func (j *Julia) SetWidthHeigh(widthHeight int) {
	j.height = widthHeight
	j.width = widthHeight
}

func (j *Julia) SetFilename(fileName string) {
	j.filename = fileName
}

func (j *Julia) SetMaxIteration(maxIteration int) {
	j.maxIteration = maxIteration
}

func (j *Julia) GenerateImg() *image.RGBA {
	x0, y0, x1, y1 := j.CX-j.Scale, j.CY-j.Scale, j.CX+j.Scale, j.CY+j.Scale
	fmt.Println(x0, y0, x1, y1)
	height, width := float64(j.height), float64(j.width)
	j.Img = image.NewRGBA(
		image.Rect(0, 0, j.width, j.height),
	)

	rangeX, rangeY := x1-x0, y1-y0
	iX, iY := rangeX/width, rangeY/height

	var wg sync.WaitGroup
	wg.Add(j.width * 2)

	for x := 0.0; x < width; x += 1 {

		// first half
		go func(x float64) {
			var i int

			for y := 0.0; y < height/2.0; y += 1 {
				// x, y represent pixel in picture

				// i iteration, apply to color
				i = j.maxIteration

				// zx, zy represent coordinate respect to pixel
				zx := x0 + (x * iX)
				zy := y0 + (y * iY)

				for zx*zx+zy*zy < 4.0 && i > 0 {
					tmp := zx*zx - zy*zy + j.CX
					zy = 2.0*zx*zy + j.CY
					zx = tmp
					i--
				}

				i = int(float32(i) / float32(j.maxIteration) * 256.0)
				if i > 255 {
					i = 255
				}
				j.Img.Set(int(x), int(y), Plan9[i])
			}
			wg.Done()
		}(x)

		// second half
		go func(x float64) {
			var i int

			for y := height / 2; y < height; y += 1 {
				// x, y represent pixel in picture

				// i iteration, apply to color
				i = j.maxIteration

				// zx, zy represent coordinate respect to pixel
				zx := x0 + (x * iX)
				zy := y0 + (y * iY)

				for zx*zx+zy*zy < 4.0 && i > 0 {
					tmp := zx*zx - zy*zy + j.CX
					zy = 2.0*zx*zy + j.CY
					zx = tmp
					i--
				}

				i = int(float32(i) / float32(j.maxIteration) * 256.0)
				if i > 255 {
					i = 255
				}
				j.Img.Set(int(x), int(y), Plan9[i])
			}
			wg.Done()
		}(x)
	}
	wg.Wait()

	return j.Img
}

func (j *Julia) CreateImg() error {

	if j.Img == nil {
		_ = j.GenerateImg()
	}

	buf := new(bytes.Buffer)

	if err := jpeg.Encode(buf, j.Img, nil); err != nil {
		return err
	}

	imgFile, err := os.Create("output/" + j.filename)
	if err != nil {
		return err
	}
	defer imgFile.Close()

	if err := png.Encode(imgFile, j.Img); err != nil {
		return err
	}
	return nil

}
