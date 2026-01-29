package gifProvider

import "testing"

var giphyJsonResponse = `{
  "data": [
    {
      "type": "gif",
      "id": "erePhJFWkfYMwTpNT8",
      "url": "https://giphy.com/gifs/happy-business-rudinihadi-erePhJFWkfYMwTpNT8",
      "slug": "happy-business-rudinihadi-erePhJFWkfYMwTpNT8",
      "bitly_gif_url": "https://gph.is/g/ZrAjep1",
      "bitly_url": "https://gph.is/g/ZrAjep1",
      "embed_url": "https://giphy.com/embed/erePhJFWkfYMwTpNT8",
      "username": "rudinihadi",
      "source": "",
      "title": "Happy Leonardo Dicaprio GIF",
      "rating": "g",
      "content_url": "",
      "source_tld": "",
      "source_post_url": "",
      "is_sticker": 0,
      "import_datetime": "2023-06-27 14:42:45",
      "trending_datetime": "0000-00-00 00:00:00",
      "images": {
        "original": {
          "height": "480",
          "width": "320",
          "size": "552689",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.gif",
          "mp4_size": "100277",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.mp4",
          "webp_size": "223100",
          "webp": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.webp",
          "frames": "83",
          "hash": "64c85a9a99fd467bede507d43139f2d0"
        },
        "downsized": {
          "height": "480",
          "width": "320",
          "size": "552689",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.gif"
        },
        "downsized_large": {
          "height": "480",
          "width": "320",
          "size": "552689",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.gif"
        },
        "downsized_medium": {
          "height": "480",
          "width": "320",
          "size": "552689",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.gif"
        },
        "downsized_small": {
          "height": "480",
          "width": "320",
          "mp4_size": "100277",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy-downsized-small.mp4"
        },
        "downsized_still": {
          "height": "480",
          "width": "320",
          "size": "552689",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy_s.gif"
        },
        "fixed_height": {
          "height": "200",
          "width": "134",
          "size": "138531",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200.gif",
          "mp4_size": "34189",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200.mp4",
          "webp_size": "94288",
          "webp": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200.webp"
        },
        "fixed_height_downsampled": {
          "height": "200",
          "width": "134",
          "size": "22232",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200_d.gif",
          "webp_size": "17446",
          "webp": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200_d.webp"
        },
        "fixed_height_small": {
          "height": "100",
          "width": "66",
          "size": "49030",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100.gif",
          "mp4_size": "14196",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100.mp4",
          "webp_size": "42232",
          "webp": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100.webp"
        },
        "fixed_height_small_still": {
          "height": "100",
          "width": "66",
          "size": "6574",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100_s.gif"
        },
        "fixed_height_still": {
          "height": "200",
          "width": "134",
          "size": "16535",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200_s.gif"
        },
        "fixed_width": {
          "height": "300",
          "width": "200",
          "size": "257493",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200w.gif",
          "mp4_size": "56434",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200w.mp4",
          "webp_size": "112860",
          "webp": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200w.webp"
        },
        "fixed_width_downsampled": {
          "height": "300",
          "width": "200",
          "size": "46766",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200w_d.gif",
          "webp_size": "32916",
          "webp": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200w_d.webp"
        },
        "fixed_width_small": {
          "height": "150",
          "width": "100",
          "size": "88320",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100w.gif",
          "mp4_size": "23732",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100w.mp4",
          "webp_size": "60472",
          "webp": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100w.webp"
        },
        "fixed_width_small_still": {
          "height": "150",
          "width": "100",
          "size": "11079",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100w_s.gif"
        },
        "fixed_width_still": {
          "height": "300",
          "width": "200",
          "size": "30125",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/200w_s.gif"
        },
        "looping": {
          "mp4_size": "266243",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy-loop.mp4"
        },
        "original_still": {
          "height": "480",
          "width": "320",
          "size": "59775",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy_s.gif"
        },
        "original_mp4": {
          "height": "480",
          "width": "320",
          "mp4_size": "100277",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.mp4"
        },
        "preview": {
          "height": "150",
          "width": "99",
          "mp4_size": "10504",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy-preview.mp4"
        },
        "preview_gif": {
          "height": "100",
          "width": "66",
          "size": "49030",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100.gif"
        },
        "preview_webp": {
          "height": "100",
          "width": "66",
          "size": "42232",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/100.webp"
        },
        "hd": {
          "height": "1080",
          "width": "720",
          "mp4_size": "433828",
          "mp4": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy-hd.mp4"
        },
        "480w_still": {
          "height": "720",
          "width": "480",
          "size": "552689",
          "url": "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/480w_s.jpg"
        }
      },
      "user": {
        "avatar_url": "https://media4.giphy.com/avatars/rudinihadi/TCCVpfeJRDLH.gif",
        "banner_image": "",
        "banner_url": "",
        "profile_url": "https://giphy.com/channel/rudinihadi/",
        "username": "rudinihadi",
        "display_name": "rudinihadi",
        "description": "Calm is Superpower",
        "instagram_url": "https://instagram.com/rudinihadi",
        "website_url": "http://sticker.ly/user/rudinihadi",
        "is_verified": false
      },
      "analytics_response_payload": "e=ZXZlbnRfdHlwZT1HSUZfU0VBUkNIJmNpZD01MjFlNTg4YjZ3OWtqNmhpam1wbjRtem04NnpxZm41c3huaWxiODNiYm13bWRrZWImZ2lmX2lkPWVyZVBoSkZXa2ZZTXdUcE5UOCZjdD1n",
      "analytics": {
        "onload": {
          "url": "https://giphy-analytics.giphy.com/v2/pingback_simple?analytics_response_payload=e%3DZXZlbnRfdHlwZT1HSUZfU0VBUkNIJmNpZD01MjFlNTg4YjZ3OWtqNmhpam1wbjRtem04NnpxZm41c3huaWxiODNiYm13bWRrZWImZ2lmX2lkPWVyZVBoSkZXa2ZZTXdUcE5UOCZjdD1n&action_type=SEEN"
        },
        "onclick": {
          "url": "https://giphy-analytics.giphy.com/v2/pingback_simple?analytics_response_payload=e%3DZXZlbnRfdHlwZT1HSUZfU0VBUkNIJmNpZD01MjFlNTg4YjZ3OWtqNmhpam1wbjRtem04NnpxZm41c3huaWxiODNiYm13bWRrZWImZ2lmX2lkPWVyZVBoSkZXa2ZZTXdUcE5UOCZjdD1n&action_type=CLICK"
        },
        "onsent": {
          "url": "https://giphy-analytics.giphy.com/v2/pingback_simple?analytics_response_payload=e%3DZXZlbnRfdHlwZT1HSUZfU0VBUkNIJmNpZD01MjFlNTg4YjZ3OWtqNmhpam1wbjRtem04NnpxZm41c3huaWxiODNiYm13bWRrZWImZ2lmX2lkPWVyZVBoSkZXa2ZZTXdUcE5UOCZjdD1n&action_type=SENT"
        }
      },
      "alt_text": "",
      "is_low_contrast": false
    }
  ],
  "meta": {
    "status": 200,
    "msg": "OK",
    "response_id": "6w9kj6hijmpn4mzm86zqfn5sxnilb83bbmwmdkeb"
  },
  "pagination": {
    "total_count": 500,
    "count": 1,
    "offset": 0
  }
}
`

func TestGetGifFromGiphyJson(t *testing.T) {
	stringBytes := []byte(giphyJsonResponse)
	expectedResult := "https://media1.giphy.com/media/v1.Y2lkPTUyMWU1ODhiNnc5a2o2aGlqbXBuNG16bTg2enFmbjVzeG5pbGI4M2JibXdtZGtlYiZlcD12MV9naWZzX3NlYXJjaCZjdD1n/erePhJFWkfYMwTpNT8/giphy.gif"
	result, err := getGifFromGiphyJson(stringBytes)
	if err != nil {
		t.Errorf("error getting gif from giphy json: %v", err)
		return
	} else if result != expectedResult {
		t.Errorf("expected gif url to be %v, got %v", expectedResult, result)
	}
}
