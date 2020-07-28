package audio_mixer

import "fmt"

//audioMixer interface describes audio handling
type AudioMixer interface {
	DoSomeAction(int64) (*string, error)
}

type audioMix struct {
	name string
}

func (v *audioMix) DoSomeAction(someNumber int64) (*string, error) {
	str := fmt.Sprintf("Hello, %s! - %d", v.name, someNumber)
	return &str, nil
}

//newAudioMixer creates private implementation of audioMixer interface
func NewAudioMixer() AudioMixer {
	return &audioMix{
		name: "audioMixer",
	}
}
