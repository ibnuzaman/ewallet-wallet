package external

import (
	"bytes"
	"context"
	"encoding/json"
	"ewallet-wallet/helpers"
	"net/http"

	"github.com/pkg/errors"
)

type Wallet struct {
	UserID int `json:"user_id"`
}

func (e *Wallet) CreateWallet(ctx context.Context) error {
	payload, err := json.Marshal(e)
	if err != nil {
		return err
	}

	url := helpers.GetEnv("WALLET_HOST", "") + helpers.GetEnv("WALLET_ENDPOINT_CREATE", "")

	httpReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return errors.Wrap(err, "failed to create http request")
	}

	client := &http.Client{}

	resp, err := client.Do(httpReq)
	if err != nil {
		return errors.Wrap(err, "failed to send http request")
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("got error response from wallet service")
	}

	return nil
}
