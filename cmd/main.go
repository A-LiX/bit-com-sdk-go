package main

import (
	"fmt"
	"time"

	"github.com/bitcom-exchange/bitcom-go-api/cmd/examples"
)

func hello() {
	fmt.Println("helllllllllllllllllo")
}

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

	// order client

	order_id := examples.PlaceNewOrderExample()
	//examples.PlaceNewBatchOrderExample()
	//order_id := "315233500"
	fmt.Println("ordrid-----------------------------", order_id)

	t1 := time.Now() //获取本地现在时间
	examples.CancelOrderExample(order_id)
	go examples.GetOrdersExample()
	go hello()
	t2 := time.Now()
	d := t2.Sub(t1)
	fmt.Println(d)

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
	//examples.PrivateSubscribeExample()
}
