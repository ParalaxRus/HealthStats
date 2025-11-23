package storage

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/paralaxrus/health-project/dbsvc/internal/storage/model"
	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/require"
)

func NewMockUserDataSource(poolMock pgxmock.PgxPoolIface) *UserDataSource {
	return &UserDataSource{db: poolMock}
}

func TestFindUser(t *testing.T) {
	mock, err := pgxmock.NewPool()
	require.NoError(t, err)
	defer mock.Close()

	ds := NewMockUserDataSource(mock)
	now := time.Now()

	tests := []struct {
		name         string
		query        string
		email        string
		expectedErr  error
		expectedUser *model.User
		expectedRows *pgxmock.Rows
	}{
		{
			"existing",
			"SELECT \\* FROM users WHERE email = \\$1  LIMIT 1",
			"test@gmail.com",
			nil,
			&model.User{Name: "test", Email: "test@gmail.com", Created: now, Password: "test-pass"},
			mock.
				NewRows([]string{"id", "email", "name", "created_at", "password"}).
				AddRow(1, "test@gmail.com", "test", now, "test-pass"),
		},
		{
			"non-existing",
			"SELECT \\* FROM users WHERE email = \\$1  LIMIT 1",
			"test2@gmail.com",
			ErrUserNotFound,
			nil,
			nil,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			query := mock.ExpectQuery(tt.query).WithArgs(tt.email)
			if tt.expectedRows != nil {
				query.WillReturnRows(tt.expectedRows)
			}

			actual, err := ds.FindUser(context.Background(), NewIndex(-1, tt.email))
			if tt.expectedErr != nil {
				errors.Is(err, err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, tt.expectedUser, actual)

			require.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
