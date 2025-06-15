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

// CreateTeacher создаёт нового преподавателя
func (h *Handler) CreateTeacher(c *gin.Context) {
    var teacher models.Teacher
    if err := c.ShouldBindJSON(&teacher); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Create(&teacher).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, teacher)
}

// UpdateTeacher обновляет данные преподавателя
func (h *Handler) UpdateTeacher(c *gin.Context) {
    id := c.Param("id")
    var teacher models.Teacher
    if err := h.DB.First(&teacher, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
        return
    }
    if err := c.ShouldBindJSON(&teacher); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Save(&teacher).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, teacher)
}

// DeleteTeacher удаляет преподавателя
func (h *Handler) DeleteTeacher(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Teacher{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted"})
}

// Аналогичные методы для Subject, Group и TeacherSubjectGroup
// GetSubjects
func (h *Handler) GetSubjects(c *gin.Context) {
    var subjects []models.Subject
    if err := h.DB.Find(&subjects).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, subjects)
}

// CreateSubject
func (h *Handler) CreateSubject(c *gin.Context) {
    var subject models.Subject
    if err := c.ShouldBindJSON(&subject); err != nil {

        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Create(&subject).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, subject)

}
// UpdateSubject
func (h *Handler) UpdateSubject(c *gin.Context) {
    id := c.Param("id")
    var subject models.Subject
    if err := h.DB.First(&subject, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
        return
    }
    if err := c.ShouldBindJSON(&subject); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Save(&subject).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, subject)
}

// DeleteSubject
func (h *Handler) DeleteSubject(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Subject{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Subject deleted"})
}

// GetGroups
func (h *Handler) GetGroups(c *gin.Context) {
    var groups []models.Group
    if err := h.DB.Find(&groups).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, groups)
}

// CreateGroup
func (h *Handler) CreateGroup(c *gin.Context) {
    var group models.Group
    if err := c.ShouldBindJSON(&group); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Create(&group).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, group)
}

// UpdateGroup
func (h *Handler) UpdateGroup(c *gin.Context) {
    id := c.Param("id")
    var group models.Group
    if err := h.DB.First(&group, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
        return
    }
    if err := c.ShouldBindJSON(&group); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Save(&group).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, group)
}

// DeleteGroup
func (h *Handler) DeleteGroup(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.Group{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Group deleted"})
}

// GetTeacherSubjectGroups
func (h *Handler) GetTeacherSubjectGroups(c *gin.Context) {
    var tsgs []models.TeacherSubjectGroup
    if err := h.DB.Preload("Teacher").Preload("Subject").Preload("Group").Find(&tsgs).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tsgs)
}

// CreateTeacherSubjectGroup
func (h *Handler) CreateTeacherSubjectGroup(c *gin.Context) {
    var tsg models.TeacherSubjectGroup
    if err := c.ShouldBindJSON(&tsg); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Create(&tsg).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, tsg)
}

// UpdateTeacherSubjectGroup
func (h *Handler) UpdateTeacherSubjectGroup(c *gin.Context) {
    id := c.Param("id")
    var tsg models.TeacherSubjectGroup
    if err := h.DB.First(&tsg, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "TeacherSubjectGroup not found"})
        return
    }
    if err := c.ShouldBindJSON(&tsg); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }
    if err := h.DB.Save(&tsg).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tsg)
}

// DeleteTeacherSubjectGroup
func (h *Handler) DeleteTeacherSubjectGroup(c *gin.Context) {
    id := c.Param("id")
    if err := h.DB.Delete(&models.TeacherSubjectGroup{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "TeacherSubjectGroup deleted"})
}