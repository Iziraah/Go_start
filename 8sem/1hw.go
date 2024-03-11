/*
Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля: имя, возраст и массив друзей. Пользователя необходимо сохранять в мапу. Пример запроса:
POST /create HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"name":"some name","age":"24","friends":[]}
**/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
    Name string 				`json:"name"`
    Age  int    				`json:"age"`
    Friends map[string]*User 	`json:"friends"`
}

func (u*User) toString()string{
	return fmt.Sprintf("name is %s, age is %d and has %d friend(s) n ", u.Name, u.Age, len(u.Friends))
}

type service struct{
	store map[string]*User
}

func main() {
	mux := http.NewServeMux()
	srv := service{make(map[string]*User)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)


	http.ListenAndServe("localhost:8080", mux)
}

func (s *service) Create(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var u *User
		if  err = json.Unmarshal(content, &u); err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		for friendName := range u.Friends {
    	if _, ok := s.store[friendName]; !ok {
        	w.WriteHeader(http.StatusBadRequest)
        	w.Write([]byte("Friend " + friendName + " doesn't exist in the store"))
        	return
    }
  }

		s.store[u.Name] = u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created " + u.Name))
		return 
	}

	w.WriteHeader(http.StatusBadRequest)
}

func (s *service) GetAll(w http.ResponseWriter, r * http.Request) {

	if r.Method == "GET"{

		responce := ""
		for _, user := range s.store{
			responce += user.toString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(responce))
		return 

	}
	w.WriteHeader(http.StatusBadRequest)
}