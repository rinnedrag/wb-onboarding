package facade

import "fmt"

//audioMixer interface describes audio handling
type audioMixer interface {
	doSomeAction(int64) (*string, error)
}

type audioMix struct {
	name string
}

func (v *audioMix) doSomeAction(someNumber int64) (*string, error) {
	str := fmt.Sprintf("Hello, %s! - %d", v.name, someNumber)
	return &str, nil
}

//newAudioMixer creates private implementation of audioMixer interface
func newAudioMixer() audioMixer {
	return &audioMix{
		name: "audioMixer",
	}
}
