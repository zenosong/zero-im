// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameAutoMessage = "auto_messages"

// AutoMessage mapped from table <auto_messages>
type AutoMessage struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Type      string    `gorm:"column:type" json:"type"`
	Content   string    `gorm:"column:content" json:"content"`
	GroupID   int64     `gorm:"column:group_id" json:"group_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName AutoMessage's table name
func (*AutoMessage) TableName() string {
	return TableNameAutoMessage
}
