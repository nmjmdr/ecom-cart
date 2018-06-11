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
```


Currently the project has been implemented with the following limitations:
The project has the following limitations. These limitations can be addressed in future

1. Promotion rules are specified within the code. All thouugh this limitation can be easily done away with, as the rules themselves can be easily represented using JSON
2. Ideally promotion rules should be managed by promotion service
    3.1 The changes in promotions can be communicated as events to cart service
3. Once promotions have been applied on cart, the resultant cart (also called as promofied cart) can be cached (unless an applied promotion changes or a new item is added, or deleted)
4. Currently all carts are held in memory. But this funcationaly is encapsulated behind a "repository" interface. A redis based repository needs to be implemented
5. No logging or units tests are implemented


