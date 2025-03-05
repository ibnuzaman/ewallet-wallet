package repository

import (
	"context"
	"ewallet-wallet/internal/models"

	"gorm.io/gorm"
)

type WalletRepo struct {
	DB *gorm.DB
}

func (r *WalletRepo) CreateWallet(ctx context.Context, wallet *models.Wallet) error {
	return r.DB.Create(wallet).Error
}
