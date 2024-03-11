/*
4. Сделайте обработчик, который возвращает всех друзей пользователя. Пример запроса:
GET /friends/user_id HTTP/1.1
Host: localhost:8080
Connection: close
После /friends/ указывается id пользователя, друзей которого мы хотим увидеть.
5. Сделайте обработчик, который обновляет возраст пользователя. Пример запроса:
PUT /user_id HTTP/1.1
Content-Type: application/json; charset=utf-8
Host: localhost:8080
{"new age":"28"}
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

type UpdateAgeRequest struct {
    UserID string 				`json:"user_id"`
    NewAge int    				`json:"new_age"`
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
	mux.HandleFunc("/get_friends", srv.GetFriends)
	mux.HandleFunc("/update_age", srv.UpdateAge)

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

func (s *service) GetFriends(w http.ResponseWriter, r *http.Request) {
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

  		user, ok := s.store[fr.SourceID]
  
  		if !ok {
   			w.WriteHeader(http.StatusBadRequest)
   			w.Write([]byte("User doesn't exist in the store"))
   			return
  			}
  
  		response := ""
  		for _, friend := range user.Friends{
   			response += friend.Name + "\n"
  			}

  		w.WriteHeader(http.StatusOK)
  		w.Write([]byte(response))
  		return 
 		}

 		w.WriteHeader(http.StatusBadRequest)
}

func (s *service) UpdateAge(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
  		content, err := ioutil.ReadAll(r.Body)
  		if err != nil {
   			w.WriteHeader(http.StatusInternalServerError)
   			w.Write([]byte(err.Error()))
   			return
  			}
  		defer r.Body.Close()

  		var updateReq *UpdateAgeRequest
  		if  err = json.Unmarshal(content, &updateReq); err != nil{
   			w.WriteHeader(http.StatusInternalServerError)
   			w.Write([]byte(err.Error()))
   			return
  			}

  		user, ok := s.store[updateReq.UserID]
  
  		if !ok {
   			w.WriteHeader(http.StatusBadRequest)
   			w.Write([]byte("User doesn't exist in the store"))
   			return
  			}
  
  		user.Age = updateReq.NewAge

  		w.WriteHeader(http.StatusOK)
  		w.Write([]byte("User " + user.Name + "'s age was updated to " + fmt.Sprint(updateReq.NewAge)))
  		return 
	}

 	w.WriteHeader(http.StatusBadRequest)
}