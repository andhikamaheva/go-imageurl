package imageurl

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// ImageURL ...
type ImageURL struct {
	URI string
}

// SetURI ...
func (i *ImageURL) SetURI(uri string) *ImageURL {
	i.URI = uri
	return i
}

// GetFileName ...
func (i *ImageURL) GetFileName() string {
	fmt.Println(i.URI)
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

	segments := strings.Split(resp.Header["Content-Type"][0], "/")

	imageType := segments[len(segments)-1]

	if imageType == string(JPG) || imageType == string(JPEG) || imageType == string(PNG) || imageType == string(GIF) || imageType == string(BMP) || imageType == string(TIFF) {
		return ImageType(imageType)
	}
	return Unknown

}
