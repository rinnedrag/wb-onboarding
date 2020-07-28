package facade

import "fmt"

//VideoConverter provides the simplifying interface to convert video
type VideoConverter interface {
	Do(string) (*string, error)
}

type videoConvert struct {
	audioMix  audioMixer
	videoFile videoFiler
}

func (v *videoConvert) Do(someString string) (*string, error) {
	amStr, err := v.audioMix.doSomeAction(1)
	if err != nil {
		return nil, err
	}
	vfStr, err := v.videoFile.doSomeAction(2)
	if err != nil {
		return nil, err
	}
	resultStr := fmt.Sprintf("%s + %s + %s", *amStr, *vfStr, someString)
	return &resultStr, nil
}

//NewVideoConverter creates private implementation of VideoConverter interface
func NewVideoConverter() VideoConverter {
	return &videoConvert{
		audioMix:  newAudioMixer(),
		videoFile: newVideoFile(),
	}
}
