package main

import (
	"fmt"
	"github.com/ciaranRoche/power_of_greyscale/pkg/imageprocessing"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	originalImage = "/tmp/original.jpg"
	greyImage = "/tmp/grey.jpg"
)

func main() {
	resp, err := http.Get("https://ciaransimages.s3-eu-west-1.amazonaws.com/IMG_8282.jpg")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	img, err := os.Create("/tmp/original.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer img.Close()

	_, err = io.Copy(img, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	imageFile, err := os.Open(originalImage)
	if err != nil {
		log.Fatal(err)
	}

	image, err := jpeg.Decode(imageFile)
	if err != nil {
		log.Fatal(err)
	}
	imageFile.Close()

	processorPipeline := imageprocessing.NewProcessorPipeline()

	greyScaleAction := imageprocessing.NewActionGreyScale()
	processorPipeline.AddAction(greyScaleAction)

	processedImage, err := processorPipeline.Transform(image)
	if err != nil {
		log.Fatal(err)
	}

	createOutput, err := os.Create(fmt.Sprintf(greyImage))
	if err != nil {
		log.Fatal(err)
	}
	defer createOutput.Close()

	jpeg.Encode(createOutput, processedImage, nil)
}
