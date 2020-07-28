package facade

import "fmt"

//videoFiler interface describes video handling
type videoFiler interface {
	doSomeAction(int64) (*string, error)
}

type videoFile struct {
	name string
}

func (v *videoFile) doSomeAction(someNumber int64) (*string, error) {
	str := fmt.Sprintf("Hello, %s! - %d", v.name, someNumber)
	return &str, nil
}

//NewVideoFile creates private implementation of VideoFiler interface
func newVideoFile() videoFiler {
	return &videoFile{
		name: "videoFile",
	}
}
