package models

type Bancos struct {
	IdBanco       string           `json:"idBanco,omitempty"`
	Nombre        string           `json:"Nombre,omitempty"`
	idSeguridad   *SeguridadStruct `json:"idSeguridad,omitempty"`
	idCliente     *Clientes        `json:"cliente,omitempty"`
	idTransaccion *Transacciones   `json:"idTransaccion,omitempty"`
}