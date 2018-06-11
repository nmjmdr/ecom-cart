package main

//import "promocalc"
import "router"
//import "models"

func main() {
	router.Start()
	select {}
  /*
  var p promocalc.Calculator
  p = &promocalc.PromoCalculator{}

  // buy 2 shirts and get 1 free
  buys := make([]models.Buy,0)
  buys = append(buys, models.Buy{ Category: "shirts", Count: 2 })
  gets := make([]models.Get, 0)
  gets = append(gets, models.Get{ Category: "shirts", Count: 1, All: false, Off: models.Off{ Discount: &models.Discount{ Percentage: 100 } } })
  promo := models.Promo { Id: "promo1",  Buys: buys, Gets: gets }

  cart := &models.Cart { Items: []models.Item{
    models.Item { Category: "shirts", Price: float32(100) }, models.Item { Category: "shirts", Price: float32(200) }, models.Item { Category: "shirts", Price: float32(130) } } }

  p.Calculate(promo, cart)
  */

}
