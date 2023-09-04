package configuration

import (
	"path/filepath"
)

// Database contains the configurations of drivers supported by the database.
type Database struct {
	driver string
	CSV    CsvDB
}

// Driver returns the configured database driver type, e.g. csv, json, postgres, etc..
func (d Database) Driver() string {
	return d.driver
}

// CsvDB holds and retrieves the configuration properties of the CSV database.
type CsvDB struct {
	fileName string
	dataDir  string
}

// NewCsvDB returns a new CsvDB implementation.
func NewCsvDB(fileName, dataDir string) CsvDB {
	return CsvDB{
		fileName: fileName,
		dataDir:  dataDir,
	}
}

// FileName returns the file name of the CSV database.
func (c CsvDB) FileName() string {
	return c.fileName
}

// DataDir returns the base data directory of the CSV database.
func (c CsvDB) DataDir() string {
	return c.dataDir
}

// FilePath returns the full path of the CSV database file.
// The file path is built by prefixing the data directory to the file name.
func (c CsvDB) FilePath() string {
	return filepath.Join(c.dataDir, c.fileName)
}
