package service

import (
	"fmt"
	"strings"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
)

const (
	invalidFltr    cocktailFilter = "invalid"
	idFltr         cocktailFilter = "id"
	nameFltr       cocktailFilter = "name"
	alcoholicFltr  cocktailFilter = "alcoholic"
	categoryFltr   cocktailFilter = "category"
	ingredientFltr cocktailFilter = "ingredient"
	glassFltr      cocktailFilter = "glass"

	noChangesDBStatus        = "no changes"
	successfulUpdateDBStatus = "database updated successfully"
)

var _ fmt.Stringer = cocktailFilter("")

type cocktailFilter string

func (f cocktailFilter) String() string {
	return string(f)
}

// newCocktailFilter returns the cocktailFilter associated to the given filter.
func newCocktailFilter(filter string) cocktailFilter {
	switch strings.ToLower(filter) {
	case idFltr.String():
		return idFltr
	case nameFltr.String():
		return nameFltr
	case alcoholicFltr.String():
		return alcoholicFltr
	case categoryFltr.String():
		return categoryFltr
	case ingredientFltr.String():
		return ingredientFltr
	case glassFltr.String():
		return glassFltr
	default:
		return invalidFltr
	}
}

func filterCocktailsById(id int, recs []entity.Cocktail) []entity.Cocktail {
	cocktails := make([]entity.Cocktail, 0)
	for _, rec := range recs {
		if id == rec.ID {
			cocktails = append(cocktails, rec)
		}
	}
	return cocktails
}

func filterCocktailsByName(name string, recs []entity.Cocktail) []entity.Cocktail {
	cocktails := make([]entity.Cocktail, 0)
	if name == "" {
		return cocktails
	}
	for _, rec := range recs {
		if strings.Contains(
			strings.ToLower(rec.Name),
			strings.ToLower(name),
		) {
			cocktails = append(cocktails, rec)
		}
	}
	return cocktails
}

func filterCocktailsByAlcoholic(name string, recs []entity.Cocktail) []entity.Cocktail {
	cocktails := make([]entity.Cocktail, 0)
	if name == "" {
		return cocktails
	}
	for _, rec := range recs {
		if strings.EqualFold(rec.Alcoholic, name) {
			cocktails = append(cocktails, rec)
		}
	}
	return cocktails
}

func filterCocktailsByCategory(name string, recs []entity.Cocktail) []entity.Cocktail {
	cocktails := make([]entity.Cocktail, 0)
	if name == "" {
		return cocktails
	}
	for _, rec := range recs {
		if strings.Contains(
			strings.ToLower(rec.Category),
			strings.ToLower(name),
		) {
			cocktails = append(cocktails, rec)
		}
	}
	return cocktails
}

func filterCocktailsByIngredient(name string, recs []entity.Cocktail) []entity.Cocktail {
	cocktails := make([]entity.Cocktail, 0)
	if name == "" {
		return cocktails
	}
	for _, rec := range recs {
		for _, ingr := range rec.Ingredients {
			if strings.Contains(
				strings.ToLower(ingr.Name),
				strings.ToLower(name),
			) {
				cocktails = append(cocktails, rec)
			}
		}
	}
	return cocktails
}

func filterCocktailsByGlass(name string, recs []entity.Cocktail) []entity.Cocktail {
	cocktails := make([]entity.Cocktail, 0)
	if name == "" {
		return cocktails
	}
	for _, rec := range recs {
		if strings.Contains(
			strings.ToLower(rec.Glass),
			strings.ToLower(name),
		) {
			cocktails = append(cocktails, rec)
		}
	}
	return cocktails
}

// findCocktail checks if the cocktail record exists in the given records list.
// if exists, returns its index and true.
func findCocktail(id int, recs []entity.Cocktail) (index int, found bool) {
	if len(recs) == 0 {
		return 0, false
	}
	for i, rec := range recs {
		if rec.ID == id {
			return i, true
		}
	}
	return 0, false
}

// cocktailsEqual compares two entity.Cocktail instances.
// If any field value not match returns false.
func cocktailsEqual(c1, c2 entity.Cocktail) bool {
	if c1.Name != c2.Name {
		return false
	}
	if c1.Alcoholic != c2.Alcoholic {
		return false
	}
	if c1.Category != c2.Category {
		return false
	}
	if c1.Instructions != c2.Instructions {
		return false
	}
	if c1.Glass != c2.Glass {
		return false
	}
	if c1.IBA != c2.IBA {
		return false
	}
	if c1.ImgAttribution != c2.ImgAttribution {
		return false
	}
	if c1.ImgSrc != c2.ImgSrc {
		return false
	}
	if c1.Tags != c2.Tags {
		return false
	}
	if c1.Thumb != c2.Thumb {
		return false
	}
	if c1.Video != c2.Video {
		return false
	}
	if len(c1.Ingredients) != len(c2.Ingredients) {
		return false
	}
	for _, iC1 := range c1.Ingredients {
		exists := false
		for _, iC2 := range c2.Ingredients {
			if iC1 == iC2 {
				exists = true
				break
			}
		}
		if !exists {
			return false
		}
	}

	return true
}
