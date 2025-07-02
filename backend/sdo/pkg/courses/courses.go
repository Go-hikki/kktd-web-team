package courses

import (
    "context"
    "time"
    "github.com/jmoiron/sqlx"
    "github.com/m1sty32/sdo/pkg/models"
)

type CourseService struct {
    db *sqlx.DB
}

func NewCourseService(db *sqlx.DB) *CourseService {
    return &CourseService{db: db}
}

type CreateCourseRequest struct {
    Title       string `json:"title"`
    Description string `json:"description"`
    TeacherID   *int   `json:"teacher_id"`
}

func (s *CourseService) CreateCourse(ctx context.Context, req CreateCourseRequest) error {
    course := models.Course{
        Title:       req.Title,
        Description: req.Description,
        TeacherID:   req.TeacherID,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }

    query := `INSERT INTO courses (title, description, teacher_id, created_at, updated_at) 
              VALUES (:title, :description, :teacher_id, :created_at, :updated_at)`
    _, err := s.db.NamedExecContext(ctx, query, course)
    return err
}

func (s *CourseService) GetCourses(ctx context.Context) ([]models.Course, error) {
    var courses []models.Course
    query := `SELECT id, title, description, teacher_id, created_at, updated_at FROM courses`
    err := s.db.SelectContext(ctx, &courses, query)
    return courses, err
}