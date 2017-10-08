package crawler

//Links of each link
type Links struct {
	Links []LinkMetadata `json:"links"`
	Count int            `json:"count"`
	Page  string         `json:"page"`
}

//LinkMetadata of each link
type LinkMetadata struct {
	Domain             string `json:"domain"`
	TLD                string `json:"tld"`
	OriginalURL        string `json:"original_url"`
	OriginalURLHash    string `json:"original_url_hash"`
	Redirects          bool   `json:"redirects"`
	DestinationURL     string `json:"destination_url"`
	DestinationURLHash string `json:"destination_url_hash"`
}
