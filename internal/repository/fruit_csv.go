package repository

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

var _ Fruit = &fruitCsv{}

// NewFruitCsv returns a new Fruit repository implementation with CSV database support.
func NewFruitCsv(cfg configuration.CsvDB) (Fruit, error) {
	if err := validateDataFile(cfg.FilePath()); err != nil {
		return nil, &CsvErr{err}
	}

	logger.Log().Debug().
		Str("repo", "FruitCsv").
		Str("file", cfg.FilePath()).
		Msg("created repository")
	return &fruitCsv{
		cfg:       cfg,
		numFields: reflect.TypeOf(entity.Fruit{}).NumField(),
	}, nil
}

type fruitCsv struct {
	cfg       configuration.CsvDB
	numFields int
}

// ReadAll returns all entity.Fruit records from the CSV database
func (f fruitCsv) ReadAll() (entity.Fruits, error) {
	filePath := f.cfg.FilePath()

	fd, err := os.Open(filePath)
	if err != nil {
		logger.Log().Error().Err(err).
			Str("file", filePath).
			Msg("ReadAll: open csv file failed")
		return nil, &CsvErr{err}
	}
	defer func() {
		if err := fd.Close(); err != nil {
			logger.Log().Error().Err(err).
				Str("file", filePath).
				Msg("ReadAll: close csv file failed")
		}
	}()

	r := csv.NewReader(fd)
	fruits := entity.Fruits{}
	for line := 1; true; line++ {
		rec, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			logger.Log().Warn().Err(err).
				Int("line", line).
				Str("file_path", filePath).
				Str("record", strings.Join(rec[:], ",")).
				Msgf("ReadAll: reading csv record failed, discarded record")
			continue
		}

		fruit, err := f.parseFruit(rec)
		if err != nil {
			logger.Log().Warn().Err(err).
				Int("line", line).
				Str("file_path", filePath).
				Str("record", strings.Join(rec[:], ",")).
				Msg("ReadAll: parsing record failed, discarded record")
			continue
		}
		fruits = append(fruits, fruit)
	}

	return fruits, nil
}

// Create adds a new entity.Fruit record into the CSV database.
func (f fruitCsv) Create(_ entity.Fruit) error {
	// TODO: implement me
	return nil
}

// CreateAll adds a list of entity.Fruit records into the CSV database.
func (f fruitCsv) CreateAll(_ entity.Fruits) error {
	// TODO: implement me
	return nil
}

// parseFruit converts a Fruit CSV record to a valid entity.Fruit
func (f fruitCsv) parseFruit(record []string) (entity.Fruit, error) {
	if len(record) == 0 {
		return entity.Fruit{}, ErrCSVRecordEmpty
	}

	if len(record) != f.numFields {
		logger.Log().Warn().
			Str("required", fmt.Sprintf("%d/%d", len(record), f.numFields)).
			Str("record", strings.Join(record[:], ",")).
			Msg("parseFruit: wrong number of record fields")
	}

	// Preallocate a new Fruit CSV record with the same number of fields as the entity.Fruit
	rec := make([]string, f.numFields)
	copy(rec, record)

	recID, err := strconv.Atoi(rec[0])
	if err != nil {
		logger.Log().Error().Err(err).
			Str("id", rec[0]).
			Msgf("parseFruit: error parsing id field")
		return entity.Fruit{}, err
	}

	expDate, err := parseUnixEpoch(rec[5])
	if err != nil {
		logger.Log().Error().Err(err).Str("expiration_date", rec[5]).
			Msgf("error parsing expiration_date field")
		return entity.Fruit{}, err
	}

	createdAt, err := parseUnixEpoch(rec[6])
	if err != nil {
		logger.Log().Error().Err(err).Str("created_at", rec[6]).
			Msgf("error parsing created_at field")
		return entity.Fruit{}, err
	}

	updatedAt, err := parseUnixEpoch(rec[7])
	if err != nil {
		logger.Log().Error().Err(err).Str("updated_at", rec[7]).
			Msgf("error parsing updated_at field")
		return entity.Fruit{}, err
	}

	return entity.Fruit{
		ID:             recID,
		Name:           rec[1],
		Description:    rec[2],
		Color:          rec[3],
		Country:        rec[4],
		ExpirationDate: expDate,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}, nil
}
