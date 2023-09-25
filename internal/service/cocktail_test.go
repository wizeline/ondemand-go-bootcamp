package service

import (
	"errors"
	"strconv"
	"testing"
	"time"

	ct "github.com/marcos-wz/capstone-go-bootcamp/internal/customtype"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/service/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCocktail(t *testing.T) {
	repo := mocks.NewCocktailRepo()
	require.NotNil(t, repo)
	out := NewCocktail(repo)
	assert.IsType(t, Cocktail{}, out)
}

func TestCocktail_GetFiltered(t *testing.T) {
	type args struct {
		filter string
		value  string
	}
	type repo struct {
		resp []entity.Cocktail
		err  error
	}
	tests := []struct {
		name string
		args args
		exp  []entity.Cocktail
		err  error
		repo repo
	}{
		{
			name: "Repository error",
			args: args{
				filter: idFltr.String(),
				value:  "4",
			},
			exp: nil,
			err: testRepoErr,
			repo: repo{
				resp: nil,
				err:  testRepoErr,
			},
		},
		{
			name: "Value empty",
			args: args{
				filter: idFltr.String(),
				value:  "",
			},
			exp:  nil,
			err:  &FilterErr{ErrFltrValueEmpty},
			repo: repo{},
		},
		{
			name: "Arbitrary",
			args: args{filter: "foo", value: "foo-value"},
			exp:  nil,
			err:  &FilterErr{ErrFltrInvalid},
			repo: repo{},
		},
		{
			name: "Value not found",
			args: args{filter: idFltr.String(), value: "123456"},
			exp:  []entity.Cocktail{},
			err:  nil,
			repo: repo{
				resp: []entity.Cocktail{},
				err:  nil,
			},
		},
		{
			name: "Bad ID",
			args: args{filter: idFltr.String(), value: "foo-id"},
			exp:  []entity.Cocktail{},
			err:  &FilterErr{&strconv.NumError{}},
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
		},
		{
			name: "ID",
			args: args{filter: idFltr.String(), value: "2"},
			exp: []entity.Cocktail{
				{ID: 2, Name: "Bar", Alcoholic: "Non alcoholic", Category: "Some Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "water", Measure: "50ml"}}},
			},
			err: nil,
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
		},
		{
			name: "Name",
			args: args{filter: nameFltr.String(), value: "bAz"},
			exp: []entity.Cocktail{
				{ID: 3, Name: "Baz", Alcoholic: "Alcoholic", Category: "Some Category", Glass: "Cocktail glass", Ingredients: []entity.Ingredient{{Name: "soda", Measure: "100ml"}}},
			},
			err: nil,
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
		},
		{
			name: "Alcoholic",
			args: args{filter: alcoholicFltr.String(), value: "non alcoholic"},
			exp: []entity.Cocktail{
				{ID: 2, Name: "Bar", Alcoholic: "Non alcoholic", Category: "Some Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "water", Measure: "50ml"}}},
			},
			err: nil,
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
		},
		{
			name: "Category",
			args: args{filter: categoryFltr.String(), value: "some"},
			exp: []entity.Cocktail{
				{ID: 2, Name: "Bar", Alcoholic: "Non alcoholic", Category: "Some Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "water", Measure: "50ml"}}},
				{ID: 3, Name: "Baz", Alcoholic: "Alcoholic", Category: "Some Category", Glass: "Cocktail glass", Ingredients: []entity.Ingredient{{Name: "soda", Measure: "100ml"}}},
			},
			err: nil,
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
		},
		{
			name: "Ingredients",
			args: args{filter: ingredientFltr.String(), value: "soda"},
			exp: []entity.Cocktail{
				{ID: 1, Name: "Foo", Alcoholic: "Alcoholic", Category: "Foo Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "soda", Measure: "80ml"}}},
				{ID: 3, Name: "Baz", Alcoholic: "Alcoholic", Category: "Some Category", Glass: "Cocktail glass", Ingredients: []entity.Ingredient{{Name: "soda", Measure: "100ml"}}},
			},
			err: nil,
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
		},
		{
			name: "Glass",
			args: args{filter: glassFltr.String(), value: "shot"},
			exp: []entity.Cocktail{
				{ID: 1, Name: "Foo", Alcoholic: "Alcoholic", Category: "Foo Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "soda", Measure: "80ml"}}},
				{ID: 2, Name: "Bar", Alcoholic: "Non alcoholic", Category: "Some Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "water", Measure: "50ml"}}},
			},
			err: nil,
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mRepo := mocks.NewCocktailRepo()
			mRepo.On("ReadAll").Return(tt.repo.resp, tt.repo.err)
			svc := NewCocktail(mRepo)
			require.NotEqual(t, Cocktail{}, svc)

			out, err := svc.GetFiltered(tt.args.filter, tt.args.value)
			if tt.err != nil {
				require.NotNil(t, err)
				require.Nil(t, out)
				assert.IsType(t, tt.err, err)
				if errW := errors.Unwrap(tt.err); errW != nil {
					assert.IsType(t, errW, errors.Unwrap(err))
				}
				return
			}
			require.Nil(t, err)
			require.NotNil(t, out)
			assert.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestCocktail_GetAll(t *testing.T) {
	type repo struct {
		resp []entity.Cocktail
		err  error
	}
	tests := []struct {
		name    string
		exp     []entity.Cocktail
		err     error
		repo    repo
		wantErr bool
	}{
		{
			name: "Repository Error",
			exp:  nil,
			err:  testRepoErr,
			repo: repo{
				resp: nil,
				err:  testRepoErr,
			},
			wantErr: true,
		},
		{
			name: "Not Found",
			exp:  []entity.Cocktail{},
			err:  nil,
			repo: repo{
				resp: []entity.Cocktail{},
				err:  nil,
			},
			wantErr: false,
		},
		{
			name: "All records",
			exp:  testCocktailsAll,
			err:  nil,
			repo: repo{
				resp: testCocktailsAll,
				err:  nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mRepo := mocks.NewCocktailRepo()
			mRepo.On("ReadAll").Return(tt.repo.resp, tt.repo.err)
			svc := NewCocktail(mRepo)
			require.NotNil(t, svc)

			out, err := svc.GetAll()
			if tt.wantErr {
				require.NotNil(t, err)
				assert.Nil(t, out)
				assert.IsType(t, tt.err, err)
				return
			}
			require.Nil(t, err)
			require.NotNil(t, out)
			assert.Len(t, tt.exp, len(out))
			assert.Equal(t, tt.exp, out)
		})
	}
}

func TestCocktail_UpdateDB(t *testing.T) {
	type repo struct {
		createArg []entity.Cocktail
		createErr error
		fetchResp []entity.Cocktail
		fetchErr  error
		readResp  []entity.Cocktail
		readErr   error
	}
	tests := []struct {
		name    string
		repo    repo
		exp     ct.DBOpsSummary
		err     error
		wantErr bool
	}{
		{
			name: "Read All error",
			repo: repo{
				readResp: nil,
				readErr:  testRepoErr,
			},
			exp:     ct.DBOpsSummary{},
			err:     testRepoErr,
			wantErr: true,
		},

		{
			name: "Fetch Data error",
			repo: repo{
				readResp:  testCocktailsAll,
				readErr:   nil,
				fetchResp: nil,
				fetchErr:  testRepoErr,
			},
			exp:     ct.DBOpsSummary{},
			err:     testRepoErr,
			wantErr: true,
		},
		{
			name: "Replace Data error",
			repo: repo{
				readResp:  []entity.Cocktail{},
				readErr:   nil,
				fetchResp: testCocktailsAll,
				fetchErr:  nil,
				createArg: []entity.Cocktail{
					{ID: 1, Name: "Foo", Alcoholic: "Alcoholic", Category: "Foo Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "soda", Measure: "80ml"}}, CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
					{ID: 2, Name: "Bar", Alcoholic: "Non alcoholic", Category: "Some Category", Glass: "Shot glass", Ingredients: []entity.Ingredient{{Name: "water", Measure: "50ml"}}, CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
					{ID: 3, Name: "Baz", Alcoholic: "Alcoholic", Category: "Some Category", Glass: "Cocktail glass", Ingredients: []entity.Ingredient{{Name: "soda", Measure: "100ml"}}, CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
				},
				createErr: testRepoErr,
			},
			exp:     ct.DBOpsSummary{},
			err:     testRepoErr,
			wantErr: true,
		},
		{
			name: "No changes",
			repo: repo{
				readResp:  testCocktailsAll,
				readErr:   nil,
				fetchResp: testCocktailsAll,
				fetchErr:  nil,
				createArg: []entity.Cocktail{},
				createErr: nil,
			},
			exp: ct.DBOpsSummary{
				Status:       noChangesDBStatus,
				NewRecs:      0,
				ModifiedRecs: 0,
				TotalOps:     0,
				TotalRecs:    3,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "Same source date, but different record",
			repo: repo{
				readResp: []entity.Cocktail{
					{ID: 1, Name: "foo"},
					{ID: 2, Name: "bar", Category: "some-category"},
					{ID: 3, Name: "baz"},
				},
				readErr: nil,
				fetchResp: []entity.Cocktail{
					{ID: 1, Name: "foo"},
					{ID: 2, Name: "bar", Category: "other-category"},
					{ID: 3, Name: "baz"},
				},
				fetchErr: nil,
				createArg: []entity.Cocktail{
					{ID: 1, Name: "foo"},
					{ID: 2, Name: "bar", Category: "other-category", UpdatedAt: dateTimeNow()},
					{ID: 3, Name: "baz"},
				},
				createErr: nil,
			},
			exp: ct.DBOpsSummary{
				Status:       successfulUpdateDBStatus,
				NewRecs:      0,
				ModifiedRecs: 1,
				TotalOps:     1,
				TotalRecs:    3,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "One updated",
			repo: repo{
				readResp: []entity.Cocktail{
					{ID: 1, Name: "foo", SrcDate: dateTimeNow()},
					{ID: 2, Name: "bar"},
					{ID: 3, Name: "baz"},
				},
				readErr: nil,
				fetchResp: []entity.Cocktail{
					{ID: 1, Name: "foo", Category: "fooCategory", SrcDate: dateTimeNow().Add(1 * time.Hour)},
					{ID: 2, Name: "bar"},
					{ID: 3, Name: "baz"},
				},
				fetchErr: nil,
				createArg: []entity.Cocktail{
					{ID: 1, Name: "foo", Category: "fooCategory", SrcDate: dateTimeNow().Add(1 * time.Hour), UpdatedAt: dateTimeNow()},
					{ID: 2, Name: "bar"},
					{ID: 3, Name: "baz"},
				},
				createErr: nil,
			},
			exp: ct.DBOpsSummary{
				Status:       successfulUpdateDBStatus,
				NewRecs:      0,
				ModifiedRecs: 1,
				TotalOps:     1,
				TotalRecs:    3,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "Two new",
			repo: repo{
				readResp: []entity.Cocktail{
					{ID: 1, Name: "foo"},
				},
				readErr: nil,
				fetchResp: []entity.Cocktail{
					{ID: 1, Name: "foo"},
					{ID: 2, Name: "bar"},
					{ID: 3, Name: "baz"},
				},
				fetchErr: nil,
				createArg: []entity.Cocktail{
					{ID: 1, Name: "foo"},
					{ID: 2, Name: "bar", CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
					{ID: 3, Name: "baz", CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
				},
				createErr: nil,
			},
			exp: ct.DBOpsSummary{
				Status:       successfulUpdateDBStatus,
				NewRecs:      2,
				ModifiedRecs: 0,
				TotalOps:     2,
				TotalRecs:    3,
			},
			err:     nil,
			wantErr: false,
		},
		{
			name: "All new",
			repo: repo{
				readResp: []entity.Cocktail{},
				readErr:  nil,
				fetchResp: []entity.Cocktail{
					{ID: 1, Name: "foo"},
					{ID: 2, Name: "bar"},
					{ID: 3, Name: "baz"},
				},
				fetchErr: nil,
				createArg: []entity.Cocktail{
					{ID: 1, Name: "foo", CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
					{ID: 2, Name: "bar", CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
					{ID: 3, Name: "baz", CreatedAt: dateTimeNow(), UpdatedAt: dateTimeNow()},
				},
				createErr: nil,
			},
			exp: ct.DBOpsSummary{
				Status:       successfulUpdateDBStatus,
				NewRecs:      3,
				ModifiedRecs: 0,
				TotalOps:     3,
				TotalRecs:    3,
			},
			err:     nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mRepo := mocks.NewCocktailRepo()
			mRepo.On("ReadAll").Return(tt.repo.readResp, tt.repo.readErr)
			mRepo.On("ReplaceData", tt.repo.createArg).Return(tt.repo.createErr)
			mRepo.On("FetchData").Return(tt.repo.fetchResp, tt.repo.fetchErr)
			svc := NewCocktail(mRepo)
			require.NotNil(t, svc)

			out, err := svc.UpdateDB()
			if tt.wantErr {
				require.NotNil(t, err)
				assert.Equal(t, ct.DBOpsSummary{}, out)
				assert.IsType(t, tt.err, err)
				return
			}
			require.Nil(t, err)
			assert.Equal(t, tt.exp.Status, out.Status)
			assert.Equal(t, tt.exp.NewRecs, out.NewRecs)
			assert.Equal(t, tt.exp.ModifiedRecs, out.ModifiedRecs)
			assert.Equal(t, tt.exp.TotalOps, out.TotalOps)
			assert.Equal(t, tt.exp.TotalRecs, out.TotalRecs)
		})
	}
}
