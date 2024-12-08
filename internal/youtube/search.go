package youtube

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type (
	mediaType   string
	searchParam string
)

const (
	songType              mediaType = "SONG"
	videoType             mediaType = "VIDEO"
	albumType             mediaType = "ALBUM"
	artistType            mediaType = "ARTIST"
	communityPlaylistType mediaType = "COMMUNITY_PLAYLIST"
	featuredPlaylistType  mediaType = "FEATURED_PLAYLIST"
	podcastType           mediaType = "PODCAST"

	songSearchParam              searchParam = "EgWKAQIIAWoKEAkQBRAKEAMQBA%3D%3D"
	videoSearchParam             searchParam = "EgWKAQIQAWoKEAkQChAFEAMQBA%3D%3D"
	albumSearchParam             searchParam = "EgWKAQIYAWoKEAkQChAFEAMQBA%3D%3D"
	artistSearchParam            searchParam = "EgWKAQIgAWoKEAkQChAFEAMQBA%3D%3D"
	communityPlaylistSearchParam searchParam = "EgeKAQQoAEABagoQAxAEEAoQCRAF"
	featuredPlaylistSearchParam  searchParam = "EgeKAQQoADgBagwQDhAKEAMQBRAJEAQ%3D"
	podcastSearchParam           searchParam = "EgWKAQJQAWoIEBAQERADEBU%3D"
)

var searchParams = map[mediaType]searchParam{
	songType:              songSearchParam,
	videoType:             videoSearchParam,
	albumType:             albumSearchParam,
	artistType:            artistSearchParam,
	communityPlaylistType: communityPlaylistSearchParam,
	featuredPlaylistType:  featuredPlaylistSearchParam,
	podcastType:           podcastSearchParam,
}

type Artist struct {
	Name string
	Id   string
}

type Album struct {
	Name string
	Id   string
}

type Song struct {
	Name     string
	Id       string
	Album    Album
	Artists  []Artist
	Plays    string
	Duration string
}

const searchEndpoint string = "https://music.youtube.com/youtubei/v1/search?prettyPrint=false"

func (y Client) Search(ctx context.Context, input string) ([]Song, error) {
	payload := y.getSearchPayload(input, songSearchParam)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, searchEndpoint, payload)
	if err != nil {
		return []Song{}, err
	}

	y.setHeaders(req)

	resp, err := y.HttpClient.Do(req)
	if err != nil {
		return []Song{}, err
	}

	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return []Song{}, err
		}
	default:
		reader = resp.Body
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return []Song{}, err
	}

	var response searchResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Song{}, err
	}

	if response.Error.Code != 0 {
		return []Song{}, fmt.Errorf("err: %v with code %v", response.Error.Message, response.Error.Code)
	}

	songs := y.parseSearchResponse(response)

	return songs, nil
}

func (y Client) parseSearchResponse(s searchResponse) []Song {
	contents := s.Contents
	songs := []Song{}
	for _, t := range contents.TabbedSearchResultsRenderer.Tabs {
		for _, c := range t.TabRenderer.Content.SectionListRenderer.Contents {
			for _, mc := range c.MusicShelfRenderer.Contents {
				song := Song{}
				for i, fc := range mc.MusicResponsiveListItemRenderer.FlexColumns {
					for _, rn := range fc.MusicResponsiveListItemFlexColumnRenderer.Text.Runs {
						if i == 0 {
							song.Name = rn.Text
							song.Id = rn.NavigationEndpoint.WatchEndpoint.VideoID
							continue
						}

						if rn.NavigationEndpoint.BrowseEndpoint.BrowseEndpointContextSupportedConfigs.BrowseEndpointContextMusicConfig.PageType == "MUSIC_PAGE_TYPE_ARTIST" {
							artist := Artist{}
							artist.Name = rn.Text
							artist.Id = rn.NavigationEndpoint.BrowseEndpoint.BrowseID
							song.Artists = append(song.Artists, artist)
						}

						if rn.NavigationEndpoint.BrowseEndpoint.BrowseEndpointContextSupportedConfigs.BrowseEndpointContextMusicConfig.PageType == "MUSIC_PAGE_TYPE_ALBUM" {
							album := Album{}
							album.Name = rn.Text
							album.Id = rn.NavigationEndpoint.BrowseEndpoint.BrowseID
							song.Album = album
						}

						if i == len(mc.MusicResponsiveListItemRenderer.FlexColumns)-2 {
							song.Plays = rn.Text
						}

						if i == len(mc.MusicResponsiveListItemRenderer.FlexColumns)-1 {
							song.Duration = rn.Text
						}

					}
				}
				songs = append(songs, song)

			}
		}
	}

	return songs
}

func (y Client) getSearchPayload(input string, param searchParam) io.Reader {
	payload := fmt.Sprintf(`
{
  "context": {
    "client": {
      "hl": "en",
      "userAgent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
      "clientName": "WEB_REMIX",
      "clientVersion": "1.20240709.02.00",
      "originalUrl": "https://music.youtube.com/",
      "musicAppInfo": {
        "pwaInstallabilityStatus": "PWA_INSTALLABILITY_STATUS_UNKNOWN",
        "webDisplayMode": "WEB_DISPLAY_MODE_BROWSER",
        "storeDigitalGoodsApiSupportStatus": {
          "playStoreDigitalGoodsApiSupportStatus": "DIGITAL_GOODS_API_SUPPORT_STATUS_UNSUPPORTED"
        }
      }
    },
    "user": {"lockedSafetyMode": false},
    "request": {
      "useSsl": true,
      "internalExperimentFlags": [],
      "consistencyTokenJars": []
    }
  },
  "query": "%v",
  "params": "%v"
}
`, input, param)
	return bytes.NewBuffer([]byte(payload))
}

const searchSuggestionsEndpoint string = "https://music.youtube.com/youtubei/v1/music/get_search_suggestions?prettyPrint=true"

func (y Client) GetSearchSuggestions(ctx context.Context, input string) ([]string, error) {
	payload := y.getSearchSuggestionPayload(input)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, searchSuggestionsEndpoint, payload)
	if err != nil {
		return []string{}, err
	}

	y.setHeaders(req)

	resp, err := y.HttpClient.Do(req)
	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return []string{}, err
		}
	default:
		reader = resp.Body
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return []string{}, err
	}

	var response SearchSuggestionResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return []string{}, err
	}

	if response.Error.Code != 0 {
		return []string{}, fmt.Errorf("err: %v with code %v", response.Error.Message, response.Error.Code)
	}

	suggestions := y.parseSearchSuggestionsResults(response)

	return suggestions, nil
}

func (y Client) parseSearchSuggestionsResults(res SearchSuggestionResponse) []string {
	parsedSuggestions := []string{}
	for _, content := range res.Contents {
		for _, suggestionContent := range content.SearchSuggestionsSectionRenderer.Contents {
			suggestion := suggestionContent.SearchSuggestionRenderer.Suggestion
			concatenatedText := ""

			for _, run := range suggestion.Runs {
				concatenatedText += run.Text
			}

			parsedSuggestions = append(parsedSuggestions, concatenatedText)

		}
	}

	return parsedSuggestions
}

func (y Client) getSearchSuggestionPayload(input string) io.Reader {
	return bytes.NewBuffer([]byte(`{ "input":"` + input + `",
		"context":{ "client":{
				"hl":"en",
				"userAgent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
				"clientName":"WEB_REMIX",
				"clientVersion":"1.20240709.02.00",
				"originalUrl":"https://music.youtube.com/",
				"musicAppInfo":{
					"pwaInstallabilityStatus":"PWA_INSTALLABILITY_STATUS_UNKNOWN",
					"webDisplayMode":"WEB_DISPLAY_MODE_BROWSER",
					"storeDigitalGoodsApiSupportStatus":{
						"playStoreDigitalGoodsApiSupportStatus":"DIGITAL_GOODS_API_SUPPORT_STATUS_UNSUPPORTED"
					}
				}
			},
			"user":{
				"lockedSafetyMode":false
			},
			"request":{
				"useSsl":true,
				"internalExperimentFlags":[],
				"consistencyTokenJars":[]
			}
		}
	}`))
}
