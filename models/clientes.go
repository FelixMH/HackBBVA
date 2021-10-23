package models

type Clientes struct {
	IdCliente     string           `json:"idCliente,omitempty"`
	IdBanco       *Bancos          `json:"idBanco,omitempty"`
	IdTerminal    *Terminales      `json:"idTerminal,omitempty"`
	IdTransaccion *Transacciones   `json:"idTransaccion,omitempty"`
	IdSeguridad   *SeguridadStruct `json:"idSeguridad,omitempty"`
}