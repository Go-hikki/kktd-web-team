package main

   import (
       "fmt"
       "github.com/gofiber/fiber/v2"
       "github.com/sirupsen/logrus"
       "github.com/swaggo/fiber-swagger"
       "github.com/m1sty32/sdo/pkg/api"
       "github.com/m1sty32/sdo/pkg/auth"
       "github.com/m1sty32/sdo/pkg/config"
       "github.com/m1sty32/sdo/pkg/db"
       "github.com/m1sty32/sdo/pkg/middlewares"
       "github.com/m1sty32/sdo/pkg/courses"
       _ "github.com/m1sty32/sdo/docs"
   )

   // @title SDO API
   // @version 1.0
   // @description API for Student Data Organizer
   // @host localhost:8080
   // @BasePath /
   func main() {
       cfg, err := config.LoadConfig()
       if err != nil {
           logrus.Fatal(err)
       }

       dbConn, err := db.NewPostgres(cfg)
       if err != nil {
           logrus.Fatal(err)
       }

       if err := db.RunMigrations(dbConn); err != nil {
           logrus.Fatal(err)
       }

       redisConn, err := db.NewRedis(cfg)
       if err != nil {
           logrus.Fatal(err)
       }

       authService := auth.NewAuthService(dbConn, redisConn, cfg.SecretKey)
       handler := api.Handler{AuthService: authService}
       courseService := courses.NewCourseService(dbConn)
       courseHandler := api.CourseHandler{CourseService: courseService}

       app := fiber.New()

       app.Get("/health", func(c *fiber.Ctx) error {
           return c.JSON(fiber.Map{"status": "ok"})
       })

       app.Get("/swagger/*", fiberSwagger.WrapHandler)

       // Auth routes
       authGroup := app.Group("/auth")
       authGroup.Post("/register", handler.Register)
       authGroup.Post("/login", handler.Login)

       // Protected routes
       protectedGroup := app.Group("/protected", middlewares.JWTProtected(cfg.SecretKey))
       protectedGroup.Get("/profile", func(c *fiber.Ctx) error {
           return c.JSON(fiber.Map{
               "user_id": c.Locals("user_id"),
               "email":   c.Locals("email"),
               "role":    c.Locals("role"),
           })
       })

       // Courses routes
       coursesGroup := app.Group("/courses", middlewares.JWTProtected(cfg.SecretKey))
       coursesGroup.Post("/", courseHandler.CreateCourse)
       coursesGroup.Get("/", courseHandler.GetCourses)

       logrus.Fatal(app.Listen(fmt.Sprintf(":%d", cfg.Server.Port)))
   }