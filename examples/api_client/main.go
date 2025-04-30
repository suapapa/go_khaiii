package main

import (
	"bytes"
	"cmp"
	"encoding/json"
	"net/http"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/suapapa/go_khaiii/pkg/khaiiitype"
)

var (
	text            = "사랑은 모든것을 덮어주고 모든것을 믿으며 모든것을 바라고 모든것을 견디어냅니다"
	khaiiiAnalyzeEP = "http://homin.dev/khaiii-api/v1/analyze"
	// khaiiiAnalyzeEP = "http://localhost:8082/v1/analyze"
	secret = cmp.Or(os.Getenv("KHAIII_API_TOKEN"), "")
)

func main() {
	if len(os.Args) > 1 {
		text = os.Args[1]
	}

	data, err := json.Marshal(map[string]string{
		"Text": text,
	})
	if err != nil {
		panic(err)
	}
	dataReader := bytes.NewReader(data)

	req, err := http.NewRequest("POST", khaiiiAnalyzeEP, dataReader)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	if secret != "" {
		req.Header.Set("Authorization", "Bearer "+secret)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	type RespData struct {
		Data khaiiitype.AnalyzeResult `json:"data"`
	}
	var respData RespData
	err = yaml.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		panic(err)
	}

	yamlRespData, err := yaml.Marshal(respData)
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(yamlRespData)
}
