package chat

// Meta stores information about a current view of a request.
type Meta struct {
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
	FirstPageURL    string `json:"first_page_url"`
	PreviousPageURL string `json:"previous_page_url"`
	URL             string `json:"url"`
	NextPageURL     string `json:"next_page_url"`
	Key             string `json:"key"`
}
