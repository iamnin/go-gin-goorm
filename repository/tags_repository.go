package repository

import "go-gin-goorm/models"

type TagsRepository interface {
	Save(tags models.Tags)
	Update(tags models.Tags)
	Delete(tagId int)
	FindById(tagsId int) (tags models.Tags, err error)
	FindAll() []models.Tags
}
