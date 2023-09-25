package repository

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/config"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	testCocktailRecs = []byte(`{
    "drinks": [
        {
            "idDrink": "1",
            "strDrink": "Acapulco",
            "strDrinkAlternate": null,
            "strTags": null,
            "strVideo": null,
            "strCategory": "Ordinary Drink",
            "strIBA": null,
            "strAlcoholic": "Alcoholic",
            "strGlass": "Old-fashioned glass",
            "strInstructions": "Combine and shake all ingredients (except mint) with ice and strain into an old-fashioned glass over ice cubes. Add the sprig of mint and serve.",
            "strInstructionsES": "Mezcle y agite todos los ingredientes (excepto la menta) con hielo y cuélelos en un vaso de rocas sobre cubitos de hielo. Añadir una ramita de menta y servir.",
            "strInstructionsDE": "Alle Zutaten (außer Minze) mit Eis mischen und schütteln und in ein old-fashioned Glas über Eiswürfel abseihen. Den Minzzweig dazugeben und servieren.",
            "strInstructionsFR": null,
            "strInstructionsIT": "Unire e scuotere tutti gli ingredienti (tranne la menta) con ghiaccio e filtrare in un bicchiere vecchio stile su cubetti di ghiaccio.Aggiungere il rametto di menta e servire.",
            "strInstructionsZH-HANS": null,
            "strInstructionsZH-HANT": null,
            "strDrinkThumb": "https://www.thecocktaildb.com/images/media/drink/il9e0r1582478841.jpg",
            "strIngredient1": "Light rum",
            "strIngredient2": "Triple sec",
            "strIngredient3": "Lime juice",
            "strIngredient4": "Sugar",
            "strIngredient5": "Egg white",
            "strIngredient6": "Mint",
            "strIngredient7": null,
            "strIngredient8": null,
            "strIngredient9": null,
            "strIngredient10": null,
            "strIngredient11": null,
            "strIngredient12": null,
            "strIngredient13": null,
            "strIngredient14": null,
            "strIngredient15": null,
            "strMeasure1": "1 1/2 oz ",
            "strMeasure2": "1 1/2 tsp ",
            "strMeasure3": "1 tblsp ",
            "strMeasure4": "1 tsp ",
            "strMeasure5": "1 ",
            "strMeasure6": "1 ",
            "strMeasure7": null,
            "strMeasure8": null,
            "strMeasure9": null,
            "strMeasure10": null,
            "strMeasure11": null,
            "strMeasure12": null,
            "strMeasure13": null,
            "strMeasure14": null,
            "strMeasure15": null,
            "strImageSource": null,
            "strImageAttribution": null,
            "strCreativeCommonsConfirmed": "Yes",
            "dateModified": "2016-09-02 11:26:16"
        },
        {
            "idDrink": "2",
            "strDrink": "Afterglow",
            "strDrinkAlternate": null,
            "strTags": null,
            "strVideo": null,
            "strCategory": "Cocktail",
            "strIBA": null,
            "strAlcoholic": "Non alcoholic",
            "strGlass": "Highball Glass",
            "strInstructions": "Mix. Serve over ice.",
            "strInstructionsES": "Mezcla. Servir con hielo.",
            "strInstructionsDE": "Mischen. Auf Eis servieren.",
            "strInstructionsFR": null,
            "strInstructionsIT": "Servire con ghiaccio.Mescolare.",
            "strInstructionsZH-HANS": null,
            "strInstructionsZH-HANT": null,
            "strDrinkThumb": "https://www.thecocktaildb.com/images/media/drink/vuquyv1468876052.jpg",
            "strIngredient1": "Grenadine",
            "strIngredient2": "Orange juice",
            "strIngredient3": "Pineapple juice",
            "strIngredient4": null,
            "strIngredient5": null,
            "strIngredient6": null,
            "strIngredient7": null,
            "strIngredient8": null,
            "strIngredient9": null,
            "strIngredient10": null,
            "strIngredient11": null,
            "strIngredient12": null,
            "strIngredient13": null,
            "strIngredient14": null,
            "strIngredient15": null,
            "strMeasure1": "1 part ",
            "strMeasure2": "4 parts ",
            "strMeasure3": "4 parts ",
            "strMeasure4": null,
            "strMeasure5": null,
            "strMeasure6": null,
            "strMeasure7": null,
            "strMeasure8": null,
            "strMeasure9": null,
            "strMeasure10": null,
            "strMeasure11": null,
            "strMeasure12": null,
            "strMeasure13": null,
            "strMeasure14": null,
            "strMeasure15": null,
            "strImageSource": null,
            "strImageAttribution": null,
            "strCreativeCommonsConfirmed": "No",
            "dateModified": "2016-07-18 22:07:32"
        },
        {
            "idDrink": "3",
            "strDrink": "Americano",
            "strDrinkAlternate": null,
            "strTags": "IBA,Classic",
            "strVideo": "https://www.youtube.com/watch?v=TmeUJ2g3ogM",
            "strCategory": "Ordinary Drink",
            "strIBA": "Unforgettables",
            "strAlcoholic": "Alcoholic",
            "strGlass": "Collins glass",
            "strInstructions": "Pour the Campari and vermouth over ice into glass, add a splash of soda water and garnish with half orange slice.",
            "strInstructionsES": "Vierta el Campari y el vermut con hielo en el vaso. Añadir un poco de agua con gas y decorar con media rodaja de naranja.",
            "strInstructionsDE": "Den Campari und den Wermut über Eis in ein Glas gießen, einen Spritzer Sodawasser hinzufügen und mit einer halben Orangenscheibe garnieren.",
            "strInstructionsFR": null,
            "strInstructionsIT": "Versare Campari e vermut su ghiaccio in un bicchiere, aggiungere un goccio di acqua di seltz e guarnire con mezza fetta d'arancia.",
            "strInstructionsZH-HANS": null,
            "strInstructionsZH-HANT": null,
            "strDrinkThumb": "https://www.thecocktaildb.com/images/media/drink/709s6m1613655124.jpg",
            "strIngredient1": "Campari",
            "strIngredient2": "Sweet Vermouth",
            "strIngredient3": "Lemon peel",
            "strIngredient4": "Orange peel",
            "strIngredient5": null,
            "strIngredient6": null,
            "strIngredient7": null,
            "strIngredient8": null,
            "strIngredient9": null,
            "strIngredient10": null,
            "strIngredient11": null,
            "strIngredient12": null,
            "strIngredient13": null,
            "strIngredient14": null,
            "strIngredient15": null,
            "strMeasure1": "1 oz ",
            "strMeasure2": "1 oz red ",
            "strMeasure3": "Twist of ",
            "strMeasure4": "Twist of ",
            "strMeasure5": null,
            "strMeasure6": null,
            "strMeasure7": null,
            "strMeasure8": null,
            "strMeasure9": null,
            "strMeasure10": null,
            "strMeasure11": null,
            "strMeasure12": null,
            "strMeasure13": null,
            "strMeasure14": null,
            "strMeasure15": null,
            "strImageSource": "https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg",
            "strImageAttribution": "Author - Cher37 https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg",
            "strCreativeCommonsConfirmed": "Yes",
            "dateModified": "2016-11-04 09:52:06"
        }
    ]
}`)

	testCocktailRecWithNoIngrs = []byte(`{
    "drinks": [
        {
            "idDrink": "1",
            "strDrink": "Acapulco",
            "strDrinkAlternate": null,
            "strTags": null,
            "strVideo": null,
            "strCategory": "Ordinary Drink",
            "strIBA": null,
            "strAlcoholic": "Alcoholic",
            "strGlass": "Old-fashioned glass",
            "strInstructions": "Combine and shake all ingredients (except mint) with ice and strain into an old-fashioned glass over ice cubes. Add the sprig of mint and serve.",
            "strInstructionsES": "Mezcle y agite todos los ingredientes (excepto la menta) con hielo y cuélelos en un vaso de rocas sobre cubitos de hielo. Añadir una ramita de menta y servir.",
            "strInstructionsDE": "Alle Zutaten (außer Minze) mit Eis mischen und schütteln und in ein old-fashioned Glas über Eiswürfel abseihen. Den Minzzweig dazugeben und servieren.",
            "strInstructionsFR": null,
            "strInstructionsIT": "Unire e scuotere tutti gli ingredienti (tranne la menta) con ghiaccio e filtrare in un bicchiere vecchio stile su cubetti di ghiaccio.Aggiungere il rametto di menta e servire.",
            "strInstructionsZH-HANS": null,
            "strInstructionsZH-HANT": null,
            "strDrinkThumb": "https://www.thecocktaildb.com/images/media/drink/il9e0r1582478841.jpg",
            "strIngredient1": null,
            "strIngredient2": null,
            "strIngredient3": null,
            "strIngredient4": null,
            "strIngredient5": null,
            "strIngredient6": null,
            "strIngredient7": null,
            "strIngredient8": null,
            "strIngredient9": null,
            "strIngredient10": null,
            "strIngredient11": null,
            "strIngredient12": null,
            "strIngredient13": null,
            "strIngredient14": null,
            "strIngredient15": null,
            "strMeasure1": null,
            "strMeasure2": null,
            "strMeasure3": null,
            "strMeasure4": null,
            "strMeasure5": null,
            "strMeasure6": null,
            "strMeasure7": null,
            "strMeasure8": null,
            "strMeasure9": null,
            "strMeasure10": null,
            "strMeasure11": null,
            "strMeasure12": null,
            "strMeasure13": null,
            "strMeasure14": null,
            "strMeasure15": null,
            "strImageSource": null,
            "strImageAttribution": null,
            "strCreativeCommonsConfirmed": "Yes",
            "dateModified": "2016-09-02 11:26:16"
        },
        {
            "idDrink": "2",
            "strDrink": "Afterglow",
            "strDrinkAlternate": null,
            "strTags": null,
            "strVideo": null,
            "strCategory": "Cocktail",
            "strIBA": null,
            "strAlcoholic": "Non alcoholic",
            "strGlass": "Highball Glass",
            "strInstructions": "Mix. Serve over ice.",
            "strInstructionsES": "Mezcla. Servir con hielo.",
            "strInstructionsDE": "Mischen. Auf Eis servieren.",
            "strInstructionsFR": null,
            "strInstructionsIT": "Servire con ghiaccio.Mescolare.",
            "strInstructionsZH-HANS": null,
            "strInstructionsZH-HANT": null,
            "strDrinkThumb": "https://www.thecocktaildb.com/images/media/drink/vuquyv1468876052.jpg",
            "strIngredient1": "Grenadine",
            "strIngredient2": "Orange juice",
            "strIngredient3": "Pineapple juice",
            "strIngredient4": null,
            "strIngredient5": null,
            "strIngredient6": null,
            "strIngredient7": null,
            "strIngredient8": null,
            "strIngredient9": null,
            "strIngredient10": null,
            "strIngredient11": null,
            "strIngredient12": null,
            "strIngredient13": null,
            "strIngredient14": null,
            "strIngredient15": null,
            "strMeasure1": "1 part ",
            "strMeasure2": "4 parts ",
            "strMeasure3": "4 parts ",
            "strMeasure4": null,
            "strMeasure5": null,
            "strMeasure6": null,
            "strMeasure7": null,
            "strMeasure8": null,
            "strMeasure9": null,
            "strMeasure10": null,
            "strMeasure11": null,
            "strMeasure12": null,
            "strMeasure13": null,
            "strMeasure14": null,
            "strMeasure15": null,
            "strImageSource": null,
            "strImageAttribution": null,
            "strCreativeCommonsConfirmed": "No",
            "dateModified": "2016-07-18 22:07:32"
        },
        {
            "idDrink": "3",
            "strDrink": "Americano",
            "strDrinkAlternate": null,
            "strTags": "IBA,Classic",
            "strVideo": "https://www.youtube.com/watch?v=TmeUJ2g3ogM",
            "strCategory": "Ordinary Drink",
            "strIBA": "Unforgettables",
            "strAlcoholic": "Alcoholic",
            "strGlass": "Collins glass",
            "strInstructions": "Pour the Campari and vermouth over ice into glass, add a splash of soda water and garnish with half orange slice.",
            "strInstructionsES": "Vierta el Campari y el vermut con hielo en el vaso. Añadir un poco de agua con gas y decorar con media rodaja de naranja.",
            "strInstructionsDE": "Den Campari und den Wermut über Eis in ein Glas gießen, einen Spritzer Sodawasser hinzufügen und mit einer halben Orangenscheibe garnieren.",
            "strInstructionsFR": null,
            "strInstructionsIT": "Versare Campari e vermut su ghiaccio in un bicchiere, aggiungere un goccio di acqua di seltz e guarnire con mezza fetta d'arancia.",
            "strInstructionsZH-HANS": null,
            "strInstructionsZH-HANT": null,
            "strDrinkThumb": "https://www.thecocktaildb.com/images/media/drink/709s6m1613655124.jpg",
            "strIngredient1": "Campari",
            "strIngredient2": "Sweet Vermouth",
            "strIngredient3": "Lemon peel",
            "strIngredient4": "Orange peel",
            "strIngredient5": null,
            "strIngredient6": null,
            "strIngredient7": null,
            "strIngredient8": null,
            "strIngredient9": null,
            "strIngredient10": null,
            "strIngredient11": null,
            "strIngredient12": null,
            "strIngredient13": null,
            "strIngredient14": null,
            "strIngredient15": null,
            "strMeasure1": "1 oz ",
            "strMeasure2": "1 oz red ",
            "strMeasure3": "Twist of ",
            "strMeasure4": "Twist of ",
            "strMeasure5": null,
            "strMeasure6": null,
            "strMeasure7": null,
            "strMeasure8": null,
            "strMeasure9": null,
            "strMeasure10": null,
            "strMeasure11": null,
            "strMeasure12": null,
            "strMeasure13": null,
            "strMeasure14": null,
            "strMeasure15": null,
            "strImageSource": "https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg",
            "strImageAttribution": "Author - Cher37 https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg",
            "strCreativeCommonsConfirmed": "Yes",
            "dateModified": "2016-11-04 09:52:06"
        }
    ]
}`)
)

type CocktailTestSuite struct {
	suite.Suite
	workDir string
	dataCSV []byte
	data    []entity.Cocktail
}

func TestCocktailTestSuite(t *testing.T) {
	suite.Run(t, new(CocktailTestSuite))
}

func (s *CocktailTestSuite) SetupSuite() {
	workDir := "testdata"
	if _, err := os.Stat(workDir); errors.Is(err, os.ErrNotExist) {
		require.NoError(s.T(), os.Mkdir(workDir, os.ModePerm),
			"create the work dataDir is mandatory")
	} else {
		require.Nil(s.T(), err)
	}
	s.workDir = workDir
	s.dataCSV = []byte(`1,foo,,,"[{""name"":""fooIngr"",""measure"":""someMeasure""}]",foo instructions,,,,,,,,0001-01-01 00:00:00,0001-01-01 00:00:00,0001-01-01 00:00:00
2,bar,,,"[{""name"":""fooIngr"",""measure"":""someMeasure""}]",bar instructions,,,,,,,,0001-01-01 00:00:00,0001-01-01 00:00:00,0001-01-01 00:00:00
3,baz,,,"[{""name"":""fooIngr"",""measure"":""someMeasure""}]",baz instructions,,,,,,,,0001-01-01 00:00:00,0001-01-01 00:00:00,0001-01-01 00:00:00
`)
	s.data = []entity.Cocktail{
		{ID: 1, Name: "foo", Instructions: "foo instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
		{ID: 2, Name: "bar", Instructions: "bar instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
		{ID: 3, Name: "baz", Instructions: "baz instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
	}
}

func (s *CocktailTestSuite) TearDownSuite() {
	assert.NoError(s.T(), os.RemoveAll(s.workDir),
		"remove the work dataDir is mandatory")
}

func (s *CocktailTestSuite) TearDownTest() {
	var count int
	files, err := os.ReadDir(s.workDir)
	assert.Nil(s.T(), err)
	for _, f := range files {
		if err := os.RemoveAll(filepath.Join(s.workDir, f.Name())); err == nil {
			count++
		}
	}
}

func (s *CocktailTestSuite) TestNewCocktail() {
	type args struct {
		csv     config.CsvDB
		dataAPI config.DataAPI
	}

	tests := []struct {
		name string
		args args
		exp  Cocktail
		err  error
	}{
		{
			name: "Invalid endpoint",
			args: args{
				dataAPI: config.NewDataAPI("https://foo.com/some-endpoint"),
				csv:     config.NewCsv("foo.csv", s.workDir),
			},
			exp: Cocktail{},
			err: &DataApiErr{},
		},
		{
			name: "Invalid data dataFile",
			args: args{
				dataAPI: config.NewDataAPI("https://thecocktaildb.com/api/json/v1/1/search.php?f=a"),
				csv:     config.NewCsv("", s.workDir),
			},
			exp: Cocktail{},
			err: &CsvErr{},
		},
		{
			name: "Valid",
			args: args{
				dataAPI: config.NewDataAPI("https://thecocktaildb.com/api/json/v1/1/search.php?f=a"),
				csv:     config.NewCsv("foo.csv", s.workDir),
			},
			exp: Cocktail{
				dataAPI: config.NewDataAPI("https://thecocktaildb.com/api/json/v1/1/search.php?f=a"),
				csv:     config.NewCsv("foo.csv", s.workDir),
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			cfg := config.Config{
				HTTP:     config.HTTP{DataAPI: tt.args.dataAPI},
				Database: config.Database{Csv: tt.args.csv},
			}

			out, err := NewCocktail(cfg)
			if tt.err != nil {
				require.NotNil(t, err)
				assert.Equal(t, Cocktail{}, out)
				assert.IsType(t, tt.err, err)
				return
			}
			require.Nil(t, err)
			assert.NotEqual(s.T(), Cocktail{}, out)
			assert.Equal(t, tt.exp.csv, out.csv)
			assert.Equal(t, tt.exp.dataAPI, out.dataAPI)
		})
	}
}

func (s *CocktailTestSuite) TestReadAll() {
	type file struct {
		name string
		mode os.FileMode
		data []byte
	}
	tests := []struct {
		name     string
		exp      []entity.Cocktail
		err      error
		file     file
		wantFile bool
	}{
		{
			name:     "Invalid file",
			exp:      nil,
			err:      &CsvErr{&fs.PathError{}},
			file:     file{name: "cocktail_foo.csv"},
			wantFile: false,
		},
		{
			name:     "Data Empty",
			exp:      []entity.Cocktail{},
			err:      nil,
			file:     file{name: "cocktail_empty.csv", mode: dataFileMode, data: nil},
			wantFile: true,
		},
		{
			name:     "Valid",
			exp:      s.data,
			err:      nil,
			file:     file{name: "cocktail_valid.csv", mode: dataFileMode, data: s.dataCSV},
			wantFile: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			csvCfg := config.NewCsv(tt.file.name, s.workDir)
			if tt.wantFile {
				require.NoError(t, os.WriteFile(csvCfg.FilePath(), tt.file.data, tt.file.mode),
					fmt.Sprintf("create the test file %q is mandatory", csvCfg.FilePath()))
			}
			repo := Cocktail{csv: csvCfg}

			out, err := repo.ReadAll()
			if tt.err != nil {
				require.NotNil(t, err)
				assert.Nil(t, out)
				assert.IsType(t, tt.err, err)
				if errWrp := errors.Unwrap(tt.err); errWrp != nil {
					assert.IsType(t, errWrp, errors.Unwrap(err))
				}
				return
			}
			require.Nil(t, err)
			assert.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}

}

func (s *CocktailTestSuite) TestReplaceDB() {
	type file struct {
		name string
		mode os.FileMode
		data []byte
	}
	tests := []struct {
		name     string
		args     []entity.Cocktail
		exp      []entity.Cocktail
		err      error
		file     file
		wantFile bool
	}{
		{
			name:     "Invalid CSV file",
			args:     s.data,
			exp:      nil,
			err:      &CsvErr{&fs.PathError{}},
			file:     file{name: "foo.csv"},
			wantFile: false,
		},
		{
			name: "Discarded record because of a parseCocktail error.",
			file: file{name: "cocktail_parse_errs.csv", mode: dataFileMode, data: nil},
			args: []entity.Cocktail{
				{ID: 1, Name: "foo", Instructions: "foo instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
				{ID: 2, Name: "bar", Instructions: "bar instructions"},
				{ID: 3, Name: "baz", Instructions: "baz instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
			},
			exp: []entity.Cocktail{
				{ID: 1, Name: "foo", Instructions: "foo instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
				{ID: 3, Name: "baz", Instructions: "baz instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
			},
			err:      nil,
			wantFile: true,
		},
		{
			name:     "Valid with no data",
			file:     file{name: "create_valid_empty.csv", mode: dataFileMode, data: nil},
			args:     s.data,
			exp:      s.data,
			err:      nil,
			wantFile: true,
		},
		{
			name: "Valid with data",
			file: file{name: "create_valid_full.csv", mode: dataFileMode, data: s.dataCSV},
			args: []entity.Cocktail{
				{ID: 2, Name: "bar", Instructions: "bar instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
				{ID: 3, Name: "baz", Instructions: "baz instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
			},
			exp: []entity.Cocktail{
				{ID: 2, Name: "bar", Instructions: "bar instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
				{ID: 3, Name: "baz", Instructions: "baz instructions", Ingredients: []entity.Ingredient{{Name: "fooIngr", Measure: "someMeasure"}}},
			},
			err:      nil,
			wantFile: true,
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			csvCfg := config.NewCsv(tt.file.name, s.workDir)
			if tt.wantFile {
				require.NoError(t, os.WriteFile(csvCfg.FilePath(), tt.file.data, tt.file.mode),
					fmt.Sprintf("creation of the test file %q is mandatory", csvCfg.FilePath()))
			}
			repo := Cocktail{csv: csvCfg}

			err := repo.ReplaceData(tt.args)
			if tt.err != nil {
				require.NotNil(s.T(), err)
				assert.IsType(t, tt.err, err)
				if errWrp := errors.Unwrap(tt.err); errWrp != nil {
					assert.IsType(t, errWrp, errors.Unwrap(err))
				}
				return
			}
			require.Nil(s.T(), err)

			recs, err := repo.ReadAll()
			require.Nil(s.T(), err)
			assert.Len(s.T(), tt.exp, len(recs))
			assert.Equal(t, tt.exp, recs)
		})
	}
}

func (s *CocktailTestSuite) TestFetchData() {
	type resp struct {
		code int
		body []byte
		err  error
	}
	tests := []struct {
		name string
		url  string
		exp  []entity.Cocktail
		err  error
		resp resp
	}{
		{
			name: "URL empty",
			url:  "",
			exp:  nil,
			err:  &DataApiErr{&url.Error{}},
			resp: resp{
				code: http.StatusBadRequest,
				body: nil,
				err:  &url.Error{},
			},
		},
		{
			name: "Bad URL",
			url:  "https://foo,.com/some-endpoint",
			exp:  nil,
			err:  &DataApiErr{&url.Error{}},
			resp: resp{
				code: 0,
				body: nil,
				err:  &url.Error{},
			},
		},
		{
			name: "HTTP Client error",
			url:  "https://foo,.com/some-endpoint",
			err:  &DataApiErr{},
			exp:  nil,
			resp: resp{
				code: 0,
				body: nil,
				err:  errors.New("http client error"),
			},
		},
		{
			name: "Bad Code",
			url:  "https://foo.com/api/v1/some-endpoint",
			exp:  nil,
			err:  &DataApiErr{ErrInvalidRespCode},
			resp: resp{
				code: http.StatusForbidden,
				body: []byte(`{ "message"": "forbidden"}`),
				err:  nil,
			},
		},
		{
			name: "Bad JSON",
			url:  "https://foo.com/api/v1/some-endpoint",
			exp:  nil,
			err:  &DataApiErr{&json.SyntaxError{}},
			resp: resp{
				code: http.StatusOK,
				body: []byte(`{`),
				err:  nil,
			},
		},
		{
			name: "No records",
			url:  "https://foo.com/api/v1/some-endpoint",
			exp:  []entity.Cocktail{},
			err:  nil,
			resp: resp{
				code: http.StatusOK,
				body: []byte(`{	"drinks": []}`),
				err:  nil,
			},
		},
		{
			name: "Parse error",
			url:  "https://foo.com/api/v1/some-endpoint",
			exp: []entity.Cocktail{
				{ID: 2, Name: "Afterglow", Alcoholic: "Non alcoholic", Category: "Cocktail", Ingredients: []entity.Ingredient{{Name: "Grenadine", Measure: "1 part "}, {Name: "Orange juice", Measure: "4 parts "}, {Name: "Pineapple juice", Measure: "4 parts "}}, Instructions: "Mix. Serve over ice.", Glass: "Highball Glass", IBA: "", ImgAttribution: "", ImgSrc: "", Tags: "", Thumb: "https://www.thecocktaildb.com/images/media/drink/vuquyv1468876052.jpg", Video: "", SrcDate: time.Date(2016, time.July, 18, 22, 7, 32, 0, time.UTC), CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 3, Name: "Americano", Alcoholic: "Alcoholic", Category: "Ordinary Drink", Ingredients: []entity.Ingredient{{Name: "Campari", Measure: "1 oz "}, {Name: "Sweet Vermouth", Measure: "1 oz red "}, {Name: "Lemon peel", Measure: "Twist of "}, {Name: "Orange peel", Measure: "Twist of "}}, Instructions: "Pour the Campari and vermouth over ice into glass, add a splash of soda water and garnish with half orange slice.", Glass: "Collins glass", IBA: "Unforgettables", ImgAttribution: "Author - Cher37 https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg", ImgSrc: "https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg", Tags: "IBA,Classic", Thumb: "https://www.thecocktaildb.com/images/media/drink/709s6m1613655124.jpg", Video: "https://www.youtube.com/watch?v=TmeUJ2g3ogM", SrcDate: time.Date(2016, time.November, 4, 9, 52, 6, 0, time.UTC), CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
			err: nil,
			resp: resp{
				code: http.StatusOK,
				body: testCocktailRecWithNoIngrs,
				err:  nil,
			},
		},
		{
			name: "All records",
			url:  "https://foo.com/api/v1/some-endpoint",
			exp: []entity.Cocktail{
				{ID: 1, Name: "Acapulco", Alcoholic: "Alcoholic", Category: "Ordinary Drink", Ingredients: []entity.Ingredient{{Name: "Light rum", Measure: "1 1/2 oz "}, {Name: "Triple sec", Measure: "1 1/2 tsp "}, {Name: "Lime juice", Measure: "1 tblsp "}, {Name: "Sugar", Measure: "1 tsp "}, {Name: "Egg white", Measure: "1 "}, {Name: "Mint", Measure: "1 "}}, Instructions: "Combine and shake all ingredients (except mint) with ice and strain into an old-fashioned glass over ice cubes. Add the sprig of mint and serve.", Glass: "Old-fashioned glass", IBA: "", ImgAttribution: "", ImgSrc: "", Tags: "", Thumb: "https://www.thecocktaildb.com/images/media/drink/il9e0r1582478841.jpg", Video: "", SrcDate: time.Date(2016, time.September, 2, 11, 26, 16, 0, time.UTC), CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 2, Name: "Afterglow", Alcoholic: "Non alcoholic", Category: "Cocktail", Ingredients: []entity.Ingredient{{Name: "Grenadine", Measure: "1 part "}, {Name: "Orange juice", Measure: "4 parts "}, {Name: "Pineapple juice", Measure: "4 parts "}}, Instructions: "Mix. Serve over ice.", Glass: "Highball Glass", IBA: "", ImgAttribution: "", ImgSrc: "", Tags: "", Thumb: "https://www.thecocktaildb.com/images/media/drink/vuquyv1468876052.jpg", Video: "", SrcDate: time.Date(2016, time.July, 18, 22, 7, 32, 0, time.UTC), CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
				{ID: 3, Name: "Americano", Alcoholic: "Alcoholic", Category: "Ordinary Drink", Ingredients: []entity.Ingredient{{Name: "Campari", Measure: "1 oz "}, {Name: "Sweet Vermouth", Measure: "1 oz red "}, {Name: "Lemon peel", Measure: "Twist of "}, {Name: "Orange peel", Measure: "Twist of "}}, Instructions: "Pour the Campari and vermouth over ice into glass, add a splash of soda water and garnish with half orange slice.", Glass: "Collins glass", IBA: "Unforgettables", ImgAttribution: "Author - Cher37 https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg", ImgSrc: "https://commons.wikimedia.org/wiki/File:Martini_Americano.jpg", Tags: "IBA,Classic", Thumb: "https://www.thecocktaildb.com/images/media/drink/709s6m1613655124.jpg", Video: "https://www.youtube.com/watch?v=TmeUJ2g3ogM", SrcDate: time.Date(2016, time.November, 4, 9, 52, 6, 0, time.UTC), CreatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC), UpdatedAt: time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)},
			},
			err: nil,
			resp: resp{
				code: http.StatusOK,
				body: testCocktailRecs,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		s.T().Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, tt.url, nil)
			require.Nil(t, err)
			mClient := mocks.NewHttpClient()
			mClient.On("Do", req).Return(&http.Response{
				StatusCode: tt.resp.code,
				Body:       io.NopCloser(bytes.NewReader(tt.resp.body)),
			}, tt.resp.err)

			repo := Cocktail{
				dataAPI:    config.NewDataAPI(tt.url),
				httpClient: mClient,
			}

			out, err := repo.FetchData()
			if tt.err != nil {
				require.NotNil(t, err)
				assert.Nil(t, out)
				assert.IsType(t, tt.err, err)
				if errWrp := errors.Unwrap(tt.err); errWrp != nil {
					assert.IsType(t, errWrp, errors.Unwrap(err))
				}
				return
			}
			require.Nil(t, err)
			assert.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}
}
