package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoint"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	campaignService := campaign.ServiceImpl{Repository: &database.CampaignRepository{}}

	handler := endpoint.Handler{CampaignService: &campaignService}

	router.Post("/campaigns", endpoint.HandleError(handler.PostCampaign))
	router.Get("/campaigns", endpoint.HandleError(handler.GetCampaigns))

	http.ListenAndServe(":3000", router)
}
