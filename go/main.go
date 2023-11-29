package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
	"todo/internal/activity"
	"todo/internal/shared/database"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/utils"
)

var cfg *AppConfig

func main() {
	cfg = ParseConfig()

	if cfg == nil {
		log.Fatal("[x] app config required")
	}

	// init database
	db := initDatabase()

	app := fiber.New(fiber.Config{
		AppName:     cfg.AppName,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(cache.New(cache.Config{
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return utils.CopyString(c.OriginalURL())
		},
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Status(http.StatusOK)
		return c.JSON(map[string]string{
			"message": "hello world",
		})
	})

	// init activity
	activity := activity.New(&activity.Dependency{
		DB: db,
	})
	activity.Route(app)

	// start
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	var errChan = make(chan error, 1)

	go func() {
		err := app.Listen(":3030")
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		log.Fatal(err)
	case <-sigChan:
		var shutDownWg sync.WaitGroup

		// close database connection
		go func() {
			shutDownWg.Add(1)
			defer shutDownWg.Done()

			log.Println("closing database connection")
			db.Close()
		}()

		// shutdown the app
		go func() {
			shutDownWg.Add(1)
			defer shutDownWg.Done()

			log.Println("shutting down the app")
			app.ShutdownWithTimeout(30 * time.Second)
		}()

		shutDownWg.Wait()
		log.Println("done!")
	}
}

func initDatabase() *sql.DB {
	if cfg == nil {
		log.Fatal("[x] app config required")
	}
	conn, err := database.Connect(&database.MysqlConfig{
		Host:     cfg.MySQLHost,
		Port:     cfg.MySQLPort,
		User:     cfg.MySQLUser,
		Password: cfg.MySQLPassword,
		DBName:   cfg.MySQLDB,
	})

	if err != nil {
		log.Fatal("[x] failed to connect to database")
	}

	return conn
}
