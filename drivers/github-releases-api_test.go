package drivers

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var repoNames = []struct{
	name string
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
func TestIsValidGithubRepoName(t *testing.T){
	for _, data := range repoNames {
		t.Run("Validating_"+data.name, func(t *testing.T) {
			assert.Equal(t, data.valid, isValidGithubRepoName(data.name))
		})
	}
}