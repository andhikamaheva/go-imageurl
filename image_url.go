package imageurl

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

// ImageURL ...
type ImageURL struct {
	URI string
}

// Init ...
func Init() *ImageURL {
	imageURL := &ImageURL{}

	return imageURL
}

// SetURI ...
func SetURI(uri string) *ImageURL {
	imageURL := &ImageURL{
		URI: uri,
	}

	return imageURL
}

// GetFileName ...
func (i *ImageURL) GetFileName() string {
	fileURL, _ := url.Parse(i.URI)

	path := fileURL.Path
	segments := strings.Split(path, "/")

	fileName := segments[len(segments)-1]

	return fileName
}

// GetImageType ...
func (i *ImageURL) GetImageType() ImageType {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, _ := client.Get(i.URI)

	defer resp.Body.Close()

	segments := strings.Split(resp.Header["Content-Type"][0], "/")

	imageType := segments[len(segments)-1]

	if segments := strings.Split(imageType, "+"); segments[0] == "svg" {
		imageType = "svg"
	}

	if imageType == string(JPG) || imageType == string(JPEG) || imageType == string(PNG) || imageType == string(GIF) || imageType == string(BMP) || imageType == string(TIFF) || imageType == string(SVG) {
		return ImageType(imageType)
	}
	return Unknown
}

// GetImageSize ...
func (i *ImageURL) GetImageSize() int32 {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, _ := client.Get(i.URI)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return int32(len(body))
}

// SaveImage ...
func (i *ImageURL) SaveImage(dir string, fileName string) *os.File {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	resp, _ := client.Get(i.URI)

	defer resp.Body.Close()

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}

	path := filepath.Join(dir, fileName)
	if file, err := os.Create(path); err != nil {
		panic(err)
	} else {
		_, err := io.Copy(file, resp.Body)

		if err != nil {
			panic(err)
		}

		return file
	}

}
