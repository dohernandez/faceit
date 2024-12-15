# How to add an endpoint

This guide will show you how to add a new endpoint to the service.

Let's say we want to add a new endpoint to the service that will return the use info based on the given id. The id will be provided as a query parameter `v1/users/26ef0140-c436-4838-a271-32652c72f6f2`.

### Step 1: Add the endpoint to the proto file

To add the new endpoint to the proto file. Open the `resources/proto/user.proto` file and add the following code:

```proto
...

// The API to manages users.
service FaceitService {
    ...
    // Get user by id.
    rpc GetUser(UserID) returns (GetUserResponse) {
      option (google.api.http) = {
        get: "/v1/users/{id}"
      };
      option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
        responses: {
          key: "200"
          value: {
            description: "User information."
            schema: {
              json_schema: {
                ref: ".api.faceit.User"
              }
            }
          }
        }
        responses: {
          key: "404"
          value: {
            description: "User not found."
            schema: {
              json_schema: {
                ref: ".google.protobuf.Empty"
              }
            }
          }
        }
      };
    }
    ...
    }
```

### Step 2: Generate the code

Run the following command to generate the code based on the proto file(s) definition:

- gRPC server and client interfaces
- gRPC-gateway
- Swagger documentation

```shell
make proto-gen
```

### Step 3: Implement the use case

Create a new file `internal/domain/usecase/get_user.go` and add the following code:

- Define the finder interface to decouple the storage.
- Use `go:generate mockery` to generate the mock for the finder interface, usefully for testing.

```go
package usecase

import (
    "context"

	"github.com/bool64/ctxd"
    "github.com/faceit/go-grpc-service/internal/domain/model"
)

//go:generate mockery --name=UserFinder --outpkg=mocks --output=mocks --filename=user_finder.go --with-expecter

// UserFinder provides the use case to find a user.
type UserFinder interface {
	FindUser(ctx context.Context, id model.UserID) (*model.User, error)
}

// GetUser is a use case to get a user.
type GetUser struct {
	finder    UserFinder

	logger ctxd.Logger
}

// NewGetUser creates a new GetUser use case.
func NewGetUser(finder UserFinder, logger ctxd.Logger) *GetUser {
    return &GetUser{
        finder:    finder,
        logger: logger,
    }
}

// GetUser returns the user information based on the given id.
func (u *GetUser) GetUser(ctx context.Context, id model.UserID) (*model.User, error) {
	ctx = ctxd.AddFields(ctx, "use_case", "GetUser", "user_id", id)
	
    user, err := u.finder.FindUser(ctx, id)
    if err != nil {
        return nil, ctxd.WrapError(ctx, err, "find user") // error contains the context fields added
    }

    return user, nil
}
```

### Step 4: Implement the storage

Add a new functionality to `internal/platform/storage/user.go` to find a user by id:

```go
package storage

import (
    "context"
	"errors"
	
    ...

    "github.com/bool64/sqluct"
    "github.com/faceit/go-grpc-service/internal/domain/model"
	"github.com/dohernandez/go-grpc-service/database"
    "github.com/dohernandez/go-grpc-service/database/pgx"
)

// UserTable is the table name for users.
const UserTable = "users"

// User represents a User repository.
type User struct {
	storage *sqluct.Storage

	// col names for users table search
	colID      string
	colCountry string
}

...

// FindUser finds a user by id.
func (s *User) FindUser(ctx context.Context, id model.UserID) (*model.User, error) {
    var user model.User

    q := s.storage.SelectStmt(UserTable, model.User{}).
		Where(squirrel.Eq{s.colID: id})

	err := s.storage.Select(ctx, q, &user)
	if err != nil {
		if pgx.IsNoRows(err) {
            return nil, database.ErrNotFound
        }
		
		return nil, err
	}

    return &user, nil
}
```

### Step 5: Implement the service

Create a new file to `internal/platform/service/get_user.go` and add the new functionality to handle the new endpoint:

- Define the service use case interface to decouple the use case.
- Use `protovalidate` to validate the request.
- Use `github.com/dohernandez/servers` to handle the errors.

```go
package service

import (
    "context"
	"errors"

	"github.com/bufbuild/protovalidate-go"
    "github.com/faceit/go-grpc-service/internal/domain/model"
	api "github.com/dohernandez/faceit/internal/platform/service/pb"
	"github.com/dohernandez/go-grpc-service/database"
	"github.com/dohernandez/servers"
	"google.golang.org/grpc/codes"
)

// GetUser defines the use case to get user.
type GetUser interface {
	GetUser(ctx context.Context, id model.UserID) (*model.User, error)
}

// GetUser returns the user information based on the given id.
func (s *FaceitService) GetUser(ctx context.Context, req *api.UserID) (*api.User, error) {
	// Validate request.
	val, err := protovalidate.New(
		protovalidate.WithMessages(
			&api.UserID{},
		),
	)
	if err != nil {
		return nil, servers.WrapError(codes.Internal, err, "create proto validator")
	}

	fieldMsgErrs, ok := isUserValid(req, val, true)
	if !ok {
		return nil, servers.Error(codes.InvalidArgument, "validation error", fieldMsgErrs)
	}
	
	user, err := s.deps.GetUser().GetUser(ctx, model.UserID(req.Id))
    if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			return nil, servers.WrapError(codes.NotFound, err, "user not found")
		}
		
        return nil, servers.WrapError(codes.Internal, err, "ups, something went wrong!")
    }

    return &api.User{
        Id:       user.ID.String(),
		PasswordHash: user.PasswordHash,
		Email:    user.Email,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Nickname: user.Nickname,
		Country:  user.Country,
    }, nil
}
```

- Add the new use case to `deps` in the `app.go` file.

```go
package app

// FaceitServiceDeps holds the dependencies for the FaceitService.
type FaceitServiceDeps interface {
	...
	
	GetUser() GetUser
}
```

### Step 6: Integrate the new endpoint

Add the new endpoint by updating the `internal/platform/app/app.go` file:

- Add the new use case to the locator.
- Create the new use case deps in the locator (`func (l *Locator) setupUsecaseDependencies()`).
- Add the new use case to the locator (`func (l *Locator) GetUser() service.GetUser`).

```go
package app

import (
    ...

    "github.com/faceit/go-grpc-service/internal/domain/usecase"
    "github.com/faceit/go-grpc-service/internal/platform/service"
)

// Locator defines application resources.
type Locator struct {
	...

	// use cases
	...
	uGetUser usecase.GetUser
}

// setupUsecaseDependencies sets up use case dependencies (domain).
func (l *Locator) setupUsecaseDependencies() {
	...
	l.uGetUser = usecase.NewGetUser(l.storageUser, l.CtxdLogger())
}


// GetUser returns the usecase.GetUser use case.
func (l *Locator) GetUser() service.GetUser {
	return l.uGetUser
}
```

### Step 7: Test the new endpoint

All independent components should be tested (`usecase.GetUser` and `storage.User`). `FaceitService.GetUser` can be tested with the integration test.

- Add the test for the new use case to `internal/domain/usecase/get_user_test.go`. It is recommended to at least test the happy path and the error path.


```feature
Feature: GetUser
  As a user, I want to see the user information.

  Scenario: Get user by id
    Given these rows are stored in table "users" of database "postgres":
      | id                                   | first_name | last_name | nickname | password_hash                                                    | email         | country |
      | 26ef0140-c436-4838-a271-32652c72f6f2 | Alice      | Bob       |          | f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d | alice@bob.com | UK      |

    When I request HTTP endpoint with method "GET" and URI "/v1/users/26ef0140-c436-4838-a271-32652c72f6f2"
    
    Then I should have response with status "OK"
    And I should have response with header "Content-Type: application/json"
    And I should have response with body:
      """
      {
        "id": "26ef0140-c436-4838-a271-32652c72f6f2",
        "first_name": "Alice",
        "last_name": "Bob",
        "nickname": "",
        "password_hash": "f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d",
        "email": "alice@bob.com",
        "country": "UK"
      }
      """
```

### Step 8: Make sure the service test suite passes

Run the following command to run the test suite:

```shell
make check
```

### Step 9: Benchmark the new endpoint (optional)

In case the new endpoint is performance-critical, it is recommended to benchmark the new endpoint. Add the benchmark test to `benchmark_test.go` file:

- Add the benchmark test case for the new endpoint.

```go
... 

        TestCases: []service.BenchmarkCases{
			{
				Name:         "get list by country",
				Uri:          "/v1/users?country=UK",
				...
			},
			{
				Name:         "get user by id",
				Uri:          "/v1/users/26ef0140-c436-4838-a271-32652c72f6f2",
				ResponseCode: http.StatusOK,
				Data: map[string]any{
					storage.UserTable: model.User{
						ID: uuid.MustParse("26ef0140-c436-4838-a271-32652c72f6f2"),
						UserState: model.UserState{
							FirstName:    "Alice",
							LastName:     "Bob",
							PasswordHash: "f6b7e19e0d867de6c0391879050e8297165728d89d7c4e9e8839972b356c4d9d",
							Email:        "alice@bob.com",
							Country:      "UK",
						},
					},
				},
			},
		},
```

Run the following command to run the benchmark test:

```shell
make bench
```