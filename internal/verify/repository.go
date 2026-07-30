package verify

type UserParams struct {
	Email  string `json:"email"`
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Status string `json:"status"`
}
