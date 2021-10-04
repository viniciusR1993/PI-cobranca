package models

import (
	"fatec/db"
	"fmt"
	"log"
	"strconv"
)

//Struct criado para refletir a tabela do BD
type Client struct {
	Id_client    int
	Name         string
	Doc          int
	Name_fantasy string
}

func AllClient() map[int]Client {
	//Conecta com Postgres
	db := db.ConectBD()

	//Executa uma  query ono postgres
	selectAllClient, err := db.Query("select * from client order by id_client asc")
	if err != nil {
		panic(err.Error())
	}

	//Lê a query e armazena num slice
	c := Client{}
	mapClient := make(map[int]Client)
	for selectAllClient.Next() {
		var id_client int
		var name, name_fantasy, doc string

		err := selectAllClient.Scan(&id_client, &name, &doc, &name_fantasy)
		if err != nil {
			panic(err.Error())
		}

		docConvertedForInt, err := strconv.Atoi(doc)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		c.Id_client = id_client
		c.Name = name
		c.Doc = docConvertedForInt
		c.Name_fantasy = name_fantasy

		mapClient[c.Id_client] = c
	}
	defer db.Close()
	return mapClient
}

func NewClient(client Client) int {
	db := db.ConectBD()

	insertDataBank, err := db.Prepare("insert into client(name, doc, fantasy_name) values($1, $2, $3)")
	if err != nil {
		log.Println("Erro in client table exemption: ", err)
		return 1
	}

	insertDataBank.Exec(client.Name, client.Doc, client.Name_fantasy)
	defer db.Close()
	return 0
}

func DeleteClient(id_client int) int {
	db := db.ConectBD()

	deletClient, err := db.Prepare("delete from client where id_client=$1")
	if err != nil {
		fmt.Println("Error deleting value from table client: ", err)
		return 1
	}

	deletClient.Exec(id_client)
	defer db.Close()
	return 0
}

func GetClient(id_client int) Client {
	db := db.ConectBD()

	clientBank, err := db.Query("select * from client where id_client=$1", id_client)
	if err != nil {
		panic(err.Error())
	}

	client := Client{}
	for clientBank.Next() {
		var id_client int
		var name, name_fantasy, doc string

		err = clientBank.Scan(&id_client, &name, &doc, &name_fantasy)
		if err != nil {
			panic(err.Error())
		}
		docConvertedForInt, err := strconv.Atoi(doc)
		if err != nil {
			log.Println("Conversion error: ", err)
		}

		client.Id_client = id_client
		client.Name = name
		client.Doc = docConvertedForInt
		client.Name_fantasy = name_fantasy
	}
	defer db.Close()
	return client
}

func UpdateClient(client Client) int {
	db := db.ConectBD()

	UpdateClient, err := db.Prepare("update client set name=$1, doc=$2, fantasy_name=$3 where id_client=$4")
	if err != nil {
		return 1
	}

	UpdateClient.Exec(client.Name, client.Doc, client.Name_fantasy, client.Id_client)
	defer db.Close()
	return 0
}
