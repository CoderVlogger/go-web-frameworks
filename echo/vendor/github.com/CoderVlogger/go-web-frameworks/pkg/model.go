package pkg

type (
	TextResponse struct {
		Message string `json:"message"`
	}

	Entity struct {
		ID          string     `json:"id"`
		Type        EntityType `json:"type"`
		Name        string     `json:"name"`
		Description string     `json:"description"`
	}

	EntityType string
)

const (
	UknownEntityType   EntityType = "unknown"
	PersonEntityType   EntityType = "person"
	CompanyEntityType  EntityType = "company"
	PlaceEntityType    EntityType = "place"
	BookEntityType     EntityType = "book"
	MovieEntityType    EntityType = "movie"
	TvSeriesEntityType EntityType = "tv_series"
	GameEntityType     EntityType = "game"
	AlbumEntityType    EntityType = "album"
	SongEntityType     EntityType = "song"
)
