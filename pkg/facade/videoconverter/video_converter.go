package videoconverter

import (
	"fmt"
)

type audioMixer interface {
	DoSomeAction(int64) (string, error)
}

type videoFiler interface {
	DoSomeAction(int64) (string, error)
}

//VideoConverter provides the simplifying interface to convert video
type VideoConverter interface {
	Do(string) (string, error)
}

type videoConvert struct {
	audioMix  audioMixer
	videoFile videoFiler
}

func (v *videoConvert) Do(someString string) (resultStr string, err error) {
	amStr, err := v.audioMix.DoSomeAction(1)
	if err != nil {
		return
	}
	vfStr, err := v.videoFile.DoSomeAction(2)
	if err != nil {
		return
	}
	resultStr = fmt.Sprintf("%s + %s + %s", amStr, vfStr, someString)
	return
}

//NewVideoConverter creates private implementation of VideoConverter interface
func NewVideoConverter(audioMixer audioMixer, videoFiler videoFiler) VideoConverter {
	return &videoConvert{
		audioMix:  audioMixer,
		videoFile: videoFiler,
	}
}
