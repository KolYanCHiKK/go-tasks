package verify

import (
	"net/http"
	"project/go-tasks/configs"
	"project/go-tasks/pkg/responce"
	"strings"
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
		resp := SendResponse{
			Status:  "Success",
			Links:   h.Email,
			Address: h.Address,
		}

		responce.Json(w, resp, 201)
	}
}

func (h *Handler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if strings.Trim(req.PathValue("hash"), " ") == "" {
			resp := VerifyErrorResponse{
				Status:  "Error",
				Message: "Hash token is missing from the request",
			}
			responce.Json(w, resp, 400)
			return
		}

		resp := VerifyResponse{Status: "Success", Links: h.Email}
		responce.Json(w, resp, 201)
	}
}
