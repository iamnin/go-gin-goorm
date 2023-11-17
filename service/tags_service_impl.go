package service

import (
	"github.com/go-playground/validator/v10"
	"go-gin-goorm/data/request"
	"go-gin-goorm/data/response"
	"go-gin-goorm/helper"
	"go-gin-goorm/models"
	"go-gin-goorm/repository"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
	Validate       *validator.Validate
}

func (t *TagsServiceImpl) Create(tags request.CreateTagRequest) {
	//TODO implement me
	err := t.Validate.Struct(tags)
	helper.ErrorPanic(err)
	tagModel := models.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}

func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	//TODO implement me
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)

	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}

func (t *TagsServiceImpl) Delete(tagsId int) {
	//TODO implement me
	t.TagsRepository.Delete(tagsId)

}

func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	//TODO implement me
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}

func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	//TODO implement me
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository, validate *validator.Validate) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
		Validate:       validate,
	}
}
