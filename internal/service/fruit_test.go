package service

import (
	"testing"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository"
)

const testDataDir = "../../test/data"

// TODO: implement mock tests

// TODO: implement integration tests
func TestFruit_Get_Int(t *testing.T) {
	cfg := configuration.NewCsvDB("fruits_valid.csv", testDataDir)
	repo := repository.NewFruitCsv(cfg)
	svc := NewFruit(repo)

	out, err := svc.Get("color", "green")
	t.Logf("ERROR: %v", err)
	t.Logf("OUT: %v", out)
}
