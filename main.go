package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_"github.com/go-sql-driver/mysql"
)




func main() {

	router := mux.NewRouter()

	router.HandleFunc("/clientes", GetClientesEndpoint).Methods("GET")
	//router.HandleFunc("/clientes/{id}", GetClienteEndpoint).Methods("GET")
	router.HandleFunc("/clientes", CreateClienteEndpoint).Methods("POST")
	//router.HandleFunc("/clientes/{id}", DeleteClienteEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4201", router))
}




/*
teste inclusão
{"tipopessoa":"fisica","cpfcnpj":"00168828065","nomecompleto":"Evandro Pezzi","nomefantasia":"Evandro","cep":"88048358","endereco":"srv. Augusto Buss","numero":"363","complemento":"","bairro":"Campeche","cidade":"Florianópolis","siglaestado":"SC"}
*/