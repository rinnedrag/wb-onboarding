package video_converter

import (
	"fmt"
	"wb-onboarding/pkg/facade/audio_mixer"
	"wb-onboarding/pkg/facade/video_file"
)

//VideoConverter provides the simplifying interface to convert video
type VideoConverter interface {
	Do(string) (*string, error)
}

type videoConvert struct {
	audioMix  audio_mixer.AudioMixer
	videoFile video_file.VideoFiler
}

func (v *videoConvert) Do(someString string) (*string, error) {
	amStr, err := v.audioMix.DoSomeAction(1)
	if err != nil {
		return nil, err
	}
	vfStr, err := v.videoFile.DoSomeAction(2)
	if err != nil {
		return nil, err
	}
	resultStr := fmt.Sprintf("%s + %s + %s", *amStr, *vfStr, someString)
	return &resultStr, nil
}

//NewVideoConverter creates private implementation of VideoConverter interface
func NewVideoConverter() VideoConverter {
	return &videoConvert{
		audioMix:  audio_mixer.NewAudioMixer(),
		videoFile: video_file.NewVideoFile(),
	}
}
