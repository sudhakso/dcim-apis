package main

import (
	"net/http"

	"github.com/dcim-apis/pkg/dcim"
	"github.com/dcim-apis/pkg/handlers"
	"github.com/dcim-apis/pkg/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func newHandler(db *gorm.DB) *handlers.DBHandler {
	db.AutoMigrate(&models.ColoProvider{})
	// db.AutoMigrate(&models.Deployment{})
	// db.AutoMigrate(&models.Device{})
	// db.AutoMigrate(&models.Rack{})

	return &handlers.DBHandler{db}
}

func main() {
	// establish DB
	var err error
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	//DB Handler
	dbh := newHandler(db)
	s := handlers.DcimServer{dbh}

	// Register Http handler
	h := dcim.Handler(s)

	http.ListenAndServe(":3000", h)
}
