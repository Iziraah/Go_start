/*
Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали двух пользователей и нам вернулись их ID, то в запросе мы можем указать ID пользователя, который инициировал запрос на дружбу, и ID пользователя, который примет инициатора в друзья. Пример запроса:
POST /make_friends HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"source_id":"1","target_id":"2"}
Данный запрос должен возвращать статус 200 и сообщение «username_1 и username_2 теперь друзья».
**/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct{
	ID string 					`json:"id"`
	Name string 				`json:"name"`
	Age  int    				`json:"age"`
	Friends map[string]*User 	`json:"friends"`
}

type FriendRequest struct {
 	SourceID string 			`json:"source_id"`
 	TargetID string 			`json:"target_id"`
}

func (u*User) toString()string{
	return fmt.Sprintf("id is %s, name is %s, age is %d and has %d friend(s) n ", u.ID, u.Name, u.Age, len(u.Friends))
}

type service struct{
	store map[string]*User
}

func main() {
	mux := http.NewServeMux()
	srv := service{make(map[string]*User)}
	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)
	mux.HandleFunc("/make_friends", srv.MakeFriends)


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

func (s *service) MakeFriends(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
  		content, err := ioutil.ReadAll(r.Body)
  		if err != nil {
   			w.WriteHeader(http.StatusInternalServerError)
   			w.Write([]byte(err.Error()))
   			return
  		}
  		defer r.Body.Close()

  	var fr *FriendRequest
 	if  err = json.Unmarshal(content, &fr); err != nil{
   		w.WriteHeader(http.StatusInternalServerError)
   		w.Write([]byte(err.Error()))
   		return
  		}

  	sourceUser, ok1 := s.store[fr.SourceID]
  	targetUser, ok2 := s.store[fr.TargetID]
  
  	if !ok1 || !ok2 {
   		w.WriteHeader(http.StatusBadRequest)
   		w.Write([]byte("Either source or target user doesn't exist in the store"))
   		return
  	}

 	sourceUser.Friends[targetUser.ID] = targetUser
  	targetUser.Friends[sourceUser.ID] = sourceUser

  	w.WriteHeader(http.StatusOK)
  	w.Write([]byte(sourceUser.Name + " and " + targetUser.Name + " are now friends"))
  	return 
 	}

 	w.WriteHeader(http.StatusBadRequest)
}