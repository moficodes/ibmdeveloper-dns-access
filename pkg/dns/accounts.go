package dns

import (
	"context"
	"errors"
	"fmt"

	"github.com/dnsimple/dnsimple-go/dnsimple"
)

// GetAccountID return 1st accountid
func GetAccountID(client *dnsimple.Client) (*string, error) {
	accounts, err := client.Accounts.ListAccounts(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	if len(accounts.Data) <= 0 {
		return nil, errors.New("No account associated")
	}

	accountID := fmt.Sprintf("%d", accounts.Data[0].ID)
	return &accountID, nil
}
