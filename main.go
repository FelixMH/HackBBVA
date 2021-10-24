package main

import (
	"github.com/FelixMH/HackBBVA/config"
	"github.com/FelixMH/HackBBVA/handler"
	"github.com/FelixMH/HackBBVA/models"
	_ "github.com/go-sql-driver/mysql"
)

// var _ = godotenv.Load(".env") // carga el archivo .env que contiene las credenciales
// var ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", //define la conexion a la bd
// 		os.Getenv("user"),
// 		os.Getenv("pass"),
// 		os.Getenv("host"),
// 		os.Getenv("port"),
// 		os.Getenv("db_name"))


const AllowedCORSDomain = "http://localhost" //AHORA SI PUSE LA INICIAL DE LA VARIABLE EN MAYUSCULA AAAAAAAAAAAAA 👺

// func getDB() (*sql.DB, error) { //conecta mysql con go
// 	return sql.Open("mysql", ConnectionString)
// }

//CRUD de terminales 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢
func InsertTerminalBD(Terminal models.Terminales) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO terminal (idTerminal, CodPaquete, Estado, Marca,Modelo,BancoProcedencia,idCliente) VALUES (?,?,?,?,?,?,?)", Terminal.IdTerminal, Terminal.CodPaquete, Terminal.Estado, Terminal.Marca, Terminal.Modelo, Terminal.BancoProcedencia, Terminal.IdCliente)
	return err
}

func DeleteTerminalBD(Id string) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM terminal WHERE idTerminal = ?", Id)
	return err
}

func UpdateTerminalBD(Terminal models.Terminales) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE terminal set CodPaquete = ?, Estado = ?,Marca = ?,Modelo = ?,BancoProcedencia = ?, idCliente = ? WHERE idTerminal = ?", Terminal.CodPaquete, Terminal.Estado, Terminal.Marca, Terminal.Modelo, Terminal.BancoProcedencia, Terminal.IdCliente, Terminal.IdTerminal)
	return err
}

func SelectTerminales() ([]models.Terminales, error) { // trae todas las terminales
	//se declara un arreglo , si hay un error lo regresa vacio
	terminales := []models.Terminales{}
	bd, err := config.GetDB()
	if err != nil {
		return terminales, err
	}
	//agarra las filas
	rows, err := bd.Query("SELECT idTerminal,CodPaquete,Estado,Marca,Modelo,BancoProcedencia,idCliente FROM terminal")
	if err != nil {
		return terminales, err
	}
	//crea los renglones
	for rows.Next() {
		//cada paso es una linea
		var terminal models.Terminales
		err = rows.Scan(&terminal.IdTerminal, &terminal.CodPaquete, &terminal.Estado, &terminal.Marca, &terminal.Modelo, &terminal.BancoProcedencia, &terminal.IdCliente)
		if err != nil {
			return terminales, err
		}
		terminales = append(terminales, terminal)
	}
	return terminales, nil
}

func SelectTerminalIdDB(id string) (models.Terminales, error) { //solo trae una terminal por id
	var terminal models.Terminales
	bd, err := config.GetDB()
	if err != nil {
		return terminal, err
	}
	row := bd.QueryRow("SELECT * FROM terminal where idTerminal = ?", id)
	err = row.Scan(&terminal.IdTerminal, &terminal.CodPaquete, &terminal.Estado, &terminal.Marca, &terminal.Modelo, &terminal.BancoProcedencia, &terminal.IdCliente)
	if err != nil {
		return terminal, err
	}
	//si todo sale bien
	return terminal, nil
}

// crud banco 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢

func InsertBancoBD(Banco models.Bancos) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}

	_, err = bd.Exec("INSERT INTO banco (idBanco, idCliente, idSeguridad, idTransaccion,Nombre) VALUES (?,?,?,?,?)", Banco.IdBanco, Banco.IdCliente, Banco.IdSeguridad, Banco.IdTransaccion, Banco.Nombre)
	return err
}

func DeleteBancoBD(Id string) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM banco WHERE idBanco = ?", Id)
	return err
}

func UpdateBancoBD(Banco models.Bancos) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE banco set idCliente = ?, idSeguridad = ?, idTransaccion = ?,Nombre = ? WHERE idBanco = ?", Banco.IdCliente, Banco.IdSeguridad, Banco.IdTransaccion, Banco.Nombre, Banco.IdBanco)
	return err
}

func SelectBancos() ([]models.Bancos, error) { // trae todas las terminales
	//se declara un arreglo , si hay un error lo regresa vacio
	bancos := []models.Bancos{}
	bd, err := config.GetDB()
	if err != nil {
		return bancos, err
	}
	//agarra las filas
	rows, err := bd.Query("SELECT * FROM banco")
	if err != nil {
		return bancos, err
	}
	//crea los renglones
	for rows.Next() {
		//cada paso es una linea
		var banco models.Bancos
		err = rows.Scan(&banco.IdBanco, &banco.IdCliente, &banco.IdSeguridad, &banco.IdTransaccion, &banco.Nombre)
		if err != nil {
			return bancos, err
		}
		bancos = append(bancos, banco)
	}
	return bancos, nil
}

//crud Clientes 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢

func InsertClienteBD(Cliente models.Clientes) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO cliente (idClinete, idBanco, idTerminal, idTransaccion,idSeguridad) VALUES (?,?,?,?,?)", Cliente.IdCliente, Cliente.IdBanco, Cliente.IdTerminal, Cliente.IdTransaccion, Cliente.IdSeguridad)
	return err
}

func DeleteClienteBD(Id string) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM cliente WHERE idCliente = ?", Id)
	return err
}

func UpdateClienteBD(Cliente models.Clientes) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE cliente set idBanco = ?, idTerminal = ?, idTransaccion = ?,idSeguridad = ? WHERE idCliente = ?", Cliente.IdBanco, Cliente.IdTerminal, Cliente.IdTerminal, Cliente.IdTransaccion, Cliente.IdSeguridad, Cliente.IdCliente)
	return err
}

func SelectClientes() ([]models.Clientes, error) { // trae todas las terminales
	//se declara un arreglo , si hay un error lo regresa vacio
	clientes := []models.Clientes{}
	bd, err := config.GetDB()
	if err != nil {
		return clientes, err
	}
	//agarra las filas
	rows, err := bd.Query("SELECT * FROM cliente")
	if err != nil {
		return clientes, err
	}
	//crea los renglones
	for rows.Next() {
		//cada paso es una linea
		var cliente models.Clientes
		err = rows.Scan(&cliente.IdCliente, &cliente.IdBanco, &cliente.IdTerminal, &cliente.IdTransaccion, &cliente.IdSeguridad)
		if err != nil {
			return clientes, err
		}
		clientes = append(clientes, cliente)
	}
	return clientes, nil
}

func SelectClienteIdDB(id string) (models.Clientes, error) { //solo trae una terminal por id
	var cliente models.Clientes
	bd, err := config.GetDB()
	if err != nil {
		return cliente, err
	}
	row := bd.QueryRow("SELECT * FROM cliente where idCliente = ?", id)
	err = row.Scan()
	if err != nil {
		return cliente, err
	}
	//si todo sale bien
	return cliente, nil
}

//crud seguridad 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢
func InsertSeguridadBD(Seguridad models.SeguridadStruct) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO seguridad (idSeguridad, idBanco, idCliente, idTerminal,idTransaccion,CodigoPaquete,Error) VALUES (?,?,?,?,?,?,?)", Seguridad.IdSeguridad, Seguridad.IdBanco, Seguridad.IdCliente, Seguridad.IdTerminal, Seguridad.IdTransaccion, Seguridad.CodigoPaquete, Seguridad.Error)
	return err
}

func DeleteSeguridadBD(Id string) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM seguridad WHERE idSeguridad = ?", Id)
	return err
}

func UpdateSeguridadBD(Seguridad models.SeguridadStruct) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE seguridad set idBanco = ?, idCliente = ?, idTerminal = ?,idTransaccion = ?,codigoPaquete = ?, Error = ? WHERE idSeguridad = ?", Seguridad.IdBanco, Seguridad.IdCliente, Seguridad.IdTerminal, Seguridad.IdTransaccion, Seguridad.CodigoPaquete, Seguridad.Error, Seguridad.IdSeguridad)
	return err
}

func SelectSeguridad() ([]models.SeguridadStruct, error) { // trae todas las terminales
	//se declara un arreglo , si hay un error lo regresa vacio
	seguros := []models.SeguridadStruct{}
	bd, err := config.GetDB()
	if err != nil {
		return seguros, err
	}
	//agarra las filas
	rows, err := bd.Query("SELECT * FROM seguridad")
	if err != nil {
		return seguros, err
	}
	//crea los renglones
	for rows.Next() {
		//cada paso es una linea
		var seguridad models.SeguridadStruct
		err = rows.Scan(&seguridad.IdSeguridad, &seguridad.IdBanco, &seguridad.IdCliente, &seguridad.IdTransaccion, &seguridad.IdTerminal, &seguridad.CodigoPaquete, &seguridad.Error)
		if err != nil {
			return seguros, err
		}
		seguros = append(seguros, seguridad)
	}
	return seguros, nil
}

func SelectSeguridadIdDB(id string) (models.SeguridadStruct, error) { //solo trae una terminal por id
	var seguridad models.SeguridadStruct
	bd, err := config.GetDB()
	if err != nil {
		return seguridad, err
	}
	row := bd.QueryRow("SELECT * FROM seguridad where idSeguridad = ?", id)
	err = row.Scan()
	if err != nil {
		return seguridad, err
	}
	//si todo sale bien
	return seguridad, nil
}

//crud transacciones 🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢🟢
func InsertTransaccionBD(Transaccion models.Transacciones) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO Transaccion (idTransaccion, EstadoTransaccion, DatosTarjeta,DatosComercio,FechaTransaccion,idBanco) VALUES (?,?,?,?,?,?)", Transaccion.IdTransaccion, Transaccion.EstadoTransaccion, Transaccion.DatosTarjeta, Transaccion.DatosComercio, Transaccion.FechaTransaccion, Transaccion.IdBanco)
	return err
}

func DeleteTransaccionBD(Id string) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM Transaccion WHERE idTransaccion = ?", Id)
	return err
}

func UpdateTransaccionBD(Transaccion models.Transacciones) error {
	bd, err := config.GetDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE cliente set EstadoTransaccion = ?, DatosTarjeta = ?, DatosComercio = ?,FechaTransaccion = ?, idBanco WHERE idTransaccion = ?", Transaccion.EstadoTransaccion, Transaccion.DatosTarjeta, Transaccion.DatosComercio, Transaccion.FechaTransaccion, Transaccion.IdBanco)
	return err
}

func SelectTransacciones() ([]models.Transacciones, error) { // trae todas las terminales
	//se declara un arreglo , si hay un error lo regresa vacio
	transacciones := []models.Transacciones{}
	bd, err := config.GetDB()
	if err != nil {
		return transacciones, err
	}
	//agarra las filas
	rows, err := bd.Query("SELECT * FROM transacciones")
	if err != nil {
		return transacciones, err
	}
	//crea los renglones
	for rows.Next() {
		//cada paso es una linea
		var transaccion models.Transacciones
		err = rows.Scan(&transaccion.IdTransaccion, &transaccion.EstadoTransaccion, &transaccion.DatosTarjeta, &transaccion.DatosComercio, &transaccion.FechaTransaccion, &transaccion.IdBanco)
		if err != nil {
			return transacciones, err
		}
		transacciones = append(transacciones, transaccion)
	}
	return transacciones, nil
}

func SelectTransaccionIdDB(id string) (models.Transacciones, error) { //solo trae una terminal por id
	var transaccion models.Transacciones
	bd, err := config.GetDB()
	if err != nil {
		return transaccion, err
	}
	row := bd.QueryRow("SELECT * FROM Transaccion where idTransaccion = ?", id)
	err = row.Scan()
	if err != nil {
		return transaccion, err
	}
	//si todo sale bien
	return transaccion, nil
}

//querys 🟡🟡🟡🟡🟡🟡🟡🟡🟡🟡🟡🟡🟡🟡

func SelectModeloTerminalDB(id string) (models.Terminales, error) { //selecciona el modelo de la terminal
	var terminal models.Terminales
	bd, err := config.GetDB()
	if err != nil {
		return terminal, err
	}
	row := bd.QueryRow("SELECT Modelo FROM terminal where idTerminal = ?", id)
	err = row.Scan(&terminal.IdTerminal, &terminal.CodPaquete, &terminal.Estado, &terminal.Marca, &terminal.Modelo, &terminal.BancoProcedencia, &terminal.IdCliente)
	if err != nil {
		return terminal, err
	}
	//si todo sale bien
	return terminal, nil
}

func SelectMarcaTerminalDB(id string) (models.Terminales, error) { //selecciona la marca de la terminal
	var terminal models.Terminales
	bd, err := config.GetDB()
	if err != nil {
		return terminal, err
	}
	row := bd.QueryRow("SELECT Marca FROM terminal where idTerminal = ?", id)
	err = row.Scan(&terminal.IdTerminal, &terminal.CodPaquete, &terminal.Estado, &terminal.Marca, &terminal.Modelo, &terminal.BancoProcedencia, &terminal.IdCliente)
	if err != nil {
		return terminal, err
	}
	//si todo sale bien
	return terminal, nil
}

func SelectNumTransaccionTerminalDB(id string) (models.Terminales, error) { //selecciona el codigo de paquete de terminal
	var terminal models.Terminales
	bd, err := config.GetDB()
	if err != nil {
		return terminal, err
	}
	row := bd.QueryRow("SELECT CodPaquete FROM terminal where idTerminal = ?", id)
	err = row.Scan(&terminal.IdTerminal, &terminal.CodPaquete, &terminal.Estado, &terminal.Marca, &terminal.Modelo, &terminal.BancoProcedencia, &terminal.IdCliente)
	if err != nil {
		return terminal, err
	}
	//si todo sale bien
	return terminal, nil
}

func main() {

	// Chequeo de la conexión
	//if ...
	//conexion a la bd

	handler.Manejadores()

}
