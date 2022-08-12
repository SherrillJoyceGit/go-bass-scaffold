package dao

import "github.com/SherrillJoyceGit/go-bass-scaffold/model"

type FishDao struct {
	//db *gorm.DB
}

func NewFishDao() (dao *FishDao) {
	return &FishDao{}
}

func (dao *FishDao) HelloPostFish(p *model.CastBaitParam) (*model.Fish, error) {
	return &model.Fish{
		FishName: p.FishName,
		Feeling:  "I am " + p.BaitStatus + "! And Happy !",
	}, nil
}
