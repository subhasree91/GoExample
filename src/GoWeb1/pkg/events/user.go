package events

import (
	db "GoWeb1/database"
	"encoding/json"
	"fmt"
)

func (userEvents Event) Create() error {

	return db.Create(EventCollection, userEvents)

}

func (events Event) Update(update map[string]interface{}, find map[string]interface{}) error {
	fmt.Println("update:", update)
	err := db.Update(EventCollection, update, find)

	if err != nil {
		fmt.Println("Error in db.Update:", err)
	}
	return err

}

func (events Event) GetDoc(filter map[string]interface{}) (Event, error) {

	var output Event
	// s := db.GetSession()
	// c := s.DB("goapp").C(EventCollection)
	// err := c.Find(nil).All(&output)
	byteArr, err := db.GetDoc(EventCollection, filter)
	if err != nil {
		fmt.Println("Error in db.Get:", err)
		return output, err
	}
	err = json.Unmarshal(byteArr, &output)

	// if err != nil{
	// 	return nil, err
	// }

	return output, err

}
func (events Event) Get() ([]Event, error) {
	var output []Event
	// s := db.GetSession()
	// c := s.DB("goapp").C(EventCollection)
	// err := c.Find(nil).All(&output)
	byteArr, err := db.Get(EventCollection, nil)
	if err != nil {
		fmt.Println("Error in db.Get:", err)
		return nil, err
	}
	err = json.Unmarshal(byteArr, &output)

	// if err != nil{
	// 	return nil, err
	// }

	return output, err

}
