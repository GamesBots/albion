package albion

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

var (
	itemsURL   = "https://raw.githubusercontent.com/broderickhyman/ao-bin-dumps/master/formatted/items.json"
	apiURL     = "https://www.albion-online-data.com"
	priceRoute = "/api/v2/stats/prices/{itemID}"
)

// ItemPrice from albion data online
type ItemPrice struct {
	ItemID           string `json:"item_id"`
	City             string `json:"city"`
	Quality          int    `json:"quality"`
	SellPriceMin     int    `json:"sell_price_min"`
	SellPriceMinDate myTime `json:"sell_price_min_date"`
	SellPriceMax     int    `json:"sell_price_max"`
	SellPriceMaxDate myTime `json:"sell_price_max_date"`
	BuyPriceMin      int    `json:"buy_price_min"`
	BuyPriceMinDate  myTime `json:"buy_price_min_date"`
	BuyPriceMax      int    `json:"buy_price_max"`
	BuyPriceMaxDate  myTime `json:"buy_price_max_date"`
}

type myTime struct {
	time.Time
}

func (t *myTime) UnmarshalJSON(buf []byte) error {
	tt, err := time.Parse(time.RFC3339, strings.Trim(string(buf), `"`)+"Z")
	if err != nil {
		return err
	}
	t.Time = tt
	return nil
}

// GetPrice for items in args
func GetPrice(item string) {
	log.Info("getting price for item: ", item)
	var prices []ItemPrice
	client := resty.New()

	client.R().SetPathParams(map[string]string{
		"itemID": item,
	}).
		SetResult(&prices).
		Get(apiURL + priceRoute)
	fmt.Printf("price object: %s\n", item)
	for _, price := range prices {
		fmt.Printf("\nquality: %d\n", price.Quality)
		fmt.Printf("location: %s\nsell:\tmin: %-10d max: %-10d\nbuy :\tmin: %-10d max: %-10d\n", price.City, price.SellPriceMin, price.SellPriceMax, price.BuyPriceMin, price.BuyPriceMax)
	}
}
