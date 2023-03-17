package endpoint

import (
	"emailn/internal/internalerror"
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type EndpointFunc func(writer http.ResponseWriter, request *http.Request) (interface{}, int, error)

func HandleError(endpointFunc EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		object, status, err := endpointFunc(writer, request)

		if err != nil {
			if errors.Is(err, internalerror.InternalServerError) {
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
