package configuration

import (
	"emailn/internal/core/service"
	"emailn/internal/entrypoint/controller"
	"emailn/internal/entrypoint/handler"
	infrarepository "emailn/internal/infra/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func BuildRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	buildCampaignRoutes(router)

	return router
}

func buildCampaignRoutes(router *chi.Mux) {
	campaignService := service.CampaignServiceImpl{CampaignRepository: &infrarepository.CampaignRepositoryImpl{}}
	campaignController := controller.CampaignController{CampaignService: &campaignService}

	router.Post("/campaigns", handler.HandleResponse(campaignController.Create))
	router.Get("/campaigns", handler.HandleResponse(campaignController.FindAll))
}
