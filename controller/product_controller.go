package controller

import (
	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (pc ProductController) GetAllProducts(c *gin.Context) {
	// products, err := ps.service.GetAll()
	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Products cannot show: " + err.Error()})
	// 	return
	// }
	// c.Header("Content-Type", "application/json")
	// c.IndentedJSON(http.StatusOK, gin.H{"response": "Request received"})
}

func (pc ProductController) GetProductById(c *gin.Context) {
	// str_id := c.Param("id")
	// product, err := ps.service.GetById(str_id)
	// if err != nil {
	// 	if errors.Is(err, service.ErrIDIsNotValid) {
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 		return
	// 	} else if errors.Is(err, service.ErrProductNotFound) {
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.Header("Content-Type", "application/json")
	// c.IndentedJSON(http.StatusOK, product)
}

func (pc ProductController) CreateProduct(c *gin.Context) {
	// var product models.Product
	// err := c.ShouldBindJSON(&product)

	// if err != nil {
	// 	c.IndentedJSON(400, gin.H{
	// 		"error": "cannot bind JSON: " + err.Error(),
	// 	})
	// 	return
	// }

	// err = ps.service.Create(&product)

	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotAcceptable, gin.H{
	// 		"error": "cannot create product: " + err.Error(),
	// 	})
	// 	return
	// }

	// c.IndentedJSON(http.StatusCreated, gin.H{"message": "Product has been created"})
}

func (pc ProductController) EditProduct(c *gin.Context) {
	// var product models.Product
	// err := c.ShouldBindJSON(&product)

	// if err != nil {
	// 	c.IndentedJSON(400, gin.H{
	// 		"error": "cannot bind JSON: " + err.Error(),
	// 	})
	// 	return
	// }

	// err = ps.service.Edit(&product)

	// if err != nil {
	// 	c.IndentedJSON(http.StatusNotAcceptable, gin.H{
	// 		"error": "cannot edit product: " + err.Error(),
	// 	})
	// 	return
	// }

	// c.IndentedJSON(http.StatusCreated, gin.H{"message": "Product has been edited", "product_id": product.Id})
}

func (pc ProductController) DeleteProduct(c *gin.Context) {
	// str_id := c.Param("id")
	// err := ps.service.Delete(str_id)
	// if err != nil {
	// 	if errors.Is(err, service.ErrIDIsNotValid) {
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "id is not valid" + err.Error()})
	// 		return
	// 	} else if errors.Is(err, service.ErrProductNotFound) {
	// 		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Product cannot be found" + err.Error()})
	// 		return
	// 	}
	// 	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Product has been deleted", "product_id": str_id})
}
