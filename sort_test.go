package charlie

import (
	"testing"

	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
)

func TestSortByVersionAsc(t *testing.T) {
	releases := []model.Release{
		{Version: model.Version{Major: "3", Minor: "2"}},
		{Version: model.Version{Major: "3", Minor: "1"}},
		{Version: model.Version{Major: "4"}},
		{Version: model.Version{Major: "1"}},
		{Version: model.Version{Major: "3", Minor: "2", Patch: "1"}},
	}

	SortByVersion(releases)
	assert.Equal(t, model.Version{Major: "1"}, releases[0].Version)
	assert.Equal(t, model.Version{Major: "3", Minor: "1"}, releases[1].Version)
	assert.Equal(t, model.Version{Major: "3", Minor: "2"}, releases[2].Version)
	assert.Equal(t, model.Version{Major: "3", Minor: "2", Patch: "1"}, releases[3].Version)
	assert.Equal(t, model.Version{Major: "4"}, releases[4].Version)
}

func TestSortByVersionDesc(t *testing.T) {
	releases := []model.Release{
		{Version: model.Version{Major: "3", Minor: "2"}},
		{Version: model.Version{Major: "3", Minor: "1"}},
		{Version: model.Version{Major: "4"}},
		{Version: model.Version{Major: "1"}},
		{Version: model.Version{Major: "3", Minor: "2", Patch: "1"}},
	}

	SortByVersionDesc(releases)
	assert.Equal(t, model.Version{Major: "4"}, releases[0].Version)
	assert.Equal(t, model.Version{Major: "3", Minor: "2", Patch: "1"}, releases[1].Version)
	assert.Equal(t, model.Version{Major: "3", Minor: "2"}, releases[2].Version)
	assert.Equal(t, model.Version{Major: "3", Minor: "1"}, releases[3].Version)
	assert.Equal(t, model.Version{Major: "1"}, releases[4].Version)
}
