package ntnswap

import (
	"fmt"
	"math/big"

	"github.com/aviate-labs/agent-go/candid/idl"
	"github.com/niccolofant/ic-arb/core/icp"
)

func (n *ntnswap) Withdraw(token icp.Token, amount *big.Int) (*big.Int, error) {
	botPrincipal := n.api.Agent().Sender()
	tokenPrincipalRaw := token.CanisterID().Raw()

	commandResult, err := n.api.Icrc55Command(BatchCommandRequest{
		Controller: Controller{
			Owner: botPrincipal.Raw(),
		},
		Commands: []Command{
			{
				Transfer: &TransferRequest{
					To: struct {
						Account         *Account "ic:\"account,variant\" json:\"account,omitempty\""
						ExternalAccount *struct {
							Ic    *Account "ic:\"ic,variant\" json:\"ic,omitempty\""
							Icp   *[]byte  "ic:\"icp,variant\" json:\"icp,omitempty\""
							Other *[]byte  "ic:\"other,variant\" json:\"other,omitempty\""
						} "ic:\"external_account,variant\" json:\"external_account,omitempty\""
						Node *struct {
							EndpointIdx EndpointIdx "ic:\"endpoint_idx\" json:\"endpoint_idx\""
							NodeId      LocalNodeId "ic:\"node_id\" json:\"node_id\""
						} "ic:\"node,variant\" json:\"node,omitempty\""
						NodeBilling *LocalNodeId "ic:\"node_billing,variant\" json:\"node_billing,omitempty\""
						Temp        *struct {
							Id        uint32      "ic:\"id\" json:\"id\""
							SourceIdx EndpointIdx "ic:\"source_idx\" json:\"source_idx\""
						} "ic:\"temp,variant\" json:\"temp,omitempty\""
					}{
						ExternalAccount: &struct {
							Ic    *Account "ic:\"ic,variant\" json:\"ic,omitempty\""
							Icp   *[]byte  "ic:\"icp,variant\" json:\"icp,omitempty\""
							Other *[]byte  "ic:\"other,variant\" json:\"other,omitempty\""
						}{
							Ic: &Account{
								Owner: botPrincipal.Raw(),
							},
						},
					},
					From: struct {
						Account *Account "ic:\"account,variant\" json:\"account,omitempty\""
						Node    *struct {
							EndpointIdx EndpointIdx "ic:\"endpoint_idx\" json:\"endpoint_idx\""
							NodeId      LocalNodeId "ic:\"node_id\" json:\"node_id\""
						} "ic:\"node,variant\" json:\"node,omitempty\""
					}{
						Account: &Account{
							Owner: botPrincipal.Raw(),
						},
					},
					Ledger: SupportedLedger{
						Ic: &tokenPrincipalRaw,
					},
					Amount: idl.NewBigNat(amount),
				},
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw %s %s tokens: %w", amount, token, err)
	}

	if commandResult.Err != nil {
		return nil, fmt.Errorf("failed to withdraw %s %s tokens: %w", amount, token, commandResult.Err.Decode())
	}

	withdrawResult := commandResult.Ok.Commands[0].Transfer
	if withdrawResult.Err != nil {
		return nil, fmt.Errorf("failed to withdraw %s %s tokens: %s", amount, token, *withdrawResult.Err)
	}

	txToFind := Tx{
		Token:  token,
		To:     n.api.Agent().Sender(),
		Amount: amount,
	}

	tx, err := n.waitForTxCompletion(txToFind)
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw: %w", err)
	}

	withdrawnAmount := new(big.Int).Sub(tx.Amount, token.Metadata().Fee)

	return withdrawnAmount, nil
}
