package interfaces

import (
	"context"
	"ewallet-wallet/cmd/proto/tokenvalidation"
	"ewallet-wallet/helpers"
)

type ITokenValidationHandler interface {
	TokenValidationHandler(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error)
}

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
