package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"sistema-streaming/database"
	"sistema-streaming/models"
)

func main() {

	database.ConectarBD()

	router := gin.Default()

	// Ruta principal de la API
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Bienvenido a la API del Sistema de Streaming LYON",
			"estado":  "Servidor funcionando correctamente",
			"version": "1.0",
		})
	})

	// =========================
	// CRUD USUARIOS
	// =========================

	router.GET("/usuarios", func(c *gin.Context) {
		var usuarios []models.Usuario

		if err := database.DB.Find(&usuarios).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, usuarios)
	})

	router.POST("/usuarios", func(c *gin.Context) {
		var usuario models.Usuario

		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		if err := database.DB.Create(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"mensaje": "Usuario registrado correctamente",
			"usuario": usuario,
		})
	})

	router.PUT("/usuarios/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var usuario models.Usuario

		if err := database.DB.First(&usuario, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		var datos models.Usuario

		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		usuario.Nombre = datos.Nombre
		usuario.Correo = datos.Correo

		if err := database.DB.Save(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Usuario actualizado correctamente",
			"usuario": usuario,
		})
	})

	router.DELETE("/usuarios/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var usuario models.Usuario

		if err := database.DB.First(&usuario, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		if err := database.DB.Delete(&usuario).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Usuario eliminado correctamente",
		})
	})

	// =========================
	// CRUD CONTENIDOS
	// =========================

	router.GET("/contenidos", func(c *gin.Context) {
		var contenidos []models.Contenido

		if err := database.DB.Find(&contenidos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, contenidos)
	})

	router.POST("/contenidos", func(c *gin.Context) {
		var contenido models.Contenido

		if err := c.ShouldBindJSON(&contenido); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		if err := database.DB.Create(&contenido).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"mensaje":   "Contenido registrado correctamente",
			"contenido": contenido,
		})
	})

	router.PUT("/contenidos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var contenido models.Contenido

		if err := database.DB.First(&contenido, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado"})
			return
		}

		var datos models.Contenido

		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		contenido.Titulo = datos.Titulo
		contenido.Genero = datos.Genero
		contenido.Tipo = datos.Tipo

		if err := database.DB.Save(&contenido).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje":   "Contenido actualizado correctamente",
			"contenido": contenido,
		})
	})

	router.DELETE("/contenidos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var contenido models.Contenido

		if err := database.DB.First(&contenido, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contenido no encontrado"})
			return
		}

		if err := database.DB.Delete(&contenido).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Contenido eliminado correctamente",
		})
	})

	// =========================
	// CRUD SUSCRIPCIONES
	// =========================

	router.GET("/suscripciones", func(c *gin.Context) {
		var suscripciones []models.Suscripcion

		if err := database.DB.Find(&suscripciones).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, suscripciones)
	})

	router.POST("/suscripciones", func(c *gin.Context) {
		var suscripcion models.Suscripcion

		if err := c.ShouldBindJSON(&suscripcion); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		if err := database.DB.Create(&suscripcion).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"mensaje":     "Suscripción registrada correctamente",
			"suscripcion": suscripcion,
		})
	})

	router.PUT("/suscripciones/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var suscripcion models.Suscripcion

		if err := database.DB.First(&suscripcion, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Suscripción no encontrada"})
			return
		}

		var datos models.Suscripcion

		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		suscripcion.NombrePlan = datos.NombrePlan
		suscripcion.Precio = datos.Precio
		suscripcion.Descripcion = datos.Descripcion

		if err := database.DB.Save(&suscripcion).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje":     "Suscripción actualizada correctamente",
			"suscripcion": suscripcion,
		})
	})

	router.DELETE("/suscripciones/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var suscripcion models.Suscripcion

		if err := database.DB.First(&suscripcion, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Suscripción no encontrada"})
			return
		}

		if err := database.DB.Delete(&suscripcion).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"mensaje": "Suscripción eliminada correctamente",
		})
	})

	// =========================
	// ESTADÍSTICAS REALES
	// =========================

	router.GET("/estadisticas", func(c *gin.Context) {
		var totalUsuarios int64
		var totalContenidos int64
		var totalSuscripciones int64

		database.DB.Model(&models.Usuario{}).Count(&totalUsuarios)
		database.DB.Model(&models.Contenido{}).Count(&totalContenidos)
		database.DB.Model(&models.Suscripcion{}).Count(&totalSuscripciones)

		c.JSON(http.StatusOK, gin.H{
			"total_usuarios":      totalUsuarios,
			"total_contenidos":    totalContenidos,
			"total_suscripciones": totalSuscripciones,
			"estado_del_proyecto": "API REST conectada a MySQL mediante GORM",
			"base_de_datos":       "streaming",
		})
	})

	router.Run(":8080")
}
