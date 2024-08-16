package route

import (
	"new_ip_data_api/controller"

	"github.com/gin-gonic/gin"
)

func Endpoints(server *gin.Engine, controller *controller.IpDataController) {
	server.POST("/store_ip/", controller.StoreIpData)
	server.GET("/total_query_by_ip/", controller.GetTotalSearchByIP)
	server.Run(":8080")
}
