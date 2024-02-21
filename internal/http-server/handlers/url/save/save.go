package save

import (
	"errors"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	resp "github.com/lallison21/to-do/internal/lib/api/response"
	"github.com/lallison21/to-do/internal/lib/logger/sl"
	"log/slog"
	"net/http"
)

type Request struct {
	RoleName    string `json:"role_name" validate:"required"`
	AccessLevel int    `json:"access_level" validate:"required"`
}

type Response struct {
	resp.Response
	RoleName string `json:"role_name,omitempty"`
}

type RoleSaver interface {
	CreateRole(roleName string, accessLevel int) (int64, error)
}

func New(log *slog.Logger, roleSaver RoleSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.url.save.New"

		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			log.Error("failed to decode request body", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request"))

			return
		}

		log.Info("request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			var validateErr validator.ValidationErrors
			errors.As(err, &validateErr)

			log.Error("invalid request", sl.Err(err))

			render.JSON(w, r, resp.ValidationError(validateErr))

			return
		}

		id, err := roleSaver.CreateRole(req.RoleName, req.AccessLevel)
		if err != nil {
			log.Info("failed to create new role", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to create new role"))

			return
		}

		log.Info("new role created", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Response: resp.OK(),
			RoleName: req.RoleName,
		})
	}
}
