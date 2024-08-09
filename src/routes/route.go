package routes

import (
	"github.com/MRizki28/go-simple-rest/src/repositories"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/v1/mahasiswa", mahasiswarepositories.GetAllData)
	r.POST("/v1/mahasiswa/create", mahasiswarepositories.CreateData)
	r.GET("/v1/mahasiswa/get/:id", mahasiswarepositories.GetDataById)

	return r
}