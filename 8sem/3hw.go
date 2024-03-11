/*
Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя и удаляет его из хранилища, а также стирает его из массива friends у всех его друзей. Пример запроса:
DELETE /user HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"target_id":"1"}
Данный запрос должен возвращать 200 и имя удалённого пользователя.
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
 	mux.HandleFunc("/delete_user", srv.DeleteUser)


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

func (s *service) DeleteUser(w http.ResponseWriter, r *http.Request) {
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

 		targetUser, ok := s.store[fr.TargetID]
  
  		if !ok {
   			w.WriteHeader(http.StatusBadRequest)
   			w.Write([]byte("User doesn't exist in the store"))
   			return
  			}
  
  		// Remove from friends list
  		for friendID, _ := range targetUser.Friends {
  		 	delete(s.store[friendID].Friends, fr.TargetID)
		}

  		delete(s.store, fr.TargetID)

  		w.WriteHeader(http.StatusOK)
  		w.Write([]byte("User " + targetUser.Name + " was deleted"))
  		return 
 		}

 	w.WriteHeader(http.StatusBadRequest)
}