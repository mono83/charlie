package db

import (
	"github.com/mono83/charlie/model"
)

// ProjectRepository represents the projects's repository contract
type ProjectRepository interface {
	GetByID(id int64) (*model.Project, error)
	GetByName(title string) (*model.Project, error)
}
