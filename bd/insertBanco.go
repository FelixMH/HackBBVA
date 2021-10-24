package bd

import (
	"github.com/FelixMH/HackBBVA/config"
	"github.com/FelixMH/HackBBVA/models"
)

func InsertBancoBD(Banco models.Bancos) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}

	_, err = bd.Exec("INSERT INTO banco (idBanco, idCliente, idSeguridad, idTransaccion,Nombre) VALUES (?,?,?,?,?)", Banco.IdBanco, Banco.IdCliente, Banco.IdSeguridad, Banco.IdTransaccion, Banco.Nombre)
	return err
}