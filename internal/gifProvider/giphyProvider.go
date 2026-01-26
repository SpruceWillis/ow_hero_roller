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
	gifUrl, ok := gifObject["images"].(map[string]any)["downsized"].(map[string]any)["url"].(string)
	if !ok {
		return "", fmt.Errorf("no gif url found in json bytes %v", string(jsonBytes[:]))
	}

	return gifUrl, nil
}

// TODO: fix this
/*
2026/01/25 19:05:40 http: panic serving 169.254.169.126:51406: interface conversion: interface {} is map[string]interface {}, not []interface {}
goroutine 33 [running]:
net/http.(*conn).serve.func1()
	/usr/local/go/src/net/http/server.go:1943 +0xd3
panic({0x7af8e0?, 0xc00022d470?})
	/usr/local/go/src/runtime/panic.go:783 +0x132
internal/gifProvider.getGifFromGiphyJson({0xc000266000, 0x2769, 0x3000})
	/usr/src/app/internal/gifProvider/giphyProvider.go:75 +0x354
internal/gifProvider.(*GiphyProvider).GetGifUrl(0xc0000fe380, {0xc0000129d0?, 0x824cd1?})
	/usr/src/app/internal/gifProvider/giphyProvider.go:50 +0x491
main.handleRollCommand(0xc0000100d8, 0x671?, {0x8b4e28, 0xc0000fe380})
	/usr/src/app/main.go:189 +0x44d
main.main.func1({0x8b6d50, 0xc000162000}, 0xc0000c6dc0)
	/usr/src/app/main.go:338 +0x28d
net/http.HandlerFunc.ServeHTTP(0xb41b20?, {0x8b6d50?, 0xc000162000?}, 0xc000476b58?)
	/usr/local/go/src/net/http/server.go:2322 +0x29
net/http.(*ServeMux).ServeHTTP(0x418925?, {0x8b6d50, 0xc000162000}, 0xc0000c6dc0)
	/usr/local/go/src/net/http/server.go:2861 +0x1c7
net/http.serverHandler.ServeHTTP({0x8b50a8?}, {0x8b6d50?, 0xc000162000?}, 0x6?)
	/usr/local/go/src/net/http/server.go:3340 +0x8e
net/http.(*conn).serve(0xc000114ab0, {0x8b73d0, 0xc000101500})
	/usr/local/go/src/net/http/server.go:2109 +0x665
created by net/http.(*Server).Serve in goroutine 1
	/usr/local/go/src/net/http/server.go:3493 +0x485
*/
