//go:build integration
// +build integration

// Contract (black-box) tests that run the same behavioural spec against every
// hour.Repository implementation.
//
// These tests need real infrastructure:
//   - Firestore emulator reachable via GCP_PROJECT / FIRESTORE_EMULATOR_HOST
//   - MySQL reachable via MYSQL_ADDR / MYSQL_USER / MYSQL_PASSWORD / MYSQL_DATABASE
//
// They are excluded from `make test` (which is unit-test only) and are run by
// `make test-integration`, which passes `-tags=integration` to `go test`.

package main_test

import (
	"context"
	"math/rand"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	main "github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer"
	"github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/internal/trainer/domain/hour"
	"github.com/stretchr/testify/require"
)

// TestRepositoryContract runs every testXxx sub-test against every real
// repository implementation. The in-memory repository is also included here
// so the contract is still exercised end-to-end in integration runs (it is
// separately covered as a pure unit test in hour_repository_test.go).
func TestRepositoryContract(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())

	repositories := createRepositories(t)

	for i := range repositories {
		// When you are looping over slice and later using iterated value in goroutine (here because of t.Parallel()),
		// you need to always create variable scoped in loop body!
		// More info here: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
		r := repositories[i]

		t.Run(r.Name, func(t *testing.T) {
			// It's always a good idea to build all non-unit tests to be able to work in parallel.
			// Thanks to that, your tests will be always fast and you will not be afraid to add more tests because of slowdown.
			t.Parallel()

			t.Run("testUpdateHour", func(t *testing.T) {
				t.Parallel()
				testUpdateHour(t, r.Repository)
			})
			t.Run("testUpdateHour_parallel", func(t *testing.T) {
				t.Parallel()
				testUpdateHour_parallel(t, r.Repository)
			})
			t.Run("testHourRepository_update_existing", func(t *testing.T) {
				t.Parallel()
				testHourRepository_update_existing(t, r.Repository)
			})
			t.Run("testUpdateHour_rollback", func(t *testing.T) {
				t.Parallel()
				testUpdateHour_rollback(t, r.Repository)
			})
		})
	}
}

type Repository struct {
	Name       string
	Repository hour.Repository
}

func createRepositories(t *testing.T) []Repository {
	return []Repository{
		{
			Name:       "Firebase",
			Repository: newFirebaseRepository(t, context.Background()),
		},
		{
			Name:       "MySQL",
			Repository: newMySQLRepository(t),
		},
		{
			Name:       "memory",
			Repository: main.NewMemoryHourRepository(testHourFactory),
		},
	}
}

func newFirebaseRepository(t *testing.T, ctx context.Context) *main.FirestoreHourRepository {
	firebaseClient, err := firestore.NewClient(ctx, os.Getenv("GCP_PROJECT"))
	require.NoError(t, err)

	return main.NewFirestoreHourRepository(firebaseClient, testHourFactory)
}

func newMySQLRepository(t *testing.T) *main.MySQLHourRepository {
	db, err := main.NewMySQLConnection()
	require.NoError(t, err)

	return main.NewMySQLHourRepository(db, testHourFactory)
}
