package imageurl

// ImageType : represents of image type that can detect by image-url package, if image type is not detected will return unkown
type ImageType string

const (

	// JPG ...
	JPG ImageType = "jpg"

	// JPEG ...
	JPEG ImageType = "jpeg"

	// PNG ...
	PNG ImageType = "png"

	// GIF ...
	GIF ImageType = "gif"

	// BMP ...
	BMP ImageType = "bmp"

	// TIFF ...
	TIFF ImageType = "tiff"

	// SVG ...
	SVG ImageType = "svg"

	// Unknown ...
	Unknown ImageType = "unknown"
)
