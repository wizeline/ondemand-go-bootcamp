package configuration

import (
	"path/filepath"
)

var (
	_ CsvDB = &csvDB{}
)

// CsvDB is a getter interface that retrieves the configuration properties of the CSV database.
type CsvDB interface {
	FileName() string
	DataDir() string
	FilePath() string
}

// Database contains the configurations of drivers supported by the database.
type Database struct {
	driver string
	CSV    CsvDB
}

// Driver returns the configured database driver type, e.g. csv, json, postgres, etc..
func (d Database) Driver() string {
	return d.driver
}

type csvDB struct {
	fileName string
	dataDir  string
}

// NewCsvDB returns a new CsvDB implementation.
func NewCsvDB(fileName, dataDir string) CsvDB {
	return &csvDB{
		fileName: fileName,
		dataDir:  dataDir,
	}
}

// FileName returns the file name of the CSV database.
func (c csvDB) FileName() string {
	return c.fileName
}

// DataDir returns the base data directory of the CSV database.
func (c csvDB) DataDir() string {
	return c.dataDir
}

// FilePath returns the full path of the CSV database file.
// The file path is built by prefixing the data directory to the file name.
func (c csvDB) FilePath() string {
	return filepath.Join(c.dataDir, c.fileName)
}
