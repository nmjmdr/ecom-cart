package promocalc

//import "fmt"
import "models"

type Calculator interface {
	ApplyPromos(promo []models.Promo, cart *models.Cart) []models.MarkedItem
}

type appliedBuys struct {
	applied      bool
	groupedItems map[string]([]models.MarkedItem)
}

type appliedBuy struct {
	applied      bool
	groupedItems map[string]([]models.MarkedItem)
}

func markBuyItems(groupedItems map[string]([]models.MarkedItem), buy models.Buy, promo models.Promo) map[string]([]models.MarkedItem) {
	markCount := 0
	for i := 0; i < len(groupedItems[buy.Category]); i++ {
		item := groupedItems[buy.Category][i]
		if markCount == buy.Count {
			break
		}
		_, isMarkedBuy := item.MarkedBuys[promo.Id]
		_, isMarkedGet := item.MarkedGets[promo.Id]
		if !isMarkedBuy && !isMarkedGet {
			groupedItems[buy.Category][i].MarkedBuys[promo.Id] = true
			markCount = markCount + 1
		}
	}
	return groupedItems
}

func applyBuy(buy models.Buy, groupedItems map[string]([]models.MarkedItem), promo models.Promo) appliedBuy {
	_, ok := groupedItems[buy.Category]
	if !ok {
		return appliedBuy{applied: false, groupedItems: groupedItems}
	}
	var matchedItems []models.MarkedItem
	for _, item := range groupedItems[buy.Category] {
		_, isMarkedBuy := item.MarkedBuys[promo.Id]
		_, isMarkedGet := item.MarkedGets[promo.Id]
		if !isMarkedBuy && !isMarkedGet {
			matchedItems = append(matchedItems, item)
		}
	}
	if len(matchedItems) < buy.Count {
		return appliedBuy{applied: false, groupedItems: groupedItems}
	}
	groupedItems = markBuyItems(groupedItems, buy, promo)
	return appliedBuy{applied: true, groupedItems: groupedItems}
}

func applyBuys(groupedItems map[string]([]models.MarkedItem), promo models.Promo) appliedBuys {
	applied := true
	for _, buy := range promo.Buys {
		appliedBuy := applyBuy(buy, groupedItems, promo)
		applied = applied && appliedBuy.applied
	}
	return appliedBuys{applied: applied, groupedItems: groupedItems}
}

func computeOffPrice(price float32, off models.Off) float32 {
	if off.Discount != nil {
		return price - (price * off.Discount.Percentage / 100)
	} else {
		return off.Fixed.Price
	}
}

func markGetItems(groupedItems map[string]([]models.MarkedItem), get models.Get, promo models.Promo) map[string]([]models.MarkedItem) {
	markCount := 0
	for i := 0; i < len(groupedItems[get.Category]); i++ {
		item := groupedItems[get.Category][i]
		if get.All == false && markCount == get.Count {
			break
		}
		_, isMarkedBuy := item.MarkedBuys[promo.Id]
		_, isMarkedGet := item.MarkedGets[promo.Id]
		if !isMarkedBuy && !isMarkedGet {
			groupedItems[get.Category][i].MarkedGets[promo.Id] = computeOffPrice(item.Item.Price, get.Off)
			markCount = markCount + 1
		}
	}
	return groupedItems
}

func applyGet(get models.Get, groupedItems map[string]([]models.MarkedItem), promo models.Promo) map[string]([]models.MarkedItem) {
	_, ok := groupedItems[get.Category]
	if !ok {
		return groupedItems
	}
	groupedItems = markGetItems(groupedItems, get, promo)
	return groupedItems
}

func applyGets(groupedItems map[string]([]models.MarkedItem), promo models.Promo) map[string]([]models.MarkedItem) {
	for _, get := range promo.Gets {
		groupedItems = applyGet(get, groupedItems, promo)
	}
	return groupedItems
}

type PromoCalculator struct {
}

func (p *PromoCalculator) applyPromo(promo models.Promo, groupedItems map[string]([]models.MarkedItem)) map[string]([]models.MarkedItem) {
	for ;true; {
		var appliedBuys = applyBuys(groupedItems, promo)
		if !appliedBuys.applied {
      break;
    }
		groupedItems = appliedBuys.groupedItems
		groupedItems = applyGets(groupedItems, promo)
	}
  return groupedItems
}

func (p *PromoCalculator) ApplyPromos(promos []models.Promo, cart *models.Cart) []models.MarkedItem {
  groupedItems := make(map[string]([]models.MarkedItem))
	for _, item := range cart.Items {
		groupedItems[item.Category] = append(groupedItems[item.Category],
			models.MarkedItem{Item: item, MarkedBuys: make(map[string]bool), MarkedGets: make(map[string]float32)})
	}
  for _, promo := range promos {
    groupedItems = p.applyPromo(promo, groupedItems)
  }
  items := make([]models.MarkedItem, 0)
  for _, item := range groupedItems {
    items = append(items, item...)
  }
  return items
}

func NewCalculator() Calculator {
  p := &PromoCalculator{}
  return p
}
