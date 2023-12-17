package productHandlers

import (
	middleware "ecommerce/api/middlewares/user"
	productServices "ecommerce/services/product"
	"ecommerce/types"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductApiHandlers struct {
	service productServices.ProductServices
}

func NewProductApiHandler(database *mongo.Database) *ProductApiHandlers {
	return &ProductApiHandlers{
		service: *productServices.NewProductService(database),
	}
}

func SetupProductRouter(p *ProductApiHandlers) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.Use(middleware.ApiKeyCheck(), middleware.CheckAccessTokenAuth())

	router.GET("/products", p.getProducts)                                    // customer access and dealer access
	router.POST("/products", middleware.CheckIfDealerRole(), p.createProduct) // create product only dealer access
	router.PUT("/products", middleware.CheckIfDealerRole())                   // update product only dealer access
	// router.GET("/products/:id")                                               // customer access and dealer access
	return router
}

func (p ProductApiHandlers) getProducts(ctx *gin.Context) {
	// var products []*types.Product
	name, isPresent := ctx.GetQuery("name")
	if !isPresent {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Provide Name"})

	}
	products, err := p.service.FindProductsByName(name)
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
	}
	ctx.JSON(200, products)
}

func (p ProductApiHandlers) createProduct(ctx *gin.Context) {
	var product *types.Product
	ctx.BindJSON(&product)
	userId := ctx.MustGet("userId")
	product.DealerId = userId.(types.DealerID)
	log.Println(product)
	result, err := p.service.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"id": result})
	return
}
