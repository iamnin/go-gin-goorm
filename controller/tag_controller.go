package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-goorm/data/request"
	"go-gin-goorm/data/response"
	"go-gin-goorm/helper"
	"go-gin-goorm/service"
	"net/http"
	"strconv"
)

type TagController struct {
	tagsService service.TagsService
}

func NewsTagController(service service.TagsService) *TagController {
	return &TagController{
		tagsService: service,
	}
}

func (controller *TagController) Create(ctx *gin.Context) {
	createTagsRequest := request.CreateTagRequest{}
	err := ctx.ShouldBindJSON(&createTagsRequest)
	helper.ErrorPanic(err)

	controller.tagsService.Create(createTagsRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) Update(ctx *gin.Context) {
	updateTagsRequest := request.UpdateTagsRequest{}
	err := ctx.ShouldBindJSON(&updateTagsRequest)
	helper.ErrorPanic(err)

	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	updateTagsRequest.Id = id

	controller.tagsService.Update(updateTagsRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TagController) Delete(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)
	controller.tagsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *TagController) FindById(ctx *gin.Context) {
	tagId := ctx.Param("tagId")
	id, err := strconv.Atoi(tagId)
	helper.ErrorPanic(err)

	tagResponse := controller.tagsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *TagController) FindAll(ctx *gin.Context) {
	tagResponse := controller.tagsService.FindAll()

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "ok",
		Data:   tagResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
