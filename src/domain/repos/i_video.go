package repos

import "famyun/src/domain/models"

type IVideo interface {
	GetById(int64) (*models.Video, error)
	FindByName(string) ([]*models.Video, error)
	Save(...*models.Video) error
	Update(*models.Video) error
}
