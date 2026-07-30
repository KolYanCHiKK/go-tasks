package verify

import (
	"encoding/json"
	"errors"
	"net/http"
	"project/go-tasks/configs"
	"project/go-tasks/pkg/response"
	"project/go-tasks/pkg/validations"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Handler struct {
	*configs.Config
}

type HandlerDeps struct {
	*configs.Config
}

func NewHandler(router *http.ServeMux, dep *HandlerDeps) {
	h := &Handler{Config: dep.Config}
	router.HandleFunc("POST /send", h.Send())
	router.HandleFunc("GET /verify/{hash}", h.Verify())
}

func (h *Handler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var reqBody SendRequest
		err := json.NewDecoder(req.Body).Decode(&reqBody)
		if err != nil {
			resp := SendErrorResponse{
				Error:   RequestError,
				Message: err.Error(),
			}
			response.Json(w, resp, 400)
			return
		}

		validate := validator.New()
		err = validate.RegisterValidation("domen", validations.GmailEmailValidate)
		if err != nil {
			res := SendErrorResponse{
				Error:   RequestError,
				Message: err.Error(),
			}
			response.Json(w, res, 400)
			return
		}

		err = validate.Struct(reqBody)
		var errorsSlice validator.ValidationErrors
		checkErr := errors.As(err, &errorsSlice)
		if checkErr {
			for _, value := range errorsSlice {
				key := value.Field() + "." + value.Tag()
				message := SendErrMessages[key]

				resp := SendErrorResponse{
					Error:   RequestError,
					Message: message,
				}
				response.Json(w, resp, 400)
				return
			}
		}
		link, err := SendEmail(h.Host, h.ReserverToken, h.File, reqBody.Email)
		if err != nil {
			resp := SendErrorResponse{
				Error:   SendingError,
				Message: err.Error(),
			}
			response.Json(w, resp, 400)
			return
		}

		resp := SendResponse{
			Status:  "Success",
			Links:   link,
			Address: reqBody.Email,
		}

		response.Json(w, resp, 201)
	}
}

func (h *Handler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		hash := strings.Trim(req.PathValue("hash"), " ")
		if hash == "" {
			resp := VerifyErrorResponse{
				Status:  "Error",
				Message: "Hash token is missing from the request",
			}
			response.Json(w, resp, 400)
			return
		}

		verify, err := VerifyLink(hash, h.File)
		if verify && err != nil && err.Error() == "HASH IS VERIFIED" {
			resp := VerifyResponse{
				Status:  "Success",
				Hash:    hash,
				Message: "Hash has already been verified",
			}
			response.Json(w, resp, 200)
			return
		}
		if err != nil {
			resp := VerifyErrorResponse{
				Status:  "Error",
				Message: err.Error(),
			}
			response.Json(w, resp, 500)
			return
		}
		if !verify {
			resp := VerifyErrorResponse{
				Status:  "Error",
				Message: "Hash was not found verify",
			}
			response.Json(w, resp, 404)
			return
		}

		resp := VerifyResponse{Status: "Success", Hash: req.Host + req.URL.String()}
		response.Json(w, resp, 200)
	}
}
