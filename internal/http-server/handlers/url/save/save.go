package save

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	resp "github.com/malinatrash/url-shortener/internal/lib/api/response"
	"github.com/malinatrash/url-shortener/internal/lib/logger/sl"
	"log/slog"
	"net/http"
)

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	resp.Response
	Alias string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave string, alias string) (int64, error)
}

func New(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handler.url.save.New"
		log := log.With(
			slog.String("op", op),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request
		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request", sl.Err(err))
			render.JSON(w, r, resp.Error("failed to decode request"))
			return
		}

		log.Info("request body decoded", slog.Any("body", req))

		// TODO: implement this
		//if err := validator.Validate(req); err != nil {
		//	log.
		//}
	}
}
