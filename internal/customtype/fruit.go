package customtype

import "fmt"

var _ fmt.Stringer = FilterFruitType("")

const (
	UndefinedFilterFruit FilterFruitType = "undefined"
	IdFilterFruit        FilterFruitType = "id"
	NameFilterFruit      FilterFruitType = "name"
	ColorFilterFruit     FilterFruitType = "color"
)

type FilterFruitType string

func (f FilterFruitType) String() string {
	return string(f)
}

type FilterFruit struct {
	Filter FilterFruitType
	Value  string
}

func NewFilterFruitType(filter string) FilterFruitType {
	switch filter {
	case IdFilterFruit.String():
		return IdFilterFruit
	case NameFilterFruit.String():
		return NameFilterFruit
	case ColorFilterFruit.String():
		return ColorFilterFruit
	default:
		return UndefinedFilterFruit
	}
}
