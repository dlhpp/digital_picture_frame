package internal

// ImageStore holds the list of image file paths
type ImageStore struct {
	Images         []string
	ImageSubscript int
}

type FlagSettings struct {
	Browser string
	Launch  bool
	Random  bool
	Rest    []string
}
