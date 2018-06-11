package promocalc

import "fmt"
import "models"


type Calculator interface {
	Calculate(promo models.Promo, cart *models.Cart)
}

type markedItem struct {
	item       models.Item
	markedBuys map[string]bool
	markedGets map[string]float32
}

type appliedBuys struct {
	applied      bool
	groupedItems map[string]([]markedItem)
}

type appliedBuy struct {
	applied      bool
	groupedItems map[string]([]markedItem)
}

func markBuyItems(groupedItems map[string]([]markedItem), buy models.Buy, promo models.Promo) map[string]([]markedItem) {
	markCount := 0
	for i := 0; i < len(groupedItems[buy.Category]); i++ {
		item := groupedItems[buy.Category][i]
		if markCount == buy.Count {
			break
		}
		_, isMarkedBuy := item.markedBuys[promo.Id]
		_, isMarkedGet := item.markedGets[promo.Id]
		if !isMarkedBuy && !isMarkedGet {
			groupedItems[buy.Category][i].markedBuys[promo.Id] = true
			markCount = markCount + 1
		}
	}
	return groupedItems
}

func applyBuy(buy models.Buy, groupedItems map[string]([]markedItem), promo models.Promo) appliedBuy {
	_, ok := groupedItems[buy.Category]
	if !ok {
		return appliedBuy{applied: false, groupedItems: groupedItems}
	}
	var matchedItems []markedItem
	for _, item := range groupedItems[buy.Category] {
		_, isMarkedBuy := item.markedBuys[promo.Id]
		_, isMarkedGet := item.markedGets[promo.Id]
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

func applyBuys(groupedItems map[string]([]markedItem), promo models.Promo) appliedBuys {
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

func markGetItems(groupedItems map[string]([]markedItem), get models.Get, promo models.Promo) map[string]([]markedItem) {
	markCount := 0
	for i := 0; i < len(groupedItems[get.Category]); i++ {
		item := groupedItems[get.Category][i]
		if markCount == get.Count {
			break
		}
		_, isMarkedBuy := item.markedBuys[promo.Id]
		_, isMarkedGet := item.markedGets[promo.Id]
		if !isMarkedBuy && !isMarkedGet {
			groupedItems[get.Category][i].markedGets[promo.Id] = computeOffPrice(item.item.Price, get.Off)
			markCount = markCount + 1
		}
	}
	return groupedItems
}

func applyGet(get models.Get, groupedItems map[string]([]markedItem), promo models.Promo) map[string]([]markedItem) {
	_, ok := groupedItems[get.Category]
	if !ok {
		return groupedItems
	}
	groupedItems = markGetItems(groupedItems, get, promo)
	return groupedItems
}

func applyGets(groupedItems map[string]([]markedItem), promo models.Promo) map[string]([]markedItem) {
	for _, get := range promo.Gets {
		groupedItems = applyGet(get, groupedItems, promo)
	}
	return groupedItems
}

type PromoCalculator struct {
}

func (p *PromoCalculator) Calculate(promo models.Promo, cart *models.Cart) {
	groupedItems := make(map[string]([]markedItem))
	for _, item := range cart.Items {
		groupedItems[item.Category] = append(groupedItems[item.Category],
			markedItem{item: item, markedBuys: make(map[string]bool), markedGets: make(map[string]float32)})
	}
	var count = 0
	for count < 1 {
		var appliedBuys = applyBuys(groupedItems, promo)
		count = count + 1
		groupedItems = appliedBuys.groupedItems
		groupedItems = applyGets(groupedItems, promo)
		fmt.Println(groupedItems)
	}
}
