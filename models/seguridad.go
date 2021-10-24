package models

type SeguridadStruct struct {
	IdSeguridad   string         `json:"idSeguridad,omitempty"`
	IdBanco       *Bancos        `json:"idBanco,omitempty"`
	IdCliente     *Clientes      `json:"idCliente,omitempty"`
	IdTransaccion *Transacciones `json:"idTransaccion,omitempty"`
	IdTerminal    *Terminales    `json:"idTerminal,omitempty"`
	CodigoPaquete string         `json:"CodigoPaquete,omitempty"`
	Error         string         `json:"Error,omitempty"`
}
