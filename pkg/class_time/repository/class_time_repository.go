package class_time

import (
	"uni_app/models"

	"gorm.io/gorm"
)

type ClassTimeRepository struct {
	db *gorm.DB
}

func NewClassTimeRepository(db *gorm.DB) *ClassTimeRepository {
	return &ClassTimeRepository{db: db}
}

func (r *ClassTimeRepository) Create(classTime *models.ClassTime) error {
	return r.db.Create(classTime).Error
}

func (r *ClassTimeRepository) Update(classTime *models.ClassTime) error {
	return r.db.Save(classTime).Error
}

func (r *ClassTimeRepository) Delete(id uint) error {
	return r.db.Delete(&models.ClassTime{}, id).Error
}

func (r *ClassTimeRepository) FindByID(id uint) (*models.ClassTime, error) {
	var classTime models.ClassTime
	err := r.db.Preload("StudentCurrentLesson").First(&classTime, id).Error
	if err != nil {
		return nil, err
	}
	return &classTime, nil
}

func (r *ClassTimeRepository) FindByStudentCurrentLessonID(studentCurrentLessonID uint) ([]*models.ClassTime, error) {
	var classTimes []*models.ClassTime
	err := r.db.Where("student_current_lesson_id = ?", studentCurrentLessonID).Find(&classTimes).Error
	if err != nil {
		return nil, err
	}
	return classTimes, nil
}

func (r *ClassTimeRepository) FindByTeacherID(teacherID uint) ([]*models.ClassTime, error) {
	var classTimes []*models.ClassTime
	err := r.db.Joins("JOIN student_current_lessons ON student_current_lessons.id = class_times.student_current_lesson_id").
		Where("student_current_lessons.teacher_id = ?", teacherID).
		Preload("StudentCurrentLesson").
		Find(&classTimes).Error
	if err != nil {
		return nil, err
	}
	return classTimes, nil
}

func (r *ClassTimeRepository) FindByTeacherAndDay(teacherID uint, dayOfWeek int) ([]*models.ClassTime, error) {
	var classTimes []*models.ClassTime
	err := r.db.Joins("JOIN student_current_lessons ON student_current_lessons.id = class_times.student_current_lesson_id").
		Where("student_current_lessons.teacher_id = ? AND class_times.day_of_week = ?", teacherID, dayOfWeek).
		Preload("StudentCurrentLesson").
		Find(&classTimes).Error
	if err != nil {
		return nil, err
	}
	return classTimes, nil
}

func (r *ClassTimeRepository) FindSelectedByStudentCurrentLessonID(studentCurrentLessonID uint) (*models.ClassTime, error) {
	var classTime models.ClassTime
	err := r.db.Where("student_current_lesson_id = ? AND is_selected = ?", studentCurrentLessonID, true).
		First(&classTime).Error
	if err != nil {
		return nil, err
	}
	return &classTime, nil
}

func (r *ClassTimeRepository) UpdateSelectedTime(studentCurrentLessonID uint, classTimeID uint) error {
	// First, unselect all times for this lesson
	err := r.db.Model(&models.ClassTime{}).
		Where("student_current_lesson_id = ?", studentCurrentLessonID).
		Update("is_selected", false).Error
	if err != nil {
		return err
	}

	// Then, select the new time
	return r.db.Model(&models.ClassTime{}).
		Where("id = ?", classTimeID).
		Update("is_selected", true).Error
}

func (r *ClassTimeRepository) GetTeacherWeeklySchedule(teacherID uint) (map[int][]*models.ClassTime, error) {
	var classTimes []*models.ClassTime
	err := r.db.Joins("JOIN student_current_lessons ON student_current_lessons.id = class_times.student_current_lesson_id").
		Where("student_current_lessons.teacher_id = ?", teacherID).
		Preload("StudentCurrentLesson").
		Find(&classTimes).Error
	if err != nil {
		return nil, err
	}

	// Organize by day of week
	schedule := make(map[int][]*models.ClassTime)
	for _, ct := range classTimes {
		schedule[ct.DayOfWeek] = append(schedule[ct.DayOfWeek], ct)
	}

	return schedule, nil
}

func (r *ClassTimeRepository) GetTeacherClassroomSchedule(teacherID uint) (map[string][]*models.ClassTime, error) {
	var classTimes []*models.ClassTime
	err := r.db.Joins("JOIN student_current_lessons ON student_current_lessons.id = class_times.student_current_lesson_id").
		Where("student_current_lessons.teacher_id = ?", teacherID).
		Preload("StudentCurrentLesson").
		Find(&classTimes).Error
	if err != nil {
		return nil, err
	}

	// Organize by classroom
	schedule := make(map[string][]*models.ClassTime)
	for _, ct := range classTimes {
		key := ct.Building + "-" + ct.ClassroomNumber
		schedule[key] = append(schedule[key], ct)
	}

	return schedule, nil
}
