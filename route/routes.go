package route

import (
	"new_ip_data_api/controller"

	"github.com/gin-gonic/gin"
)

func setTrustedProxies(server *gin.Engine) {
	err := server.SetTrustedProxies([]string{"192.168.1.1", "192.168.1.2"})
	if err != nil {
		panic(err)
	}
}

func Endpoints(server *gin.Engine, controller *controller.IpDataController) {

	setTrustedProxies(server)

	server.POST("/store_ip/", controller.StoreIpData)
	server.GET("/total_search_by_ip/", controller.GetTotalSearchByIP)
	server.GET("/total_search_by_country/", controller.GetTotalSearchByCountry)
	server.DELETE("/delete_entries_by_ip/", controller.DeleteIpDataByIp)
	server.GET("/nearest_to_se_square/", controller.DistanciaPcaSe)
	server.Run(":8080")
}
