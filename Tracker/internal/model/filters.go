package model

type FilterParams struct {
	Search   string
	Location string

	CreationFrom int
	CreationTo   int

	AlbumFrom int
	AlbumTo   int

	Members []int

	Locations []string

	LocationCount int
	ConcertCount  int
}