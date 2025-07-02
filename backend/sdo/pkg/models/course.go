package models

import "time"

type Course struct {
    ID          int       `db:"id"`
    Title       string    `db:"title"`
    Description string    `db:"description"`
    TeacherID   *int      `db:"teacher_id"` // Nullable
    CreatedAt   time.Time `db:"created_at"`
    UpdatedAt   time.Time `db:"updated_at"`
}