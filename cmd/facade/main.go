package main

import (
	"fmt"
	"log"
	"wb-onboarding/pkg/facade"
)

func main() {
	videoConverter := facade.NewVideoConverter()
	str, err := videoConverter.Do("test string")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*str)
}
