# Capstone - Fruit Stock
Fruit stock is a system to manage fruit storage in a fruit shop.


# System Features
- Runs a Rest API over a http.Server instance with chi.Mux support.
- The API configuration is set by system environment variables with the `CAPSTONE` word prefixed. If configuration is missed, set up the default one.
- Easy design to support more database drivers, e.g., JSON, PostgreSQL, MongoDB, etc.
- Builds, launch, and test operations can be done by make statements. Moreover, it is useful for CI/CD implementations.
- Full logging support through the zerolog library.
- Dates are based on `Epoch Unix Timestamp`.
- Clean Architecture implementation:
  - Entity
  - Controller
  - Service
  - Repository
- Patterns:
  - Singleton pattern
  - Dependency Injection
  - Repository pattern
- Handling Errors:
  - Error propagation.
  - Custom errors wrap the native error into it, useful for error type validations into the packages.
  - The controller implements an error handler to determine the HTTP error code, response status and native error description.

# How it works
You can retrieve all the fruits, a specific one by id and a filter list by Name,Color,Country.
- Filter name and filter values are case insensitive.
- Any wrong csv record is discarded.

To get all fruits in the storage, run a request to the following endpoint:
```
http://localhost:8080/api/v0/fruits
```
You can get a specific fruit by id and a filter list of them. The following filters are supported:

Filtering fruits by ID:
```
http://localhost:8080/api/v0/fruit/id/1
```
Filtering fruits by NAME:
```
http://localhost:8080/api/v0/fruit/name/apple
```
Filtering fruits by COLOR:
```
http://localhost:8080/api/v0/fruit/color/green
```
Filtering fruits by COUNTRY:
```
http://localhost:8080/api/v0/fruit/country/mexico
```

# Run
The API listens by default on the port `8080`. To run the API, execute the following command:
```
make run
```

# Testing
The controller, service, and repository layers implement unit tests with mock support.
To run tests of each layer of the clean architecture, use the following commands.
```
make test-controller
make test-service
make test-repository
```
Optional. Run all tests of the architecture.
```
make test-clean-architecture     
```
Run all tests of the system.
```
make test-all     
```

