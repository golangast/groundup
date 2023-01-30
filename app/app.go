

package main

	import (
		"fmt"
		"html/template"
		"io"
		"net/http"
		_ "net/http/pprof"
		"os"
		"path/filepath"
		_"time"
		"context"
		"strings"
		."app/db"
."app/getdata"
. "app/db/createuser"
. "app/db/createusers"
//#import
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
	
	func main() {
		Createuser() 
Createusers() 
//#createdb
		e := echo.New()
		t, err := ParseDirectory("templates/")
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
		// Start server
		e.Logger.Fatal(e.Start(":3000"))
		
	
		
		
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
	


func Routes(e *echo.Echo) {
	e.GET("/", Home)
	e.GET("/route/:routes", List)
	
}

func Home(c echo.Context) error {

	return c.Render(http.StatusOK, "home.html", map[string]interface{}{})

}

func List(c echo.Context) error {
	var data []Data
	routes := c.Param("routes")
	nospaceroutes := strings.ReplaceAll(routes, " ", "")
	nospaceroutesnoslash := strings.ReplaceAll(nospaceroutes, "/", "")
	db, err := DbConnection()
	if err != nil {
		fmt.Println(err)
	}
	var exists bool    
	type Data struct {
		ID     string "param:'id' query:'id' form:'id'"
		Datas    string "param:'datas' query:'datas' form:'datas'"
	}

	d:=Data{}
	const create string = "CREATE TABLE IF NOT EXISTS data ( id INTEGER NOT NULL PRIMARY KEY, datas DATETIME NOT NULL);"

	if _, err := db.Exec(create); err != nil {
		fmt.Println(err)
	   }

	   stmts := db.QueryRowContext(context.Background(), "SELECT EXISTS(SELECT 1 FROM data WHERE datas=?)", d.Datas)
	err = stmts.Scan(&exists)
	if err != nil {
		fmt.Println(err)
	}

	//prepare the statement to ensure no sql injection
	stmt, err := db.Prepare("INSERT INTO data(datas) VALUES(?)")
	if err != nil {
		fmt.Println(err)
	}

	//actually make the execution of the query
	_, err = stmt.Exec(d.Datas)
	if err != nil {
		fmt.Println(err)
	}

	data=Getalldata(db)
	

	return c.Render(http.StatusOK, nospaceroutesnoslash+".html", map[string]interface{}{
		"data":data,
	})

}


