package model

import "time"

type UserRole struct { 
	ID    int64 `gorm:"primaryKey"`
	RoleID    int64 `gorm:"index"`
	UserID    int64 `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// TableName overrides the table name used by User to `profiles`
func (User UserRole) TableName() string {
  return "user_role"
}