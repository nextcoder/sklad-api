package controllers

import (
	"main/src/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetOneProduct(c *gin.Context) {
	// Получаем ID продукта из параметров запроса
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Если ID не может быть преобразован в число, возвращаем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid product ID", "data": nil})
		return
	}

	// Получаем продукт по ID
	product, err := models.GetOneProduct(uint(id))
	if err != nil {
		// Если продукт не найден или произошла другая ошибка
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Product not found", "data": nil})
		return
	}

	// Возвращаем найденный продукт
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product retrieved successfully", "data": product})
}

func GetProducts(c *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error(), "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Products retrieved successfully", "data": products})
}

func CreateProduct(c *gin.Context) {
	var product models.Product // Используйте структуру Product из вашего пакета models

	// Пробуем получить JSON из тела запроса и десериализовать его в структуру
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	// Передаем десериализованный продукт в функцию добавления продукта
	addedProduct, err := models.AddNewProduct(product) // Убедитесь, что AddNewProduct принимает объект Product и возвращает Product, error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	// Возвращаем добавленный продукт в ответе
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product created successfully", "data": addedProduct})
}

func DeleteProduct(c *gin.Context) {
	// Получаем ID продукта из параметров запроса
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// Если ID не может быть преобразован в число, возвращаем ошибку
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": "Invalid product ID", "data": nil})
		return
	}

	// Получаем продукт по ID
	product, err := models.DeleteProduct(uint(id))
	if err != nil {
		// Если продукт не найден или произошла другая ошибка
		c.JSON(http.StatusNotFound, gin.H{"status": "failed", "message": "Product not found", "data": nil})
		return
	}

	// Возвращаем найденный продукт
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Product deleted successfully", "data": product})
}

func UpdateProduct(c *gin.Context) {
	// Получаем ID продукта из параметра пути
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	// Десериализуем JSON из тела запроса в структуру Product
	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Вызываем функцию обновления продукта из модели, передавая ID и новые данные
	product, err := models.UpdateProduct(uint(id), updatedProduct)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем обновлённый продукт
	c.JSON(http.StatusOK, gin.H{"product": product})
}
