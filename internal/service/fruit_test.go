package service

// TODO: implement mock tests

//func TestFruit_Get_Int(t *testing.T) {
//	cfg := configuration.NewCsvDB("fruits_valid.csv", testDataDir)
//	repo := repository.NewFruitCsv(cfg)
//	svc := NewFruit(repo)
//
//	out, err := svc.Get("color", "green")
//	t.Logf("ERROR: %v", err)
//	t.Logf("OUT: %v", out)
//}

//{
//	name:    "Configuration Empty",
//	cfg:     configuration.NewCsvDB("", ""),
//	err:     Err(""),
//	wantErr: true,
//},
//{
//	name:     "Arbitrary data file",
//	fileName: "foo.csv",
//	//err:      &Err{},
//	wantErr: true,
//},
//{
//	name: "Invalid data file ",
//},
//{
//	name: "Filter by ID",
//	cfg:  configuration.NewCsvDB("fruits_valid.csv", testDataDir),
//	args: args{filter: "id", value: "5"},
//	exp: []entity.Fruit{
//		{ID: 5, Name: "orange", Color: "orange"},
//	},
//	err:     nil,
//	wantErr: false,
//},
//{
//	name: "Filter by Name",
//	cfg:  configuration.NewCsvDB("fruits_valid.csv", testDataDir),
//	args: args{filter: "name", value: "apple"},
//	exp: []entity.Fruit{
//		{ID: 1, Name: "apple", Color: "red"},
//		{ID: 2, Name: "apple", Color: "green"},
//	},
//	err:     nil,
//	wantErr: false,
//},
//{
//	name: "Filter by Color",
//	cfg:  configuration.NewCsvDB("fruits_valid.csv", testDataDir),
//	args: args{filter: "color", value: "green"},
//	exp: []entity.Fruit{
//		{ID: 2, Name: "apple", Color: "green"},
//		{ID: 3, Name: "pear", Color: "green"},
//		{ID: 8, Name: "lime", Color: "green"},
//		{ID: 9, Name: "grape", Color: "green"},
//	},
//	err:     nil,
//	wantErr: false,
//},
