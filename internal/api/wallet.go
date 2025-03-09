package api

import (
	"ewallet-wallet/helpers"
	"ewallet-wallet/internal/constants"
	"ewallet-wallet/internal/interfaces"
	"ewallet-wallet/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletApi struct {
	WalletService interfaces.IWalletService
}

func (api *WalletApi) Create(c *gin.Context) {
	var (
		log = helpers.Logger
		req models.Wallet
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Info("failed to parse request :", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, nil)
		return
	}

	if req.UserID == 0 {
		log.Info("user_id is empty")
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedBadParseRequest, nil)
		return
	}

	err := api.WalletService.Create(c.Request.Context(), &req)
	if err != nil {
		log.Info("failed to create wallet :", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, req)

}
