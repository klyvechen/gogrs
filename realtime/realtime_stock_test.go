package realtime

import (
	"fmt"
	"testing"
	"time"
)

func TestStockRealTime(*testing.T) {
	r := &StockRealTime{
		No: "2618",
		//Date:      time.Now(),
		Date:     time.Date(2015, 4, 1, 0, 0, 0, 0, time.Local),
		Exchange: "tse",
	}

	r.URL()
	v, _ := r.Get()
	fmt.Println(v.BestAskPrice)
	fmt.Println(v.BestBidPrice)
	fmt.Println(v.BestAskVolume)
	fmt.Println(v.BestBidVolume)
	fmt.Println(v)
	fmt.Println("UnixMapData", r.UnixMapData)
}

func TestStockRealTimeOTC(*testing.T) {
	r := &StockRealTime{
		No: "8446",
		//Date:      time.Now(),
		Date:     time.Date(2015, 4, 1, 0, 0, 0, 0, time.Local),
		Exchange: "otc",
	}

	r.URL()
	v, _ := r.Get()
	fmt.Println(v.BestAskPrice)
	fmt.Println(v.BestBidPrice)
	fmt.Println(v.BestAskVolume)
	fmt.Println(v.BestBidVolume)
	fmt.Println(v)
	fmt.Println("UnixMapData", r.UnixMapData)
}

func BenchmarkGet(b *testing.B) {
	r := &StockRealTime{
		No: "2618",
		//Date:      time.Now(),
		Date:     time.Date(2015, 4, 1, 0, 0, 0, 0, time.Local),
		Exchange: "tse",
	}

	for i := 0; i <= b.N; i++ {
		r.Get()
	}
}

func ExampleStockRealTime() {
	r := StockRealTime{
		No:       "2618",
		Date:     time.Date(2014, 12, 26, 0, 0, 0, 0, time.Local),
		Exchange: "tse",
	}

	data, _ := r.Get()
	fmt.Printf("%v", data)
}
