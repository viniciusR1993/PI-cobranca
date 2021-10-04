package models

import (
	"fatec/db"
	"fmt"
	"log"
	"strconv"
)

//Struct criado para refletir a tabela do BD
type Transaction struct {
	Id_transaction int
	Maturity_date  string
	Issue_date     string
	Invoice        string
}

func AllTransaction() map[int]Transaction {
	//Conecta com Postgres
	db := db.ConectBD()

	//Executa uma  query ono postgres
	selectAllTransaction, err := db.Query("select * from transaction order by id_transaction asc")
	if err != nil {
		panic(err.Error())
	}

	//Lê a query e armazena num slice
	t := Transaction{}
	mapTransaction := make(map[int]Transaction)
	for selectAllTransaction.Next() {
		var id_transaction, maturity_date, invoice, issue_date string

		err := selectAllTransaction.Scan(&id_transaction, &maturity_date, &invoice, &issue_date)
		if err != nil {
			panic(err.Error())
		}

		id_transactionConvertedForInt, err := strconv.Atoi(id_transaction)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		t.Id_transaction = id_transactionConvertedForInt
		t.Maturity_date = maturity_date
		t.Invoice = invoice
		t.Issue_date = issue_date
		mapTransaction[t.Id_transaction] = t
	}
	defer db.Close()
	return mapTransaction
}

func NewTransaction(transaction Transaction) {
	db := db.ConectBD()

	insertDataTransaction, err := db.Prepare("insert into Transaction(maturity_date, invoice, issue_date) values($1, $2,$3)")
	if err != nil {
		panic(err.Error())
	}

	insertDataTransaction.Exec(transaction.Maturity_date, transaction.Invoice, transaction.Issue_date)
	defer db.Close()
}

func DeleteTransaction(transaction Transaction) int {
	db := db.ConectBD()

	deletTransaction, err := db.Prepare("delete from Transaction where id_transaction=$1")
	if err != nil {
		return 1
	}

	_, err = deletTransaction.Exec(transaction.Id_transaction)
	if err != nil {
		return 1
	}

	defer db.Close()
	return 0
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func GetTransaction(id_transaction int) Transaction {
	db := db.ConectBD()

	transactions, err := db.Query("select * from Transaction where id_transaction=$1", id_transaction)
	if err != nil {
		panic(err.Error())
	}

	transaction := Transaction{}
	for transactions.Next() {
		var id_transaction, maturity_date, invoice, issue_date string

		err = transactions.Scan(&id_transaction, &maturity_date, &invoice, &issue_date)
		if err != nil {
			panic(err.Error())
		}

		id_transactionConvertedForInt, err := strconv.Atoi(id_transaction)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		transaction.Id_transaction = id_transactionConvertedForInt
		transaction.Maturity_date = maturity_date
		transaction.Invoice = invoice
		transaction.Issue_date = issue_date
	}
	defer db.Close()
	return transaction
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func UpdateTransaction(transaction Transaction) int {
	db := db.ConectBD()

	UpdateBanks, err := db.Prepare("update Transaction set maturity_date = $1, invoice=$2, issue_date = $3 where id_transaction=$4")
	if err != nil {
		fmt.Println("Error in updating the bank table: ", err)
		return 1
	}

	UpdateBanks.Exec(transaction.Maturity_date, transaction.Invoice, transaction.Issue_date, transaction.Id_transaction)
	defer db.Close()
	return 0
}
