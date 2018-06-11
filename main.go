package main

import "promocalc"
//import "router"
import "models"
import "fmt"

func main() {
	//router.Start()
	//select {}

  var p promocalc.Calculator
  p = &promocalc.PromoCalculator{}

  promo1 := models.Promo { Id: "promo1",
  Description: "if you buy 2 or more trousers, you get 15 percent off belts and shoes",
  Buys: []models.Buy{ models.Buy{ Category: "trousers", Count: 2 }  },
  Gets: []models.Get{ models.Get{ Category: "belts", All: true, Off: models.Off{ Discount: &models.Discount{Percentage: 15} } },
   models.Get{ Category: "shoes", All: true, Off: models.Off{ Discount: &models.Discount{Percentage: 15} } } }}

  cart := &models.Cart { Items: []models.Item{
    models.Item { Category: "trousers", Price: float32(100) }, models.Item { Category: "trousers", Price: float32(200) }, models.Item { Category: "belts", Price: float32(130) } } }

  r := p.ApplyPromos([]models.Promo{promo1}, cart)
  fmt.Println(r)
  
}
