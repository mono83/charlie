package drivers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var repoNames = []struct {
	name  string
	valid bool
}{
	{"mono83", true},
	{"Mono83", true},
	{"mo-no83", true},
	{"mono--83", true},
	{"-mono83", true},
	{"mono83-", true},
	{"mon..o83", true},
	{"mono!83", false},
	{"mo#no83", false},
	{"mo@no83", false},
	{"mo$no83", false},
	{"mo%no83", false},
	{"mo^no83", false},
	{"mo&no83", false},
	{"mo*o83", false},
	{"mo(no83", false},
	{"mo)no83", false},
	{"mo+no83", false},
	{"mo=no83", false},
	{"moРусскиЙno83", false},
	{"mo[no83", false},
	{"mo]no83", false},
	{"mo}no83", false},
	{"mo{no83", false},
	{"mo|no83", false},
	{"mo\no83", false},
}

func TestIsValidGithubRepoName(t *testing.T) {
	for _, data := range repoNames {
		t.Run("Validating_"+data.name, func(t *testing.T) {
			assert.Equal(t, data.valid, isValidGithubRepoName(data.name))
		})
	}
}

var userNames = []struct {
	name  string
	valid bool
}{
	{"mono83", true},
	{"Mono83", true},
	{"mo-no83", true},
	{"mono--83", false},
	{"-mono83", false},
	{"mono83-", false},
	{"mon..o83", false},
	{"mono!83", false},
	{"mo#no83", false},
	{"mo@no83", false},
	{"mo$no83", false},
	{"mo%no83", false},
	{"mo^no83", false},
	{"mo&no83", false},
	{"mo*o83", false},
	{"mo(no83", false},
	{"mo)no83", false},
	{"mo+no83", false},
	{"mo=no83", false},
	{"moРусскиЙno83", false},
	{"mo[no83", false},
	{"mo]no83", false},
	{"mo}no83", false},
	{"mo{no83", false},
	{"mo|no83", false},
	{"mo\no83", false},
}

func TestIsValidGithubUserName(t *testing.T) {
	for _, data := range userNames {
		t.Run("Validating_"+data.name, func(t *testing.T) {
			assert.Equal(t, data.valid, isValidGithubUserName(data.name))
		})
	}
}

var repositories = []struct {
	name  string
	valid bool
}{
	{"mono83/charlie", true},
	{"mono83//charlie", false},
	{"mono83/charlie/barlie", false},
}

func TestIsValidGithubRepository(t *testing.T) {
	for _, data := range repositories {
		t.Run("Validating_"+data.name, func(t *testing.T) {
			assert.Equal(t, data.valid, isValidGithubRepository(data.name))
		})
	}
}
