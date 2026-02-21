package v2

import (
	"github.com/s4mn0v/bitget/internal"
	"github.com/s4mn0v/bitget/internal/common"
)

type SpotOrderClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *SpotOrderClient) Init() *SpotOrderClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

// PlaceOrder Function
func (p *SpotOrderClient) PlaceOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/place-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) BatchPlaceOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/batch-orders", postBody)
	return resp, err
}

func (p *SpotOrderClient) CancelOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/cancel-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) BatchCancelOrders(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/batch-cancel-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) OrdersHistory(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/trade/history-orders", params)
	return resp, err
}

func (p *SpotOrderClient) OrdersPending(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/trade/unfilled-orders", params)
	return resp, err
}

func (p *SpotOrderClient) Fills(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/trade/fills", params)
	return resp, err
}

// PlacePlanOrder Function
func (p *SpotOrderClient) PlacePlanOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/place-plan-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) CancelPlanOrder(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/cancel-plan-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) OrdersPlanPending(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/trade/current-plan-order", params)
	return resp, err
}

func (p *SpotOrderClient) OrdersPlanHistory(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/trade/history-plan-order", params)
	return resp, err
}

// --- Trade Extentions ---

func (p *SpotOrderClient) CancelReplaceOrder(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/cancel-replace-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) BatchCancelReplaceOrder(params map[string]interface{}) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/batch-cancel-replace-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) CancelSymbolOrder(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/cancel-symbol-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) OrderInfo(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/trade/orderInfo", params)
	return resp, err
}

// --- Plan Oreder Extentions ---

func (p *SpotOrderClient) ModifyPlanOrder(params map[string]string) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/modify-plan-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) PlanSubOrder(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/spot/trade/plan-sub-order", params)
	return resp, err
}

func (p *SpotOrderClient) BatchCancelPlanOrder(params map[string]interface{}) (string, error) {
	postBody, _ := internal.ToJson(params)
	resp, err := p.BitgetRestClient.DoPost("/api/v2/spot/trade/batch-cancel-plan-order", postBody)
	return resp, err
}

func (p *SpotOrderClient) TraderOrderCloseTracking(params map[string]string) (string, error) {
	postBody, jsonErr := internal.ToJson(params)
	if jsonErr != nil {
		return "", jsonErr
	}
	resp, err := p.BitgetRestClient.DoPost("/api/v2/copy/spot-trader/order-close-tracking", postBody)
	return resp, err
}

func (p *SpotOrderClient) TraderOrderCurrentTrack(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/copy/spot-trader/order-current-track", params)
	return resp, err
}

func (p *SpotOrderClient) TraderOrderHistoryTrack(params map[string]string) (string, error) {
	resp, err := p.BitgetRestClient.DoGet("/api/v2/copy/spot-trader/order-history-track", params)
	return resp, err
}
