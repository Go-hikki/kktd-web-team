package parser

import (
    "fmt"
    "schedule-parser/internal/models"
    "strings"
    "github.com/xuri/excelize/v2"
)

type ParsedData struct {
    Teachers          map[string]bool
    Subjects          map[string]bool
    Groups            map[string]bool
    TeacherSubjectGroup []models.TeacherSubjectGroup
}

func ParseExcel(filePath string) (*ParsedData, error) {
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        return nil, fmt.Errorf("failed to open Excel file: %v", err)
    }
    defer f.Close()

    data := &ParsedData{
        Teachers:          make(map[string]bool),
        Subjects:          make(map[string]bool),
        Groups:            make(map[string]bool),
        TeacherSubjectGroup: []models.TeacherSubjectGroup{},
    }

    // Читаем лист с расписанием
    rows, err := f.GetRows("Основное расписание")
    if err != nil {
        return nil, fmt.Errorf("failed to read sheet: %v", err)
    }

    // Находим строку с группами
    var groupRow int
    for i, row := range rows {
        for _, cell := range row {
            if strings.Contains(cell, "Группа:") {
                groupRow = i
                break
            }
        }
        if groupRow != 0 {
            break
        }
    }

    if groupRow == 0 {
        return nil, fmt.Errorf("group row not found")
    }

    // Извлекаем группы
    groups := make(map[int]string)
    for col, cell := range rows[groupRow] {
        if strings.Contains(cell, "09.02.07") || strings.Contains(cell, "29.02") || strings.Contains(cell, "38.02") {
            groups[col] = cell
            data.Groups[cell] = true
        }
    }

    // Парсим строки с предметами и преподавателями
    for i := groupRow + 1; i < len(rows)-1; i += 2 {
        // Проверяем, что следующая строка (с преподавателями) существует
        if i+1 >= len(rows) {
            continue
        }

        for col, group := range groups {
            // Проверяем, что столбец существует в текущей строке
            if col >= len(rows[i]) || col >= len(rows[i+1]) {
                continue
            }

            subject := strings.TrimSpace(rows[i][col])
            teacher := strings.TrimSpace(rows[i+1][col])

            if subject != "" && teacher != "" && !strings.Contains(subject, "КЛАССНЫЙ ЧАС") {
                data.Subjects[subject] = true
                data.Teachers[teacher] = true
                data.TeacherSubjectGroup = append(data.TeacherSubjectGroup, models.TeacherSubjectGroup{
                    Teacher: models.Teacher{FullName: teacher},
                    Subject: models.Subject{Name: subject},
                    Group:   models.Group{Name: group},
                })
            }
        }
    }

    return data, nil
}