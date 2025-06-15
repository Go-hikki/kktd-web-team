package models

import (
    "gorm.io/gorm"
)

// Teacher представляет преподавателя
type Teacher struct {
    gorm.Model
    FullName string `gorm:"unique;not null"`
}

// Subject представляет учебный предмет
type Subject struct {
    gorm.Model
    Name string `gorm:"unique;not null"`
}

// Group представляет учебную группу
type Group struct {
    gorm.Model
    Name string `gorm:"unique;not null"`
}

// TeacherSubjectGroup связывает преподавателя, предмет и группу
type TeacherSubjectGroup struct {
    gorm.Model
    TeacherID uint
    SubjectID uint
    GroupID   uint
    Teacher   Teacher `gorm:"foreignKey:TeacherID"`
    Subject   Subject `gorm:"foreignKey:SubjectID"`
    Group     Group   `gorm:"foreignKey:GroupID"`
}