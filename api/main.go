package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

type errorResp struct {
	Statuscode int
	Message    string
}

var users = map[string]User{}

func main() {
	http.HandleFunc("/createuser", adduser)
	http.HandleFunc("/user", getUsers)

	fmt.Println("users are :", users)
     fmt.Println("server started ")
    log.Fatalf("server not started :%v", http.ListenAndServe(":8080", nil))

}

// create /add new user record in map

func adduser(f http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		f.WriteHeader(http.StatusMethodNotAllowed)
		return
}
user :=User{}
err :=json.NewDecoder(r.Body).Decode(&user)
if err!=nil{
		f.WriteHeader(http.StatusBadRequest)

		errr:=errorResp{
			Statuscode: http.StatusBadRequest,
			Message:    "payload is not right ",
		}
		 json.NewEncoder(f).Encode(errr)
		return
}
users[user.Name]=user                                                        //*ask
               f.WriteHeader(http.StatusCreated)

			   json.NewEncoder(f).Encode(user)
			   fmt.Println("users are :", users)

return

}





// return user  record from map

func getUsers(f http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		f.WriteHeader(http.StatusMethodNotAllowed)
		return

	}

	// jsonenc := json.NewEncoder(f)
	// err := jsonenc.Encode(users)
	err := json.NewEncoder(f).Encode(users)
	if err != nil {
		f.WriteHeader(http.StatusBadRequest)
		err := errorResp{

			Statuscode: http.StatusBadRequest,
			Message:    "payload ban nhi paya ,err =" + err.Error(),
		}
		json.NewEncoder(f).Encode(err)

		return
	}

	f.WriteHeader(http.StatusOK)

	json.NewEncoder(f).Encode(users)
	
	
	return

}
