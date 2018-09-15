package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"time"
	"runtime"
	"sync"
)

// Parameter
// 0: path to input file
// 1: path to output file
func main() {
	// src:= "./color.jpg"
	// dest : = "./grey.jpg"
	args := os.Args[1:]
	src := args[0]
	dest := args[1]

	t := time.Now()
	convert(src, dest)
	fmt.Println(time.Now().Sub(t).Seconds())
}


func convert(src, dest string) {
	existingImageFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer existingImageFile.Close()

	// We only need this because we already read from the file
	// We have to reset the file pointer back to beginning
	existingImageFile.Seek(0, 0)
	loadedImage, err := jpeg.Decode(existingImageFile)

	if err != nil {
		fmt.Println(err)
	}
	bounds := loadedImage.Bounds()
	m := image.NewNRGBA(loadedImage.Bounds())
	

	coreNum := runtime.NumCPU()
	if coreNum > 5 { coreNum = 5}
	// coreNum = 2
	sizeY := bounds.Max.Y/coreNum;
	Y := sizeY
	minY := bounds.Min.Y
	var wg sync.WaitGroup
	wg.Add(coreNum);
	for core := 0; core < coreNum; core++ {
		go func(a, b int) {
			defer wg.Done()
			for y := a; y <= b; y++ {
				for x := bounds.Min.X; x < bounds.Max.X; x++ {
					r, g, b, _ := loadedImage.At(x, y).RGBA()
					z := (r + g + b) / 795 // 3*255
					c := color.Gray{uint8(z)}
					m.Set(x, y, c)
				}
			}

		}(minY, Y)
		minY = Y
		Y +=sizeY
		if Y > bounds.Max.Y {
			Y = bounds.Max.Y
		}
	}
	
	wg.Wait()

	f, errWrite := os.Create(dest)
	if errWrite != nil {
		fmt.Println("Can not create file ", errWrite)
	}
	defer f.Close()

	jpeg.Encode(f, m, nil)
}
func convert2(src, dest string) {
	existingImageFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	}
	defer existingImageFile.Close()

	existingImageFile.Seek(0, 0)

	loadedImage, err := jpeg.Decode(existingImageFile)

	if err != nil {
		fmt.Println(err)
	}

	bounds := loadedImage.Bounds()
	m := image.NewNRGBA(loadedImage.Bounds())
	
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := loadedImage.At(x, y).RGBA()
			z := (r + g + b) / 795
			c := color.Gray{uint8(z)}
			m.Set(x, y, c)
		}
	}


	f, errWrite := os.Create(dest)
	if errWrite != nil {
		fmt.Println("Can not create file ", errWrite)
	}
	defer f.Close()

	jpeg.Encode(f, m, nil)
}