package audiomixer

import "fmt"

//AudioMixer interface describes audio handling
type AudioMixer interface {
	DoSomeAction(int64) (string, error)
}

type audioMix struct {
	name string
}

func (v *audioMix) DoSomeAction(someNumber int64) (str string, err error) {
	str = fmt.Sprintf("Hello, %s! - %d", v.name, someNumber)
	return
}

//NewAudioMixer creates private implementation of audioMixer interface
func NewAudioMixer() AudioMixer {
	return &audioMix{
		name: "audioMixer",
	}
}
