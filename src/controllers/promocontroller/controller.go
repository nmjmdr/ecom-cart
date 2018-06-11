// Currently these are predefined applyPromos, but this service can easily be extended to provide
// the functionality:
// -> Add new promos, where the definition of a promo can be defined using JSON
// -> change values of promos
// Ideally This would be its own service, with changes in promotions communicated as events
// This architecture will be explained in greater detail in ReadMe

package promocontroller

import (
	"encoding/json"
	"net/http"
	"promocache"
	"utils"
)

func GetPromos(w http.ResponseWriter, r *http.Request) {
	promos := promocache.GetPromoCache().GetAll()
	json, _ := json.Marshal(promos)
	utils.SendResult(w, r, json)
}
