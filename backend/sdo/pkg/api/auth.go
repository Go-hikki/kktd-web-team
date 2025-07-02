package api

import (
    "errors"
    "github.com/gofiber/fiber/v2"
    "github.com/m1sty32/sdo/pkg/auth"
)

// Handler holds dependencies for HTTP handlers
type Handler struct {
    AuthService *auth.AuthService
}

// Register handles user registration
// @Summary Register a new user
// @Description Register a new user with email, password, and role
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.RegisterRequest true "User registration data"
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Failure 409 {string} string "Conflict"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/register [post]
func (h *Handler) Register(c *fiber.Ctx) error {
    var req auth.RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request",
        })
    }

    err := h.AuthService.Register(c.Context(), req)
    if err != nil {
        if errors.Is(err, errors.New("user with this email already exists")) {
            return c.Status(fiber.StatusConflict).JSON(fiber.Map{
                "error": err.Error(),
            })
        }
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "OK",
    })
}

// Login handles user authentication
// @Summary Login user
// @Description Login user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body auth.LoginRequest true "User login data"
// @Success 200 {object} auth.LoginResponse
// @Failure 400 {string} string "Bad Request"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /auth/login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
    var req auth.LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Invalid request",
        })
    }

    resp, err := h.AuthService.Login(c.Context(), req)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(resp)
}
