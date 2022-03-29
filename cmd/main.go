package main

import (
	"fmt"
	"time"

	"github.com/bitcom-exchange/bitcom-go-api/cmd/examples"
)

func main() {
	// system client
	//examples.GetSystemVersionExample()
	//examples.GetSystemTimestampExample()
	//examples.GetCancelOnlyStatusExample()

	// market client
	//examples.GetIndexExample()
	//examples.GetInstrumentsExample()
	//examples.GetTickerExample()
	//examples.GetOrderBookExample()
	//examples.GetMarketTradeExample()
	//examples.GetKlinesExample()
	//examples.GetDeliveryInfoExample()
	//examples.GetMarketSummaryExample()
	//examples.GetFundingRateExample()
	//examples.GetFundingRateHistoryExample()
	//examples.GetTotalVolumeExample()

	// account client
	//examples.GetAccountsExample()
	//examples.GetPositionsExample()
	//examples.GetTransactionLogsExample()
	//examples.GetUserDeliveriesExample()
	//examples.GetUserSettlementsExample()
	//examples.ConfigCodExample()
	//examples.GetCodConfigExample()
	//examples.GetMmpStateExample()
	//examples.UpdateMmpConfigExample()
	//examples.ResetMmpStateExample()

	order_id := examples.PlaceNewOrderExample()
	fmt.Println("ordrid==================================================================", order_id)

	var t_temp time.Time
	// order client
	go examples.PrivateSubscribeExample(order_id, &t_temp)

	time.Sleep(time.Second * 10)

	//examples.PlaceNewBatchOrderExample()
	//order_id := "315233500"

	t1 := time.Now() //获取本地现在时间
	t_temp = t1
	examples.CancelOrderExample(order_id, &t_temp)
	time.Sleep(time.Second * 20)

	//cancel_status := false

	// go func() {
	// 	for {
	//order_id := "123456789"
	//_ = examples.GetOrdersExample(order_id)
	// 		if cancel_status == true {
	// 			t3 := time.Now()
	// 			d2 := t3.Sub(t1)
	// 			fmt.Println("t3-t1=", d2)
	// 			break
	// 		} else {
	// 			fmt.Println(order_id, " cancel operation uncomplete")
	// 		}
	// 	}
	// }()
	//examples.AmendOrderExample()
	//examples.AmendBatchOrdersExample()
	//examples.ClosePositionsExample()
	//examples.GetOpenOrdersExample()

	//examples.GetOrdersExample()
	//examples.GetStopOrdersExample()
	//examples.GetUserTradesExample()
	//examples.GetEstMarginsExample()

	// Block Trade
	//examples.NewParadigmBlockOrdersExample()
	//examples.QueryParadigmBlockOrdersExample()
	//examples.QueryParadigmBlockOrdersByPlatformExample()

	// WebSocket
	//examples.PublicSubscribeExample()
}
