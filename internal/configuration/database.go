package configuration

import (
	"fmt"
)

var (
	_ CsvDB = &csvDB{}
)

type CsvDB interface {
	FileName() string
	DataDir() string
	FilePath() string
}

type Database struct {
	driver string
	CSV    CsvDB
}

func (d Database) Driver() string {
	return d.driver
}

type csvDB struct {
	fileName string
	dataDir  string
}

func NewCsvDB(fileName, dataDir string) CsvDB {
	return &csvDB{
		fileName: fileName,
		dataDir:  dataDir,
	}
}

func (c csvDB) FileName() string {
	return c.fileName
}

func (c csvDB) DataDir() string {
	return c.dataDir
}

func (c csvDB) FilePath() string {
	return fmt.Sprintf("%s/%s", c.dataDir, c.fileName)
}
