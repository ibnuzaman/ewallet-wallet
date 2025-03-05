package services

import (
	"context"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/models"
)

type WalletService struct {
	WalletRepo interfaces.IWalletRepo
}

func (s *WalletService) Create(ctx context.Context, userID int) error {
	return s.WalletRepo.CreateWallet(ctx, &models.Wallet{UserID: userID})
}
