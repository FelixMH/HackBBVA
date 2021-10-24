package models

type SeguridadStruct struct {
	IdSeguridad   string
	IdBanco       *Bancos
	IdCliente     *Clientes
	IdTransaccion *Transacciones
	IdTerminal    *Terminales
	CodigoPaquete string
	Error         string
}
