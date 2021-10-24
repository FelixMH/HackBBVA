package models

type Terminales struct {
	IdTerminal       string    `json:"idTerminal,omitempty"`
	CodPaquete       string    `json:"codPaquete,omitempty"`
	Estado           string    `json:"Estado,omitempty"`
	Marca            string    `json:"Marca,omitempty"`
	Modelo           string    `json:"Modelo,omitempty"`
	BancoProcedencia string    `json:"BancoProcedencia,omitempty"`
	IdCliente        *Clientes `json:"idCliente,omitempty"`
}
