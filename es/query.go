package es

import (
	"bytes"
	"encoding/json"
	"errors"

	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
	esapiv6 "github.com/elastic/go-elasticsearch/v6/esapi"
	"github.com/robherley/lumberjack/util"
)

// SearchOpts are the options to pass to elasticsearch's client search
type SearchOpts struct {
	// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-search.html
	Params []func(*esapiv6.SearchRequest)
	// https://www.elastic.co/guide/en/elasticsearch/reference/6.8/search-request-body.html
	Body *map[string]interface{}
}

// Search is a slight wrapper around elastic's search function
func Search(client *elasticsearch6.Client, opts *SearchOpts) (response *map[string]interface{}, err error) {
	searchParams := opts.Params
	if opts.Body != nil {
		var buf bytes.Buffer
		if err := json.NewEncoder(&buf).Encode(opts.Body); err != nil {
			return nil, err
		}
		searchParams = append(searchParams, client.Search.WithBody(&buf))
	}

	res, err := client.Search(searchParams...)
	defer res.Body.Close()

	if err != nil {
		return nil, err
	}

	if res.IsError() {
		var errResponse map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&errResponse); err != nil {
			return nil, err
		}
		util.Log.Warnln(errResponse)
		return nil, errors.New("error/invalid response from elasticsearch")
	}

	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	return response, nil
}
