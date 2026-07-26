package verify

type SendResponse struct {
	Status  string `json:"status"`
	Links   string `json:"links,omitempty"`
	Address string `json:"address,omitempty"`
}

type VerifyResponse struct {
	Status string `json:"status"`
	Links  string `json:"links"`
}

type VerifyErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
