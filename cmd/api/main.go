package main

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"emailn/internal/infrastructure/database"
	"emailn/internal/internalerrors"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	service := campaign.Service{Repository: &database.CampaignRepository{}}

	router.Post("/campaigns", func(writer http.ResponseWriter, request *http.Request) {
		var newCampaignDto contract.NewCampaignDto
		render.DecodeJSON(request.Body, &newCampaignDto)

		id, err := service.Create(newCampaignDto)

		if err != nil {
			if errors.Is(err, internalerrors.InternalServerError) {
				render.Status(request, 500)
			} else {
				render.Status(request, 422)
			}

			render.JSON(writer, request, map[string]string{"error": err.Error()})

			return
		}

		render.Status(request, 201)
		render.JSON(writer, request, map[string]string{"id": id})
	})

	http.ListenAndServe(":3000", router)
}
