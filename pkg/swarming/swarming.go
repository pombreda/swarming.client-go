// Copyright 2014 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package swarming

import (
	"net/http"
	"net/url"
	"sync"

	"github.com/maruel/swarming.client-go/pkg/common"
)

// Swarming defines a Swarming client.
type Swarming struct {
	host   string
	client *http.Client
}

func (s *Swarming) getJSON(resource string, v interface{}) error {
	return common.GetJSON(s.client, s.host+resource, v)
}

// NewSwarming returns a new Swarming client with oauth2 token preloaded from
// the Cache if possible.
func NewSwarming(host string, c *common.Cache) (*Swarming, error) {
	if c == nil {
		c = common.LoadCache("")
	}
	client, err := c.ConnectToAuthService(host)
	if err != nil {
		return nil, err
	}
	return &Swarming{host, client}, nil
}

// RequestByID returns the TaskRequest.
func (s *Swarming) RequestByID(id string) (TaskRequest, error) {
	out := TaskRequest{}
	err := s.getJSON("/swarming/api/v1/client/task/"+id+"/request", &out)
	return out, err
}

// TasksByTag returns the results corresponding to a tag.
func (s *Swarming) TasksByTag(out chan<- TaskResult, tags ...string) error {
	// TODO(maruel): Cursor support.
	u := "/swarming/api/v1/client/tasks?limit=500"
	for _, tag := range tags {
		u += "&tag=" + url.QueryEscape(tag)
	}
	v := &struct {
		Items []TaskResult `json:"items"`
	}{}
	err := s.getJSON(u, v)
	if err == nil {
		var wg sync.WaitGroup
		for _, item := range v.Items {
			wg.Add(1)
			go func(item TaskResult) {
				defer wg.Done()
				item.TaskRequest, err = s.RequestByID(item.ID)
				if err == nil {
					out <- item
				}
			}(item)
		}
		wg.Wait()
	}
	return err
}

// TaskRequestProperties describes the idempotent properties of a task.
type TaskRequestProperties struct {
	Commands             [][]string        `json:"commands"`
	Data                 [][]string        `json:"data"`
	Dimensions           map[string]string `json:"dimensions"`
	Env                  map[string]string `json:"env"`
	ExecutionTimeoutSecs int               `json:"execution_timeout_secs"`
	Idempotent           bool              `json:"idempotent"`
	IoTimeoutSecs        int               `json:"io_timeout_secs"`
}

// TaskRequest describes a complete request.
type TaskRequest struct {
	//"created_ts": "2014-10-24 00:00:00",
	//"expiration_ts": "2014-10-24 00:00:00",
	Name           string                `json:"name"`
	Priority       int                   `json:"priority"`
	Properties     TaskRequestProperties `json:"properties"`
	PropertiesHash string                `json:"properties_hash"`
	Tags           []string              `json:"tags"`
	User           string                `json:"user"`
}

// TaskResult describes the results of a task.
type TaskResult struct {
	TaskRequest TaskRequest
	//"abandoned_ts": "2014-10-24 00:00:00",
	BotID      string `json:"bot_id"`
	BotVersion string `json:"bot_version"`
	//"completed_ts": "2014-10-24 00:00:00",
	//"created_ts": "2014-10-24 00:00:00",
	DedupedFrom     string    `json:"deduped_from"`
	Durations       []float64 `json:"durations"`
	ExitCodes       []int     `json:"exit_codes"`
	Failure         bool      `json:"failure"`
	ID              string    `json:"id"`
	InternalFailure bool      `json:"internal_failure"`
	//"modified_ts": "2014-10-24 00:00:00",
	Name           string   `json:"name"`
	PropertiesHash string   `json:"properties_hash"`
	ServerVersions []string `json:"server_versions"`
	// "started_ts": "2014-10-24 00:00:00",
	State     int    `json:"state"`
	TryNumber int    `json:"try_number"`
	User      string `json:"user"`
}

// Duration returns the total duration of a task.
func (s *TaskResult) Duration() (out float64) {
	for _, d := range s.Durations {
		out += d
	}
	return
}
