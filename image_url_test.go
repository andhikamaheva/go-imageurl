package imageurl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cheekybits/is"
)

/* imageURL.URI = "" */

func TestGetFileName(t *testing.T) {
	is := is.New(t)

	fileName := SetURI("https://dummyimage.com/600x400/000/fff.jpg").GetFileName()

	is.Equal("fff.jpg", fileName)
}

func TestJPGImage(t *testing.T) {
	is := is.New(t)

	imageType := SetURI("https://dummyimage.com/600x400/000/fff.jpg").GetImageType()

	is.Equal("jpg", imageType)
}

// TestJPEGImage ...
func TestJPEGImage(t *testing.T) {
	is := is.New(t)

	imageType := SetURI("https://dummyimage.com/600x400/000/fff.jpeg").GetImageType()

	is.Equal("jpeg", imageType)
}

// TestPNGImage ...
func TestPNGImage(t *testing.T) {
	is := is.New(t)

	imageType := SetURI("https://dummyimage.com/600x400/000/fff.png").GetImageType()

	is.Equal("png", imageType)
}

// TestGIFImage ...
func TestGIFImage(t *testing.T) {
	is := is.New(t)

	imageType := SetURI("https://dummyimage.com/600x400/000/fff.gif").GetImageType()

	is.Equal("gif", imageType)
}

// TestBMPImage ...
func TestBMPImage(t *testing.T) {
	is := is.New(t)

	imageType := SetURI("http://eeweb.poly.edu/~yao/EL5123/image/lena_gray.bmp").GetImageType()

	is.Equal("bmp", imageType)
}

// TestTIFFImage ...
func TestTIFFImage(t *testing.T) {
	is := is.New(t)

	imageType := SetURI("http://eeweb.poly.edu/~yao/EL5123/image/lena_color.tiff").GetImageType()

	is.Equal("tiff", imageType)
}

// TestSVGImage ...
func TestSVGImage(t *testing.T) {
	is := is.New(t)

	imageType := SetURI("https://placeholder.pics/svg/300").GetImageType()

	is.Equal("svg", imageType)
}

// TestGetImageSize ...
func TestGetImageSize(t *testing.T) {
	is := is.New(t)

	imageSize := SetURI("https://placeholder.pics/svg/300").GetImageSize()

	is.Equal(322, imageSize)
}

// TestGetContentType ...
func TestGetContentType(t *testing.T) {
	is := is.New(t)

	contentType := SetURI("https://placeholder.pics/svg/300").GetContentType()

	is.Equal("image/svg+xml", contentType)
}

// TestSaveImage ...
func TestSaveImage(t *testing.T) {

	is := is.New(t)

	image := SetURI("https://dummyimage.com/600x400/000/fff.png")

	file := image.SaveImage("public/images", image.GetFileName())

	file.Close()

	path := filepath.Join("public/images", image.GetFileName())

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Error(err)
	} else {
		is.OK("File Exist")
	}
}
