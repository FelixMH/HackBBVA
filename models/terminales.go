package models

type Terminales struct {
	IdTerminal       string
	CodPaquete       string
	Estado           string
	Marca            string
	Modelo           string
	BancoProcedencia string
	IdCliente        *Clientes
}
