package v2

import (
	"github.com/s4mn0v/bitget/internal"
	"github.com/s4mn0v/bitget/internal/common"
)

type SpotWalletAPI struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotWalletAPI) Transfer(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/transfer", postBody)
	return resp, err
}

func (p *SpotWalletAPI) DepositAddress(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-address", params)
	return resp, err
}

func (p *SpotWalletAPI) Withdrawal(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/withdrawal", postBody)
	return resp, err
}

func (p *SpotWalletAPI) WithdrawalRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/withdrawal-records", params)
	return resp, err
}

func (p *SpotWalletAPI) DepositRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/deposit-records", params)
	return resp, err
}

func (p *SpotWalletAPI) ModifyDepositAccount(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/modify-deposit-account", postBody)
	return resp, err
}

func (p *SpotWalletAPI) TransferCoinInfo(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/transfer-coin-info", params)
	return resp, err
}

func (p *SpotWalletAPI) SubAccountTransfer(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/subaccount-transfer", postBody)
	return resp, err
}

func (p *SpotWalletAPI) SubMainTransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/sub-main-trans-record", params)
	return resp, err
}

func (p *SpotWalletAPI) SwitchBgbDeduct(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/account/switch-deduct", postBody)
	return resp, err
}

func (p *SpotWalletAPI) SubAccountDepositAddress(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/subaccount-deposit-address", params)
	return resp, err
}

func (p *SpotWalletAPI) CancelWithdrawal(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/wallet/cancel-withdrawal", postBody)
	return resp, err
}

func (p *SpotWalletAPI) SubAccountDepositRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/wallet/subaccount-deposit-records", params)
	return resp, err
}
