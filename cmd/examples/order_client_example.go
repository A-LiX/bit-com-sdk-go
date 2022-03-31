package examples

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bitcom-exchange/bitcom-go-api/config"
	"github.com/bitcom-exchange/bitcom-go-api/logging/applogger"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/client/restclient"
	"github.com/bitcom-exchange/bitcom-go-api/pkg/model"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

func PlaceNewOrderExample() string {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["qty"] = "10"
	paramMap["side"] = "sell"
	paramMap["price"] = "55000.00"
	paramMap["order_type"] = "limit"

	resp, err := orderClient.NewOrder(paramMap)
	if err != nil {
		applogger.Error("Place new order error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		order_id := gjson.Get(respJson, "order_id")

		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			if jsonErr != nil {
				applogger.Info("Place new order: \n%s", pretty.Pretty([]byte(respJson)))
			}
		}
		return order_id.Str
	}
	ret_str := "null"
	return ret_str
}

func PlaceNewBatchOrderExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	orderMapOne := make(map[string]interface{})

	orderMapOne["instrument_id"] = "BTC-PERPETUAL"
	orderMapOne["qty"] = "1000"
	orderMapOne["side"] = "buy"
	orderMapOne["price"] = "10380.00"
	orderMapOne["order_type"] = "limit"

	orderMapTwo := make(map[string]interface{})

	orderMapTwo["instrument_id"] = "BTC-PERPETUAL"
	orderMapTwo["qty"] = "1000"
	orderMapTwo["side"] = "sell"
	orderMapTwo["price"] = "10410.00"
	orderMapTwo["order_type"] = "limit"

	orderSlice := []map[string]interface{}{orderMapOne, orderMapTwo}

	paramMap := map[string]interface{}{
		"orders_data": orderSlice,
	}

	resp, err := orderClient.NewBatchOrders(paramMap)
	if err != nil {
		applogger.Error("Place batch new orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Place batch new orders: \n%s", pretty.Pretty([]byte(respJson)))

		}
	}
}

func CancelOrderExample(order_id *string, t1 time.Time, file12 *os.File) {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["order_id"] = *order_id

	resp, err := orderClient.CancelOrders(paramMap)
	if err != nil {
		applogger.Error("Cancel orders error: %s", err)
	} else {
		t2 := time.Now()
		d1 := t2.Sub(t1)
		fmt.Println("t2-t1=", d1)

		str0 := []byte(t_cancel.Format("15:04:05.000"))
		str1 := []byte(*order_id)
		str2 := []byte(",")
		str3 := []byte(d1.String())
		str4 := []byte("\n")

		str0 = append(str0, str2...)
		str0 = append(str0, str1...)
		str0 = append(str0, str2...)
		str0 = append(str0, str3...)
		str0 = append(str0, str4...)

		_, err = file12.Write([]byte(str0))
		if err == nil {
			fmt.Printf("writed to file12: %s\n", str0)
		}

		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		}
		applogger.Info("Cancel orders: \n%s", pretty.Pretty([]byte(respJson)))
	}
}

func AmendOrderExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["order_id"] = "1373134"
	paramMap["price"] = "10600.00"

	resp, err := orderClient.AmendOrder(paramMap)
	if err != nil {
		applogger.Error("Amend order error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Amend order: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func AmendBatchOrdersExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	orderMapOne := make(map[string]interface{})

	orderMapOne["order_id"] = "1373574"
	orderMapOne["price"] = "10701.00"

	orderMapTwo := make(map[string]interface{})

	orderMapTwo["order_id"] = "1373134"
	orderMapTwo["price"] = "10601.00"

	orderSlice := []map[string]interface{}{orderMapOne, orderMapTwo}

	paramMap := map[string]interface{}{
		"orders_data": orderSlice,
	}

	resp, err := orderClient.AmendBatchOrders(paramMap)
	if err != nil {
		applogger.Error("Amend batch orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Amend batch orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func ClosePositionsExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["order_type"] = "limit"
	paramMap["price"] = "10286.50"

	resp, err := orderClient.ClosePositions(paramMap)
	if err != nil {
		applogger.Error("Close positions error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Close positions: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetOpenOrdersExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"

	resp, err := orderClient.GetOpenOrders(paramMap)
	if err != nil {
		applogger.Error("Get open orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get open orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetOrdersExample(order_id string) bool {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["category"] = "future"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := orderClient.GetOrders(paramMap)
	if err != nil {
		applogger.Error("Get orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			//var jArray []Jinfo
			ris := [](map[string]interface{}){}
			err := json.Unmarshal([]byte(respJson), &ris)
			if err != nil {
				fmt.Println(err)
			}
			applogger.Info("Get  orders: \n%s", pretty.Pretty([]byte(respJson)))
			//fmt.Println("1:", order_id, "   2:", ris[0]["order_id"])
			if strings.Compare(order_id, ris[0]["order_id"].(string)) == 0 {
				if strings.Compare("cancelled", ris[0]["status"].(string)) == 0 {
					return true
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return false
}

func GetStopOrdersExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := orderClient.GetStopOrders(paramMap)
	if err != nil {
		applogger.Error("Get stop orders error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get stop orders: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetUserTradesExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["currency"] = "BTC"
	paramMap["instrument_id"] = "BTC-PERPETUAL"

	resp, err := orderClient.GetUserTrades(paramMap)
	if err != nil {
		applogger.Error("Get user trades error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get user trades: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}

func GetEstMarginsExample() {
	orderClient := new(restclient.OrderClient).Init(config.User1Host, config.User1AccessKey, config.User1SecretKey)

	paramMap := make(map[string]interface{})
	paramMap["instrument_id"] = "BTC-PERPETUAL"
	paramMap["price"] = "10318.00"
	paramMap["qty"] = "2000"

	resp, err := orderClient.GetEstMargins(paramMap)
	if err != nil {
		applogger.Error("Get estimated margins error: %s", err)
	} else {
		respJson, jsonErr := model.ToJson(resp.Data)
		if jsonErr != nil {
			applogger.Error("Marshal response error: %s", jsonErr)
		} else {
			applogger.Info("Get estimated margins: \n%s", pretty.Pretty([]byte(respJson)))
		}
	}
}
