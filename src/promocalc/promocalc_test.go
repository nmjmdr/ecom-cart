package promocalc

import (
	"models"
	"testing"
)

func Test_When_BuyXGetYPromo_Is_Applied_It_Should_Apply_To_All_Eligble_Items(t *testing.T) {
	promo := models.Promo{Id: "promo4",
		Description: "buy 2 shirts and get 1 shirt free",
		Buys:        []models.Buy{models.Buy{Category: "shirts", Count: 2}},
		Gets:        []models.Get{models.Get{Category: "shirts", Count: 1, Off: models.Off{Discount: &models.Discount{Percentage: 100}}}}}

	cart := &models.Cart{Items: []models.Item{
		models.Item{Category: "shirts", Price: float32(100)}, models.Item{Category: "shirts", Price: float32(200)}, models.Item{Category: "shirts", Price: float32(130)}, models.Item{Category: "shirts", Price: float32(50)}}}

	p := NewCalculator()
	promofiedCart := p.ApplyPromos([]models.Promo{promo}, cart)
	if promofiedCart.TotalPrice != 480 {
		t.Fatal("Does not match the expected total price")
	}
	if promofiedCart.TotalOffPrice != 350 {
		t.Fatal("Does not match the expected total off price")
	}
}

func Test_When_An_Item_Is_Elgible_For_Multiple_Promos_It_Should_Consider_Lowest_Off_Price(t *testing.T) {
	promo1 := models.Promo{Id: "promo4",
		Description: "buy 2 shirts and get 1 shirt free",
		Buys:        []models.Buy{models.Buy{Category: "shirts", Count: 2}},
		Gets:        []models.Get{models.Get{Category: "shirts", Count: 1, Off: models.Off{Discount: &models.Discount{Percentage: 100}}}}}

	//if you buy two shirts, each additional shirt costs only 45$
	promo2 := models.Promo{Id: "promo2",
		Description: "if you buy two shirts, each additional shirt costs only 45$",
		Buys:        []models.Buy{models.Buy{Category: "shirts", Count: 2}},
		Gets:        []models.Get{models.Get{Category: "shirts", All: true, Off: models.Off{Fixed: &models.Fixed{Price: 45}}}}}

	cart := &models.Cart{Items: []models.Item{
		models.Item{Category: "shirts", Price: float32(100)}, models.Item{Category: "shirts", Price: float32(200)}, models.Item{Category: "shirts", Price: float32(130)}, models.Item{Category: "shirts", Price: float32(50)}}}

	p := NewCalculator()
	promofiedCart := p.ApplyPromos([]models.Promo{promo1, promo2}, cart)
	if promofiedCart.TotalPrice != 480 {
		t.Fatal("Does not match the expected total price")
	}
	// 100 + 120 + 0 (130 => 0) + 45
	if promofiedCart.TotalOffPrice != 345 {
		t.Fatal("Does not match the expected total off price")
	}
}
