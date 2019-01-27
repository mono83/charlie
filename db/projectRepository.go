package db

import (
	"github.com/mono83/charlie/model"
)

// ProjectRepository represents repository contract for Project models
type ProjectRepository interface {
	// GetByID returns Project by its database identifier
	GetByID(id int64) (*model.Project, error)
	// GetByName returns Project by its name
	GetByName(name string) (*model.Project, error)
}
