package repository

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/config"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

// Cocktail represents the Cocktail repository.
type Cocktail struct {
	csv        config.CsvDB
	dataAPI    config.DataAPI
	httpClient HttpClient
}

// NewCocktail returns a new Cocktail repository implementation.
func NewCocktail(cfg config.Config) (Cocktail, error) {
	dataAPI := cfg.HTTP.DataAPI
	csvDB := cfg.Database.Csv
	if err := checkEndpoint(dataAPI.URL()); err != nil {
		return Cocktail{}, &DataApiErr{err}
	}
	if err := createDataFile(csvDB.FileName(), csvDB.DataDir()); err != nil {
		return Cocktail{}, &CsvErr{err}
	}

	logger.Log().Debug().
		Str("csv_file", csvDB.FilePath()).
		Str("data_api", dataAPI.URL()).
		Msg("created Cocktail repository")
	return Cocktail{
		csv:        csvDB,
		dataAPI:    dataAPI,
		httpClient: &http.Client{},
	}, nil
}

// ReadAll returns all entity.Cocktail records from the CSV data file.
func (r Cocktail) ReadAll() ([]entity.Cocktail, error) {
	file := r.csv.FilePath()

	fd, err := os.Open(file)
	if err != nil {
		logger.Log().Error().Err(err).Str("method", "ReadAll").Str("file", file).
			Msg("open csv file failed")
		return nil, &CsvErr{err}
	}
	defer func() {
		if err := fd.Close(); err != nil {
			logger.Log().Error().Err(err).Str("method", "ReadAll").Str("file", file).
				Msg("close csv file failed")
		}
	}()

	reader := csv.NewReader(fd)
	cocktails := make([]entity.Cocktail, 0)
	for line := 1; true; line++ {
		var rec cocktailCSVRec
		rec, err = reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			logger.Log().Warn().Err(err).Str("method", "ReadAll").Int("line", line).Str("file", file).Str("record", strings.Join(rec[:], ",")).
				Msgf("read record failed, discarded record")
			continue
		}

		cocktail, err := rec.parseCocktail()
		if err != nil {
			logger.Log().Error().Err(err).Str("method", "ReadAll").Int("line", line).Str("file", file).Str("record", strings.Join(rec[:], ",")).
				Msg("parsing record failed, discarded record")
			continue
		}
		cocktails = append(cocktails, cocktail)
	}

	return cocktails, nil
}

// ReplaceData replace all data with the given entity.Cocktail records.
func (r Cocktail) ReplaceData(cocktails []entity.Cocktail) error {
	file := r.csv.FilePath()

	f, err := os.OpenFile(file, os.O_TRUNC|os.O_WRONLY, dataFileMode)
	if err != nil {
		logger.Log().Error().Err(err).Str("method", "ReplaceData").Str("file", file).
			Msg("open csv file failed")
		return &CsvErr{err}
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			logger.Log().Error().Err(err).Str("method", "ReplaceData").Str("file", file).
				Msg("close csv file failed")
		}
	}(f)

	writer := csv.NewWriter(f)
	for i, c := range cocktails {
		rec, errP := parseCocktailCsvRec(c)
		if errP != nil {
			logger.Log().Error().Err(err).Int("index", i).Str("cocktail", fmt.Sprintf("ID: %d, Name: %v", c.ID, c.Name)).
				Msg("ReplaceData: parsing cocktail to csv record failed, discarded record")
			continue
		}
		if err := writer.Write(rec); err != nil {
			logger.Log().Error().Err(err).Str("method", "ReplaceData").Int("index", i).Str("file", file).Str("record", strings.Join(rec[:], ",")).
				Msg("writing record failed, discarded record")
		}

	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		logger.Log().Error().Err(err).Str("method", "ReplaceData").
			Msg("flush writer failed")
		return &CsvErr{err}
	}

	return nil
}

// FetchData returns the fetched entity.Cocktail records from the data API.
func (r Cocktail) FetchData() ([]entity.Cocktail, error) {
	req, err := http.NewRequest(http.MethodGet, r.dataAPI.URL(), nil)
	if err != nil {
		return nil, &DataApiErr{err}
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, &DataApiErr{err}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log().Error().Err(err).Str("method", "FetchData").
				Msg("close response body failed")
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		logger.Log().Error().Err(err).Str("method", "FetchData").Int("code", resp.StatusCode).
			Msg("bad status code, expected 200")
		return nil, &DataApiErr{ErrInvalidRespCode}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, &DataApiErr{err}
	}

	data := extCocktails{}
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, &DataApiErr{err}
	}

	cocktails := make([]entity.Cocktail, 0)
	for i, rec := range data.Drinks {
		cocktail, errN := rec.parseCocktail()
		if errN != nil {
			logger.Log().Error().Err(errN).Str("method", "FetchData").Int("record", i).Str("cocktail", rec.DrinkName).
				Msg("parsing cocktail failed, external record skipped")
			continue
		}
		cocktails = append(cocktails, cocktail)
	}

	return cocktails, nil
}
