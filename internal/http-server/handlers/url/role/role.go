package role

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	resp "github.com/lallison21/school-project-server/internal/lib/api/response"
	"github.com/lallison21/school-project-server/internal/lib/logger/sl"
	"log/slog"
	"net/http"
	"strconv"
)

type IRole interface {
	GetRoles() ([]Role, error)
	GetRoleById(id int) (Role, error)
	CreateRole(roleName string, accessLevel int) (Role, error)
}

type Role struct {
	Id         int    `json:"id,omitempty"`
	RoleName   string `json:"role_name,omitempty"`
	AccessLeve int    `json:"access_leve,omitempty"`
}

type Response struct {
	resp.Response
	Role Role `json:"role,omitempty"`
}

type ResponseRoles struct {
	resp.Response
	Roles []Role `json:"roles,omitempty"`
}

type CreateRoleRequest struct {
	RoleName    string `json:"role_name" validate:"required"`
	AccessLevel int    `json:"access_level" validate:"required"`
}

func CreateRole(log *slog.Logger, role IRole) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.url.role.CreateRole"

		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req CreateRoleRequest

		if err := render.DecodeJSON(r.Body, &req); err != nil {
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

		res, err := role.CreateRole(req.RoleName, req.AccessLevel)
		if err != nil {
			log.Info("failed to create new role", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to create new role"))

			return
		}

		log.Info("new role created", slog.Any("created role", res))

		render.JSON(w, r, Response{
			Response: resp.OK(),
			Role:     res,
		})
	}
}

func GetRoleById(log *slog.Logger, role IRole) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.url.role.GetRoleById"

		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			log.Error("failed to decode request params", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to decode request params"))

			return
		}

		log.Info("request params decoded", slog.Int("request id", id))

		res, err := role.GetRoleById(id)
		if err != nil {
			log.Info("failed to get role by id", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get role by id"))

			return
		}

		log.Info("role by id founded", slog.Any("founded role", res))
		render.JSON(w, r, Response{
			Response: resp.OK(),
			Role:     res,
		})

	}
}

func GetRoles(log *slog.Logger, role IRole) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "handlers.url.role.GetRoles"

		log = log.With(
			slog.String("fn", fn),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		res, err := role.GetRoles()
		if err != nil {
			log.Info("failed to get roles", sl.Err(err))

			render.JSON(w, r, resp.Error("failed to get roles"))

			return
		}

		log.Info("all roles was founded", slog.Any("founded roles", res))
		render.JSON(w, r, ResponseRoles{
			Response: resp.OK(),
			Roles:    res,
		})
	}
}
