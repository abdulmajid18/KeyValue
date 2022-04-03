package main

import (
	"KeyValueStore/other/helper"
	"fmt"
)

func main() {

	path := "/home/rozz/Desktop/mydb/key.db"

	db, err := helper.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	err = db.Put("Name", "Abdul-Majid")
	if err != nil {
		fmt.Println(err)
	}

	value, state, err := db.Get("Name")
	if err != nil {
		fmt.Println(err)
	}
	if state {
		fmt.Printf("The value is %s ", value)
	}

}
