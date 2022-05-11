package serv

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"

	. "github.com/golangast/groundup/dashboard/dbsql/dbutil"
	. "github.com/golangast/groundup/dashboard/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

var err error

func Serv() {
	exists, err := Exists("db/urls.db")
	if err != nil {
		fmt.Println(err)
	}
	if !exists {
		if err := os.MkdirAll("db", os.ModeSticky|os.ModePerm); err != nil {
			fmt.Println("Directory(ies) successfully created with sticky bits and full permissions")
		} else {
			fmt.Println("Whoops, could not create directory(ies) because", err)
		}
		_, err := os.Create("db/database.db")
		if err != nil {
			fmt.Println(err)
		}
	}
	GenerateTable()
	e := echo.New()
	t, err := ParseDirectory("dashboard/templates/")
	if err != nil {
		fmt.Println(err)
	}

	renderer := &TemplateRenderer{
		templates: template.Must(t, err),
	}

	e.Renderer = renderer

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	Routes(e)

	// Route
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "SAMEORIGIN",
		HSTSMaxAge:            3600,
		ContentSecurityPolicy: "",
	}))
	e.Use(middleware.BodyLimit("3M"))
	e.IPExtractor = echo.ExtractIPDirect()
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(30)))
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":5002"))
}
func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func ParseDirectory(dirpath string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}
