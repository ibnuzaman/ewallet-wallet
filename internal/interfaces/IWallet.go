package interfaces

import (
	"context"
	"ewallet-wallet/internal/models"

	"github.com/gin-gonic/gin"
)

type IWalletRepo interface {
	CreateWallet(ctx context.Context, wallet *models.Wallet) error
}

type IWalletService interface {
	Create(ctx context.Context, wallet *models.Wallet) error
}
type IWalletAPI interface {
	Create(*gin.Context)
}
