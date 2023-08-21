package controller

//const testDataDir = "../../test/data"
//
//func TestGet_Int(t *testing.T) {
//	cfg := configuration.NewCsvDB("fruits_valid.csv", testDataDir)
//	require.NotNil(t, cfg)
//	repo, err := repository.NewFruitCsv(cfg)
//	require.Nil(t, err)
//	require.NotNil(t, repo)
//	svc := service.NewFruit(repo)
//	require.NotNil(t, svc)
//	ctrl := NewFruitHTTP(svc)
//	require.NotNil(t, ctrl)
//
//	type params struct {
//		filter string
//		value  string
//	}
//	type resp struct {
//		code int
//		//body string // TODO: implement body validations
//	}
//	tests := []struct {
//		name    string
//		params  params
//		resp    resp
//		exp     []entity.Fruit
//		wantErr bool
//	}{
//		{
//			name:    "Empty",
//			params:  params{},
//			resp:    resp{code: http.StatusMovedPermanently},
//			wantErr: true,
//		},
//		{
//			name:    "Arbitrary",
//			params:  params{filter: "bad-filter", value: "foo"},
//			resp:    resp{code: http.StatusBadRequest},
//			wantErr: true,
//		},
//		{
//			name:    "Not Found",
//			params:  params{filter: "id", value: "1234"},
//			resp:    resp{code: http.StatusOK},
//			wantErr: true,
//		},
//		{
//			name:    "ID",
//			params:  params{filter: "id", value: "1"},
//			resp:    resp{code: http.StatusOK},
//			wantErr: false,
//		},
//		{
//			name:    "Name",
//			params:  params{filter: "name", value: "apple"},
//			resp:    resp{code: http.StatusOK},
//			wantErr: false,
//		},
//		{
//			name:    "Color",
//			params:  params{filter: "color", value: "green"},
//			resp:    resp{code: http.StatusOK},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			// Request
//			path := fmt.Sprintf(
//				"http://localhost:8080/api/v%d/fruit/%v/%v",
//				configuration.GetInstance().AppVersion.MajorVersion(),
//				tt.params.filter, tt.params.value)
//			req, err := http.NewRequest("GET", path, nil)
//			require.Nil(t, err)
//			require.NotNil(t, req)
//
//			// Server instance
//			rr := httptest.NewRecorder()
//			router := mux.NewRouter()
//			router.HandleFunc("/api/v0/fruit/{filter}/{value}", ctrl.GetFruit)
//			router.ServeHTTP(rr, req)
//
//			//t.Logf("CODE: %v", rr.Code)
//			//t.Logf("Body: %v", rr.Body.String())
//
//			assert.Equal(t, tt.resp.code, rr.Code)
//			if tt.wantErr {
//				// TODO: validate error
//				assert.NotNil(t, rr.Body)
//			}
//		})
//	}
//}
