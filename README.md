# Bitget Golang SDK (Modified Version)

Bitget Golang SDK, fixed some errors from the original SDK and migrated to v2 of Bitget API.

For the original version, please visit the link below:
[Original version](https://github.com/BitgetLimited/v3-bitget-api-sdk/blob/master/bitget-golang-sdk-api/README_EN.md)

---

# Modifications Made

- **Enhanced Trading Operations (v2)**
  - Added support for **Modify Order** (Price/Size) for active limit orders.
  - Added **Flash Close Position** to exit positions at market price immediately.
  - Added **Reversal (Click Backhand)** functionality to flip positions.
  - Added specific endpoints for **TP/SL Plan Orders** (Place and Modify).
  > Ref: [Modify Order](https://www.bitget.com/api-doc/contract/trade/Modify-Order), [Flash Close](https://www.bitget.com/api-doc/contract/trade/Flash-Close-Position)
  > Modified Code: `pkg/client/v2/mixorderclient.go`

- **Advanced Market Data (v2)**
  - Added **Current & History Funding Rate** queries.
  - Added **Open Interest** and **Mark/Index/Market Prices** real-time data.
  - Added **Position Tier** (Leverage brackets) query.
  - Added **Historical Candles** for Index and Mark prices.
  > Ref: [Market Tickers](https://www.bitget.com/api-doc/contract/market/Get-All-Symbol-Ticker), [Funding Rate](https://www.bitget.com/api-doc/contract/market/Get-Current-Funding-Rate)
  > Modified Code: `pkg/client/v2/mixmarketclient.go`

- **Account & Sub-Account Management (v2)**
  - Added **Estimated Open Count** to calculate max affordable size before ordering.
  - Added **Account Bills** with support for the last 90 days.
  - Added **Sub-account Assets** query for master accounts.
  > Ref: [Est. Open Count](https://www.bitget.com/api-doc/contract/account/Est-Open-Count), [Sub-Account Assets](https://www.bitget.com/api-doc/contract/account/Get-Sub-Account-Contract-Assets)
  > Modified Code: `pkg/client/v2/mixaccountclient.go`

- **FUTURES History & Fills**
  - Added new model for listing FILLED historical transaction data.
  - Added func for get HISTORY POSITION of all symbols.
  > Ref: [Get Fill History](https://www.bitget.com/api-doc/contract/trade/Get-Fill-History), [Get History Position](https://www.bitget.com/api-doc/contract/position/Get-History-Position)
  > Modified Code: `internal/model/history.go`, `pkg/client/v2/mixorderclient.go`, `pkg/client/v2/mixaccountclient.go`

---

# Usage examples

### 1. Get Futures Fill History

```go
func (a *App) GetFuturesFillHistory(symbol string, productType string) (interface{}, error) {
 params := map[string]string{
  "symbol":      symbol,
  "productType": productType,
  "limit":       "50",
 }

 resp, err := a.mixOrderClient.FillHistory(params)
 if err != nil {
  return nil, err
 }

 var result map[string]interface{}
 if err := json.Unmarshal([]byte(resp), &result); err != nil {
  return nil, err
 }

 if data, ok := result["data"].(map[string]interface{}); ok {
  if fillList, ok := data["fillList"]; ok {
   return fillList, nil
  }
@ }

 return []interface{}{}, nil
}

```

---

### 2. Flash Close a Position

```go
func (a *App) ClosePositionFast(symbol string, holdSide string) (string, error) {
 params := map[string]string{
  "symbol":      symbol,
  "productType": "USDT-FUTURES",
  "holdSide":    holdSide, // "long" or "short"
 }

 return a.mixOrderClient.FlashClosePositions(params)
}
```

---

### 3. Get Estimated Openable Quantity

```go
func (a *App) GetMaxOpenSize(symbol string, price string, leverage string) (string, error) {
 params := map[string]string{
  "symbol":      symbol,
  "productType": "USDT-FUTURES",
  "marginCoin":  "USDT",
  "openPrice":   price,
  "openAmount":  "1000", // Margin to use
  "leverage":    leverage,
 }

 return a.mixAccountClient.OpenCount(params)
}
```

