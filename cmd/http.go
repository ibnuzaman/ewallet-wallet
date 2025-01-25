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

	userV1 := r.Group("/v1/user")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)

	userV1withAuth := userV1.Use()
	userV1withAuth.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.Logout)
	userV1withAuth.POST("/refresh-token", dependency.MiddlewareValidateAuth, dependency.RefreshTokenAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}

}

type Dependency struct {
	UserRepository interfaces.IUserRepository

	HealthchekAPI   interfaces.IHealthcheckHandler
	RegisterAPI     interfaces.IRegisHandler
	LoginAPI        interfaces.ILoginHandler
	LogoutAPI       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler

	TokenValidation *api.TokenValidationHandler
}

func dependencyInject() Dependency {
	healthcheckSvc := &services.Healthcheck{}
	healthchekcAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	regisSvc := &services.RegisterService{
		RegisterRepo: userRepo,
	}

	regisAPI := &api.RegisterHandler{
		RegisterService: regisSvc,
	}

	loginSvc := &services.LoginService{
		UserRepo: userRepo,
	}

	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	logoutSvc := &services.LogoutService{
		UserRepo: userRepo,
	}

	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	refreshTokenSvc := &services.RefreshTokenService{
		UserRepo: userRepo,
	}

	refreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokenSvc,
	}

	tokenValidationSvc := &services.TokenValidationService{
		UserRepo: userRepo,
	}

	tokenValidationAPI := &api.TokenValidationHandler{
		TokenValidationService: tokenValidationSvc,
	}

	return Dependency{
		UserRepository:  userRepo,
		HealthchekAPI:   healthchekcAPI,
		RegisterAPI:     regisAPI,
		LoginAPI:        loginAPI,
		LogoutAPI:       logoutAPI,
		RefreshTokenAPI: refreshTokenAPI,
		TokenValidation: tokenValidationAPI,
	}

}
