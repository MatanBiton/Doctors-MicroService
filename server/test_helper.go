package main

import (
	"context"
	"database/sql"
	"log"

	ms "github.com/TekClinic/MicroService-Lib"
	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3" // Import SQLite driver
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"

	// Use the correct package based on `ms.Claims` requirements
	sets "k8s.io/apimachinery/pkg/util/sets"
)

// mockClaims is a mock implementation of ms.Claims interface for testing.
type mockClaims struct {
	userID string
	roles  []string
}

// HasRole checks if the mock claims contain the specified role.
func (c *mockClaims) HasRole(role string) bool {
	for _, r := range c.roles {
		if r == role {
			return true
		}
	}
	return false
}

// GetRoles returns all roles in the correct set type.
func (c *mockClaims) GetRoles() sets.Set[string] {
	roleSet := sets.New[string]() // ✅ Correctly initializes a set
	for _, r := range c.roles {
		roleSet.Insert(r)
	}
	return roleSet
}

// UserID returns the user ID.
func (c *mockClaims) UserID() string {
	return c.userID
}

// Mock VerifyToken to avoid real authentication
func (server *doctorsServer) VerifyToken(ctx context.Context, token string) (ms.Claims, error) {
	return &mockClaims{
		userID: "test_user",
		roles:  []string{"admin"},
	}, nil
}

// createTestDoctorsServer initializes an in-memory SQLite database for testing.
func createTestDoctorsServer() (*doctorsServer, error) {
	// Open an in-memory SQLite database
	sqlDB, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
		return nil, err
	}

	// Initialize Bun with SQLite dialect
	db := bun.NewDB(sqlDB, sqlitedialect.New())

	// Verify the database connection
	if db == nil {
		log.Fatalf("Bun DB is nil! Database not initialized")
		return nil, err
	}

	// Initialize BaseServiceServer (mocked if real one fails)
	baseService, err := ms.CreateBaseServiceServer()
	if err != nil {
		log.Printf("Skipping BaseServiceServer initialization for tests: %v", err)
		baseService = nil // ✅ Fix: Don't initialize as a struct if it's an interface
	}

	// Create the test schema
	err = createTestSchema(context.Background(), db)
	if err != nil {
		log.Fatalf("Failed to create test schema: %v", err)
		return nil, err
	}

	// Return the test server
	return &doctorsServer{
		BaseServiceServer: baseService, // ✅ Fix: Use `nil` instead of `{}` if interface
		db:                db,
		validate:          validator.New(validator.WithRequiredStructEnabled()),
	}, nil
}

// createTestSchema sets up the SQLite schema for testing without modifying production schema.
func createTestSchema(ctx context.Context, db *bun.DB) error {
	models := []interface{}{
		(*Doctor)(nil),
	}

	// Create tables
	for _, model := range models {
		if _, err := db.NewCreateTable().IfNotExists().Model(model).Exec(ctx); err != nil {
			return err
		}
	}

	return nil
}
