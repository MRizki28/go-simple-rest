package mahasiswarepositories

import (
	"net/http"

	"github.com/MRizki28/go-simple-rest/src/config"
	"github.com/MRizki28/go-simple-rest/src/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func GetAllData(c *gin.Context) {
	var mahasiswa []models.TbMahasiswa
	config.DB.Find(&mahasiswa)

	if len(mahasiswa) <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "not Found",
			"message": "Data not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"data":    mahasiswa,
		})
	}
}

func CreateData(c *gin.Context) {

	mahasiswa := models.TbMahasiswa{
		Nama: c.PostForm("nama"),
		Alamat: c.PostForm("alamat"),
	}

	if  err := validate.Struct(mahasiswa); err != nil {
		errors := make(map[string]string)
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				field := fieldError.Field()
				tag := fieldError.Tag()
				errors[field] = field + " is required."
				if tag == "required" {
					errors[field] = field + " is required."
				}
			}
		}
		c.JSON(http.StatusUnprocessableEntity , gin.H{
			"message": "not validate",
			"error": errors,
		})

		return
	}

	if  err := config.DB.Create(&mahasiswa); err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed",
			"status": "Data not created",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"status": "Data has been created",
		"data": mahasiswa,
	})
}

func GetDataById(c *gin.Context) {
	var mahasiswa models.TbMahasiswa
	id := c.Param("id")
	config.DB.First(&mahasiswa, id)

	if mahasiswa.Id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status": "not Found",
			"message": "Data not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"data":    mahasiswa,
		})
	}
}