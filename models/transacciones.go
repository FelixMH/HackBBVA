package models

type Transacciones struct {
	IdTransaccion     string  `json:"idTransaccion,omitempty"`
	EstadoTransaccion string  `json:"EstadoTransaccion,omitempty"`
	DatosTarjeta      string  `json:"DatosTarjeta,omitempty"`
	DatosComercio     string  `json:"DatosComercio,omitempty"`
	FechaTransaccion  string  `json:"FechaTransaccion,omitempty"`
	IdBanco           *Bancos `json:"idBanco,omitempty"`
}
