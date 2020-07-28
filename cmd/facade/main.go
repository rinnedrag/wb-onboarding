package main

import (
	"fmt"
	"log"
	"wb-onboarding/pkg/facade/audiomixer"
	"wb-onboarding/pkg/facade/videoconverter"
	"wb-onboarding/pkg/facade/videofile"
)

func main() {
	sl := make([]int, 1000)
	sl = append(sl, 1)
	videoConverter := videoconverter.NewVideoConverter(audiomixer.NewAudioMixer(), videofile.NewVideoFile())
	str, err := videoConverter.Do("test string")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
