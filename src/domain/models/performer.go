package models

import "time"

type Gender int64

const Male = Gender(1)
const FeMale = Gender(2)

type Performer struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`   //姓名
	Gender   Gender    `json:"gender"` //性别
	Birthday time.Time `json:"-"`      //出生日期
	Photo    string    `json:"avatar"` //照片
	Age      int64     `json:"age"`    //年龄
	DbTime   DBTime    `xorm:"extends"`
}

func NewPerformer(cfg ...performerConfiguration) *Performer {
	performer := &Performer{}
	performerConfigurations(cfg).apply(performer)
	return performer
}

type performerConfiguration func(p *Performer)
type performerConfigurations []performerConfiguration

func (cfg performerConfigurations) apply(p *Performer) {
	for _, c := range cfg {
		c(p)
	}
}
func WithPerformerName(arg string) performerConfiguration {
	return func(p *Performer) {
		p.Name = arg
	}
}

func WithPerformerGender(arg Gender) performerConfiguration {
	return func(p *Performer) {
		p.Gender = arg
	}
}

func WithPerformerBirthday(arg time.Time) performerConfiguration {
	return func(p *Performer) {
		p.Birthday = arg
	}
}

func WithPerformerPhoto(arg string) performerConfiguration {
	return func(p *Performer) {
		p.Photo = arg
	}
}

func WithPerformerAge(arg int64) performerConfiguration {
	return func(p *Performer) {
		p.Age = arg
	}
}
