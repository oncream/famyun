package models

import "time"

type Video struct {
	Id           int64     `json:"id"`          //唯一表示
	CnName       string    `json:"cnName"`      //中文名字
	EnName       string    `json:"enName"`      //英文名字
	Synopsis     string    `json:"synopsis"`    //剧情介绍
	DirectorId   int64     `json:"-"`           //导演ID
	PerformerIds []int64   `json:"-"`           //演员ID
	Cover        string    `json:"cover"`       //封面链接
	Address      string    `json:"address"`     //片源
	ReleaseTime  time.Time `json:"releaseTime"` //上映时间
	DbTime       DBTime    `xorm:"extends"`
}

func NewVideo(attrs ...VideoAttrFunc) *Video {
	movie := &Video{}
	VideoAttrFuncs(attrs).apply(movie)
	return movie
}
