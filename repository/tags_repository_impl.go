package repository

import (
	"errors"
	"go-gin-goorm/data/request"
	"go-gin-goorm/helper"
	"go-gin-goorm/models"
	"gorm.io/gorm"
)

type TagsRepositoryImpl struct {
	Db *gorm.DB
}

func (t *TagsRepositoryImpl) FindById(tagsId int) (tags models.Tags, err error) {
	//TODO implement me
	var tag models.Tags

	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

func NewTagsRepositoryImpl(Db *gorm.DB) TagsRepository {
	return &TagsRepositoryImpl{Db: Db}
}

// Delete
func (t *TagsRepositoryImpl) Delete(tagsId int) {
	var tags models.Tags
	result := t.Db.Where("id= ?", tagsId).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// Find All
func (t *TagsRepositoryImpl) FindAll() []models.Tags {
	var tags []models.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById
func (t *TagsRepositoryImpl) FindByd(tagsId int) (tags models.Tags, err error) {
	var tag models.Tags

	result := t.Db.Find(&tag, tagsId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TagsRepository
func (t *TagsRepositoryImpl) Save(tags models.Tags) {
	result := t.Db.Create(&tags)
	helper.ErrorPanic(result.Error)
}

// Update implements Tags Repository
func (t *TagsRepositoryImpl) Update(tags models.Tags) {
	var updateTag = request.UpdateTagsRequest{
		Id:   tags.Id,
		Name: tags.Name,
	}
	result := t.Db.Model(&tags).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
