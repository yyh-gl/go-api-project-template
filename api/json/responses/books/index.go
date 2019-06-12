package responses

type IndexResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}
