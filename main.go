package main

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"sync"
)

const Version = "1.0.1"

type (
	Logger interface {
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}

	Driver struct {
		mutex   sync.Mutex
		mutexes map[string]*sync.Mutex
		dir     string
		log     Logger
	}
)

type Options struct {
	Logger
}

func New(dir string, options *Options) (*Driver, error) {
	dir = filepath.Clean(dir)

	opts := Options{}
}

func (d *Driver) Write() error {

}

func (d *Driver) Read() error {

}

func (d *Driver) ReadAll() {

}

func (d *Driver) Delete() error {

}

func (d *Driver) getOrCreateMutex() *sync.Mutex {

}

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"

	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "55", "23344333", "OS Intelligent", Address{"Overland Park", "Kansas", "United States of America", "410013"}},
		{"Jodi", "53", "23344333", "OS Intelligent", Address{"Pratt", "Kansas", "United States of America", "67120"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, User{
			Name:    value.Name,
			Age:     value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allusers := []User{}

	for _, f := range records {
		employeeFound := User{}
		if err := json.Unmarshal([]byte(f), &employeeFound); err != nil {
			fmt.Println("Error", err)
		}
		allusers = append(allusers, employeeFound)
	}
}

//fmt.Println((allusers))

//if err := db.Delete("user", "john")1 err != nil{
//	fmt.Println("Error, err")
//}

//if err := db.Delete("user", ""); err != nil{
//	fmt.Println(Erroe, err)
//}
