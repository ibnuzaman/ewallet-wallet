package interfaces

import (
	"context"
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(*gin.Context)
}
