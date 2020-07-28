package video_file

import "fmt"

//videoFiler interface describes video handling
type VideoFiler interface {
	DoSomeAction(int64) (*string, error)
}

type videoFile struct {
	name string
}

func (v *videoFile) DoSomeAction(someNumber int64) (*string, error) {
	str := fmt.Sprintf("Hello, %s! - %d", v.name, someNumber)
	return &str, nil
}

//NewVideoFile creates private implementation of VideoFiler interface
func NewVideoFile() VideoFiler {
	return &videoFile{
		name: "videoFile",
	}
}
