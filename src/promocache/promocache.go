package promocache

import "models"

type PromoCache interface {
	GetAll() []models.Promo
	Get(promoId string) (models.Promo, bool)
}

type DefaultPromoCache struct {
	promos map[string]models.Promo
}

var cache PromoCache

func (d *DefaultPromoCache) GetAll() []models.Promo {
	promos := make([]models.Promo, 0)
	for _, promo := range d.promos {
		promos = append(promos, promo)
	}
	return promos
}

func (d *DefaultPromoCache) Get(promoId string) (models.Promo, bool) {
	promo, ok := d.promos[promoId]
	return promo, ok
}

func getDefaultPromos() []models.Promo {
	// if you buy 2 or more trousers, you get 15% off belts and shoes
	promo1 := models.Promo{Id: "promo1",
		Description: "if you buy 2 or more trousers, you get 15 percent off belts and shoes",
		Buys:        []models.Buy{models.Buy{Category: "trousers", Count: 2}},
		Gets: []models.Get{models.Get{Category: "belts", All: true, Off: models.Off{Discount: &models.Discount{Percentage: 15}}},
			models.Get{Category: "shoes", All: true, Off: models.Off{Discount: &models.Discount{Percentage: 15}}}}}

	//if you buy two shirts, each additional shirt costs only 45$
	promo2 := models.Promo{Id: "promo2",
		Description: "if you buy two shirts, each additional shirt costs only 45$",
		Buys:        []models.Buy{models.Buy{Category: "shirts", Count: 2}},
		Gets:        []models.Get{models.Get{Category: "shirts", All: true, Off: models.Off{Fixed: &models.Fixed{Price: 45}}}}}

	//if you purchase 3 or more shirts, all ties are half price
	promo3 := models.Promo{Id: "promo3",
		Description: "if you purchase 3 or more shirts, all ties are half price",
		Buys:        []models.Buy{models.Buy{Category: "shirts", Count: 3}},
		Gets:        []models.Get{models.Get{Category: "ties", All: true, Off: models.Off{Discount: &models.Discount{Percentage: 50}}}}}

	// buy 2 shirts and get 1 shirt free
	promo4 := models.Promo{Id: "promo4",
		Description: "buy 2 shirts and get 1 shirt free",
		Buys:        []models.Buy{models.Buy{Category: "shirts", Count: 2}},
		Gets:        []models.Get{models.Get{Category: "shirts", Count: 1, Off: models.Off{Discount: &models.Discount{Percentage: 100}}}}}

	return []models.Promo{promo1, promo2, promo3, promo4}

}

func newDefaultPromoCache() PromoCache {
	c := &DefaultPromoCache{promos: make(map[string]models.Promo)}
	// for now add default promos
	defaultPromos := getDefaultPromos()
	for _, promo := range defaultPromos {
		c.promos[promo.Id] = promo
	}
	return c
}

func GetPromoCache() PromoCache {
	if cache == nil {
		cache = newDefaultPromoCache()
	}
	return cache
}
