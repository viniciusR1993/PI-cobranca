package controllers

import (
	"encoding/json"
	"fatec/models"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AllTransaction(w http.ResponseWriter, r *http.Request) {
	transactions := models.AllTransaction()
	json.NewEncoder(w).Encode(transactions)
}

func GetTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transaction

	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &transactions); err != nil {
		json.NewEncoder(w).Encode(1)
	}

	ret := models.GetTransaction(transactions.Id_transaction)

	json.NewEncoder(w).Encode(ret)
}

func UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transaction
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &transactions); err != nil {
		fmt.Print(err)
	}

	ret := models.UpdateTransaction(transactions)
	json.NewEncoder(w).Encode(ret)
}

func InsertTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transaction
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &transactions); err != nil {
		fmt.Print(err)
	}

	models.NewTransaction(transactions)
	json.NewEncoder(w).Encode(0)
}

func DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	var transactions models.Transaction
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &transactions); err != nil {
		fmt.Print(err)
	}

	models.DeleteTransaction(transactions)
	json.NewEncoder(w).Encode(0)
}
