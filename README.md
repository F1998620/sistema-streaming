# 🎬 Sistema de Gestión para Plataforma de Streaming

## 📖 Descripción

Este proyecto consiste en el desarrollo de un **Sistema de Gestión para una Plataforma de Streaming**, implementado en el lenguaje de programación **Go (Golang)**.

La aplicación integra una **API REST**, una **base de datos MySQL** y el framework **GORM**, permitiendo administrar usuarios, contenidos multimedia y planes de suscripción mediante servicios web.

---

# 🎯 Objetivo

Desarrollar una aplicación que integre los conocimientos adquiridos durante la asignatura, aplicando:

* Programación Orientada a Objetos.
* Encapsulación.
* Interfaces.
* Manejo de errores.
* Servicios Web REST.
* Serialización mediante JSON.
* Persistencia de datos con MySQL.

---

# 🛠 Tecnologías utilizadas

* Go (Golang)
* Gin Gonic
* MySQL Server
* GORM
* Visual Studio Code
* Git
* GitHub
* Thunder Client

---

# 🏗 Arquitectura del proyecto

```
Sistema Streaming

├── api
│   └── main.go
│
├── database
│   └── conexion.go
│
├── models
│   ├── usuario.go
│   ├── contenido.go
│   └── suscripcion.go
│
├── services
│
├── views
│
├── go.mod
├── go.sum
└── README.md
```

---

# 📡 Servicios REST implementados

## Usuarios

| Método | Endpoint       | Descripción                |
| ------ | -------------- | -------------------------- |
| GET    | /usuarios      | Obtener todos los usuarios |
| POST   | /usuarios      | Registrar usuario          |
| PUT    | /usuarios/{id} | Actualizar usuario         |
| DELETE | /usuarios/{id} | Eliminar usuario           |

---

## Contenidos

| Método | Endpoint         | Descripción          |
| ------ | ---------------- | -------------------- |
| GET    | /contenidos      | Obtener contenidos   |
| POST   | /contenidos      | Registrar contenido  |
| PUT    | /contenidos/{id} | Actualizar contenido |
| DELETE | /contenidos/{id} | Eliminar contenido   |

---

## Suscripciones

| Método | Endpoint            | Descripción            |
| ------ | ------------------- | ---------------------- |
| GET    | /suscripciones      | Obtener suscripciones  |
| POST   | /suscripciones      | Registrar suscripción  |
| PUT    | /suscripciones/{id} | Actualizar suscripción |
| DELETE | /suscripciones/{id} | Eliminar suscripción   |

---

## Estadísticas

| Método | Endpoint      |
| ------ | ------------- |
| GET    | /estadisticas |

---

# 🗄 Base de datos

Nombre de la base de datos:

```
streaming
```

Tablas implementadas:

* usuarios
* contenidos
* suscripciones

---

# ▶️ Ejecución del proyecto

Instalar dependencias:

```bash
go mod tidy
```

Ejecutar la API:

```bash
go run api/main.go
```

Servidor:

```
http://localhost:8080
```

---

# ✅ Funcionalidades implementadas

* CRUD de Usuarios.
* CRUD de Contenidos.
* CRUD de Suscripciones.
* Estadísticas del sistema.
* Persistencia en MySQL.
* API REST.
* Serialización JSON.
* Manejo de errores.
* Conexión mediante GORM.

---

# 📷 Evidencias del proyecto

Se recomienda agregar capturas de:

* API ejecutándose.
* Thunder Client.
* MySQL Workbench.
* GitHub.
* Estadísticas del sistema.

---

# 👨‍💻 Autor

Francis Leon

Estudiante de Ingeniería en Ciberseguridad

---

# 📅 Fecha

28 Junio 2026
