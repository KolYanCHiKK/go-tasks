package verify

type SendResponse struct {
	Status  string `json:"status"`
	Links   string `json:"links,omitempty"`
	Address string `json:"address,omitempty"`
}

type SendErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
type SendRequest struct {
	Email string `json:"email" validate:"required,email,domen=gmail"`
}

type EmailAddress struct {
	Email string `json:"email"`
	Name  string `json:"name,omitempty"`
}
type SendingRequestBody struct {
	From     EmailAddress   `json:"from"`
	To       []EmailAddress `json:"to"`
	Subject  string         `json:"subject"`
	Text     string         `json:"text"`
	Category string         `json:"category,omitempty"`
}

type SendingErrorResponseBody struct {
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}
type VerifyResponse struct {
	Status  string `json:"status"`
	Hash    string `json:"hash"`
	Message string `json:"message,omitempty"`
}

type VerifyErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
