package main

import (
	"fmt"
	"os"
	"runtime"
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

	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)
	runtime.GOMAXPROCS(cpuNum)

	file12, err := os.Create("t1_t2.csv")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	file13, err := os.Create("t1_t3.csv")
	if err != nil {
		fmt.Printf("err_file13: %v\n", err)
	}
	defer file12.Close()
	defer file13.Close()

	var t_temp time.Time

	t_temp = time.Now()

	var order_id *string
	oid := "000000000"
	order_id = &oid
	go examples.PrivateSubscribeExample(order_id, &t_temp, file13)
	time.Sleep(time.Second * 10)

	for {
		*order_id = examples.PlaceNewOrderExample()
		fmt.Println("place_order_id:", *order_id)
		t1 := time.Now() //获取本地现在时间
		t_temp = t1
		examples.CancelOrderExample(order_id, t_temp, file12)
		time.Sleep(time.Second * 10)

		fmt.Println("==================================================================================================")
	}

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
