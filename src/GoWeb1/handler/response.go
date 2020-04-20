package handler

import (
	redis "GoWeb1/cache"
	event "GoWeb1/pkg/events"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeLink(w http.ResponseWriter, r *http.Request) {
	Path := html.EscapeString(r.URL.Path)
	fmt.Fprintf(w, "Welcome home!")
	fmt.Println(Path)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createevent")
	user := event.Event{}
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error occured while reading the request body :", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error occured while reading the request body " + err.Error())
		return
	}
	err = user.Create()
	if err != nil {
		fmt.Println("Error occured while reading the create function :", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error occured while creating the request " + err.Error())
		return
	}
	redis.RedisSet(user.ID+user.Name, user)
	json.NewEncoder(w).Encode("Created event ")
}

func GetEvent(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)

	filter := make(map[string]interface{})
	for k, v := range param {
		filter[k] = v
	}

	user := event.Event{}
	temp := redis.RedisGet(param["id"] + param["name"])
	if temp != nil {
		fmt.Println("redis cache")
		json.NewEncoder(w).Encode(temp)
	}

	userDoc, err := user.GetDoc(filter)

	if err != nil {
		fmt.Println("Error occured while reading the get function :", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error occured while get request " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(userDoc)

}

func GetEvents(w http.ResponseWriter, r *http.Request) {

	user := event.Event{}
	arr, err := user.Get()

	if err != nil {
		fmt.Println("Error occured while reading the get function :", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error occured while get request " + err.Error())
		return
	}
	json.NewEncoder(w).Encode(arr)

}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {

	user := event.Event{}
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error occured while reading the request body " + err.Error())
		return
	}

	err = json.Unmarshal(body, &user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Error occured while unmarshal" + err.Error())
		return
	}

	update := make(map[string]interface{})
	err = json.Unmarshal(body, &update)

	find := make(map[string]interface{})

	query := r.URL.Query()

	//valArr := query["username"]

	for key, value := range query {

		if len(value) != 0 {
			find[key] = value[0]
		}

	}

	err = user.Update(update, find)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error occured while updating query" + err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Update successful")

}
