package models

import (
	"errors"
	"uni_app/database"
)

type UserStatus string

const (
	USER_STATUS_ACTIVE   UserStatus = "active"
	USER_STATUS_INACTIVE UserStatus = "inactive"
)

type Users []*User

type User struct {
	database.Model
	UserName      string        `gorm:"uniqueIndex;not null" json:"username,omitempty"`
	FirstName     string        `gorm:"not null" json:"first_name,omitempty"`
	LastName      string        `gorm:"not null" json:"last_name,omitempty"`
	Number        string        `gorm:"uniqueIndex;not null" json:"number,omitempty"`
	PersonalCode  string        `json:"personal_code,omitempty"`
	DegreeLevel   DegreeLevel   `gorm:"not null" json:"degree_level_id,omitempty"`
	MajorID       database.PID  `gorm:"not null" json:"major_id,omitempty"`
	Major         Major         `gorm:"foreignKey:MajorID" json:"major,omitempty"`
	UniID         database.PID  `json:"uni_id,omitempty"`
	Uni           Uni           `json:"uni,omitempty" gorm:"foreignKey:UniID"`
	NationalCode  *string       `gorm:"uniqueIndex" json:"national_code,omitempty"`
	NominatedByID *database.PID `json:"nominated_by_id,omitempty"`
	NominatedBy   *User         `gorm:"foreignKey:NominatedByID;constraint:OnDelete:SET NULL;" json:"nominated_by,omitempty"`
	Email         string        `gorm:"uniqueIndex;" json:"email,omitempty"`
	Password      string        `gorm:"not null" json:"-,omitempty"`
	Token         Token         `json:"token,omitempty" gorm:"polymorphic:Owner;"`
	Status        UserStatus    `gorm:"default:'active'" json:"status,omitempty"`
	UserRoles     []*UserRole   `json:"user_roles,omitempty"`
	Ratings       []Rating      `json:"ratings,omitempty" gorm:"foreignKey:UserID"`
}

type FetchUserRequest struct {
	FetchRequest
	DegreeLevel  DegreeLevel  `json:"degree_level_id,omitempty"`
	MajorID      database.PID `json:"major_id,omitempty"`
	UniID        database.PID `json:"uni_id,omitempty"`
	NationalCode string       `json:"national_code,omitempty"`
	TeacherCode  string       `json:"teacher_code,omitempty"`
	Email        string       `json:"email,omitempty"`
	Number       string       `json:"number,omitempty"`
	PersonalCode string       `json:"personal_code,omitempty"`
}

type UserRegisterRequest struct {
	UserName     string       `json:"username"`
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	Number       string       `json:"number"`
	PersonalCode string       `json:"personal_code"`
	DegreeLevel  DegreeLevel  `json:"degree_level"`
	MajorID      database.PID `json:"major_id"`
	UniID        database.PID `json:"uni_id"`
	NationalCode string       `json:"national_code"`
	Email        string       `json:"email"`
	Password     string       `json:"password"`
}

func (request *UserRegisterRequest) IsValid() error {
	// Validate all required fields
	if request.UserName == "" {
		return errors.New("username is required")
	}
	if request.FirstName == "" {
		return errors.New("first name is required")
	}
	if request.LastName == "" {
		return errors.New("last name is required")
	}
	if request.Number == "" {
		return errors.New("number is required")
	}
	// if request.PersonalCode == "" {
	// 	return errors.New("personal code is required")
	// }

	if request.DegreeLevel != DegreeLevelBachelor && request.DegreeLevel != DegreeLevelMaster && request.DegreeLevel != DegreeLevelPhD {
		return errors.New("degree level is invalid")
	}

	if !request.MajorID.IsValid() {
		return errors.New("major ID is required")
	}

	if !request.UniID.IsValid() {
		return errors.New("university ID is required")
	}

	if request.NationalCode == "" {
		return errors.New("national code is required")
	}

	if request.Email == "" {
		return errors.New("email is required")
	}

	if request.Password == "" {
		return errors.New("password is required")
	}
	return nil
}

type UserLoginRequst struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func UserAcceptIncludes() []string {
	return []string{
		"DegreeLevel",
		"DegreeMajor",
		"DegreeUni",
		"NominatedBy",
		"Token",
		"UserRoles",
		"Ratings",
	}
}
