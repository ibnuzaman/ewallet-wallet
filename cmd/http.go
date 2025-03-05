package cmd

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/api"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/repository"
	"ewallet-wallet/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHttp() {

	dependency := dependencyInject()

	r := gin.Default()

	r.GET("/health", dependency.HealthchekAPI.HealcheckHandlerHTTP)
	walletv1 := r.Group("/wallet/v1")
	{
		walletv1.POST("/", dependency.WalletAPI.Create)
	}

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}

}

type Dependency struct {
	HealthchekAPI interfaces.IHealthchecAPI
	WalletAPI     interfaces.IWalletAPI
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthchekcAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	WalletRepo := &repository.WalletRepo{
		DB: helpers.DB,
	}

	WalletSvc := &services.WalletService{
		WalletRepo: WalletRepo,
	}

	WalletApi := &api.WalletApi{
		WalletService: WalletSvc,
	}

	return Dependency{
		HealthchekAPI: healthchekcAPI,
		WalletAPI:     WalletApi,
	}

}
