package service

import (
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

var _ Fruit = &fruit{}

type Fruit interface {
	Get(filter, value string) ([]entity.Fruit, error)
	GetAll() ([]entity.Fruit, error)
}

type FruitRepo interface {
	ReadAll() (entity.Fruits, error)
}

type fruit struct {
	repo FruitRepo
}

func NewFruit(repo FruitRepo) Fruit {
	logger.Log().Debug().
		Str("service", "Fruit").
		Msg("created service")
	return &fruit{
		repo: repo,
	}
}

func (f fruit) Get(_, _ string) ([]entity.Fruit, error) {
	return f.repo.ReadAll()
}

//func (f fruitCsv) Read(filter, value string) ([]entity.Fruit, error) {
//	if filter == "" {
//		return nil, ErrFilterEmpty
//	}
//
//	recs, err := f.getRecords()
//	if err != nil {
//		return nil, newCsvErr(err)
//	}
//
//	switch filter {
//	case "id":
//		id, err := strconv.Atoi(value)
//		if err != nil {
//			return nil, newCsvErr(err)
//		}
//		return f.ReadById(id, recs), nil
//	case "name":
//		return f.ReadByName(value, recs), nil
//	case "color":
//		return f.ReadByColor(value, recs), nil
//	case "country":
//		return f.ReadByCountry(value, recs), nil
//	}
//
//	logger.Log().Error().Str("filter", filter).
//		Str("method", "Read").
//		Msgf("invalid csv filter")
//	return nil, ErrInvalidFilter
//}

func (f fruit) GetAll() ([]entity.Fruit, error) {
	return f.repo.ReadAll()
}

//func (f fruitCsv) ReadById(id int, recs [][]string) []entity.Fruit {
//	var fruits []entity.Fruit
//	for _, rec := range recs {
//		fruit, err := f.parseFruit(rec)
//		if err != nil {
//			continue
//		}
//		if id == fruit.ID {
//			fruits = append(fruits, fruit)
//		}
//	}
//	return fruits
//}
//
//func (f fruitCsv) ReadByName(name string, recs [][]string) []entity.Fruit {
//	var fruits []entity.Fruit
//	for _, rec := range recs {
//		fruit, err := f.parseFruit(rec)
//		if err != nil {
//			continue
//		}
//		if name == fruit.Name {
//			fruits = append(fruits, fruit)
//		}
//	}
//	return fruits
//}
//
//func (f fruitCsv) ReadByColor(color string, recs [][]string) []entity.Fruit {
//	var fruits []entity.Fruit
//	for _, rec := range recs {
//		fruit, err := f.parseFruit(rec)
//		if err != nil {
//			continue
//		}
//		if color == fruit.Color {
//			fruits = append(fruits, fruit)
//		}
//	}
//	return fruits
//}
//
//func (f fruitCsv) ReadByCountry(country string, recs [][]string) []entity.Fruit {
//	var fruits []entity.Fruit
//	for _, rec := range recs {
//		fruit, err := f.parseFruit(rec)
//		if err != nil {
//			continue
//		}
//		if country == fruit.Country {
//			fruits = append(fruits, fruit)
//		}
//	}
//	return fruits
//}
//
//func (f fruitCsv) parseFruit(rec []string) (entity.Fruit, error) {
//
//	if len(rec) > f.numRecFields {
//		logger.Log().Warn().
//			Int("max", f.numRecFields).
//			Int("current", len(rec)).
//			Msg("number of record fields exceeded")
//	}
//
//	recAux := make([]string, f.numRecFields)
//	copy(recAux, rec)
//	//fmt.Println("Rec:", rec)
//	//fmt.Println("RecAux:", recAux)
//
//	recID, err := strconv.Atoi(recAux[0])
//	if err != nil {
//		logger.Log().Error().Err(err).Str("id", recAux[0]).Msgf("error parsing id field")
//		return entity.Fruit{}, err
//	}
//
//	//expDate, err := time.Parse(time.RFC3339, recAux[5])
//	//if err != nil {
//	//	logger.Log().Error().Err(err).Str("expiration_date", recAux[5]).
//	//		Msgf("error parsing expiration_date field")
//	//	return entity.Fruit{}, err
//	//}
//	//
//	//createdAt, err := time.Parse(time.RFC3339, recAux[6])
//	//if err != nil {
//	//	logger.Log().Error().Err(err).Str("created_at", recAux[6]).
//	//		Msgf("error parsing created_at field")
//	//	return entity.Fruit{}, err
//	//}
//	//
//	//updatedAt, err := time.Parse(time.RFC3339, recAux[7])
//	//if err != nil {
//	//	logger.Log().Error().Err(err).Str("updated_at", recAux[7]).
//	//		Msgf("error parsing updated_at field")
//	//	return entity.Fruit{}, err
//	//}
//
//	return entity.Fruit{
//		ID:          recID,
//		Name:        recAux[1],
//		Description: recAux[2],
//		Color:       recAux[3],
//		Country:     recAux[4],
//		//ExpirationDate: expDate,
//		//CreatedAt:      createdAt,
//		//UpdatedAt:      updatedAt,
//	}, nil
//}
