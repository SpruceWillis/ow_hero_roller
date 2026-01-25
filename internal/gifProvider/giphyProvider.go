package gifProvider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GiphyProvider struct {
	ApiKey string
}

func NewGiphyProvider(apiKey string) *GiphyProvider {
	return &GiphyProvider{
		ApiKey: apiKey,
	}
}

const giphySearchUrl = "https://api.giphy.com/v1/gifs/search"

func (g *GiphyProvider) GetGifUrl(heroName string) (string, error) {
	queryString := fmt.Sprintf("%v overwatch", heroName)
	req, err := http.NewRequest("GET", giphySearchUrl, nil)
	if err != nil {
		return "", err
	}
	queryParams := map[string]string{
		"api_key":             g.ApiKey,
		"q":                   queryString,
		"limit":               "1",
		"remove_low_contrast": "true",
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
	return getGifFromGiphyJson(resBody)
}

func (g *GiphyProvider) EmbedMessage(heroName string) string {
	return fmt.Sprintf("%v (Powered by GIPHY)", heroName)
}

func getGifFromGiphyJson(jsonBytes []byte) (string, error) {
	var result map[string]any
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return "error parsing JSON", err
	}
	data, ok := result["data"]
	if !ok {
		return "", fmt.Errorf("no data found in json bytes %v", string(jsonBytes[:]))
	}
	dataArray := data.([]any)
	if len(dataArray) == 0 {
		return "", fmt.Errorf("no data found in json bytes %v", string(jsonBytes[:]))
	}
	gifObject, ok := dataArray[0].(map[string]any)
	if !ok {
		return "", fmt.Errorf("no gif found in json bytes %v", string(jsonBytes[:]))
	}
	gifUrl, ok := gifObject["images"].([]any)[0].(map[string]any)["downsized"].(map[string]any)["url"].(string)
	if !ok {
		return "", fmt.Errorf("no gif url found in json bytes %v", string(jsonBytes[:]))
	}

	return gifUrl, nil
}
