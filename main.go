package main

import "promocalc"

func main() {
  var p promocalc.Calculator
  p = &promocalc.PromoCalculator{}

  // buy 2 shirts and get 1 free
  buys := make([]promocalc.Buy,0)
  buys = append(buys, promocalc.Buy{ Category: "shirts", Count: 2 })
  gets := make([]promocalc.Get, 0)
  gets = append(gets, promocalc.Get{ Category: "shirts", Count: 1, All: false, Off: promocalc.Off{ Discount: &promocalc.Discount{ Percentage: 100 } } })
  promo := promocalc.Promo { Id: "promo1",  Buys: buys, Gets: gets }

  cart := &promocalc.Cart { Items: []promocalc.Item{
    promocalc.Item { Category: "shirts", Price: float32(100) }, promocalc.Item { Category: "shirts", Price: float32(200) }, promocalc.Item { Category: "shirts", Price: float32(130) } } }

  p.Calculate(promo, cart)
}
