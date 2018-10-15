package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
	"io/ioutil"
	"net/http"
	"time"
)

// BasicGet is a simple wrapper over HTTP client
func BasicGet(url string) (int, string, error) {
	return Get(GetParams{URL: url})
}

// Get performs GET request according to passed parameters
func Get(params GetParams) (int, string, error) {

	if params.URL == "" {
		return 0, "", errors.New("Request URL missing")
	}

	// Making logger
	log := xray.ROOT.Fork().WithLogger("http-client").With(args.URL(params.URL))

	// Building HTTP client
	client := http.Client{}

	// Building request
	log.Trace("Making GET request to :url")
	before := time.Now()
	req, err := http.NewRequest("GET", params.URL, nil)
	if err != nil {
		log.Warning("Unable to build request. Maybe URL (:url) is incorrect - :err", args.Error{Err: err})
		return -1, "", err
	}
	req.Header.Add("User-Agent", "Charlie Changelog Agent (v0.1-alpha)")

	for k, v := range params.Headers {
		req.Header.Add(k, v)
	}

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

// Only200 takes result from HTTP get and returns error
// it HTTP status code not 200
// If there were error already, it returns existing one
func Only200(code int, body string, err error) (string, error) {
	if err != nil {
		return body, err
	}
	if code != 200 {
		return body, fmt.Errorf("expected HTTP 200 but got %d", code)
	}

	return body, err
}

// IntoJSON builds and returns lambda, that will decode from JSON
// results, obtained from functions like Only200
func IntoJSON(target interface{}) func(string, error) error {
	return func(body string, err error) error {
		if err != nil {
			return err
		}
		return json.Unmarshal([]byte(body), target)
	}
}

// GetParams contains all the parameters for get request
type GetParams struct {
	URL     string
	Headers map[string]string
}
