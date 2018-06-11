# ecom-cart

ecom-cart implements a promotion engine which can compute promotion rules such as:
1. if you buy 2 or more trousers, you get 15% off belts and shoes
2. if you buy two shirts, each additional shirt costs only 45$
3. if you purchase 3 or more shirts, all ties are half price
4. buy 2 shirts and get 1 shirt free

Notice that fourth rule (buy 2 shirts and get 1 shirt free) has to be applied in such way if the customer has purchased 7 shirts then two of them should be free.


## Design Details
Given time the project would follow the below design approach:

 

## Limitations:
The project has the following limitations. These limitations can be addressed in future

1. Currently does not implement any unit tests 
2. Promotion rules are specified within the code. All thouugh this limitation can be easily done away with, as the rules themselves can be easily represented using JSON
3. Ideally promotion rules should be managed by promotion service
    3.1 The changes in promotions can be communicated as events to cart service
4. Once promotions have been applied on cart, the resultant cart (also called as promofied cart) can be cached (unless an applied promotion changes or a new item is added, or deleted)
5. Currently all carts are held in memory. But this funcationaly is encapsulated behind a "repository" interface. A redis based repository needs to be implemented
