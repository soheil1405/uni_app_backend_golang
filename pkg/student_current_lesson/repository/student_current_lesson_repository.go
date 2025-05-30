package repository

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type StudentCurrentLessonRepository struct {
	db *gorm.DB
}

func NewStudentCurrentLessonRepository(db *gorm.DB) *StudentCurrentLessonRepository {
	return &StudentCurrentLessonRepository{db: db}
}

func (r *StudentCurrentLessonRepository) Create(studentCurrentLesson *models.StudentCurrentLesson) error {
	return r.db.Create(studentCurrentLesson).Error
}

func (r *StudentCurrentLessonRepository) Update(studentCurrentLesson *models.StudentCurrentLesson) error {
	return r.db.Save(studentCurrentLesson).Error
}

func (r *StudentCurrentLessonRepository) Delete(id uint) error {
	return r.db.Delete(&models.StudentCurrentLesson{}, id).Error
}

func (r *StudentCurrentLessonRepository) FindByID(id uint) (*models.StudentCurrentLesson, error) {
	var studentCurrentLesson models.StudentCurrentLesson
	err := r.db.Preload("Student").Preload("Lesson").Preload("Term").Preload("Teacher").Preload("ClassTimes").
		First(&studentCurrentLesson, id).Error
	if err != nil {
		return nil, err
	}
	return &studentCurrentLesson, nil
}

func (r *StudentCurrentLessonRepository) FindByTeacherID(teacherID uint) ([]*models.StudentCurrentLesson, error) {
	var studentCurrentLessons []*models.StudentCurrentLesson
	err := r.db.Preload("Student").Preload("Lesson").Preload("Term").Preload("ClassTimes").
		Where("teacher_id = ?", teacherID).Find(&studentCurrentLessons).Error
	if err != nil {
		return nil, err
	}
	return studentCurrentLessons, nil
}

func (r *StudentCurrentLessonRepository) FindByStudentID(studentID uint) ([]*models.StudentCurrentLesson, error) {
	var studentCurrentLessons []*models.StudentCurrentLesson
	err := r.db.Preload("Lesson").Preload("Term").Preload("Teacher").Preload("ClassTimes").
		Where("student_id = ?", studentID).Find(&studentCurrentLessons).Error
	if err != nil {
		return nil, err
	}
	return studentCurrentLessons, nil
}

func (r *StudentCurrentLessonRepository) FindByTermID(termID uint) ([]*models.StudentCurrentLesson, error) {
	var studentCurrentLessons []*models.StudentCurrentLesson
	err := r.db.Preload("Student").Preload("Lesson").Preload("Teacher").Preload("ClassTimes").
		Where("term_id = ?", termID).Find(&studentCurrentLessons).Error
	if err != nil {
		return nil, err
	}
	return studentCurrentLessons, nil
}

func (r *StudentCurrentLessonRepository) FindByTeacherAndTerm(teacherID, termID uint) ([]*models.StudentCurrentLesson, error) {
	var studentCurrentLessons []*models.StudentCurrentLesson
	err := r.db.Preload("Student").Preload("Lesson").Preload("ClassTimes").
		Where("teacher_id = ? AND term_id = ?", teacherID, termID).Find(&studentCurrentLessons).Error
	if err != nil {
		return nil, err
	}
	return studentCurrentLessons, nil
}

func (r *StudentCurrentLessonRepository) FindByTeacherAndLesson(teacherID, lessonID uint) ([]*models.StudentCurrentLesson, error) {
	var studentCurrentLessons []*models.StudentCurrentLesson
	err := r.db.Preload("Student").Preload("Term").Preload("ClassTimes").
		Where("teacher_id = ? AND lesson_id = ?", teacherID, lessonID).Find(&studentCurrentLessons).Error
	if err != nil {
		return nil, err
	}
	return studentCurrentLessons, nil
}

func (r *StudentCurrentLessonRepository) UpdateScore(id uint, score float64) error {
	studentCurrentLesson, err := r.FindByID(id)
	if err != nil {
		return err
	}

	studentCurrentLesson.Score = &score
	return r.Update(studentCurrentLesson)
}

func (r *StudentCurrentLessonRepository) UpdateStatus(id uint, status string) error {
	studentCurrentLesson, err := r.FindByID(id)
	if err != nil {
		return err
	}

	studentCurrentLesson.Status = status
	return r.Update(studentCurrentLesson)
}

func (r *StudentCurrentLessonRepository) GetTeacherSchedule(teacherID uint) ([]*models.StudentCurrentLesson, error) {
	var studentCurrentLessons []*models.StudentCurrentLesson
	err := r.db.Preload("Student").Preload("Lesson").Preload("Term").Preload("ClassTimes").
		Where("teacher_id = ?", teacherID).
		Order("term_id").
		Find(&studentCurrentLessons).Error
	if err != nil {
		return nil, err
	}
	return studentCurrentLessons, nil
}

func (r *StudentCurrentLessonRepository) GetTeacherStudents(teacherID uint) ([]*models.Student, error) {
	var students []*models.Student
	err := r.db.Joins("JOIN student_current_lessons ON student_current_lessons.student_id = students.id").
		Where("student_current_lessons.teacher_id = ?", teacherID).
		Distinct().
		Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}

func (r *StudentCurrentLessonRepository) GetTeacherClassDetails(teacherID uint, lessonID uint) (*models.StudentCurrentLesson, error) {
	var studentCurrentLesson models.StudentCurrentLesson
	err := r.db.Preload("Student").Preload("Lesson").Preload("Term").Preload("ClassTimes").
		Where("teacher_id = ? AND lesson_id = ?", teacherID, lessonID).
		First(&studentCurrentLesson).Error
	if err != nil {
		return nil, err
	}
	return &studentCurrentLesson, nil
}
