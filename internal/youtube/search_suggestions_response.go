package youtube

type SearchSuggestionResponse struct {
	Error    ResponseError               `json:"error"`
	Contents []SearchSuggestionsContents `json:"contents"`
}

type SearchSuggestionsContents struct {
	SearchSuggestionsSectionRenderer SearchSuggestionsSectionRenderer `json:"searchSuggestionsSectionRenderer"`
}

type SearchSuggestionsSectionRenderer struct {
	Contents []SuggestionContent `json:"contents"`
}

type SuggestionContent struct {
	SearchSuggestionRenderer SearchSuggestionRenderer `json:"searchSuggestionRenderer"`
}

type SearchSuggestionRenderer struct {
	Suggestion Suggestion `json:"suggestion"`
}

type Suggestion struct {
	Runs []SearchSuggestionsRun `json:"runs"`
}

type SearchSuggestionsRun struct {
	Text string `json:"text"`
	Bold bool   `json:"bold,omitempty"`
}
