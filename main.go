package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 1. Get Identity from Environment Variable (Default to 'Unknown')
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "UnknownService"
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// 2. Initialize SQLite Database
	dbFile := "./database/enterprise.db"
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a simple table if it doesn't exist
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY, message TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)")
	statement.Exec()

	// Create a trigger to update created_at on insert if not provided
	statement, _ = db.Prepare("CREATE TRIGGER IF NOT EXISTS set_created_at AFTER INSERT ON logs BEGIN UPDATE logs SET created_at = CURRENT_TIMESTAMP WHERE rowid = NEW.rowid AND NEW.created_at IS NULL; END;")
	statement.Exec()

	if appName == "Payroll" {
		statement, _ = db.Prepare("CREATE TABLE IF NOT EXISTS employees (id INTEGER PRIMARY KEY, first_name TEXT, last_name TEXT, address_one TEXT, address_two TEXT, postcode TEXT, city TEXT, state TEXT, date_of_birth TEXT, sex TEXT, ethnicity TEXT, position TEXT, department TEXT, joined_at TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP, deleted_at DATETIME)")
		statement.Exec()
	}

	// Run payroll-dump.sql, located at ./payroll-dump.sql
	if appName == "Payroll" {
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM employees").Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		if count == 0 { // Only run if the table has no records
			dumpFile := "./database/seeder/payroll-dump.sql"
			dump, err := os.ReadFile(dumpFile)
			if err != nil {
				log.Fatal(err)
			}
			if _, err = db.Exec(string(dump)); err != nil {
				log.Fatal(err)
			}
		}
	}

	// --- ROUTES ---

	// The "Work" Route
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Hello from the %s Service!", appName)
		createdAt := time.Now().Format(time.RFC3339)

		// Record the visit in SQLite
		_, err := db.Exec("INSERT INTO logs (message, created_at) VALUES (?, ?)", message, createdAt)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to log to DB"})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"service":    appName,
			"status":     "active",
			"message":    message,
			"created_at": createdAt,
		})
	})

	// Add a GET request to all employee's full name and position at "/employees".
	e.GET("/employees", func(c echo.Context) error {
		rows, err := db.Query("SELECT first_name, last_name, position FROM employees")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		defer rows.Close()

		type Employee struct {
			FullName string `json:"full_name"`
			Position string `json:"position"`
		}

		var employees []Employee
		for rows.Next() {
			var firstName, lastName, position string
			if err := rows.Scan(&firstName, &lastName, &position); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}
			employees = append(employees, Employee{
				FullName: fmt.Sprintf("%s %s", firstName, lastName),
				Position: position,
			})
		}
		return c.JSON(http.StatusOK, employees)
	})

	// 3. SRE Essentials: Health and Readiness Probes
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/ready", func(c echo.Context) error {
		// In a real app, you'd check if the DB connection is actually alive here
		return c.String(http.StatusOK, "READY")
	})

	// Start Server
	port := "8080"
	fmt.Printf("Starting %s on port %s...\n", appName, port)
	e.Logger.Fatal(e.Start(":" + port))
}
