package repository

import (
	"regexp"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/louistwiice/go/simplebase/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_List(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockUsers := []models.User{
		{
			ID: uuid.New().String(), Email: "mike@gmail.com", FirstName: "Mike", LastName: "Spenser", Password: "mikepassword", IsActive: false, IsStaff: false, IsSuperuser: false, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
		{
			ID: uuid.New().String(), Email: "jimmy@gmail.com", FirstName: "Jimmy", LastName: "Fallon", Password: "jimmypassword", IsActive: false, IsStaff: false, IsSuperuser: false, CreatedAt: time.Now(), UpdatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "is_active", "is_staff", "is_superuser", "created_at", "updated_at"}).
		AddRow(mockUsers[0].ID, mockUsers[0].Email, mockUsers[0].FirstName, mockUsers[0].LastName, mockUsers[0].IsActive, mockUsers[0].IsStaff, mockUsers[0].IsSuperuser, mockUsers[0].CreatedAt, mockUsers[0].UpdatedAt).
		AddRow(mockUsers[1].ID, mockUsers[1].Email, mockUsers[1].FirstName, mockUsers[1].LastName, mockUsers[1].IsActive, mockUsers[1].IsStaff, mockUsers[1].IsSuperuser, mockUsers[1].CreatedAt, mockUsers[1].UpdatedAt)

	query := regexp.QuoteMeta(`SELECT id, email, first_name, last_name, is_active, is_staff, is_superuser, created_at, updated_at FROM user`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WillReturnRows(rows)

	repo := NewUserSQL(db)
	result, err := repo.List()
	assert.NoError(t, err)
	assert.Equal(t, result[0].ID, mockUsers[0].ID)
	assert.Equal(t, result[0].FirstName, mockUsers[0].FirstName)
	assert.Equal(t, result[0].IsActive, mockUsers[0].IsActive)
}

func Test_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	id := uuid.New().String()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mockUser := models.User{
		ID: id, Email: "mike@gmail.com", FirstName: "Mike", LastName: "Spenser", Password: "mikepassword", IsActive: false, IsStaff: false, IsSuperuser: false, CreatedAt: time.Now(), UpdatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "email", "first_name", "last_name", "password", "is_active", "is_staff", "created_at", "updated_at"}).
		AddRow(id, mockUser.Email, mockUser.FirstName, mockUser.LastName, mockUser.Password, mockUser.IsActive, mockUser.IsStaff, mockUser.CreatedAt, mockUser.UpdatedAt)

	query := regexp.QuoteMeta(`SELECT id, email, first_name, last_name, password, is_active, is_staff, created_at, updated_at FROM user WHERE id=?`)
	prep := mock.ExpectPrepare(query)
	prep.ExpectQuery().WithArgs(id).WillReturnRows(rows)

	repo := NewUserSQL(db)
	result, err := repo.Get(id)
	assert.NoError(t, err)
	assert.Equal(t, result.FirstName, mockUser.FirstName)
	assert.Equal(t, result.ID, mockUser.ID)
	assert.Equal(t, result.Password, mockUser.Password)
}
