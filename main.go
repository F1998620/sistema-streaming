package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Usuario struct {
	ID     int
	Nombre string
	Correo string
}

type Contenido struct {
	ID     int
	Titulo string
	Genero string
	Tipo   string
}

type Suscripcion struct {
	ID          int
	NombrePlan  string
	Precio      string
	Descripcion string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var usuarios []Usuario
	var contenidos []Contenido
	var suscripciones []Suscripcion

	opcion := ""

	for opcion != "14" {
		fmt.Println(" PLATAFORMA DE STREAMING LYON ")
		fmt.Println(" Fecha:", time.Now().Format("02/01/2006 15:04"))
		fmt.Println("1. Registrar Usuario")
		fmt.Println("2. Ver Usuarios")
		fmt.Println("3. Buscar Usuario")
		fmt.Println("4. Actualizar Usuario")
		fmt.Println("5. Eliminar Usuario")
		fmt.Println("6. Registrar Contenido")
		fmt.Println("7. Ver Contenido")
		fmt.Println("8. Buscar Contenido")
		fmt.Println("9. Actualizar Contenido")
		fmt.Println("10. Eliminar Contenido")
		fmt.Println("11. Registrar Suscripción")
		fmt.Println("12. Ver Suscripciones")
		fmt.Println("13. Ver Estadísticas del Sistema")
		fmt.Println("14. Salir")
		fmt.Print("Seleccione una opción: ")

		opcion, _ = reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {

		case "1":
			fmt.Print("Ingrese el nombre: ")
			nombre, _ := reader.ReadString('\n')
			nombre = strings.TrimSpace(nombre)

			fmt.Print("Ingrese el correo: ")
			correo, _ := reader.ReadString('\n')
			correo = strings.TrimSpace(correo)

			nuevoUsuario := Usuario{
				ID:     len(usuarios) + 1,
				Nombre: nombre,
				Correo: correo,
			}

			usuarios = append(usuarios, nuevoUsuario)
			fmt.Println("Usuario registrado correctamente.")

		case "2":
			fmt.Println("\nLISTA DE USUARIOS REGISTRADOS")
			fmt.Println("--------------------------------")

			if len(usuarios) == 0 {
				fmt.Println("No existen usuarios registrados.")
			} else {
				for _, usuario := range usuarios {
					fmt.Println("ID:", usuario.ID)
					fmt.Println("Nombre:", usuario.Nombre)
					fmt.Println("Correo:", usuario.Correo)
					fmt.Println("--------------------------------")
				}
			}

		case "3":
			fmt.Print("Ingrese el nombre del usuario a buscar: ")
			busqueda, _ := reader.ReadString('\n')
			busqueda = strings.TrimSpace(busqueda)

			encontrado := false

			for _, usuario := range usuarios {
				if strings.Contains(strings.ToLower(usuario.Nombre), strings.ToLower(busqueda)) {
					fmt.Println("\nUsuario encontrado:")
					fmt.Println("ID:", usuario.ID)
					fmt.Println("Nombre:", usuario.Nombre)
					fmt.Println("Correo:", usuario.Correo)
					encontrado = true
				}
			}

			if !encontrado {
				fmt.Println("No se encontró ningún usuario con ese nombre.")
			}

		case "4":
			fmt.Print("Ingrese el ID del usuario a actualizar: ")
			idTexto, _ := reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)

			actualizado := false

			for i := range usuarios {
				if fmt.Sprint(usuarios[i].ID) == idTexto {
					fmt.Print("Ingrese el nuevo nombre: ")
					nuevoNombre, _ := reader.ReadString('\n')
					nuevoNombre = strings.TrimSpace(nuevoNombre)

					fmt.Print("Ingrese el nuevo correo: ")
					nuevoCorreo, _ := reader.ReadString('\n')
					nuevoCorreo = strings.TrimSpace(nuevoCorreo)

					usuarios[i].Nombre = nuevoNombre
					usuarios[i].Correo = nuevoCorreo

					fmt.Println("Usuario actualizado correctamente.")
					actualizado = true
				}
			}

			if !actualizado {
				fmt.Println("No se encontró un usuario con ese ID.")
			}

		case "5":
			fmt.Print("Ingrese el ID del usuario a eliminar: ")
			idTexto, _ := reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)

			eliminado := false

			for i := range usuarios {
				if fmt.Sprint(usuarios[i].ID) == idTexto {
					usuarios = append(usuarios[:i], usuarios[i+1:]...)
					fmt.Println("Usuario eliminado correctamente.")
					eliminado = true
					break
				}
			}

			if !eliminado {
				fmt.Println("No se encontró un usuario con ese ID.")
			}

		case "6":
			fmt.Print("Ingrese el título del contenido: ")
			titulo, _ := reader.ReadString('\n')
			titulo = strings.TrimSpace(titulo)

			fmt.Print("Ingrese el género: ")
			genero, _ := reader.ReadString('\n')
			genero = strings.TrimSpace(genero)

			fmt.Print("Ingrese el tipo (Película/Serie/Documental): ")
			tipo, _ := reader.ReadString('\n')
			tipo = strings.TrimSpace(tipo)

			nuevoContenido := Contenido{
				ID:     len(contenidos) + 1,
				Titulo: titulo,
				Genero: genero,
				Tipo:   tipo,
			}

			contenidos = append(contenidos, nuevoContenido)
			fmt.Println("Contenido registrado correctamente.")

		case "7":
			fmt.Println("\nCATÁLOGO DE CONTENIDO")
			fmt.Println("--------------------------------")

			if len(contenidos) == 0 {
				fmt.Println("No existe contenido registrado.")
			} else {
				for _, contenido := range contenidos {
					fmt.Println("ID:", contenido.ID)
					fmt.Println("Título:", contenido.Titulo)
					fmt.Println("Género:", contenido.Genero)
					fmt.Println("Tipo:", contenido.Tipo)
					fmt.Println("--------------------------------")
				}
			}

		case "8":
			fmt.Print("Ingrese el título del contenido a buscar: ")
			busqueda, _ := reader.ReadString('\n')
			busqueda = strings.TrimSpace(busqueda)

			encontrado := false

			for _, contenido := range contenidos {
				if strings.Contains(strings.ToLower(contenido.Titulo), strings.ToLower(busqueda)) {
					fmt.Println("\nContenido encontrado:")
					fmt.Println("ID:", contenido.ID)
					fmt.Println("Título:", contenido.Titulo)
					fmt.Println("Género:", contenido.Genero)
					fmt.Println("Tipo:", contenido.Tipo)
					fmt.Println("--------------------------------")
					encontrado = true
				}
			}

			if !encontrado {
				fmt.Println("No se encontró contenido con ese título.")
			}

		case "9":
			fmt.Print("Ingrese el ID del contenido a actualizar: ")
			idTexto, _ := reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)

			actualizado := false

			for i := range contenidos {
				if fmt.Sprint(contenidos[i].ID) == idTexto {
					fmt.Print("Ingrese el nuevo título: ")
					nuevoTitulo, _ := reader.ReadString('\n')
					nuevoTitulo = strings.TrimSpace(nuevoTitulo)

					fmt.Print("Ingrese el nuevo género: ")
					nuevoGenero, _ := reader.ReadString('\n')
					nuevoGenero = strings.TrimSpace(nuevoGenero)

					fmt.Print("Ingrese el nuevo tipo: ")
					nuevoTipo, _ := reader.ReadString('\n')
					nuevoTipo = strings.TrimSpace(nuevoTipo)

					contenidos[i].Titulo = nuevoTitulo
					contenidos[i].Genero = nuevoGenero
					contenidos[i].Tipo = nuevoTipo

					fmt.Println("Contenido actualizado correctamente.")
					actualizado = true
				}
			}

			if !actualizado {
				fmt.Println("No se encontró contenido con ese ID.")
			}

		case "10":
			fmt.Print("Ingrese el ID del contenido a eliminar: ")
			idTexto, _ := reader.ReadString('\n')
			idTexto = strings.TrimSpace(idTexto)

			eliminado := false

			for i := range contenidos {
				if fmt.Sprint(contenidos[i].ID) == idTexto {
					contenidos = append(contenidos[:i], contenidos[i+1:]...)
					fmt.Println("Contenido eliminado correctamente.")
					eliminado = true
					break
				}
			}

			if !eliminado {
				fmt.Println("No se encontró contenido con ese ID.")
			}

		case "11":
			fmt.Print("Ingrese el nombre del plan: ")
			nombrePlan, _ := reader.ReadString('\n')
			nombrePlan = strings.TrimSpace(nombrePlan)

			fmt.Print("Ingrese el precio del plan: ")
			precio, _ := reader.ReadString('\n')
			precio = strings.TrimSpace(precio)

			fmt.Print("Ingrese una descripción del plan: ")
			descripcion, _ := reader.ReadString('\n')
			descripcion = strings.TrimSpace(descripcion)

			nuevaSuscripcion := Suscripcion{
				ID:          len(suscripciones) + 1,
				NombrePlan:  nombrePlan,
				Precio:      precio,
				Descripcion: descripcion,
			}

			suscripciones = append(suscripciones, nuevaSuscripcion)
			fmt.Println("Suscripción registrada correctamente.")

		case "12":
			fmt.Println("\nPLANES DE SUSCRIPCIÓN")
			fmt.Println("--------------------------------")

			if len(suscripciones) == 0 {
				fmt.Println("No existen suscripciones registradas.")
			} else {
				for _, suscripcion := range suscripciones {
					fmt.Println("ID:", suscripcion.ID)
					fmt.Println("Plan:", suscripcion.NombrePlan)
					fmt.Println("Precio:", suscripcion.Precio)
					fmt.Println("Descripción:", suscripcion.Descripcion)
					fmt.Println("--------------------------------")
				}
			}

		case "13":
			fmt.Println("\nESTADÍSTICAS DEL SISTEMA")
			fmt.Println("--------------------------------")
			fmt.Println("Total de usuarios registrados:", len(usuarios))
			fmt.Println("Total de contenidos registrados:", len(contenidos))
			fmt.Println("Total de suscripciones registradas:", len(suscripciones))
			fmt.Println("Porcentaje de avance del proyecto: 75%")
			fmt.Println("Estado: Módulos principales implementados")
			fmt.Println("--------------------------------")

		case "14":
			fmt.Println("Saliendo del sistema...")

		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}
