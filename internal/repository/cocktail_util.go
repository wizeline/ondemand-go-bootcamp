package repository

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/logger"
)

// Indexes of the entity.Cocktail fields related the specific column in the CSV dataFile.
const (
	idIdx             csvColumn = 0
	nameIdx           csvColumn = 1
	alcoholicIdx      csvColumn = 2
	categoryIdx       csvColumn = 3
	ingredientsIdx    csvColumn = 4
	instructionsIdx   csvColumn = 5
	glassIdx          csvColumn = 6
	ibaIdx            csvColumn = 7
	imgAttributionIdx csvColumn = 8
	imgSrcIdx         csvColumn = 9
	tagsIdx           csvColumn = 10
	thumbIdx          csvColumn = 11
	videoIdx          csvColumn = 12
	srcDateIdx        csvColumn = 13
	createdAtIdx      csvColumn = 14
	updatedAtIdx      csvColumn = 15
)

var csvHeadersMap = map[csvColumn]string{
	idIdx:             "id",
	nameIdx:           "name",
	alcoholicIdx:      "alcoholic",
	categoryIdx:       "category",
	ingredientsIdx:    "ingredients",
	instructionsIdx:   "instructions",
	glassIdx:          "glass",
	ibaIdx:            "iba",
	imgAttributionIdx: "image_attribution",
	imgSrcIdx:         "image_source",
	tagsIdx:           "tags",
	thumbIdx:          "thumb",
	videoIdx:          "video",
	srcDateIdx:        "source_date",
	createdAtIdx:      "created_at",
	updatedAtIdx:      "updated_at",
}

// csvColumn represents the column index of the csv file.
type csvColumn int

// drink represents the json record provided by the public API.
type drink struct {
	Alcoholic        string `json:"strAlcoholic"`
	Category         string `json:"strCategory"`
	DateModified     string `json:"dateModified"`
	DrinkId          string `json:"idDrink"`
	DrinkName        string `json:"strDrink"`
	DrinkAlternate   string `json:"strDrinkAlternate"`
	DrinkThumb       string `json:"strDrinkThumb"`
	Glass            string `json:"strGlass"`
	IBA              string `json:"strIBA"`
	ImageSource      string `json:"strImageSource"`
	ImageAttribution string `json:"strImageAttribution"`
	Instructions     string `json:"strInstructions"`
	Ingredient1      string `json:"strIngredient1"`
	Ingredient2      string `json:"strIngredient2"`
	Ingredient3      string `json:"strIngredient3"`
	Ingredient4      string `json:"strIngredient4"`
	Ingredient5      string `json:"strIngredient5"`
	Ingredient6      string `json:"strIngredient6"`
	Ingredient7      string `json:"strIngredient7"`
	Ingredient8      string `json:"strIngredient8"`
	Ingredient9      string `json:"strIngredient9"`
	Ingredient10     string `json:"strIngredient10"`
	Ingredient11     string `json:"strIngredient11"`
	Ingredient12     string `json:"strIngredient12"`
	Ingredient13     string `json:"strIngredient13"`
	Ingredient14     string `json:"strIngredient14"`
	Ingredient15     string `json:"strIngredient15"`
	Measure1         string `json:"strMeasure1"`
	Measure2         string `json:"strMeasure2"`
	Measure3         string `json:"strMeasure3"`
	Measure4         string `json:"strMeasure4"`
	Measure5         string `json:"strMeasure5"`
	Measure6         string `json:"strMeasure6"`
	Measure7         string `json:"strMeasure7"`
	Measure8         string `json:"strMeasure8"`
	Measure9         string `json:"strMeasure9"`
	Measure10        string `json:"strMeasure10"`
	Measure11        string `json:"strMeasure11"`
	Measure12        string `json:"strMeasure12"`
	Measure13        string `json:"strMeasure13"`
	Measure14        string `json:"strMeasure14"`
	Measure15        string `json:"strMeasure15"`
	Tags             string `json:"strTags"`
	Video            string `json:"strVideo"`
}

// extCocktails holds all the json records provided by the public API.
type extCocktails struct {
	Drinks []drink `json:"drinks"`
}

// getIngredients returns a valid entity.Ingredient list
func (c drink) getIngredients() []entity.Ingredient {
	ingredients := make([]entity.Ingredient, 0)

	if c.Ingredient1 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient1, Measure: c.Measure1})
	}
	if c.Ingredient2 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient2, Measure: c.Measure2})
	}
	if c.Ingredient3 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient3, Measure: c.Measure3})
	}
	if c.Ingredient4 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient4, Measure: c.Measure4})
	}
	if c.Ingredient5 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient5, Measure: c.Measure5})
	}
	if c.Ingredient6 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient6, Measure: c.Measure6})
	}
	if c.Ingredient7 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient7, Measure: c.Measure7})
	}
	if c.Ingredient8 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient8, Measure: c.Measure8})
	}
	if c.Ingredient9 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient9, Measure: c.Measure9})
	}
	if c.Ingredient10 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient10, Measure: c.Measure10})
	}
	if c.Ingredient11 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient11, Measure: c.Measure11})
	}
	if c.Ingredient12 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient12, Measure: c.Measure12})
	}
	if c.Ingredient13 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient13, Measure: c.Measure13})
	}
	if c.Ingredient14 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient14, Measure: c.Measure14})
	}
	if c.Ingredient15 != "" {
		ingredients = append(ingredients, entity.Ingredient{Name: c.Ingredient15, Measure: c.Measure15})
	}
	return ingredients
}

// parseCocktail returns a valid entity.Cocktail instance
func (c drink) parseCocktail() (entity.Cocktail, error) {
	if c == (drink{}) {
		return entity.Cocktail{}, ErrJsonRecEmpty
	}

	id, err := strconv.Atoi(c.DrinkId)
	if err != nil {
		return entity.Cocktail{}, err
	}

	if c.DrinkName == "" {
		return entity.Cocktail{}, ErrCocktailNameEmpty
	}

	if c.Instructions == "" {
		return entity.Cocktail{}, ErrCocktailInstructionsEmpty
	}

	ingredients := c.getIngredients()
	if len(ingredients) == 0 {
		return entity.Cocktail{}, ErrCocktailIngredientsEmpty
	}

	srcDate, err := time.Parse(time.DateTime, c.DateModified)
	if err != nil {
		return entity.Cocktail{}, err
	}

	return entity.Cocktail{
		ID:             id,
		Name:           c.DrinkName,
		Alcoholic:      c.Alcoholic,
		Category:       c.Category,
		Ingredients:    ingredients,
		Instructions:   c.Instructions,
		Glass:          c.Glass,
		IBA:            c.IBA,
		ImgAttribution: c.ImageAttribution,
		ImgSrc:         c.ImageSource,
		Tags:           c.Tags,
		Thumb:          c.DrinkThumb,
		Video:          c.Video,
		SrcDate:        srcDate,
	}, nil
}

// cocktailCSVRec represents a cocktail csv record
type cocktailCSVRec []string

// parseCocktail returns a valid entity.Cocktail instance.
func (record cocktailCSVRec) parseCocktail() (entity.Cocktail, error) {
	if len(record) == 0 {
		return entity.Cocktail{}, ErrCSVRecEmpty
	}

	numFields := len(csvHeadersMap)
	if len(record) != numFields {
		logger.Log().Warn().Str("required", fmt.Sprintf("%d/%d", len(record), numFields)).Str("record", strings.Join(record[:], ",")).
			Msg("newCocktailFromCSVRec: wrong number of fields")
	}

	// Preallocate a csv record matching the number of entity.Cocktail fields to avoid nil errors.
	rec := make([]string, numFields)
	copy(rec, record)

	recID, err := strconv.Atoi(rec[idIdx])
	if err != nil {
		logger.Log().Error().Err(err).Str("id", rec[idIdx]).
			Msgf("newCocktailFromCSVRec: parsing id failure")
		return entity.Cocktail{}, err
	}

	if rec[nameIdx] == "" {
		logger.Log().Error().Err(ErrCocktailNameEmpty).Str("name", rec[nameIdx]).
			Msgf("newCocktailFromCSVRec: parsing name failure")
		return entity.Cocktail{}, ErrCocktailNameEmpty
	}

	var ingredients []entity.Ingredient
	if err := json.Unmarshal([]byte(rec[ingredientsIdx]), &ingredients); err != nil {
		logger.Log().Error().Err(err).Str("ingredients", rec[ingredientsIdx]).
			Msgf("newCocktailFromCSVRec: unmarshalling ingredients failure")
		return entity.Cocktail{}, err
	}
	if len(ingredients) == 0 {
		logger.Log().Error().Err(ErrCocktailIngredientsEmpty).Str("ingredients", rec[ingredientsIdx]).
			Msgf("newCocktailFromCSVRec: parsing ingredients failure")
		return entity.Cocktail{}, ErrCocktailIngredientsEmpty
	}

	if rec[instructionsIdx] == "" {
		logger.Log().Error().Err(ErrCocktailInstructionsEmpty).Str("name", rec[instructionsIdx]).
			Msgf("newCocktailFromCSVRec: parsing instructions failure")
		return entity.Cocktail{}, ErrCocktailInstructionsEmpty
	}

	srcDate, err := time.Parse(time.DateTime, rec[srcDateIdx])
	if err != nil {
		logger.Log().Error().Err(err).Str("source_date", rec[srcDateIdx]).
			Msgf("newCocktailFromCSVRec: parsing source date failure")
		return entity.Cocktail{}, err
	}

	createdAt, err := time.Parse(time.DateTime, rec[createdAtIdx])
	if err != nil {
		logger.Log().Error().Err(err).Str("created_at", rec[createdAtIdx]).
			Msgf("newCocktailFromCSVRec: parsing created_at failure")
		return entity.Cocktail{}, err
	}

	updatedAt, err := time.Parse(time.DateTime, rec[updatedAtIdx])
	if err != nil {
		logger.Log().Error().Err(err).Str("updated_at", rec[updatedAtIdx]).
			Msgf("newCocktailFromCSVRec: parsing updated_at failure")
		return entity.Cocktail{}, err
	}

	return entity.Cocktail{
		ID:             recID,
		Name:           rec[nameIdx],
		Alcoholic:      rec[alcoholicIdx],
		Category:       rec[categoryIdx],
		Instructions:   rec[instructionsIdx],
		Ingredients:    ingredients,
		Glass:          rec[glassIdx],
		IBA:            rec[ibaIdx],
		ImgAttribution: rec[imgAttributionIdx],
		ImgSrc:         rec[imgSrcIdx],
		Tags:           rec[tagsIdx],
		Thumb:          rec[thumbIdx],
		Video:          rec[videoIdx],
		SrcDate:        srcDate,
		CreatedAt:      createdAt,
		UpdatedAt:      updatedAt,
	}, nil
}

// parseCocktailCsvRec returns a valid csv record from an entity.Cocktail
func parseCocktailCsvRec(c entity.Cocktail) ([]string, error) {
	ingredients, err := json.Marshal(c.Ingredients)
	if err != nil {
		return nil, err
	}

	rec := make([]string, len(csvHeadersMap))
	rec[idIdx] = strconv.Itoa(c.ID)
	rec[nameIdx] = c.Name
	rec[alcoholicIdx] = c.Alcoholic
	rec[categoryIdx] = c.Category
	rec[ingredientsIdx] = string(ingredients)
	rec[instructionsIdx] = c.Instructions
	rec[glassIdx] = c.Glass
	rec[ibaIdx] = c.IBA
	rec[imgAttributionIdx] = c.ImgAttribution
	rec[imgSrcIdx] = c.ImgSrc
	rec[tagsIdx] = c.Tags
	rec[thumbIdx] = c.Thumb
	rec[videoIdx] = c.Video
	rec[srcDateIdx] = c.SrcDate.Format(time.DateTime)
	rec[createdAtIdx] = c.CreatedAt.Format(time.DateTime)
	rec[updatedAtIdx] = c.UpdatedAt.Format(time.DateTime)
	return rec, nil
}
