package internal

// ImageStore holds the list of image file paths
type ImageStore struct {
	Images         []string
	ImageSubscript int
}

type FlagSettings struct {
	Kiosk  bool
	Random bool
	Url    string
}
