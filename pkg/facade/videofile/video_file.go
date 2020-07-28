package videofile

import "fmt"

//VideoFiler interface describes video handling
type VideoFiler interface {
	DoSomeAction(int64) (string, error)
}

type videoFile struct {
	name string
}

func (v *videoFile) DoSomeAction(someNumber int64) (str string, err error) {
	str = fmt.Sprintf("Hello, %s! - %d", v.name, someNumber)
	return
}

//NewVideoFile creates private implementation of VideoFiler interface
func NewVideoFile() VideoFiler {
	return &videoFile{
		name: "videoFile",
	}
}
