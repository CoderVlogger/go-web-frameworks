package pkg

import (
	"fmt"
	"sync"
)

var (
	ErrEntityAlreadyExists = fmt.Errorf("entity already exists")
)

type EntityRepository interface {
	Init()
	Add(entity *Entity) error
	Update(entity *Entity) error
	Get(id string) (*Entity, error)
	List() ([]*Entity, error)
	Delete(id string) error
}

type EntityMemoryRepository struct {
	mutex    sync.RWMutex
	entities []*Entity
}

func NewEntityMemoryRepository() *EntityMemoryRepository {
	return &EntityMemoryRepository{
		entities: []*Entity{},
	}
}

func (r *EntityMemoryRepository) Init() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.entities = []*Entity{
		{
			ID:          "1",
			Type:        PersonEntityType,
			Name:        "John Doe",
			Description: "John Doe is a person",
		},
		{
			ID:          "2",
			Type:        CompanyEntityType,
			Name:        "Google",
			Description: "Google is a company",
		},
		{
			ID:          "3",
			Type:        PlaceEntityType,
			Name:        "New York",
			Description: "New York is a place",
		},
		{
			ID:          "4",
			Type:        BookEntityType,
			Name:        "The Hitchhiker's Guide to the Galaxy",
			Description: "The Hitchhiker's Guide to the Galaxy is a book",
		},
		{
			ID:          "5",
			Type:        MovieEntityType,
			Name:        "Star Wars",
			Description: "Star Wars is a movie",
		},
		{
			ID:          "6",
			Type:        TvSeriesEntityType,
			Name:        "Game of Thrones",
			Description: "Game of Thrones is a tv series",
		},
		{
			ID:          "7",
			Type:        GameEntityType,
			Name:        "Minecraft",
			Description: "Minecraft is a game",
		},
		{
			ID:          "8",
			Type:        AlbumEntityType,
			Name:        "The Beatles",
			Description: "The Beatles is an album",
		},
		{
			ID:          "9",
			Type:        SongEntityType,
			Name:        "Yesterday",
			Description: "Yesterday is a song",
		},
	}
}

func (r *EntityMemoryRepository) Add(entity *Entity) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, e := range r.entities {
		if e.ID == entity.ID {
			return ErrEntityAlreadyExists
		}
	}

	r.entities = append(r.entities, entity)

	return nil
}

func (r *EntityMemoryRepository) Update(entity *Entity) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, e := range r.entities {
		if e.ID == entity.ID {
			r.entities[i] = entity
			return nil
		}
	}

	return nil
}

func (r *EntityMemoryRepository) Get(id string) (*Entity, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, e := range r.entities {
		if e.ID == id {
			return e, nil
		}
	}

	return nil, nil
}

func (r *EntityMemoryRepository) List() ([]*Entity, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	return r.entities, nil
}

func (r *EntityMemoryRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for i, e := range r.entities {
		if e.ID == id {
			r.entities = append(r.entities[:i], r.entities[i+1:]...)
			return nil
		}
	}

	return nil
}
