 > [!NOTE]
 > I'm not an expert, and everything written here may not align with how other developers describe it. This is simply my understanding of the architecture and how I explain what is happening.

 ## How Does This Architecture Work?

At least for the __CRUD__ parts of the system, they work almost identically. Each component has the following structure:

1. Service
    - A service acts as an __interface__ to the __database__.
    - Each service contains methods to interact with the database for specific actions such as creating, reading, updating, and deleting records.
    - For example, the `user_service` contains all methods related to interacting with the _database_ for user operations.
    
2. Route
    - This is simply the __HTTP endpoint__ that exposes the service's functionality to the outside world.

Additionally, the architecture includes an __App__ and __Middlewares__ implementation.

______

### Why a App?

The __App__ structure provides flexibility. By creating an app when setting up the server, we can inject the __database__ into all services without having to pass it as an argument or use other workarounds. This approach also simplifies testing, as we can create a new app instance to isolate parts of the system during tests.

The App struct also holds the __router__ (or __server_mux__) for the same reason.

____

### Creating a service.

To understand how to create a new service in this architecture, let's walk through creating a service called __employee__.

Steps:
1. __Create a new package__ for the service inside the `services` folder.
2. __Define the service struct__ in a new file. For example:

``` go
type ServiceEmployee struct {
    App *app.Application
}
``` 

3. __Add methods__ to the service. For example:
    - `Create` / `Register`
    - `Update` 
    - And so on.
    
The function signature should reference the struct you created earlier:

``` go
func (service *ServiceEmployee) Create(name string, age int){
    // Your implementation here
}
```

> [!NOTE]
> The methods must be in the same package where the struct is defined. Otherwise, you will encounter a compile error.

#### Testing the service

When you build the application, it will search for tests in each service folder. If tests are found, they will be executed, and the application will only build if all tests pass. If no tests are found, the application will not build.

To write tests:

1. Create a `test` folder inside your service folder.
2. Use the `testing` package from Go to write your tests. For example:

``` go
func TestCreate(t *testing.T) {}
func TestCreate_EmptyFields(t *testing.T) {}
```

__Naming Tip__: Test functions should start with `Test`, followed by the method being tested (e.g., `Create`), and then a description of the test case after an underscore (`_`). For example, `TestCreate_EmptyFields` tests the `Create` method with empty input fields.

____

### Creating routes

Once you have a service, you need to create HTTP routes to expose its functionality.

Steps:

1.Define a routes struct in the same package as the service. For example:

```go
type RoutesEmployees struct {
    App *app.Application
}
```

This allows you to pass the database (or other dependencies) through the App struct.

2. __Create a function to register all routes__ for the service. For example:

```go
func RegisterRoutes(app *app.Application) {
	employeeRoutes := &RoutesEmployee{App: app}

	app.Router.Handle("/employee/create", middlewares.Chain(
		http.HandlerFunc(employeeRoutes.CreateRoute),

		middlewares.CorsMiddleware("GET"),
	))
}
```
> [!NOTE]
> For more details no middlewares, refer to the middleware documentation.

_____

#### Dependencies

- __Go dot env__ - github.com/joho/godotenv
- __Cert generator__ - golang.org/x/crypto/acme/autocert
- __Postgresql driver__ - github.com/lib/pq
- __To hash data__ golang.org/x/crypto/bcrypt