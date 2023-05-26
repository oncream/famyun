package dao

import (
	"errors"
	"famyun/src/domain/models"
	"famyun/src/domain/repos"
	"fmt"
	"xorm.io/xorm"
)

type VideoRepo struct {
	engine *xorm.Engine
}

var _ repos.IVideo = (*VideoRepo)(nil)

func NewVideoRepo(engine *xorm.Engine) repos.IVideo {
	return &VideoRepo{engine}
}

func (v *VideoRepo) GetById(i int64) (*models.Video, error) {
	video := &models.Video{}
	has, err := v.engine.ID(i).Get(video)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New(fmt.Sprintf("ID:[%d]不存在", i))
	}
	return video, nil
}

func (v *VideoRepo) FindByName(s string) ([]*models.Video, error) {
	//todo 待完善
	return nil, nil
}

func (v *VideoRepo) Save(video ...*models.Video) error {
	_, err := v.engine.Insert(video)
	return err
}

func (v *VideoRepo) Update(video *models.Video) error {
	_, err := v.engine.Update(video)
	return err
}
