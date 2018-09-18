package drivers

import (
	"errors"
	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"io/ioutil"
	"net/http"
	"time"
)

// HTTPGet is a simple wrapper over HTTP client
func HTTPGet(url string) (int, string, error) {
	// Making logger
	log := xray.ROOT.Fork().WithLogger("http-client").With(args.URL(url))

	// Building HTTP client
	client := http.Client{}

	// Building request
	log.Trace("Making GET request to :url")
	before := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Warning("Unable to build request. Maybe URL (:url) is incorrect - :err", args.Error{Err: err})
		return -1, "", err
	}
	req.Header.Add("User-Agent", "Charlie Changelog Agent (v0.1-alpha)")

	// Making request
	res, err := client.Do(req)
	if err != nil {
		log.Warning("HTTP request to :url failed - :err", args.Error{Err: err})
		return -1, "", err
	}
	defer res.Body.Close()

	// Reading response bytes
	bts, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Warning("Unable to read response body - :err", args.Error{Err: err})
		return res.StatusCode, "", err
	}
	if len(bts) == 0 {
		log.Warning("Response is empty")
		return res.StatusCode, "", errors.New("empty response")
	}
	log.InBytes(bts)
	log.Info("HTTP GET :url done in :delta", args.Delta(time.Now().Sub(before)))

	return res.StatusCode, string(bts), nil
}
