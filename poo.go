package main

import "fmt"

// Gestionable es una interfaz que obliga a implementar el método MostrarInfo.
type Gestionable interface {
	MostrarInfo()
}

// UsuarioPOO representa un usuario aplicando encapsulamiento.
type UsuarioPOO struct {
	id     int
	nombre string
	correo string
}

// Constructor
func NewUsuarioPOO(id int, nombre string, correo string) UsuarioPOO {
	return UsuarioPOO{
		id:     id,
		nombre: nombre,
		correo: correo,
	}
}

// Métodos Get y Set para aplicar encapsulamiento.
func (u *UsuarioPOO) SetNombre(nombre string) {
	u.nombre = nombre
}

func (u UsuarioPOO) GetNombre() string {
	return u.nombre
}

func (u *UsuarioPOO) SetCorreo(correo string) {
	u.correo = correo
}

func (u UsuarioPOO) GetCorreo() string {
	return u.correo
}

func (u UsuarioPOO) MostrarInfo() {
	fmt.Println("ID:", u.id)
	fmt.Println("Nombre:", u.nombre)
	fmt.Println("Correo:", u.correo)
}
