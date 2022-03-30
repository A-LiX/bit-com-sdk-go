package examples

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/wsclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model/ws"
	"github.com/tidwall/gjson"
)

var t_cancel *time.Time
var oid *string
var f13 *os.File

func responseHandlerExample(resp interface{}) {
	switch t := resp.(type) {
	case ws.SubscriptionSuccessResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive subscription success response: %s", respJson)
	case ws.SubscriptionFailResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive subscription fail response: %s", respJson)
	case ws.DepthSnapshotResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive depth snapshot response: %s", respJson)
	case ws.DepthUpdateResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive depth update response: %s", respJson)
	case ws.OrderBookResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive orderbook response: %s", respJson)
	case ws.Depth1Response:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive depth1 response: %s", respJson)
	case ws.TickerResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive ticker response: %s", respJson)
	case ws.KlineResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive kline response: %s", respJson)
	case ws.TradeResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive trade response: %s", respJson)
	case ws.MarketTradeResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive market trade response: %s", respJson)
	case ws.IndexPriceResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive index price response: %s", respJson)
	case ws.MarkPriceResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive mark price response: %s", respJson)
	case ws.AccountResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive account response: %s", respJson)
	case ws.PositionResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Receive position response: %s", respJson)
	case ws.OrderResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}

		data := gjson.Get(respJson, "data")
		//

		for _, r := range data.Array() {
			if strings.Compare(r.Map()["order_id"].Str, *oid) == 0 {
				if strings.Compare(r.Map()["status"].Str, "cancelled") == 0 {
					fmt.Println("give_order_id:", *oid)
					fmt.Println("cancel_order_id:", data.Array()[0].Map()["order_id"])
					t3 := time.Now()
					d2 := t3.Sub(*t_cancel)
					fmt.Println("t3-t1=", d2)

					str1 := []byte(*oid)
					str2 := []byte(",")
					str3 := []byte(d2.String())
					str4 := []byte("\n")
					str1 = append(str1, str2...)
					str1 = append(str1, str3...)
					str1 = append(str1, str4...)
					_, _ = f13.Write([]byte(str1))
					break
				}
			}
		}
		// mp1 := make(map[string]interface{})

		// ris := [](map[string]interface{}){}
		// err := json.Unmarshal([]byte(mp1["data"].respJson), &ris)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		//applogger.Info("Receive order response: %s", respJson)
	case ws.UserTradeResponse:
		respJson, jsonErr := model.ToJson(resp)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}

		applogger.Info("Receive user trades response: %s", respJson)
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}
}

func getWsAuthToken() string {
	wsAuthClient := new(restclient.WsAuthClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)
	resp, err := wsAuthClient.GetWsAuthToken()
	if err != nil {
		applogger.Error("Get auth token error: %s", err)
		return ""
	} else {
		return resp.Data.Token
	}
}

func PublicSubscribeExample() {

	client := new(wsclient.PublicWebsocketClient).Init(config.WsHost, 60)

	paramMap := map[string]interface{}{
		"type":        "subscribe",
		"channels":    []string{"ticker"},
		"instruments": []string{"BTC-PERPETUAL"},
		"interval":    "100ms",
	}

	client.SetHandler(
		func() {
			client.Subscribe(paramMap)
		},
		responseHandlerExample,
	)

	client.Connect(false)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	paramMap["instruments"] = []string{"BTC-PERPETUAL"}
	client.Subscribe(paramMap)
	fmt.Scanln()

	paramMap["instruments"] = []string{"BTC-PERPETUAL", "ETH-PERPETUAL"}
	client.Subscribe(paramMap)
	fmt.Scanln()

	client.Close()
	applogger.Info("Client closed")
}

func PrivateSubscribeExample(order_id *string, t1 *time.Time, file13 *os.File) {
	client := new(wsclient.PrivateWebsocketClient).Init(config.WsHost, getWsAuthToken, 60)

	paramMap := map[string]interface{}{
		"type":        "subscribe",
		"instruments": []string{"BTC-PERPETUAL"},
		"channels":    []string{"order"},
		"currencies":  []string{"BTC"},
		"categories":  []string{"future"},
		"interval":    "100ms",
	}
	client.Connect(true)
	t_cancel = t1
	oid = order_id
	f13 = file13
	client.SetHandler(
		func() {
			client.Subscribe(paramMap)
		},
		responseHandlerExample,
	)

	client.Connect(false)

	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()

	client.Close()
	applogger.Info("Client closed")
}
