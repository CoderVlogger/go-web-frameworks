package pkg

type (
	TextResponse struct {
		Message string `json:"message"`
	}

	Entity struct {
		ID          string     `json:"id"`
		Type        EntityType `json:"type" form:"type"`
		Name        string     `json:"name" form:"name"`
		Description string     `json:"description" form:"description"`
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
