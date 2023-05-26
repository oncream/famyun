package dao

import (
	"famyun/src/domain/models"
	"famyun/src/domain/repos"
	"xorm.io/builder"
	"xorm.io/xorm"
)

type PerformerRepo struct {
	engine *xorm.Engine
}

func NewPerformerRepo(engine *xorm.Engine) repos.IPerformer {
	return &PerformerRepo{engine}
}

var _ repos.IPerformer = (*PerformerRepo)(nil)

func (p *PerformerRepo) GetById(i int64) (*models.Performer, error) {
	performer := &models.Performer{}
	_, err := p.engine.ID(i).Get(performer)
	return performer, err
}

func (p *PerformerRepo) FindById(i ...int64) ([]*models.Performer, error) {
	r := make([]*models.Performer, 0, len(i))
	return r, p.engine.Where(builder.And(builder.In("id", i))).Find(&r)
}

func (p *PerformerRepo) FindByName(s string) ([]*models.Performer, error) {
	return nil, nil
}

func (p *PerformerRepo) Save(performer ...*models.Performer) error {
	_, err := p.engine.Insert(performer)
	return err
}

func (p *PerformerRepo) Update(performer *models.Performer) error {
	_, err := p.engine.Update(performer)
	return err
}
