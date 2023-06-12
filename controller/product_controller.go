package controller

import (
	"net/http"

	"github.com/aabdullahgungor/go-restapi-rabbitmq/constants"
	"github.com/aabdullahgungor/go-restapi-rabbitmq/utils"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
}

func (pc ProductController) GetAllProducts(c *gin.Context) {
	var rmqProducer utils.RMQProducer
	err := rmqProducer.PublishMessage("application/json", nil, constants.TASK1)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.Header("Content-Type", "application/json")
	c.IndentedJSON(http.StatusOK, gin.H{"response": "Request received"})
}

func (pc ProductController) GetProductById(c *gin.Context) {
}

func (pc ProductController) CreateProduct(c *gin.Context) {
}

func (pc ProductController) EditProduct(c *gin.Context) {
}

func (pc ProductController) DeleteProduct(c *gin.Context) {
}
