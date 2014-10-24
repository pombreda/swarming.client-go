// Copyright 2014 Marc-Antoine Ruel. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/maruel/ofh"
)

type cacheData struct {
	Apps map[string]ofh.InstalledApp
}

// Cache is an OAuth2 client token cache.
type Cache struct {
	path string
	data cacheData
}

// LoadCache loads a Cache from disk.
func LoadCache(path string) *Cache {
	// Load the file. Can't use ~/.isolated_oauth because it's too different.
	o := &Cache{path: path}
	if path != "" {
		readJSONFile(path, &o.data)
	}
	if o.data.Apps == nil {
		o.data.Apps = make(map[string]ofh.InstalledApp)
	}
	return o
}

// Save saves the cache back to disk.
func (c *Cache) Save() {
	if c.path != "" {
		writeJSONFile(c.path, c.data)
	}
}

// ConnectToAuthService creates a preloaded 'auth_service' compatible
// http.Client.
func (c *Cache) ConnectToAuthService(host string) (*http.Client, error) {
	app, ok := c.data.Apps[host]
	if !ok {
		v := &struct {
			ClientID          string `json:"client_id"`
			ClientNotSoSecret string `json:"client_not_so_secret"`
			PrimaryURL        string `json:"primary_url"`
		}{}
		err := GetJSON(nil, host+"/auth/api/v1/server/oauth_config", v)
		if err != nil {
			return nil, err
		}
		app = *ofh.MakeInstalledApp()
		app.ClientID = v.ClientID
		app.ClientSecret = v.ClientNotSoSecret
		// TODO(maruel): What about v.PrimaryURL?
		c.data.Apps[host] = app
	}
	client, err := app.GetClient("https://www.googleapis.com/auth/userinfo.email", nil)
	return client, err
}

// GetJSON does an HTTP GET on a JSON endpoint.
func GetJSON(c *http.Client, URL string, v interface{}) error {
	if c == nil {
		c = http.DefaultClient
	}
	log.Printf("GET %s", URL)
	resp, err := c.Get(URL)
	if err != nil {
		return fmt.Errorf("Couldn't resolve %s: %s", URL, err)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("HTTP Status %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(v); err != nil {
		return fmt.Errorf("Unexected response %s: %s", URL, err)
	}
	return nil
}

func readJSONFile(filePath string, object interface{}) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filePath, err)
	}
	defer func() {
		_ = f.Close()
	}()
	if err = json.NewDecoder(f).Decode(object); err != nil {
		return fmt.Errorf("failed to decode %s: %s", filePath, err)
	}
	return nil
}

// writeJSONFile writes object as json encoded into filePath with 2 spaces indentation.
func writeJSONFile(filePath string, object interface{}) error {
	d, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode %s: %s", filePath, err)
	}

	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("failed to open %s: %s", filePath, err)
	}
	defer func() {
		_ = f.Close()
	}()
	if _, err := f.Write(d); err != nil {
		return fmt.Errorf("failed to write %s: %s", filePath, err)
	}
	return nil
}
