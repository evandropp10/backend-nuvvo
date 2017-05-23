package main

import (
	"database/sql"
	"net/http"
	"encoding/json"
	_"github.com/go-sql-driver/mysql"
)

type Cliente struct {
	ID int `json:"id,omitempty"`
	TipoPessoa string `json:"tipoPessoa,omitempty"`
	CpfCnpj string `json:"cpfCnpj,omitempty"`
	NomeCompleto string `json:"nomeCompleto,omitempty"`
	NomeFantasia string `json:"nomeFantasia,omitempty"`
	Cep string `json:"cep,omitempty"`
	Endereco string `json:"endereco,omitempty"`
	Numero int `json:numero,omitempty`
	Complemento string `json:"complemento,omitempty"`
	Bairro string `json:"bairro,omitempty"`
	Cidade string `json:"cidade,omitempty"`
	SiglaEstado string `json:"estado,omitempty"`
}


//var clientes []Cliente

//var db, err = sql.Open("mysql", "root:setembro2016@tcp(localhost:3306)/nuvvo?charset=utf8")

/*
func GetClienteEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item := range clientes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Cliente{})
}
*/
func GetClientesEndpoint(w http.ResponseWriter, req *http.Request){
	var clientes = ListaClientesDB()

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(clientes)
}

func CreateClienteEndpoint(w http.ResponseWriter, req *http.Request){
	var cliente Cliente
	_ = json.NewDecoder(req.Body).Decode(&cliente)
	IncluirClienteDB(cliente)
	json.NewEncoder(w).Encode(cliente)
}

/*
func DeleteClienteEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, item := range clientes {
		if item.ID == params["id"] {
			clientes = append(clientes[:index], clientes[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(clientes)
}
*/

func IncluirClienteDB(cliente Cliente){

	var db, err = sql.Open("mysql", "root:setembro2016@tcp(localhost:3306)/nuvvo?charset=utf8")
	checkError(err)

	stmt, err := db.Prepare("Insert into cliente(tipoPessoa, cpfCnpj, nomeCompleto, nomeFantasia, cep, endereco, numero, complemento, bairro, cidade, siglaEstado) values(?,?,?,?,?,?,?,?,?,?,?)")
	checkError(err)

	_, err = stmt.Exec(cliente.TipoPessoa,cliente.CpfCnpj ,cliente.NomeCompleto, cliente.NomeFantasia, cliente.Cep, cliente.Endereco, cliente.Numero, cliente.Complemento, cliente.Bairro, cliente.Cidade, cliente.SiglaEstado)
	checkError(err)

	db.Close()

}

func ListaClientesDB() []Cliente {

	clientesLista := []Cliente{}

	var db, err = sql.Open("mysql", "root:setembro2016@tcp(localhost:3306)/nuvvo?charset=utf8")
	checkError(err)

	rows, err := db.Query("SELECT * FROM cliente")
	checkError(err)

	for rows.Next() {
		listacliente := Cliente{}
		
		err = rows.Scan(&listacliente.ID, &listacliente.TipoPessoa, &listacliente.CpfCnpj, &listacliente.NomeCompleto, &listacliente.NomeFantasia, &listacliente.Cep, &listacliente.Endereco, &listacliente.Numero, &listacliente.Complemento, &listacliente.Bairro, &listacliente.Cidade, &listacliente.SiglaEstado)
		
		checkError(err)
		

		clientesLista = append(clientesLista, listacliente)
	}

	db.Close()

	return clientesLista
}


func checkError (err error) {
	if err != nil {
		panic(err)
	}
}