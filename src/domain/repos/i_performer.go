package repos

import "famyun/src/domain/models"

type IPerformer interface {
	GetById(int64) (*models.Performer, error)
	FindById(...int64) ([]*models.Performer, error)
	FindByName(string) ([]*models.Performer, error)
	Save(...*models.Performer) error
	Update(*models.Performer) error
}
