package models

type SeguridadStruct struct {
	IdSeguridad   string
	idBanco       *Bancos
	idCliente     *Clientes
	idTransaccion *Transacciones
	idTerminal    *Terminales
	codigoPaquete string
	error         string
}