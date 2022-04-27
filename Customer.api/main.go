package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)


type Customer struct {
	ID				string 	`json:"id"`
	PhoneNumber 	string  `json:"phonenumber"`
	MailId 			string  `json:"mailid"`
	Name			*Name	`json:"name"`
}


type Name struct {
	Firstname 	string  `json:"firstname"`
	Lastname 	string  `json:"lastname"`
}


var customers []Customer


func getCustomers(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(customers)
	
}


func getCustomer(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item:= range customers {
		if item.ID==params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Customer{})
}

func createCustomer(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var customer Customer
	_ = json.NewDecoder(r.Body).Decode(&customer)
	customer.ID = strconv.Itoa(rand.Intn(1000000)) 
	customers = append(customers, customer)
	json.NewEncoder(w).Encode(customer)
}

func updateCustomer(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range customers {
		if item.ID == params["id"] {
		customers = append(customers[:index],customers[index+1:]...)
		var customer Customer
		_ = json.NewDecoder(r.Body).Decode(&customer)
		customer.ID = params["id"]
		customers = append(customers, customer)
		json.NewEncoder(w).Encode(customer)
		return
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func deleteCustomer(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range customers {
		if item.ID == params["id"] {
		customers = append(customers[:index],customers[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(customers)
}

func main(){
	
	r := mux.NewRouter()

	
	customers = append(customers, Customer{ID: "1", PhoneNumber:"1234567890", MailId: "viratkohli@gmail.com", Name: &Name {Firstname: "Virat", Lastname:"Kohli"}})
	customers = append(customers, Customer{ID: "2", PhoneNumber:"0987654321", MailId: "tejakasula@gamil.com", Name: &Name {Firstname: "Teja", Lastname:"Kasula"}})
	r.HandleFunc("/api/customers", getCustomer).Methods("GET")
	r.HandleFunc("/api/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/api/customers", createCustomer).Methods("POST")
	r.HandleFunc("/api/customers/{id}", updateCustomer).Methods("PUT")
	r.HandleFunc("/api/customers/{id}", deleteCustomer).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", r))
}