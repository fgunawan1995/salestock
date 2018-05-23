package main
//Imports
import (
	"github.com/gin-gonic/gin"
)
//Routing
func main() {
	r := gin.Default()

	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		//Items
		v1.POST("/items", PostItem)
		v1.GET("/items", GetItems)
		v1.GET("/items/:id", GetItem)
		v1.PUT("/items/:id", UpdateItem)
		v1.DELETE("/items/:id", DeleteItem)
		//Barang Masuk
		v1.POST("/itemsin", PostItemIn)
		v1.GET("/itemsin", GetItemsIn)
		v1.GET("/itemsin/:id", GetItemIn)
		v1.PUT("/itemsin/:id", UpdateItemIn)
		v1.DELETE("/itemsin/:id", DeleteItemIn)
		//Barang Keluar
		v1.POST("/itemsout", PostItemOut)
		v1.GET("/itemsout", GetItemsOut)
		v1.GET("/itemsout/:id", GetItemOut)
		v1.PUT("/itemsout/:id", UpdateItemOut)
		v1.DELETE("/itemsout/:id", DeleteItemOut)
		//Report
		v1.GET("/itemsreport", GetItemsReport)
		v1.GET("/salesreport", GetSalesReport)
		//Export CSV
		v1.GET("/csvitems", GetCsvItems)
		v1.GET("/csvitemsin", GetCsvItemsIn)
		v1.GET("/csvitemsout", GetCsvItemsOut)
		v1.GET("/csvitemsreport", GetCsvItemsReport)
		v1.GET("/csvsalesreport", GetCsvSalesReport)
	}
	r.Run(":8080")
}