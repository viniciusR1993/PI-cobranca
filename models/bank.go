package models

import (
	"fatec/db"
	"fmt"
	"log"
	"strconv"
)

//Struct criado para refletir a tabela do BD
type Bank struct {
	Id_bank int
	Name    string
	Cod     int
}

func AllBank() map[int]Bank {
	//Conecta com Postgres
	db := db.ConectBD()

	//Executa uma  query ono postgres
	selectAllBank, err := db.Query("select * from bank order by id_bank asc")
	if err != nil {
		panic(err.Error())
	}

	//Lê a query e armazena num slice
	b := Bank{}
	mapBank := make(map[int]Bank)
	for selectAllBank.Next() {
		var id_bank, cod, name string

		err := selectAllBank.Scan(&cod, &name, &id_bank)
		if err != nil {
			panic(err.Error())
		}

		id_bankConvertedForInt, err := strconv.Atoi(id_bank)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		codConvertedForInt, err := strconv.Atoi(cod)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		b.Id_bank = id_bankConvertedForInt
		b.Name = name
		b.Cod = codConvertedForInt

		mapBank[b.Id_bank] = b
	}
	defer db.Close()
	return mapBank
}

func NewBank(bank Bank) {
	db := db.ConectBD()

	insertDataBank, err := db.Prepare("insert into bank(name, cod) values($1, $2)")
	if err != nil {
		panic(err.Error())
	}

	insertDataBank.Exec(bank.Name, bank.Cod)
	defer db.Close()
}

func DeleteBank(bank Bank) int {
	db := db.ConectBD()

	deletBank, err := db.Prepare("delete from bank where id_bank=$1")
	if err != nil {
		return 1
	}

	_, err = deletBank.Exec(bank.Id_bank)
	if err != nil {
		return 1
	}

	defer db.Close()
	return 0
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func GetBank(id_bank int) Bank {
	db := db.ConectBD()

	banks, err := db.Query("select * from bank where id_bank=$1", id_bank)
	if err != nil {
		panic(err.Error())
	}

	bank := Bank{}
	for banks.Next() {
		var name, cod, id_bank string

		err = banks.Scan(&cod, &name, &id_bank)
		if err != nil {
			panic(err.Error())
		}

		id_bankConvertedForInt, err := strconv.Atoi(id_bank)
		if err != nil {
			log.Println("Conversion error: ", err)
		}
		codConvertedForInt, err := strconv.Atoi(cod)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		bank.Id_bank = id_bankConvertedForInt
		bank.Name = name
		bank.Cod = codConvertedForInt
	}
	defer db.Close()
	return bank
}

//Mudar a struct para os tupos de variaveis do banco e adaptar o método
func UpdateBank(bank Bank) int {
	db := db.ConectBD()

	UpdateBanks, err := db.Prepare("update bank set name=$1, cod=$2 where id_bank=$3")
	if err != nil {
		fmt.Println("Error in updating the bank table")
		return 1
	}

	UpdateBanks.Exec(bank.Name, bank.Cod, bank.Id_bank)
	defer db.Close()
	return 0
}
