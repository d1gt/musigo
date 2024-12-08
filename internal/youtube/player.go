package youtube

import (
	"bytes"
	"cmp"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"slices"
)

const playerEndpoint string = "https://music.youtube.com/youtubei/v1/player?key=AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w&prettyPrint=true"

// const playerEndpoint string = "https://music.youtube.com/youtubei/v1/player?prettyPrint=true"

type playerData struct {
	audioAdaptiveFormats []AdaptiveFormats
	videoAdaptiveFormats []AdaptiveFormats
}

func (y Client) GetPlayer(ctx context.Context, videoId string) (playerData, error) {
	payload := y.getPlayerPayload(videoId)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, playerEndpoint, payload)
	if err != nil {
		return playerData{}, err
	}

	y.setHeaders(req)

	resp, err := y.HttpClient.Do(req)
	if err != nil {
		return playerData{}, err
	}

	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return playerData{}, err
		}
	default:
		reader = resp.Body
	}

	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return playerData{}, err
	}

	var response PlayerResponse

	err = json.Unmarshal(body, &response)
	if err != nil {
		return playerData{}, err
	}

	if response.Error.Code != 0 {
		return playerData{}, fmt.Errorf("%v", response.Error)
	}

	// fmt.Println(string(body))
	playerData := y.parsePlayerResponse(response)
	return playerData, nil
}

func (c Client) parsePlayerResponse(res PlayerResponse) playerData {
	audioFormats := res.StreamingData.AdaptiveFormats
	videoFormats := res.StreamingData.AdaptiveFormats

	audioMimeType, _ := regexp.Compile(`^audio/(webm|mp4);`)

	filteredAudioStreamFormats := slices.DeleteFunc(audioFormats, func(e AdaptiveFormats) bool {
		return !audioMimeType.MatchString(e.MimeType)
	})

	slices.SortStableFunc(filteredAudioStreamFormats, func(a, b AdaptiveFormats) int {
		return cmp.Compare(b.Bitrate, a.Bitrate)
	})

	filteredVideoStreamFormats := slices.DeleteFunc(videoFormats, func(e AdaptiveFormats) bool {
		return !audioMimeType.MatchString(e.MimeType)
	})

	slices.SortStableFunc(videoFormats, func(a, b AdaptiveFormats) int {
		return cmp.Compare(b.Bitrate, a.Bitrate)
	})

	return playerData{
		audioAdaptiveFormats: filteredAudioStreamFormats,
		videoAdaptiveFormats: filteredVideoStreamFormats,
	}
}

func (y Client) getPlayerPayload(videoId string) io.Reader {
	payload := []byte(`{"context":{"client":{"hl":"en","deviceMake":"","deviceModel":"","visitorData":"","userAgent":"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36,gzip(gfe)","clientName":"WEB","clientVersion":"2.20241113.07.00","osName":"X11","osVersion":"","originalUrl":"https://www.youtube.com/watch?v=` + videoId + `","platform":"DESKTOP","clientFormFactor":"UNKNOWN_FORM_FACTOR","configInfo":{"appInstallData":""},"userInterfaceTheme":"USER_INTERFACE_THEME_LIGHT","browserName":"Chrome","browserVersion":"131.0.0.0","acceptHeader":"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8","deviceExperimentId":"","screenWidthPoints":1099,"screenHeightPoints":989,"screenPixelDensity":1,"screenDensityFloat":1,"utcOffsetMinutes":120,"memoryTotalKbytes":"1000000","clientScreen":"WATCH","mainAppWebInfo":{"graftUrl":"/watch?v=` + videoId + `","pwaInstallabilityStatus":"PWA_INSTALLABILITY_STATUS_CAN_BE_INSTALLED","webDisplayMode":"WEB_DISPLAY_MODE_BROWSER","isWebNativeShareAvailable":false}},"user":{"lockedSafetyMode":false},"request":{"useSsl":true,"internalExperimentFlags":[],"consistencyTokenJars":[]},"clickTracking":{"clickTrackingParams":"CMUBEKQwGAAiEwidr43EqOGJAxV7RnoFHRSVOcUyB3JlbGF0ZWRIm-OXqvy46qUbmgEFCAEQ-B0="},"adSignalsInfo":{"params":[{"key":"dt","value":"1731776055268"},{"key":"flash","value":"0"},{"key":"frm","value":"0"},{"key":"u_tz","value":"120"},{"key":"u_his","value":"8"},{"key":"u_h","value":"994"},{"key":"u_w","value":"1106"},{"key":"u_ah","value":"994"},{"key":"u_aw","value":"1106"},{"key":"u_cd","value":"24"},{"key":"bc","value":"31"},{"key":"bih","value":"989"},{"key":"biw","value":"1084"},{"key":"brdim","value":"1,7,1,7,1106,7,1106,994,1099,989"},{"key":"vis","value":"1"},{"key":"wgl","value":"true"},{"key":"ca_type","value":"image"}]}},"videoId":"` + videoId + `","playbackContext":{"contentPlaybackContext":{"currentUrl":"/watch?v=` + videoId + `","vis":0,"splay":false,"autoCaptionsDefaultOn":false,"autonavState":"STATE_OFF","html5Preference":"HTML5_PREF_WANTS","signatureTimestamp":20039,"referer":"https://www.youtube.com/watch?v=G0upx8VF8Zs","lactMilliseconds":"-1","watchAmbientModeContext":{"watchAmbientModeEnabled":true}}},"racyCheckOk":false,"contentCheckOk":false}`)
	return bytes.NewBuffer([]byte(payload))
}
