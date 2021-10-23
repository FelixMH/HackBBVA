package models

type Terminales struct {
	IdTerminal       string
	CodPaquete       string
	Estado           string
	Marca            string
	Modelo           string
	BancoProcedencia string
	idCliente        *Clientes
}