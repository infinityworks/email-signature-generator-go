package server

import (
	"errors"
	"io"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/go-playground/validator.v9"
	

	"github.com/infinityworks/email-signature-generator/handlers"
)

type api struct {
	server *echo.Echo
}

type TemplateRegistry struct {
	templates map[string]*template.Template
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "base.html", data)
}

func NewApi() *api {
	e := echo.New()

	templates := make(map[string]*template.Template)
	templates["signature.html"] = template.Must(template.ParseFiles("views/signature.html", "views/base.html"))
	templates["form.html"] = template.Must(template.ParseFiles("views/form.html", "views/base.html"))

	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Secure())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Routes
	e.GET("/health", handlers.Health)

	e.GET("/", handlers.GetForm)
	e.POST("/", handlers.PostForm)

	return &api{
		server: e,
	}
}

func (a *api) Start(port string) {
	a.server.Logger.Fatal(a.server.Start(port))
}
