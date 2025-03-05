package cmd

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHttp() {

	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/health", dependency.HealthchekAPI.HealcheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}

}

type Dependency struct {
	HealthchekAPI interfaces.IHealthcheckHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthchekcAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	return Dependency{

		HealthchekAPI: healthchekcAPI,
	}

}
