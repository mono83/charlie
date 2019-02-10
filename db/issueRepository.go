package db

import (
	"github.com/mono83/charlie/model"
)

// IssueRepository represents repository contract for Issue models
type IssueRepository interface {
	// GetByID returns Issue by its database identifier
	GetByID(id int64) (*model.Issue, error)
	// GetByReleaseId returns Issues by Release identifier
	GetByReleaseId(releaseId int64) ([]*model.Issue, error)
	// Store persists Issue model and returns generated issue identifier
	Store(issues []*model.Issue) (error)

}
