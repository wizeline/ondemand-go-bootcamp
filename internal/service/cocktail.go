package service

import (
	"strconv"
	"time"

	ct "github.com/marcos-wz/capstone-go-bootcamp/internal/customtype"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

// Cocktail performs the core operations for Cocktail.
type Cocktail struct {
	repo CocktailRepo
}

// CocktailRepo is the abstraction of the Cocktail repository dependency.
type CocktailRepo interface {
	ReadAll() ([]entity.Cocktail, error)
	ReplaceData(recs []entity.Cocktail) error
	FetchData() ([]entity.Cocktail, error)
}

// NewCocktail returns a new Cocktail service implementation.
func NewCocktail(repo CocktailRepo) Cocktail {
	logger.Log().Debug().Msg("created Cocktail service")
	return Cocktail{
		repo: repo,
	}
}

// GetFiltered returns a filtered list of entity.Cocktail records from the database.
func (s Cocktail) GetFiltered(filter, value string) ([]entity.Cocktail, error) {
	if filter == "" {
		return nil, &FilterErr{ErrFltrTypeEmpty}
	}
	if value == "" {
		return nil, &FilterErr{ErrFltrValueEmpty}
	}
	fltr := newCocktailFilter(filter)
	if fltr == invalidFltr {
		return nil, &FilterErr{ErrFltrInvalid}
	}

	recs, err := s.repo.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(recs) == 0 {
		return []entity.Cocktail{}, nil
	}

	switch fltr {
	case idFltr:
		id, e := strconv.Atoi(value)
		if e != nil {
			return nil, &FilterErr{e}
		}
		return filterCocktailsById(id, recs), nil
	case nameFltr:
		return filterCocktailsByName(value, recs), nil
	case alcoholicFltr:
		return filterCocktailsByAlcoholic(value, recs), nil
	case categoryFltr:
		return filterCocktailsByCategory(value, recs), nil
	case ingredientFltr:
		return filterCocktailsByIngredient(value, recs), nil
	case glassFltr:
		return filterCocktailsByGlass(value, recs), nil
	default:
		logger.Log().Error().Err(ErrFltrInvalid).Str("filter", filter).Str("value", value).
			Msgf("GetFiltered: filter not supported")
		return nil, &FilterErr{ErrFltrInvalid}
	}
}

// GetAll returns all the entity.Cocktail records from the database.
func (s Cocktail) GetAll() ([]entity.Cocktail, error) {
	cocktails := make([]entity.Cocktail, 0)
	recs, err := s.repo.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(recs) > 0 {
		return recs, nil
	}
	return cocktails, nil
}

// UpdateDB updates the database records from the public API and returns a database operations summary.
// A new record is created if the fetched one does not exist in the database.
// If the record exists but the fetched record's date is newer, the record gets updated in the database.
// If the record exists, the fetched record date is the same, and any of the values is different, the record gets updated in the database.
func (s Cocktail) UpdateDB() (ct.DBOpsSummary, error) {
	dataSet, err := s.repo.ReadAll()
	if err != nil {
		return ct.DBOpsSummary{}, err
	}

	extData, err := s.repo.FetchData()
	if err != nil {
		return ct.DBOpsSummary{}, err
	}

	status := noChangesDBStatus
	start := time.Now().UTC()
	nrCounter := 0
	mrCounter := 0
	for _, rec := range extData {
		index, found := findCocktail(rec.ID, dataSet)
		if !found {
			nrCounter++
			rec.CreatedAt = dateTimeNow()
			rec.UpdatedAt = rec.CreatedAt
			dataSet = append(dataSet, rec)
			continue
		}

		if rec.SrcDate.After(dataSet[index].SrcDate) {
			mrCounter++
			rec.UpdatedAt = dateTimeNow()
			dataSet[index] = rec
			continue
		}

		if rec.SrcDate == dataSet[index].SrcDate && !cocktailsEqual(rec, dataSet[index]) {
			mrCounter++
			rec.UpdatedAt = dateTimeNow()
			dataSet[index] = rec
		}
	}

	totalOps := nrCounter + mrCounter
	if totalOps > 0 {
		if err := s.repo.ReplaceData(dataSet); err != nil {
			return ct.DBOpsSummary{}, err
		}
		status = successfulUpdateDBStatus
	}

	end := time.Now().UTC()
	return ct.DBOpsSummary{
		Status:       status,
		StartTime:    start,
		EndTime:      end,
		Duration:     end.Sub(start).String(),
		NewRecs:      nrCounter,
		ModifiedRecs: mrCounter,
		TotalOps:     totalOps,
		TotalRecs:    len(dataSet),
	}, nil
}
