package models

import (
	"uni_app/database"
)

// UserMode ...
type UserMode uint

const (
	// USER_STATUS_ACTIVE Active
	USER_STATUS_ACTIVE UserMode = iota + 1
	// PENDING ...
	PENDING
	// USER_STATUS_INACTIVE Inactive
	USER_STATUS_INACTIVE
)

type Users []*User

type User struct {
	database.Model
	UserName      string        `gorm:"uniqueIndex;not null" json:"username,omitempty"`
	FirstName     string        `gorm:"not null" json:"first_name,omitempty"`
	LastName      string        `gorm:"not null" json:"last_name,omitempty"`
	Number        string        `gorm:"uniqueIndex;not null" json:"number,omitempty"`
	PersonalCode  string        `json:"personal_code,omitempty"`
	TeacherCode   string        `gorm:"size:50;not null;unique" json:"teacher_code,omitempty"`
	DegreeLevelID uint          `gorm:"not null" json:"degree_level_id,omitempty"`
	DegreeLevel   DegreeLevel   `gorm:"foreignKey:DegreeLevelID" json:"degree_level,omitempty"`
	DegreeMajorID uint          `gorm:"not null" json:"degree_major_id,omitempty"`
	DegreeMajor   Major         `gorm:"foreignKey:DegreeMajorID" json:"degree_major,omitempty"`
	DegreeUniID   uint          `json:"degree_uni_id,omitempty"`
	DegreeUni     Uni           `json:"degree_uni,omitempty"`
	NationalCode  *string       `gorm:"uniqueIndex" json:"national_code,omitempty"`
	NominatedByID *database.PID `json:"nominated_by_id,omitempty"`
	NominatedBy   *User         `gorm:"foreignKey:NominatedByID;constraint:OnDelete:SET NULL;" json:"nominated_by,omitempty"`
	Email         string        `gorm:"uniqueIndex;" json:"email,omitempty"`
	Password      string        `gorm:"not null" json:"-,omitempty"`
	Token         Token         `json:"token,omitempty" gorm:"polymorphic:Owner;"`
	Status        UserMode      `gorm:"default:1" json:"status,omitempty"`
	UserRoles     []*UserRole   `json:"user_roles,omitempty"`
}

type UserLoginRequst struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}
