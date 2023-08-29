package handlers

import (
	"bufio"
	"clase_03_mysql_driver/conectar"
	"clase_03_mysql_driver/modelos"
	"fmt"
	"log"
	"os"
)

func Listar() {
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes ORDER BY id DESC;"
	datos, err := conectar.Db.Query(sql)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
	/*
		clientes := modelos.Clientes{}
		for datos.Next() {
			dato := modelos.Cliente{}
			datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
			clientes = append(clientes, dato)
		}
		fmt.Println(clientes)
	*/

	fmt.Println("\n listado de clientes")
	fmt.Println("-------------")
	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | E-mail: %s | Telefono: %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
		fmt.Println("-------------------------------------------------------------")
	}

}

func ListarPorID(id int) {
	conectar.Conectar()
	sql := "SELECT id, nombre, correo, telefono FROM clientes WHERE id = ?;"
	datos, err := conectar.Db.Query(sql, id)
	if err != nil {
		fmt.Println(err)
	}
	defer conectar.CerrarConexion()
	fmt.Println("\n listado de cliente")
	fmt.Println("---------------------------------------------------------------------")
	for datos.Next() {
		var dato modelos.Cliente
		err := datos.Scan(&dato.Id, &dato.Nombre, &dato.Correo, &dato.Telefono)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Id: %v | Nombre: %v | E-mail: %s | Telefono: %s \n", dato.Id, dato.Nombre, dato.Correo, dato.Telefono)
	}
	fmt.Println("---------------------------------------------------------------------------")

}

func Insertar(cliente modelos.Cliente) {
	conectar.Conectar()
	sql := "INSERT INTO clientes VALUES (null, ?, ?, ?, ?)"
	//VALUES ('Juan Pérez', 'juan@example.com', '123-456-7890', '2023-08-24 14:00:00');"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, cliente.Fecha)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se creo el registro exitosamente")
}

func Editar(cliente modelos.Cliente, id int) {
	conectar.Conectar()
	sql := "UPDATE clientes SET nombre = ?, correo = ?, telefono = ?, fecha = ? WHERE id = ?;"
	result, err := conectar.Db.Exec(sql, cliente.Nombre, cliente.Correo, cliente.Telefono, cliente.Fecha, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("Se edito el registro exitosamente")
}

func Eliminar(id int) {
	conectar.Conectar()
	sql := "DELETE FROM clientes WHERE id = ?;"
	_, err := conectar.Db.Exec(sql, id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Se elimino el registro exitosamente")
}

// funciones de trabajo

var ID int
var nombre, correo, telefono string

func Ejecutar() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Seleccione una opción: \n\n")
	fmt.Println("1. Listar clientes \n")
	fmt.Println("2. Listar clientes por ID \n")
	fmt.Println("3. Crear cliente \n")
	fmt.Println("4. Editar cliente \n")
	fmt.Println("5. Eliminar cliente \n")
	if scanner.Scan() {
		for {
			if scanner.Text() == "1" {
				Listar()
				return
			}
			if scanner.Text() == "2" {
				fmt.Println("")
				Listar()
				return
			}
			// este lo debes terminar chavo, todo el menu perrin
		}
	}

}
