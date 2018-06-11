# ecom-cart

ecom-cart implements a promotion engine which can compute promotion rules such as:
1. if you buy 2 or more trousers, you get 15% off belts and shoes
2. if you buy two shirts, each additional shirt costs only 45$
3. if you purchase 3 or more shirts, all ties are half price
4. buy 2 shirts and get 1 shirt free

Notice that fourth rule (buy 2 shirts and get 1 shirt free) has to be applied in such way if the customer has purchased 7 shirts then two of them should be free.


## Design Details
Given time, the project would follow the below design approach:
![design](https://raw.githubusercontent.com/nmjmdr/ecom-cart/master/screenshots/design.png)

_Cart Service:_
1. Mantains carts in a distributed cache (such as Redis)
2. Gets promotions from promotions service and caches them
3. Applies the promotions on carts and caches the computation; Invalidating the cache if a cart is updated or promos are updated
3. In order to be aware of changes to promos, it listens to events from promotions service for:
     1. New promotions
     2. Deletion of promotions
     3. Update of promotions
4. Promo calculation is CPU intensive. Further tests could be done to determine if it makes sense to movie Promo calcuation as service and provide a cluster of CPU heavy nodes

### Endpoints and current design:
```
. Controllers (CartController, PromoController)
. CartController provides the following endpoints:
    - GET: /carts/{id} 
    - POST: /carts/  > Creates a new cart
    - DEL: /carts/{id}
    - POST /carts/{cartId}/promofied > Applies promos to a cart. The ids of promos that need to be applied has to be passed as an array in POST body:
    {
       promos: ["promo1","promo2"]
    }
. PromoController
    - GET: /promos : Gets all promos (as JSON objects) 
. PromoCalc : Computes and applies the promotions to a cart
. PromoCache: Provides the promos that can be applied. Currently the four promos have been hardcoded in the code. The PromoController can be easily extended to support addition of new promos
. Repository: Stores all carts. Currently an in-memory version of the repository has been immplemented. This needs to be changed to use Redis.
```
## Representation of a promo:
A promo is represented as follows:
```
{
  id: "promo-id",
  description "Buy two shirt and a trouser and get a shirt free"
  buys: [
     { 
       category: "shirts",
       count: 2
     },
     {
       category: "trousers",
       count: 1
     }
  ],
 gets: [
   {
     category: "shirts"
     count: 1,
     off : {
       discount: { percentage: 100 } 
     }
   }
 ]
```
The above promo represents `Buy two shirt and a trouser and get a shirt free`. As it can be observed, the schema to represent a promo is quite flexible and can be used to represent different promo rules.

Examples promos as GO code:
```
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
```

Currently the project has been implemented with the following limitations:
The project has the following limitations. These limitations can be addressed in future

1. Promotion rules are specified within the code. All thouugh this limitation can be easily done away with, as the rules themselves can be easily represented using JSON
2. Ideally promotion rules should be managed by promotion service
    3.1 The changes in promotions can be communicated as events to cart service
3. Once promotions have been applied on cart, the resultant cart (also called as promofied cart) can be cached (unless an applied promotion changes or a new item is added, or deleted)
4. Currently all carts are held in memory. But this funcationaly is encapsulated behind a "repository" interface. A redis based repository needs to be implemented
5. No logging or units tests are implemented


