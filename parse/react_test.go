package parse

import (
	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var reactLog = `
## 16.5.1 (September 13, 2018)

### React

* Improve the warning when 'React.forwardRef' receives an unexpected number of arguments. ([@andresroberto](https://github.com/andresroberto) in [#13636](https://github.com/facebook/react/issues/13636))

### React DOM

* Fix a regression in unstable exports used by React Native Web. ([@aweary](https://github.com/aweary) in [#13598](https://github.com/facebook/react/issues/13598))
* Fix a crash when component defines a method called 'isReactComponent'. ([@gaearon](https://github.com/gaearon) in [#13608](https://github.com/facebook/react/issues/13608))
* Fix a crash in development mode in IE9 when printing a warning. ([@link-alex](https://github.com/link-alex) in [#13620](https://github.com/facebook/react/issues/13620))
* Provide a better error message when running 'react-dom/profiling' with 'schedule/tracking'. ([@bvaughn](https://github.com/bvaughn) in [#13605](https://github.com/facebook/react/issues/13605))
* If a 'ForwardRef' component defines a 'displayName', use it in warnings. ([@probablyup](https://github.com/probablyup) in [#13615](https://github.com/facebook/react/issues/13615))

### Schedule (Experimental)

* Add a separate profiling entry point at 'schedule/tracking-profiling'. ([@bvaughn](https://github.com/bvaughn) in [#13605](https://github.com/facebook/react/issues/13605))
## 16.4.2 (August 1, 2018)

### React DOM Server

* Fix a [potential XSS vulnerability when the attacker controls an attribute name](https://reactjs.org/blog/2018/08/01/react-v-16-4-2.html) ('CVE-2018-6341'). This fix is available in the latest 'react-dom@16.4.2', as well as in previous affected minor versions: 'react-dom@16.0.1', 'react-dom@16.1.2', 'react-dom@16.2.1', and 'react-dom@16.3.3'. ([@gaearon](https://github.com/gaearon) in [#13302](https://github.com/facebook/react/pull/13302))

* Fix a crash in the server renderer when an attribute is called 'hasOwnProperty'. This fix is only available in 'react-dom@16.4.2'. ([@gaearon](https://github.com/gaearon) in [#13303](https://github.com/facebook/react/pull/13303))
`

func TestReactChangelog(t *testing.T) {
	releases, err := ReactChangelog("", reactLog)

	if assert.NoError(t, err) {
		assert.Len(t, releases, 2)

		assert.Equal(t, model.Version{Major: "16", Minor: "5", Patch: "1", Label: ""}, releases[0].Version)
		if assert.Len(t, releases[0].Issues, 7) {
			assert.Equal(t, []string{"React"}, releases[0].Issues[0].Components)
			assert.Equal(t, []string{"React DOM"}, releases[0].Issues[1].Components)
			assert.Len(t, releases[0].Issues[6].Components, 0)
		}

		summary := releases[0].SummaryType()
		if assert.Len(t, summary, 4) {
			assert.Equal(t, 1, summary[model.Info])
			assert.Equal(t, 2, summary[model.Added])
			assert.Equal(t, 1, summary[model.Performance])
			assert.Equal(t, 3, summary[model.Fixed])
		}

		date, err := time.Parse("2006-01-02", "2018-09-13")
		assert.NoError(t, err)
		assert.Equal(t, date, releases[0].Date)

		// Next release
		assert.Equal(t, model.Version{Major: "16", Minor: "4", Patch: "2", Label: ""}, releases[1].Version)
		if assert.Len(t, releases[1].Issues, 2) {
			assert.Equal(t, []string{"React DOM Server"}, releases[1].Issues[0].Components)
			assert.Equal(t, []string{"React DOM Server"}, releases[1].Issues[1].Components)
		}

		summary = releases[1].SummaryType()
		if assert.Len(t, summary, 1) {
			assert.Equal(t, 2, summary[model.Fixed])
		}

		date, err = time.Parse("2006-01-02", "2018-08-01")
		assert.NoError(t, err)
		assert.True(t, releases[1].Date.Equal(date))
	}
}
