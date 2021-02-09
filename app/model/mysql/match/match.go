package match

import (
	"fmt"
	"gorm.io/gorm"
	"match-go/app/model/mysql/common"
)

type SQL struct {
	master *gorm.DB
	table  string
}

type Model struct {
	ID           uint64 `gorm:"primarykey"`
	Name         string `gorm:"type:varchar(255);default:'';not null;comment:比赛名称"`
	Pic          string `gorm:"type:varchar(255);default:'';not null;comment:比赛图片"`
	Description  string `gorm:"type:varchar(255);default:'';not null;comment:比赛描述"`
	common.Model `gorm:"embedded"`
}

func (m *Model) TableName() string {
	return "match"
}

func NewMatchSQL(master *gorm.DB) *SQL {
	return &SQL{
		master: master,
		table:  "match",
	}
}

func (m *SQL) GetAll() []Model {
	res := make([]Model, 0)
	if err := m.master.Table(m.table).Find(&res).Error; err != nil {
		fmt.Println("error", err)
	}
	return res
}
