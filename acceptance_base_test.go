package main_test

import (
	"gorm.io/gorm"
	"learning-go-challenges"
	"os"
	"testing"
)

var dbConnection *gorm.DB

func TestMain(m *testing.M) {
	println("Running test main")
	dbConnection = main.InitDb()

	exitCode := m.Run()

	println("Finished test main")
	os.Exit(exitCode)
}
