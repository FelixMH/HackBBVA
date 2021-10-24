package models

type Bancos struct {
	IdBanco       string           `json:"idBanco,omitempty"`
	Nombre        string           `json:"Nombre,omitempty"`
	IdSeguridad   *SeguridadStruct `json:"idSeguridad,omitempty"`
	IdCliente     *Clientes        `json:"cliente,omitempty"`
	IdTransaccion *Transacciones   `json:"idTransaccion,omitempty"`
}
