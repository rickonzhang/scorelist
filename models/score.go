package models

import (
	"scorelist/dao"
)

type Score struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Score string `json:"score"`
	Status bool `json:"status"`
	Like int `json:like`
}

func CreateAScore(score *Score)(err error){
	err = dao.DB.Create(&score).Error
	return
}

func GetAllScore() (scoreList []*Score, err error){
	if err = dao.DB.Find(&scoreList).Error; err != nil{
		return nil, err
	}
	return
}

func GetAScore(id string)(score *Score, err error){
	score = new(Score)
	if err = dao.DB.Debug().Where("id=?", id).First(score).Error; err!=nil{
		return nil, err
	}
	return
}
//不带错误检查的单条查询
func GetAScore2(id string)(score *Score){
	score = new(Score)
	dao.DB.Where("id=?", id).First(score)
	return
}

func UpdateAScore(score *Score)(err error){
	err = dao.DB.Save(score).Error
	return
}

func DeleteAScore(id string)(err error){
	err = dao.DB.Where("id=?", id).Delete(&Score{}).Error
	return
}

func UpdateMsg(name string)(err error) {
	msg := Score{}
	err = dao.DB.Where("name=?",name).First(&msg).Error
	if msg.Like == 0{
		err = dao.DB.Model(&msg).Update("like",1).Error
	}else{
		err = dao.DB.Model(&msg).Update("like",0).Error
	}
	dao.DB.Save(&msg)
	return
}


