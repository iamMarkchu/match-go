package common

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Status    uint8          `gorm:"comment:状态"`
	Opr       string         `gorm:"comment:操作人"`
	CreatedAt time.Time      `gorm:"comment:创建时间"`
	UpdatedAt time.Time      `gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt `gorm:"index;comment:删新时间"`
}
