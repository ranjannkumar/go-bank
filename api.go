package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


// Represents the API server.
// Stores the listening address (e.g., ":8080").
type APIServer struct {
	listenAddr string
	store      Storage
}


//Creates and returns a new APIServer instance with the given port.
func NewAPIServer(listenAddr string,store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run(){
	router := mux.NewRouter()
	//Registers the /account route, wrapping handleAccount inside makeHTTPHandleFunc for error handling.
	router.HandleFunc("/account",makeHTTPHandleFunc(s.handleAccount))
	router.HandleFunc("/account/{id}",makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API server running on port: ",s.listenAddr)
	http.ListenAndServe(s.listenAddr,router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter,r *http.Request)error{
	if r.Method=="GET"{
		return s.handleGetAccount(w,r)
	}
	if r.Method=="POST"{
		return s.handleCreateAccount(w,r)
	}
	if r.Method=="DELETE"{
		return s.handleDeleteAccount(w,r)
	}
	return fmt.Errorf("method not allowed %s",r.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter,r *http.Request)error{
	// account := NewAccount("Ranjan","Shah")
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	return WriteJSON(w,http.StatusOK,&Account{})
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter,r *http.Request)error{
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter,r *http.Request)error{
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter,r *http.Request)error{
	return nil
}


func WriteJSON(w http.ResponseWriter,status int, v any)error{
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

//This helps in structuring API handlers that return errors.
type apiFunc func (http.ResponseWriter,*http.Request)error


//Defines a structured JSON error response.
type ApiError struct{
	Error string
}


//Middleware for Error Handling
func makeHTTPHandleFunc(f apiFunc)http.HandlerFunc{
	return func (w http.ResponseWriter,r *http.Request)  {
		if err :=f(w,r);err!=nil{
			WriteJSON(w,http.StatusBadRequest,ApiError{Error: err.Error()})
		}
	}
}