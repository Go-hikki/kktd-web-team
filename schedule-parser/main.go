package main

import (
    "fmt"
    "schedule-parser/internal/api"
    "schedule-parser/internal/config"
    "schedule-parser/internal/database"
    "schedule-parser/internal/models"
    "schedule-parser/internal/parser"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func main() {
    // Загрузка конфигурации
    cfg, err := config.LoadConfig()
    if err != nil {
        fmt.Printf("Error loading config: %v\n", err)
        return
    }

    // Подключение к БД
    db, err := database.ConnectDB(cfg)
    if err != nil {
        fmt.Printf("Error connecting to database: %v\n", err)
        return
    }

    // Выполнение миграций
    if err := database.RunMigrations(db); err != nil {
        fmt.Printf("Error running migrations: %v\n", err)
        return
    }

    // Парсинг Excel-файла
    data, err := parser.ParseExcel(cfg.ExcelFilePath)
    if err != nil {
        fmt.Printf("Error parsing Excel file: %v\n", err)
        return
    }

    // Сохранение данных в БД
    if err := saveDataToDB(db, data); err != nil {
        fmt.Printf("Error saving data to database: %v\n", err)
        return
    }

    fmt.Println("Data successfully parsed and saved to database!")

    // Запуск HTTP-сервера
    r := gin.Default()
    handler := api.NewHandler(db)

    // Определение маршрутов
    r.GET("/teachers", handler.GetTeachers)
    r.GET("/subjects", handler.GetSubjects)
    r.GET("/groups", handler.GetGroups)
    r.GET("/teacher-subject-groups", handler.GetTeacherSubjectGroups)

    // Запуск сервера на порту 8080
    if err := r.Run(":8080"); err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}

func saveDataToDB(db *gorm.DB, data *parser.ParsedData) error {
    // Сохранение преподавателей
    for teacherName := range data.Teachers {
        teacher := models.Teacher{FullName: teacherName}
        if err := db.FirstOrCreate(&teacher, models.Teacher{FullName: teacherName}).Error; err != nil {
            return fmt.Errorf("failed to save teacher %s: %v", teacherName, err)
        }
    }

    // Сохранение предметов
    for subjectName := range data.Subjects {
        subject := models.Subject{Name: subjectName}
        if err := db.FirstOrCreate(&subject, models.Subject{Name: subjectName}).Error; err != nil {
            return fmt.Errorf("failed to save subject %s: %v", subjectName, err)
        }
    }

    // Сохранение групп
    for groupName := range data.Groups {
        group := models.Group{Name: groupName}
        if err := db.FirstOrCreate(&group, models.Group{Name: groupName}).Error; err != nil {
            return fmt.Errorf("failed to save group %s: %v", groupName, err)
        }
    }

    // Сохранение связей
    for _, tsg := range data.TeacherSubjectGroup {
        var teacher models.Teacher
        var subject models.Subject
        var group models.Group

        // Находим ID преподавателя, предмета и группы
        if err := db.Where("full_name = ?", tsg.Teacher.FullName).First(&teacher).Error; err != nil {
            return fmt.Errorf("failed to find teacher %s: %v", tsg.Teacher.FullName, err)
        }
        if err := db.Where("name = ?", tsg.Subject.Name).First(&subject).Error; err != nil {
            return fmt.Errorf("failed to find subject %s: %v", tsg.Subject.Name, err)
        }
        if err := db.Where("name = ?", tsg.Group.Name).First(&group).Error; err != nil {
            return fmt.Errorf("failed to find group %s: %v", tsg.Group.Name, err)
        }

        // Сохраняем связь
        tsgEntry := models.TeacherSubjectGroup{
            TeacherID: teacher.ID,
            SubjectID: subject.ID,
            GroupID:   group.ID,
        }
        if err := db.FirstOrCreate(&tsgEntry, models.TeacherSubjectGroup{
            TeacherID: teacher.ID,
            SubjectID: subject.ID,
            GroupID:   group.ID,
        }).Error; err != nil {
            return fmt.Errorf("failed to save teacher-subject-group: %v", err)
        }
    }

    return nil
}