package entity

import (
	"time"
)

// PrimaryKeyID 自增主键ID
type PrimaryKeyID struct {
	ID uint64 `gorm:"primary_key;autoIncrement" json:"id"`
}

// SoftDelete 软删除
type SoftDelete struct {
	IsDelete  uint8     `gorm:"not null;default:0;comment:删除标识 0-未删除 1-已删除" json:"is_delete"`
	DeletedAt time.Time `gorm:"type:datetime(0);index;not null" json:"deleted_at"`
}

// Time 添加时间,修改时间
type Time struct {
	CreatedAt time.Time `gorm:"type:datetime(0);index;not null" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"type:datetime(0);not null" json:"updated_at,omitempty"`
}
