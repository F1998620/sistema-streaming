package main

import (
	"errors"
	"strings"
)

// validarTexto revisa que el dato ingresado no esté vacío.
// Esta función permite aplicar manejo de errores en entradas del sistema.
func validarTexto(valor string, campo string) error {
	if strings.TrimSpace(valor) == "" {
		return errors.New("el campo " + campo + " no puede estar vacío")
	}

	return nil
}
