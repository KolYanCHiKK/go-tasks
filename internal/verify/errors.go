package verify

var (
	SendErrMessages = map[string]string{
		"Host.required": "В запросе отсутствует email",
		"Host.email":    "Неверно указан email",
		"Host.domen":    "Отправка не может происходить на адрес с данным доменом",
	}
	RequestError = "Invalid request body"
	SendingError = "Invalid sending email"
)
