package models

import "uni_app/database"

type StudentCurrentLessons []*StudentCurrentLesson

type StudentCurrentLesson struct {
	database.Model
	StudentID   database.PID `json:"student_id,omitempty"`
	Student     *Student     `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;" json:"student,omitempty"`
	LessonID    database.PID `json:"lesson_id,omitempty"`
	Lesson      *Lesson      `gorm:"foreignKey:LessonID;constraint:OnDelete:CASCADE;" json:"lesson,omitempty"`
	TermID      database.PID `json:"term_id,omitempty"`
	Term        *Term        `gorm:"foreignKey:TermID;constraint:OnDelete:CASCADE;" json:"term,omitempty"`
	TeacherID   database.PID `json:"teacher_id,omitempty"`
	Teacher     *User        `gorm:"foreignKey:TeacherID;constraint:OnDelete:SET NULL;" json:"teacher,omitempty"`
	Score       *float64     `json:"score,omitempty"`
	Status      string       `gorm:"default:'active'" json:"status,omitempty"` // active, dropped, completed
	Description string       `json:"description,omitempty"`
	ClassTimes  []ClassTime  `json:"class_times,omitempty" gorm:"foreignKey:StudentCurrentLessonID;constraint:OnDelete:CASCADE;"`
}

type FetchStudentCurrentLessonRequest struct {
	FetchRequest
	StudentID database.PID `json:"student_id" query:"student_id"`
	TermID    database.PID `json:"term_id" query:"term_id"`
	LessonID  database.PID `json:"lesson_id" query:"lesson_id"`
	TeacherID database.PID `json:"teacher_id" query:"teacher_id"`
	Status    string       `json:"status" query:"status"`
}

func StudentCurrentLessonAcceptIncludes() []string {
	return []string{
		"Student",
		"Lesson",
		"Term",
		"Teacher",
		"ClassTimes",
	}
}
