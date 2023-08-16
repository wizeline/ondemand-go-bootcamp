package repository

import (
	"encoding/csv"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/configuration"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
	"os"
	"reflect"
	"strconv"
)

// TODO: create filter types on package custom-type

var _ Fruit = &fruitCsv{}

func NewFruitCsv(cfg configuration.CsvDB) Fruit {
	logger.Log().Debug().Str("repository", "FruitCsv").Str("file", cfg.FilePath()).
		Msg("created repository")
	return &fruitCsv{
		cfg:          cfg,
		numRecFields: reflect.TypeOf(entity.Fruit{}).NumField(),
	}
}

type fruitCsv struct {
	cfg          configuration.CsvDB
	numRecFields int
}

func (f fruitCsv) Read(filter, value string) ([]entity.Fruit, error) {
	if filter == "" {
		return nil, ErrFilterEmpty
	}

	recs, err := f.getRecords()
	if err != nil {
		return nil, newCsvErr(err)
	}

	switch filter {
	case "id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return nil, newCsvErr(err)
		}
		return f.getById(id, recs), nil
	case "name":
		return f.getByName(value, recs), nil
	case "color":
		return f.getByColor(value, recs), nil
	case "country":
		return f.getByCountry(value, recs), nil
	}

	logger.Log().Error().Str("filter", filter).
		Str("method", "Read").
		Msgf("invalid csv filter")
	return nil, ErrInvalidFilter
}

func (f fruitCsv) ReadAll() ([]entity.Fruit, error) {
	recs, err := f.getRecords()
	if err != nil {
		logger.Log().Error().Str("Method", "ReadAll").
			Msgf("getting record failed")
		return nil, newCsvErr(err)
	}
	var fruits []entity.Fruit
	for _, rec := range recs {
		fruit, err := f.parseFruit(rec)
		if err != nil {
			continue
		}
		fruits = append(fruits, fruit)
	}
	return fruits, nil
}

func (f fruitCsv) Create(_ entity.Fruit) error {
	// TODO: implement me
	return nil
}

func (f fruitCsv) CreateAll(_ []entity.Fruit) error {
	// TODO: implement me
	return nil
}

func (f fruitCsv) getById(id int, recs [][]string) []entity.Fruit {
	var fruits []entity.Fruit
	for _, rec := range recs {
		fruit, err := f.parseFruit(rec)
		if err != nil {
			continue
		}
		if id == fruit.ID {
			fruits = append(fruits, fruit)
		}
	}
	return fruits
}

func (f fruitCsv) getByName(name string, recs [][]string) []entity.Fruit {
	var fruits []entity.Fruit
	for _, rec := range recs {
		fruit, err := f.parseFruit(rec)
		if err != nil {
			continue
		}
		if name == fruit.Name {
			fruits = append(fruits, fruit)
		}
	}
	return fruits
}

func (f fruitCsv) getByColor(color string, recs [][]string) []entity.Fruit {
	var fruits []entity.Fruit
	for _, rec := range recs {
		fruit, err := f.parseFruit(rec)
		if err != nil {
			continue
		}
		if color == fruit.Color {
			fruits = append(fruits, fruit)
		}
	}
	return fruits
}

func (f fruitCsv) getByCountry(country string, recs [][]string) []entity.Fruit {
	var fruits []entity.Fruit
	for _, rec := range recs {
		fruit, err := f.parseFruit(rec)
		if err != nil {
			continue
		}
		if country == fruit.Country {
			fruits = append(fruits, fruit)
		}
	}
	return fruits
}

func (f fruitCsv) getRecords() ([][]string, error) {
	filePath := f.cfg.FilePath()
	fd, err := os.Open(filePath)
	if err != nil {
		logger.Log().Error().Err(err).Str("file", filePath).
			Msg("getRecords: open csv file failed")
		return nil, newCsvErr(err)
	}
	defer func() {
		if err := fd.Close(); err != nil {
			logger.Log().Error().Err(err).Str("file", filePath).
				Msg("getRecords: close csv file failed")
		}
	}()
	return csv.NewReader(fd).ReadAll()
}

func (f fruitCsv) parseFruit(rec []string) (entity.Fruit, error) {

	if len(rec) > f.numRecFields {
		logger.Log().Warn().
			Int("max", f.numRecFields).
			Int("current", len(rec)).
			Msg("number of record fields exceeded")
	}

	recAux := make([]string, f.numRecFields)
	copy(recAux, rec)
	//fmt.Println("Rec:", rec)
	//fmt.Println("RecAux:", recAux)

	recID, err := strconv.Atoi(recAux[0])
	if err != nil {
		logger.Log().Error().Err(err).Str("id", recAux[0]).Msgf("error parsing id field")
		return entity.Fruit{}, err
	}

	//expDate, err := time.Parse(time.RFC3339, recAux[5])
	//if err != nil {
	//	logger.Log().Error().Err(err).Str("expiration_date", recAux[5]).
	//		Msgf("error parsing expiration_date field")
	//	return entity.Fruit{}, err
	//}
	//
	//createdAt, err := time.Parse(time.RFC3339, recAux[6])
	//if err != nil {
	//	logger.Log().Error().Err(err).Str("created_at", recAux[6]).
	//		Msgf("error parsing created_at field")
	//	return entity.Fruit{}, err
	//}
	//
	//updatedAt, err := time.Parse(time.RFC3339, recAux[7])
	//if err != nil {
	//	logger.Log().Error().Err(err).Str("updated_at", recAux[7]).
	//		Msgf("error parsing updated_at field")
	//	return entity.Fruit{}, err
	//}

	return entity.Fruit{
		ID:          recID,
		Name:        recAux[1],
		Description: recAux[2],
		Color:       recAux[3],
		Country:     recAux[4],
		//ExpirationDate: expDate,
		//CreatedAt:      createdAt,
		//UpdatedAt:      updatedAt,
	}, nil
}
