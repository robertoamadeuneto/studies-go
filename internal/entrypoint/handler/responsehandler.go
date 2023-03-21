package handler

import (
	"emailn/internal"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ControllerFunc func(writer http.ResponseWriter, request *http.Request) (interface{}, int, error)

func HandleResponse(controllerFunc ControllerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		object, status, err := controllerFunc(writer, request)

		if err != nil {
			if errors.Is(err, internal.InternalServerError) {
				render.Status(request, 500)
			} else {
				render.Status(request, 422)
			}

			render.JSON(writer, request, map[string]string{"error": err.Error()})

			return
		}

		render.Status(request, status)

		if object != nil {
			render.JSON(writer, request, object)
		}
	})
}
