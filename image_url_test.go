package imageurl

import (
	"testing"

	"github.com/cheekybits/is"
)

/* imageURL.URI = "" */

func TestGetFileName(t *testing.T) {

	is := is.New(t)

	imageURL := ImageURL{}

	fileName := imageURL.SetURI("https://dummyimage.com/600x400/000/fff.jpg").GetFileName()

	is.Equal("fff.jpg", fileName)
}

func TestJPGImage(t *testing.T) {
	is := is.New(t)
	imageURL := ImageURL{}

	imageType := imageURL.SetURI("https://dummyimage.com/600x400/000/fff.jpg").GetImageType()
	is.Equal("jpg", imageType)
}
