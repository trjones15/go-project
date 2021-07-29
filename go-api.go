package main

//import (
//	"log"
//	"net/http"
//)

//type server struct{}
//
//func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	switch r.Method {
//	case "GET":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "get called"}`))
//	case "POST":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "post called"}`))
//	case "PUT":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "put called"}`))
//	case "DELETE":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "delete called"}`))
//	default:
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "not found"}`))
//	}
//}
//
//func main() {
//	s := &server{}
//	http.Handle("/", s)
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

//added functionality without creating the struct
//func home(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	switch r.Method {
//	case "GET":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "get called"}`))
//	case "POST":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "post called"}`))
//	case "PUT":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "put called"}`))
//	case "DELETE":
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "delete called"}`))
//	default:
//		w.WriteHeader(http.StatusOK)
//		w.Write([]byte(`{"message": "not found"}`))
//	}
//}
//
//func main() {
//	http.HandleFunc("/", home)
//	log.Fatal(http.ListenAndServe(":8080", nil))
//}

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", r))
}
