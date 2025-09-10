package internal

// ImageStore holds the list of image file paths
type ImageStore struct {
	Images         []string
	ImageSubscript int
	Fadetime       int
	Holdtime       int
	Title          string
}

type FlagSettings struct {
	Browser string
	Launch  bool
	Random  bool
	Rest    []string
}
