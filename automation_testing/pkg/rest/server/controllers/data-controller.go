package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sindhutrisha/intelops_compage_testing/automation_testing/pkg/rest/server/models"
	"github.com/sindhutrisha/intelops_compage_testing/automation_testing/pkg/rest/server/services"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type DataController struct {
	dataService *services.DataService
}

func NewDataController() (*DataController, error) {
	dataService, err := services.NewDataService()
	if err != nil {
		return nil, err
	}
	return &DataController{
		dataService: dataService,
	}, nil
}

func (dataController *DataController) CreateData(context *gin.Context) {
	// validate input
	var input models.Data
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger data creation
	if _, err := dataController.dataService.CreateData(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Data created successfully"})
}

func (dataController *DataController) UpdateData(context *gin.Context) {
	// validate input
	var input models.Data
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger data update
	if _, err := dataController.dataService.UpdateData(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Data updated successfully"})
}

func (dataController *DataController) FetchData(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger data fetching
	data, err := dataController.dataService.GetData(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, data)
}

func (dataController *DataController) DeleteData(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger data deletion
	if err := dataController.dataService.DeleteData(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Data deleted successfully",
	})
}

func (dataController *DataController) ListData(context *gin.Context) {
	// trigger all data fetching
	data, err := dataController.dataService.ListData()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, data)
}

func (*DataController) PatchData(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*DataController) OptionsData(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*DataController) HeadData(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
