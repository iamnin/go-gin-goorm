package service

import (
	"go-gin-goorm/data/request"
	"go-gin-goorm/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
