package main

import (
	"fmt"
	"log"
	"wb-onboarding/pkg/facade/video_converter"
)

func main() {
	videoConverter := video_converter.NewVideoConverter()
	str, err := videoConverter.Do("test string")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*str)
}
