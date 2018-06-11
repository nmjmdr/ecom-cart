package router

import "github.com/gorilla/mux"
import "net/http"
import "fmt"
import "cartcontroller"

func Start() {
	r := mux.NewRouter()
  // Ideally the routes would be defined in their own file: routes.js
	r.HandleFunc("/status", StatusHandler).Methods("GET")

	r.HandleFunc("/carts", cartcontroller.Create).Methods("POST")
  r.HandleFunc("/carts/{id}", cartcontroller.Get).Methods("GET")
  r.HandleFunc("/carts/{id}", cartcontroller.Delete).Methods("DELETE")

  // Add these later
  //r.HandleFunc("/carts/{id}/items", itemController.Create).Methods("POST")
  //r.HandleFund("/carts/{cartId}/items/{itemId}", itemController.Delete).Methods("DELETE")

	go func() {
		if err := http.ListenAndServe(":8090", r); err != nil {
			fmt.Println(err)
		}
	}()
}

// Can be enhanced to collect and report status of various parameters
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}
