package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m1sty32/sdo/pkg/courses"
)

type CourseHandler struct {
	CourseService *courses.CourseService
}

// @Summary Create a new course
// @Description Create a new course with title, description, and optional teacher_id
// @Tags Courses
// @Accept json
// @Produce json
// @Param course body courses.CreateCourseRequest true "Course data"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /courses [post]
func (h *CourseHandler) CreateCourse(c *fiber.Ctx) error {
	var req courses.CreateCourseRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	err := h.CourseService.CreateCourse(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "OK"})
}

// @Summary Get all courses
// @Description Retrieve a list of all courses
// @Tags Courses
// @Produce json
// @Success 200 {array} models.Course
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /courses [get]
func (h *CourseHandler) GetCourses(c *fiber.Ctx) error {
	courses, err := h.CourseService.GetCourses(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(courses)
}
