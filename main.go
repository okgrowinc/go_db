package main

import(
"fmt"
"os"
"encoding/json"
)

type Address struct{
	City string
	State string
	Country string
	Pincode json.Number
}

type User struct{
	Name string
	Age json.Number
	Contact string
	Company string
	Address Address
}

func main(){
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
				}
}

	employees := []User{
		{"John", "55", "23344333", "OS Intelligent", Address{"Overland Park", "Kansas", "United States of America", "410013"}},
		{"Jodi", "53", "23344333", "OS Intelligent", Address{"Pratt", "Kansas", "United States of America", "67120"}},
}

for _, value := range employees{
	db.Write("users", value.Name, User{

	})
}