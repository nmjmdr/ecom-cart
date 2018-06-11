package router
import "github.com/gorilla/mux"
import "net/http"
import "fmt"

func Start() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  go func() {
  if err := http.ListenAndServe(":8090", r); err != nil {
      fmt.Println(err)
    }
  }()
}


func HomeHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Hello")
}
