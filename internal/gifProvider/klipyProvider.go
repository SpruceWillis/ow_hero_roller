package gifProvider

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type KlipyProvider struct {
	ApiKey string
	Limit  int
}

const klipySearchUrl = "https://api.klipy.com/v2/search"

func NewKlipyProvider(apiKey string, limit int) *KlipyProvider {
	return &KlipyProvider{
		ApiKey: apiKey,
		Limit:  limit,
	}
}

func (k *KlipyProvider) GetGifUrl(heroName string) (string, error) {
	queryString := fmt.Sprintf("%v overwatch", heroName)
	req, err := http.NewRequest("GET", klipySearchUrl, nil)
	if err != nil {
		return "", err
	}
	queryParams := map[string]string{
		"q":             queryString,
		"key":           k.ApiKey,
		"country":       "US",
		"locale":        "en_US",
		"contentfilter": "off",
		"format_filter": "gif",
		"limit":         strconv.Itoa(k.Limit),
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
	return getGifFromKlipyJson(resBody)
}

func getGifFromKlipyJson(jsonBytes []byte) (string, error) {
	var result map[string]any
	err := json.Unmarshal(jsonBytes, &result)
	if err != nil {
		fmt.Println("error parsing JSON", err)
	}
	results, ok := result["results"].([]any)
	if !ok || len(results) == 0 {
		return "", fmt.Errorf("no results found in json bytes %v", string(jsonBytes[:]))
	}
	randIndex := rand.Intn(len(results))
	resultGif := results[randIndex]
	mediaFormats := resultGif.(map[string]any)["media_formats"]
	mp4Block, ok := mediaFormats.(map[string]any)["gif"]
	if !ok {
		return "", fmt.Errorf("no gif results found in first result of json bytes %v", string(jsonBytes[:]))
	}
	url, ok := mp4Block.(map[string]any)["url"]
	if !ok {
		return "", fmt.Errorf("no gif results found in first result of json bytes %v", string(jsonBytes[:]))
	}
	return url.(string), nil
}

func (k *KlipyProvider) EmbedMessage(heroName string) string {
	return fmt.Sprintf("%v (Powered by Klipy)", heroName)
}
