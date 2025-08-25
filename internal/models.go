package internal

// ImageStore holds the list of image file paths
type ImageStore struct {
	Images         []string
	ImageSubscript int
}

type FlagSettings struct {
	Browser    bool
	Fullscreen bool
	Random     bool
	Url        string
}
