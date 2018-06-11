# ecom-cart

ecom-cart implements a promotion engine which can compute promotion rules such as:
1. if you buy 2 or more trousers, you get 15% off belts and shoes
2. if you buy two shirts, each additional shirt costs only 45$
3. if you purchase 3 or more shirts, all ties are half price
4. buy 2 shirts and get 1 shirt free

Notice that fourth rule (buy 2 shirts and get 1 shirt free) has to be applied in such way if the customer has purchased 7 shirts then two of them should be free.

## Installation and Usage
This project follows a slightly different structure from the recommended GOLang guidelines. (This can be corrected later) The packages are contained within the project folder under `src` folder
The project depends upon:
1. "github.com/gorilla/mux"
2. "github.com/satori/go.uuid"
To install the dependencies run the following commands in project folder:
```
go get -u github.com/satori/go.uuid
go get -u github.com/gorilla/mux"
```
To run the project:
```
go run main.go
```
### Usage:
1. Create a cart
Ex: cart with - two trousers, three shirts, one belt, one shoe
```
curl localhost:8090/carts -X POST -d '{ "items": [{"name": "t1", "category": "trousers", "price": 100}, {"name": "t2", "category": "trousers", "price": 110}, {"name": "belt-1", "category": "belts", "price": 100}, {"name": "shoe-1", "category": "shoes", "price": 100}, {"name": "s1", "category": "shirts", "price": 120}, {"name": "s2", "category": "shirts", "price": 90}, {"name": "s3", "category": "shirts", "price": 65} ] }'

Response:
{"message":"Created","id":"06f9af8a-d312-41a7-9576-f097958cc61f","_links":[{"instance":"/carts/06f9af8a-d312-41a7-9576-f097958cc61f"}]}
```
The cart is now created and assigned a unique id

2. Get a cart
```
curl localhost:8090/carts/06f9af8a-d312-41a7-9576-f097958cc61f

Response:
{"id":"06f9af8a-d312-41a7-9576-f097958cc61f",
"items":[
{"id":"","name":"t1","category":"trousers","price":100},
{"id":"","name":"t2","category":"trousers","price":110},
{"id":"","name":"belt-1","category":"belts","price":100},
{"id":"","name":"shoe-1","category":"shoes","price":100},
{"id":"","name":"s1","category":"shirts","price":120},
{"id":"","name":"s2","category":"shirts","price":90},
{"id":"","name":"s3","category":"shirts","price":65}
]
}
```
3. Get available promos:
```
curl localhost:8090/promos

Response:

[
  {
    "id": "promo2",
    "description": "if you buy two shirts, each additional shirt costs only 45$",
    "buys": [ { "category": "shirts", "count": 2 } ],
    "gets": [ { "category": "shirts", "all": true, "count": 0, "off": { "discount": null, "fixed": { "price": 45 } } ]
  },
  {
    "id": "promo3",
    "description": "if you purchase 3 or more shirts, all ties are half price",
  ... 
```
4. Apply promos
Note: Applies the following promos:
1. promo1 = if you buy 2 or more trousers, you get 15% off belts and shoes 
2. promo2 = if you buy two shirts, each additional shirt costs only 45$
```
curl localhost:8090/carts/06f9af8a-d312-41a7-9576-f097958cc61f/promofied -X POST -d '{ "promos": ["promo1","promo2"] }'

Reponse:
{
"totalPrice":685,
"totalOffPrice":635,

"items":[
{"Item":{"id":"","name":"t1","category":"trousers","price":100},"sourceForPromos":{"promo1":true},"targetForPromos":{}},
{"Item":{"id":"","name":"t2","category":"trousers","price":110},"sourceForPromos":{"promo1":true},"targetForPromos":{}},
{"Item":{"id":"","name":"belt-1","category":"belts","price":100},"sourceForPromos":{},"targetForPromos":{"promo1":85}},
{"Item":{"id":"","name":"shoe-1","category":"shoes","price":100},"sourceForPromos":{},"targetForPromos":{"promo1":85}},
{"Item":{"id":"","name":"s1","category":"shirts","price":120},"sourceForPromos":{"promo2":true},"targetForPromos":{}},
{"Item":{"id":"","name":"s2","category":"shirts","price":90},"sourceForPromos":{"promo2":true},"targetForPromos":{}},
{"Item":{"id":"","name":"s3","category":"shirts","price":65},"sourceForPromos":{},"targetForPromos":{"promo2":45}}
]
}
```
The property `sourceForPromos` indicactes the list of promos for which the item is a source or a trigger.
he property `targetForPromos` indicactes the list of promos that have been applied for the item.

_If more than one promo applies to an item, the off-price (or discounted price) is taken to be the minimum of all promo applications_
So if for an item with sale price of 100, two promos apply, one with price = 95$ and other with prie = 85$. Then 85$ is considered as the `off price`.

5. Delete a cart
```
curl -X DELETE localhost:8090/carts/06f9af8a-d312-41a7-9576-f097958cc61f
Returns: 200 OK (probably ideally should 204 - No content)
```

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


