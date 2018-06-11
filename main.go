package main

//import "promocalc"
import "router"
//import "models"
import "fmt"

const ListenAddress = ":8090"

func main() {
	router.Start(ListenAddress)
  fmt.Printf("Server listening on: %s ...",ListenAddress)
  fmt.Println()
  // List to quit channel here later, quit gracefully
	select {}
  /*
  var p promocalc.Calculator
  p = &promocalc.PromoCalculator{}

  //if you buy two shirts, each additional shirt costs only 45$
  promo2 := models.Promo { Id: "promo2",
  Description: "if you buy two shirts, each additional shirt costs only 45$",
  Buys: []models.Buy{ models.Buy{ Category: "shirts", Count: 2 }  },
  Gets: []models.Get{ models.Get{ Category: "shirts", All: true, Off: models.Off{ Fixed: &models.Fixed{Price: 45} } } }}

  //if you purchase 3 or more shirts, all ties are half price
  promo3 := models.Promo { Id: "promo3",
  Description: "if you purchase 3 or more shirts, all ties are half price",
  Buys: []models.Buy{ models.Buy{ Category: "shirts", Count: 3 }  },
  Gets: []models.Get{ models.Get{ Category: "ties", All: true, Off: models.Off{ Discount: &models.Discount{Percentage: 50} } } }}


  cart := &models.Cart { Items: []models.Item{
    models.Item { Category: "shirts", Price: float32(100) }, models.Item { Category: "shirts", Price: float32(200) }, models.Item { Category: "shirts", Price: float32(130) }, models.Item { Category: "ties", Price: float32(50) }, models.Item { Category: "ties", Price: float32(80) } } }

  r := p.ApplyPromos([]models.Promo{promo2, promo3}, cart)
  fmt.Println(r)
  */

}
