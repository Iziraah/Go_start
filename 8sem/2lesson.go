package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type servise struct{
	store map[string] string
}

func main() {
	mux := http.NewServeMux()
	srv := servise{make(map[string]string)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAllUsers)


	http.ListenAndServe("localhost:8080", mux)
}

func (s *servise) Create(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}

		splittedContent := strings.Split(string(content), " ")
		s.store[splittedContent[0]] = string(content)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + splittedContent[0]))
		return 
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *servise) GetAllUsers(w http.ResponseWriter, r * http.Request) {

	if r.Method == "GET"{

		responce := ""
		for _, user := range s.store{
			responce += user + "\n"
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responce))
		return 

	}
	w.WriteHeader(http.StatusBadRequest)
}