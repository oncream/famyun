package models

import "time"

type VideoAttrFunc func(v *Video)
type VideoAttrFuncs []VideoAttrFunc

func WithVideoCnName(arg string) VideoAttrFunc {
	return func(v *Video) {
		v.CnName = arg
	}
}
func WithVideoEnName(arg string) VideoAttrFunc {
	return func(v *Video) {
		v.EnName = arg
	}
}

func WithVideoCover(arg string) VideoAttrFunc {
	return func(v *Video) {
		v.Cover = arg
	}
}

func WithVideoAddress(arg string) VideoAttrFunc {
	return func(v *Video) {
		v.Address = arg
	}
}

func WithVideoReleaseTime(arg time.Time) VideoAttrFunc {
	return func(v *Video) {
		v.ReleaseTime = arg
	}
}

func WithVideoDirectorId(arg int64) VideoAttrFunc {
	return func(v *Video) {
		v.DirectorId = arg
	}
}

func WithVideoPerformerIds(arg ...int64) VideoAttrFunc {
	return func(v *Video) {
		v.PerformerIds = arg
	}
}

func (this VideoAttrFuncs) apply(v *Video) {
	for _, f := range this {
		f(v)
	}
}
