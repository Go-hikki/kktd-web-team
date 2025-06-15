package main

import (
	"fmt"
	"schedule-parser/internal/api"
	"schedule-parser/internal/config"
	"schedule-parser/internal/database"
	"schedule-parser/internal/models"
	"schedule-parser/internal/parser"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}

	if err := database.RunMigrations(db); err != nil {
		fmt.Printf("Error running migrations: %v\n", err)
		return
	}

	data, err := parser.ParseExcel(cfg.ExcelFilePath)
	if err != nil {
		fmt.Printf("Error parsing Excel file: %v\n", err)
		return
	}

	if err := saveDataToDB(db, data); err != nil {
		fmt.Printf("Error saving data to database: %v\n", err)
		return
	}

	fmt.Println("Data successfully parsed and saved to database!")

	r := gin.Default()

	// Настройка CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"}, // Укажите порт фронтенда
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	handler := api.NewHandler(db)

	// Маршруты
	r.GET("/teachers", handler.GetTeachers)
	r.POST("/teachers", handler.CreateTeacher)
	r.PUT("/teachers/:id", handler.UpdateTeacher)
	r.DELETE("/teachers/:id", handler.DeleteTeacher)

	r.GET("/subjects", handler.GetSubjects)
	r.POST("/subjects", handler.CreateSubject)
	r.PUT("/subjects/:id", handler.UpdateSubject)
	r.DELETE("/subjects/:id", handler.DeleteSubject)

	r.GET("/groups", handler.GetGroups)
	r.POST("/groups", handler.CreateGroup)
	r.PUT("/groups/:id", handler.UpdateGroup)
	r.DELETE("/groups/:id", handler.DeleteGroup)

	r.GET("/teacher-subject-groups", handler.GetTeacherSubjectGroups)
	r.POST("/teacher-subject-groups", handler.CreateTeacherSubjectGroup)
	r.PUT("/teacher-subject-groups/:id", handler.UpdateTeacherSubjectGroup)
	r.DELETE("/teacher-subject-groups/:id", handler.DeleteTeacherSubjectGroup)

	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func saveDataToDB(db *gorm.DB, data *parser.ParsedData) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Сохранение преподавателей
	for teacherName := range data.Teachers {
		teacher := models.Teacher{FullName: teacherName}
		if err := tx.FirstOrCreate(&teacher, models.Teacher{FullName: teacherName}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to save teacher %s: %v", teacherName, err)
		}
	}

	// Сохранение предметов
	for subjectName := range data.Subjects {
		subject := models.Subject{Name: subjectName}
		if err := tx.FirstOrCreate(&subject, models.Subject{Name: subjectName}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to save subject %s: %v", subjectName, err)
		}
	}

	// Сохранение групп
	for groupName := range data.Groups {
		group := models.Group{Name: groupName}
		if err := tx.FirstOrCreate(&group, models.Group{Name: groupName}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to save group %s: %v", groupName, err)
		}
	}

	// Сохранение связей
	for _, tsg := range data.TeacherSubjectGroup {
		var teacher models.Teacher
		var subject models.Subject
		var group models.Group

		if err := tx.Where("full_name = ?", tsg.Teacher.FullName).First(&teacher).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to find teacher %s: %v", tsg.Teacher.FullName, err)
		}
		if err := tx.Where("name = ?", tsg.Subject.Name).First(&subject).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to find subject %s: %v", tsg.Subject.Name, err)
		}
		if err := tx.Where("name = ?", tsg.Group.Name).First(&group).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to find group %s: %v", tsg.Group.Name, err)
		}

		tsgEntry := models.TeacherSubjectGroup{
			TeacherID: teacher.ID,
			SubjectID: subject.ID,
			GroupID:   group.ID,
		}
		if err := tx.FirstOrCreate(&tsgEntry, models.TeacherSubjectGroup{
			TeacherID: teacher.ID,
			SubjectID: subject.ID,
			GroupID:   group.ID,
		}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to save teacher-subject-group: %v", err)
		}
	}

	return tx.Commit().Error
}
