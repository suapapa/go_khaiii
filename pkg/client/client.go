package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/suapapa/khaiii-api/pkg/khaiiitype"
)

type Client struct {
	BaseURL string
	ApiKey  string
	Client  *http.Client
}

type WordChunk khaiiitype.WordChunk

func (c *Client) Analyze(input string) ([]*WordChunk, error) {
	data, err := json.Marshal(map[string]string{
		"Text": input,
	})
	if err != nil {
		return nil, err
	}
	dataReader := bytes.NewReader(data)

	req, err := http.NewRequest("POST", c.BaseURL, dataReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	if c.ApiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	}

	var client *http.Client
	if c.Client != nil {
		client = c.Client
	} else {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("status code: %d", resp.StatusCode)
	}

	type RespData struct {
		Data khaiiitype.AnalyzeResult `json:"data"`
	}
	var respData RespData
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		return nil, err
	}

	ret := make([]*WordChunk, 0, len(respData.Data.WordChunks))
	for _, wc := range respData.Data.WordChunks {
		ret = append(ret, (*WordChunk)(wc))
	}

	return ret, nil
}
