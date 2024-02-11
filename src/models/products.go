package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

func GetAllProducts() ([]Product, error) {
	var products []Product
	if err := Database.Find(&products).Error; err != nil {
		// В случае ошибки возвращаем пустой список и саму ошибку
		return nil, err
	}
	// Возвращаем список продуктов без ошибки, если все прошло успешно
	return products, nil
}

func AddNewProduct(product Product) (Product, error) {
	if err := Database.Create(&product).Error; err != nil {
		// В случае ошибки возвращаем пустой продукт и саму ошибку
		return Product{}, err
	}
	// Возвращаем созданный продукт без ошибки, если все прошло успешно
	return product, nil
}

func GetOneProduct(id uint) (*Product, error) {
	var product Product
	if err := Database.First(&product, id).Error; err != nil {
		// Если продукт не найден или произошла другая ошибка
		return nil, err
	}
	// Возвращаем найденный продукт и nil в качестве ошибки
	return &product, nil
}

func DeleteProduct(id uint) (Product, error) {
	var product Product
	// Сначала находим продукт по ID
	if err := Database.First(&product, id).Error; err != nil {
		// Если продукт не найден или произошла другая ошибка
		return Product{}, err
	}

	// Теперь удаляем найденный продукт
	if err := Database.Delete(&product).Error; err != nil {
		// Если при удалении произошла ошибка
		return Product{}, err
	}

	// Возвращаем данные удаленного продукта и nil в качестве ошибки
	// Обратите внимание: в реальности, после удаления, данные продукта уже не существуют в БД,
	// но мы возвращаем копию данных до удаления.
	return product, nil
}

func UpdateProduct(id uint, updatedProduct Product) (Product, error) {
	var product Product

	// Сначала находим продукт по ID
	if err := Database.First(&product, id).Error; err != nil {
		return Product{}, err // Возвращаем пустой продукт и ошибку, если продукт не найден
	}

	// Обновляем найденный продукт данными из updatedProduct
	if err := Database.Model(&product).Updates(updatedProduct).Error; err != nil {
		return Product{}, err // Возвращаем пустой продукт и ошибку в случае ошибки обновления
	}

	return product, nil // Возвращаем обновлённый продукт без ошибки
}
