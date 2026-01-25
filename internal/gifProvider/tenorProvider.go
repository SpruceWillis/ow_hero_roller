package gifProvider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type TenorProvider struct {
	ApiKey string
}

const tenorSearchUrl = "https://tenor.googleapis.com/v2/search"

func NewTenorProvider(apiKey string) *TenorProvider {
	return &TenorProvider{
		ApiKey: apiKey,
	}
}

func (t *TenorProvider) GetGifUrl(heroName string) (string, error) {
	queryString := fmt.Sprintf("%v overwatch", heroName)
	req, err := http.NewRequest("GET", tenorSearchUrl, nil)
	if err != nil {
		return "", err
	}
	queryParams := map[string]string{
		"q":             queryString,
		"key":           t.ApiKey,
		"locale":        "en_US",
		"media_filter":  "gif",
		"limit":         "1",
		"contentFilter": "off",
		"ar_range":      "all",
		"random":        "true",
	}
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	req.Close = false
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("error requesting gif %v\n", err)
		return "", err
	}
	return getGifFromTenorJson(resBody)
}

func getGifFromTenorJson(jsonBytes []byte) (string, error) {
	var result map[string]any
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		fmt.Println("error parsing JSON", err)
	}
	results, ok := result["results"]
	if !ok {
		return "", fmt.Errorf("no results found in json bytes %v", string(jsonBytes[:]))
	}
	resultArray := results.([]interface{})
	if len(resultArray) == 0 {
		return "", fmt.Errorf("no results found in json bytes %v", string(jsonBytes[:]))
	}
	mediaFormats := resultArray[0].(map[string]any)["media_formats"]
	gifBlock, ok := mediaFormats.(map[string]any)["gif"]
	if !ok {
		return "", fmt.Errorf("no gif results found in first result of json bytes %v", string(jsonBytes[:]))
	}
	gifUrl, ok := gifBlock.(map[string]any)["url"]
	if !ok {
		return "", fmt.Errorf("no gif url found in first result of json bytes %v", string(jsonBytes[:]))
	}
	stringUrl := gifUrl.(string)
	return stringUrl, nil
}

func (t *TenorProvider) EmbedMessage(heroName string) string {
	return heroName
}
