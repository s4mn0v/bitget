package v2

import (
	"github.com/s4mn0v/bitget/internal"
	"github.com/s4mn0v/bitget/internal/common"
)

type SpotAccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotAccountClient) Init() *SpotAccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *SpotAccountClient) Info() (string, error) {
	params := internal.NewParams()
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/info", params)
	return resp, err
}

func (p *SpotAccountClient) Assets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/assets", params)
	return resp, err
}

func (p *SpotAccountClient) Bills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/bills", params)
	return resp, err
}

func (p *SpotAccountClient) TransferRecords(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/transferRecords", params)
	return resp, err
}

func (p *SpotAccountClient) SubAccountAssets(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/subaccount-assets", params)
	return resp, err
}

func (p *SpotAccountClient) BgbDeductInfo() (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/deduct-info", internal.NewParams())
	return resp, err
}

func (p *SpotAccountClient) UpgradeStatus(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/account/upgrade-status", params)
	return resp, err
}

func (p *SpotAccountClient) UpgradeAccount(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/account/upgrade", postBody)
	return resp, err
}
