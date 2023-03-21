package main

import (
	"emailn/internal/core/service"
	"emailn/internal/entrypoint/controller"
	"emailn/internal/entrypoint/handler"
	infrarepository "emailn/internal/infra/repository"
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

	campaignService := service.CampaignServiceImpl{CampaignRepository: &infrarepository.CampaignRepositoryImpl{}}
	campaignController := controller.CampaignController{CampaignService: &campaignService}

	router.Post("/campaigns", handler.HandleResponse(campaignController.Create))
	router.Get("/campaigns", handler.HandleResponse(campaignController.FindAll))

	http.ListenAndServe(":3000", router)
}
