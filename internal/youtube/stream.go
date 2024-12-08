package youtube

import (
	"cmp"
	"context"
	"io"
	"regexp"
	"slices"

	yt2 "github.com/kkdai/youtube/v2"
)

func (y Client) Stream(ctx context.Context, id string) (io.ReadCloser, int64, error) {
	videoURL := "https://www.youtube.com/watch?v=" + id

	vid, _ := y.streamClient.GetVideoContext(ctx, videoURL)

	formats := vid.Formats

	// audioMimeType, _ := regexp.Compile(`^audio/(webm|mp4);`)

	audioMimeType, _ := regexp.Compile(`^audio/webm;`)
	filteredAudioStreamFormats := slices.DeleteFunc(formats, func(e yt2.Format) bool {
		return !audioMimeType.MatchString(e.MimeType)
	})

	slices.SortStableFunc(filteredAudioStreamFormats, func(a, b yt2.Format) int {
		return cmp.Compare(b.Bitrate, a.Bitrate)
	})

	return y.streamClient.GetStreamContext(ctx, vid, &filteredAudioStreamFormats[0])
}
