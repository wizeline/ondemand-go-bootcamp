package controller

import (
	"net/http"

	ct "github.com/marcos-wz/capstone-go-bootcamp/internal/customtype"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var _ HTTP = Cocktail{}

// Cocktail configures the routes and handler functions of the Cocktail controller
type Cocktail struct {
	svc CocktailSvc
}

// CocktailSvc is the abstraction of the Cocktail service dependency.
type CocktailSvc interface {
	GetFiltered(filter, value string) ([]entity.Cocktail, error)
	GetAll() ([]entity.Cocktail, error)
	UpdateDB() (ct.DBOpsSummary, error)
}

// NewCocktail returns a new Cocktail controller implementation.
func NewCocktail(svc CocktailSvc) Cocktail {
	return Cocktail{
		svc: svc,
	}
}

// SetRoutes sets a fresh middleware stack for the handle functions and mounts the routes in the provided sub router.
func (c Cocktail) SetRoutes(r chi.Router) {
	r.Get("/cocktail/{filter}/{value}", c.getCocktail)
	r.Get("/cocktails", c.getCocktails)
	r.Get("/cocktail/updatedb", c.updateDB)
}

// getCocktail is a handler function that retrieve a list of filtered cocktails in the database in JSON format.
func (c Cocktail) getCocktail(w http.ResponseWriter, r *http.Request) {
	filter := chi.URLParam(r, "filter")
	value := chi.URLParam(r, "value")

	cocktails, err := c.svc.GetFiltered(filter, value)
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, cocktails)
}

// getCocktails is a handler function that retrieve all the cocktails in the database in JSON format.
func (c Cocktail) getCocktails(w http.ResponseWriter, r *http.Request) {
	cocktails, err := c.svc.GetAll()
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, cocktails)
}

// updateDB is a handler function that updates the database records from a public API.
func (c Cocktail) updateDB(w http.ResponseWriter, r *http.Request) {
	summary, err := c.svc.UpdateDB()
	if err != nil {
		errJSON(w, r, err)
		return
	}
	render.JSON(w, r, summary)
}
