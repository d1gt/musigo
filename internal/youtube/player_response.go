package youtube

type PlayerResponse struct {
	Error         ResponseError `json:"error"`
	StreamingData StreamingData `json:"streamingData"`
}

type StreamingData struct {
	Formats         Formats           `json:"formats"`
	AdaptiveFormats []AdaptiveFormats `json:"adaptiveFormats"`
}

type AdaptiveFormats struct {
	Itag            int    `json:"itag"`
	SignatureCipher string `json:"signatureCipher"`
	URL             string `json:"url"`
	MimeType        string `json:"mimeType"`
	Bitrate         int    `json:"bitrate"`
	Width           int    `json:"width,omitempty"`
	Height          int    `json:"height,omitempty"`
	InitRange       struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"initRange"`
	IndexRange struct {
		Start string `json:"start"`
		End   string `json:"end"`
	} `json:"indexRange"`
	LastModified     string `json:"lastModified"`
	ContentLength    string `json:"contentLength"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps,omitempty"`
	QualityLabel     string `json:"qualityLabel,omitempty"`
	ProjectionType   string `json:"projectionType"`
	AverageBitrate   int    `json:"averageBitrate"`
	ApproxDurationMs string `json:"approxDurationMs"`
	ColorInfo        struct {
		Primaries               string `json:"primaries"`
		TransferCharacteristics string `json:"transferCharacteristics"`
		MatrixCoefficients      string `json:"matrixCoefficients"`
	} `json:"colorInfo,omitempty"`
	HighReplication bool    `json:"highReplication,omitempty"`
	AudioQuality    string  `json:"audioQuality,omitempty"`
	AudioSampleRate string  `json:"audioSampleRate,omitempty"`
	AudioChannels   int     `json:"audioChannels,omitempty"`
	LoudnessDb      float64 `json:"loudnessDb,omitempty"`
	Xtags           string  `json:"xtags,omitempty"`
	IsDrc           bool    `json:"isDrc,omitempty"`
}

type Formats []struct {
	Itag             int    `json:"itag"`
	SignatureCipher  string `json:"signatureCipher"`
	URL              string `json:"url"`
	MimeType         string `json:"mimeType"`
	Bitrate          int    `json:"bitrate"`
	Width            int    `json:"width"`
	Height           int    `json:"height"`
	LastModified     string `json:"lastModified"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps"`
	QualityLabel     string `json:"qualityLabel"`
	ProjectionType   string `json:"projectionType"`
	AudioQuality     string `json:"audioQuality"`
	ApproxDurationMs string `json:"approxDurationMs"`
	AudioSampleRate  string `json:"audioSampleRate"`
	AudioChannels    int    `json:"audioChannels"`
}
