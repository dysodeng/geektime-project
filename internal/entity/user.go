package entity

// User 用户
type User struct {
	PrimaryKeyID
	Username string `gorm:"type:varchar(150);index:unique;not null;default:'';comment:用户名" json:"username"`
	Nickname string `gorm:"type:varchar(150);not null;default:'';comment:用户昵称" json:"nickname"`
	SoftDelete
	Time
}

func (User) TableName() string {
	return "user"
}
