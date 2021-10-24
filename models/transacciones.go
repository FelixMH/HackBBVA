package models

type Transacciones struct {
	IdTransaccion     string
	EstadoTransaccion string
	DatosTarjeta      string
	DatosComercio     string
	FechaTransaccion  string
	IdBanco           *Bancos
}
