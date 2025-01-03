package main

import (
	"flag"
	"fmt"
	"log"
	"main/filters"
	"sync"
)

func main() {
	inputFileName := flag.String("i", "input.mp4", "input video file name")
	outputFileName := flag.String("o", "output.mp4", "output video file name")
	fps := flag.Int("f", 30, "frames per second of the output video")
	profileId := flag.Int("p", len(filters.Profiles), "available profiles:"+filters.GetProfileNames())

	flag.Parse()

	tempFolder, err := createTemporaryFolder()
	if err != nil {
		log.Fatal(err)
	}

	frames, err := splitFrames(tempFolder, *inputFileName, *fps)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: Check really channels are needed
	channel := make(chan bool, 10)

	var wg sync.WaitGroup
	wg.Add(len(frames))

	for _, frame := range frames {
		go func(frame string) {
			channel <- true

			img, err := decodeFrame(frame)
			if err != nil {
				log.Fatal(err)
			}

			img = filters.Run(img, *profileId)

			err = saveFrame(frame, img)
			if err != nil {
				log.Fatal(err)
			}

			defer func() {
				<-channel
				wg.Done()
			}()
		}(frame)
	}

	wg.Wait()

	err = joinFrames(tempFolder, *outputFileName, *fps)
	if err != nil {
		log.Fatal(err)
	}

	err = deleteTemporaryFolder(tempFolder)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
