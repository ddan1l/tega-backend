package user_repository_test

import (
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ddan1l/tega-backend/database"
	user_dto "github.com/ddan1l/tega-backend/dto/user"
	"github.com/ddan1l/tega-backend/models"
	user_repository "github.com/ddan1l/tega-backend/repositories/user"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestUserPgRepository_FindByID_Success(t *testing.T) {
	db, mock := database.NewMockDatabase()

	rows := sqlmock.NewRows([]string{"id", "email"}).
		AddRow(1, "test@example.com")

	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE "users"."id" = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2$`).
		WithArgs(1, 1).
		WillReturnRows(rows)

	repo := user_repository.NewUserPgRepository(db)
	user, err := repo.FindById(&user_dto.FindByIdDto{ID: 1})

	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", user.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserPgRepository_FindByID_NotFound(t *testing.T) {
	db, mock := database.NewMockDatabase()

	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE .*`).
		WithArgs(999, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	repo := user_repository.NewUserPgRepository(db)
	user, err := repo.FindById(&user_dto.FindByIdDto{ID: 999})

	assert.Error(t, err)
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Nil(t, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserPgRepository_Create_Success(t *testing.T) {
	db, mock := database.NewMockDatabase()
	repo := user_repository.NewUserPgRepository(db)

	createDto := &user_dto.CreateUserDto{
		FullName:     "John Doe",
		Email:        "john@example.com",
		PasswordHash: "hashed_password",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" \(.*\) VALUES \(.*\)`).
		WithArgs(
			sqlmock.AnyArg(), // updated_at
			sqlmock.AnyArg(), // deleted_at
			createDto.FullName,
			createDto.Email,
			createDto.PasswordHash,
		).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "created_at"}).
				AddRow(1, time.Now()),
		)
	mock.ExpectCommit()

	user, err := repo.Create(createDto)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, user.ID)
	assert.Equal(t, "John Doe", user.FullName)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserPgRepository_Create_DuplicateEmail(t *testing.T) {
	db, mock := database.NewMockDatabase()
	repo := user_repository.NewUserPgRepository(db)

	createDto := &user_dto.CreateUserDto{
		FullName:     "John Doe",
		Email:        "existing@example.com",
		PasswordHash: "hashed_password",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users" \(.*\) VALUES \(.*\)`).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			createDto.FullName,
			createDto.Email,
			createDto.PasswordHash,
		).
		WillReturnError(gorm.ErrDuplicatedKey)
	mock.ExpectRollback()

	user, err := repo.Create(createDto)

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.ErrorIs(t, err, gorm.ErrDuplicatedKey)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserPgRepository_FindByEmail_Success(t *testing.T) {
	db, mock := database.NewMockDatabase()
	repo := user_repository.NewUserPgRepository(db)

	email := "test@example.com"
	expectedUser := models.User{
		ID:    1,
		Email: email,
	}

	rows := sqlmock.NewRows([]string{"id", "email"}).
		AddRow(expectedUser.ID, expectedUser.Email)

	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE "users"."email" = \$1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT \$2$`).
		WithArgs(email, 1).
		WillReturnRows(rows)

	user, err := repo.FindByEmail(&user_dto.FindByEmailDto{Email: email})

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, expectedUser.ID, user.ID)
	assert.Equal(t, expectedUser.Email, user.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserPgRepository_FindByEmail_NotFound(t *testing.T) {
	db, mock := database.NewMockDatabase()
	repo := user_repository.NewUserPgRepository(db)

	email := "notfound@example.com"

	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE .*`).
		WithArgs(email, 1).
		WillReturnError(gorm.ErrRecordNotFound)

	user, err := repo.FindByEmail(&user_dto.FindByEmailDto{Email: email})

	assert.NoError(t, err)
	assert.Nil(t, user)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUserPgRepository_FindByEmail_DatabaseError(t *testing.T) {
	db, mock := database.NewMockDatabase()
	repo := user_repository.NewUserPgRepository(db)

	email := "error@example.com"
	expectedError := errors.New("database connection failed")

	mock.ExpectQuery(`^SELECT \* FROM "users" WHERE .*`).
		WithArgs(email, 1).
		WillReturnError(expectedError)

	user, err := repo.FindByEmail(&user_dto.FindByEmailDto{Email: email})

	assert.Error(t, err)
	assert.Nil(t, user)
	assert.Equal(t, expectedError, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
