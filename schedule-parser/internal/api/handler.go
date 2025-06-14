package api

import (
    "net/http"
    "schedule-parser/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type Handler struct {
    DB *gorm.DB
}

// NewHandler создаёт новый обработчик с подключением к БД
func NewHandler(db *gorm.DB) *Handler {
    return &Handler{DB: db}
}

// GetTeachers возвращает список всех преподавателей
func (h *Handler) GetTeachers(c *gin.Context) {
    var teachers []models.Teacher
    if err := h.DB.Find(&teachers).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, teachers)
}

// GetSubjects возвращает список всех предметов
func (h *Handler) GetSubjects(c *gin.Context) {
    var subjects []models.Subject
    if err := h.DB.Find(&subjects).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, subjects)
}

// GetGroups возвращает список всех групп
func (h *Handler) GetGroups(c *gin.Context) {
    var groups []models.Group
    if err := h.DB.Find(&groups).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, groups)
}

// GetTeacherSubjectGroups возвращает связи преподаватель-предмет-группа
func (h *Handler) GetTeacherSubjectGroups(c *gin.Context) {
    var tsgs []models.TeacherSubjectGroup
    if err := h.DB.Preload("Teacher").Preload("Subject").Preload("Group").Find(&tsgs).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tsgs)
}