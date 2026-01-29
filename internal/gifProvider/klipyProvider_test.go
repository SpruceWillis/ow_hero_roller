package gifProvider

import "testing"

var klipyJsonResponse = `{
  "results": [
    {
      "id": "4551195970372378",
      "title": "Greetings: Man Waving Hello",
      "media_formats": {
        "gif": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/p30ORxL1.gif",
          "duration": 0,
          "preview": "",
          "dims": [
            498,
            498
          ],
          "size": 2614179
        },
        "mediumgif": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/OXB1QWhn.gif",
          "duration": 0,
          "preview": "",
          "dims": [
            640,
            640
          ],
          "size": 873745
        },
        "tinygif": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/NIKSFkmQ.gif",
          "duration": 0,
          "preview": "",
          "dims": [
            220,
            220
          ],
          "size": 149153
        },
        "nanogif": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/rxFdVqew.gif",
          "duration": 0,
          "preview": "",
          "dims": [
            90,
            90
          ],
          "size": 32170
        },
        "gifpreview": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/283NoyuyEhLxSJz3qa.jpg",
          "duration": 0,
          "preview": "",
          "dims": [
            640,
            640
          ],
          "size": 25496
        },
        "tinygifpreview": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/iKL40aAFcLpJB9S5.jpg",
          "duration": 0,
          "preview": "",
          "dims": [
            220,
            220
          ],
          "size": 8676
        },
        "nanogifpreview": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/FAWLnCWDS75vMHGF2X3.jpg",
          "duration": 0,
          "preview": "",
          "dims": [
            90,
            90
          ],
          "size": 2974
        },
        "webp": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/0fbcxcN4.webp",
          "duration": 0,
          "preview": "",
          "dims": [
            640,
            640
          ],
          "size": 435968
        },
        "webm": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/tPTbV6ndVZ55CS6Az.webm",
          "duration": 0,
          "preview": "",
          "dims": [
            640,
            640
          ],
          "size": 74085
        },
        "tinywebm": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/wLpNEwIiCTwSYwHnRRRs.webm",
          "duration": 0,
          "preview": "",
          "dims": [
            320,
            320
          ],
          "size": 43351
        },
        "nanowebm": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/QBc76qLbuGNOks.webm",
          "duration": 0,
          "preview": "",
          "dims": [
            150,
            150
          ],
          "size": 44283
        },
        "mp4": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/kRvr7yGzX0V876EitO2Z.mp4",
          "duration": 0,
          "preview": "",
          "dims": [
            640,
            640
          ],
          "size": 369164
        },
        "loopedmp4": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/kRvr7yGzX0V876EitO2Z.mp4",
          "duration": 0,
          "preview": "",
          "dims": [
            640,
            640
          ],
          "size": 369164
        },
        "tinymp4": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/wbRY32dCtajM2l0LCj.mp4",
          "duration": 0,
          "preview": "",
          "dims": [
            320,
            320
          ],
          "size": 115924
        },
        "nanomp4": {
          "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/RJJSBhTg2yPQd.mp4",
          "duration": 0,
          "preview": "",
          "dims": [
            150,
            150
          ],
          "size": 42419
        }
      },
      "created": 1765483200,
      "content_description": "",
      "itemurl": "https://klipy.com/gifs/greetings-PSr",
      "url": "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/p30ORxL1.gif",
      "tags": [
        "greetings"
      ],
      "flags": [],
      "hasaudio": false,
      "content_description_source": ""
    }
  ],
  "next": "eyJtb2RlIjoicHJlIiwib2Zmc2V0Ijo4fQ=="
}`

func TestGetGifFromKlipyJson(t *testing.T) {
	stringBytes := []byte(klipyJsonResponse)
	result, err := getGifFromKlipyJson(stringBytes)
	expectedResult := "https://static.klipy.com/ii/d7aec6f6f171607374b2065c836f92f4/ec/f3/kRvr7yGzX0V876EitO2Z.mp4"
	if err != nil {
		t.Errorf("error getting mp4 from klipy json: %v", err)
	} else if result != expectedResult {
		t.Errorf("expected url to be %v, got %v", expectedResult, result)
	}
}
