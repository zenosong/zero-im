// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameChatSession = "chat_sessions"

// ChatSession mapped from table <chat_sessions>
type ChatSession struct {
	ID         int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID     int64 `gorm:"column:user_id" json:"user_id"`
	QueriedAt  int64 `gorm:"column:queried_at" json:"queried_at"`
	AcceptedAt int64 `gorm:"column:accepted_at" json:"accepted_at"`
	CanceledAt int64 `gorm:"column:canceled_at" json:"canceled_at"`
	BrokeAt    int64 `gorm:"column:broke_at" json:"broke_at"`
	GroupID    int64 `gorm:"column:group_id" json:"group_id"`
	AdminID    int64 `gorm:"column:admin_id" json:"admin_id"`
	Type       int32 `gorm:"column:type" json:"type"`
}

// TableName ChatSession's table name
func (*ChatSession) TableName() string {
	return TableNameChatSession
}
