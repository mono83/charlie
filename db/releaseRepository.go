package db

import (
	"github.com/mono83/charlie/model"
)

// ReleaseRepository represents repository contract of Release models
type ReleaseRepository interface {
	// GetByID returns Release by its database identifier
	GetByID(id int64) (*model.Release, error)
	// GetByProjectID returns Releases by project identifier
	GetByProjectID(projectId int64) ([]*model.Release, error)
}
