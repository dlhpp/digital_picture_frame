package internal

// ImageStore holds the list of image file paths
type ImageStore struct {
	Images         []string
	ImageSubscript int
}

type FlagSettings struct {
	Fullscreen bool
	Screensize string
	Random     bool
	Url        string
}
