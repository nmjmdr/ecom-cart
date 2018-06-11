package router

import "github.com/gorilla/mux"
import "net/http"
import "fmt"
import "controllers/cartcontroller"
import "controllers/promocontroller"

func Start(listenAddress string) {
	r := mux.NewRouter()
	// Ideally the routes would be defined in their own file: routes.js
	r.HandleFunc("/status", StatusHandler).Methods("GET")

	r.HandleFunc("/carts", cartcontroller.Create).Methods("POST")
	r.HandleFunc("/carts/{id}", cartcontroller.Get).Methods("GET")
	r.HandleFunc("/carts/{id}", cartcontroller.Delete).Methods("DELETE")

	// Add these later
	//r.HandleFunc("/carts/{id}/items", itemController.Create).Methods("POST")
	//r.HandleFunc("/carts/{cartId}/items/{itemId}", itemController.Delete).Methods("DELETE")

	// This would be a CPU intensive request
	// Could be move to its own service, if it needs to be scaled
	// It can be scaled easily as it does not hold state
	r.HandleFunc("/carts/{cartId}/promofied", cartcontroller.ApplyPromos).Methods("POST")

	// Ideally the below controller would its own service
	r.HandleFunc("/promos", promocontroller.GetPromos).Methods("GET")

	go func() {
		if err := http.ListenAndServe(listenAddress, r); err != nil {
			fmt.Println(err)
		} else {
      fmt.Println("Listening on port: 8090")
    }
	}()
}

// Can be enhanced to collect and report status of various parameters
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
