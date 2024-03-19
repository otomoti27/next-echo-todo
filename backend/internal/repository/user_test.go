package repository_test

import (
	"echo-api/domain"
	"echo-api/internal/repository"
	"echo-api/internal/repository/mocks"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserRepository_Create(t *testing.T) {
	gormDB, sqlMock, err := mocks.NewMockSQL()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}

	db, err := gormDB.DB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(gormDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta("INSERT INTO `users` (`email`,`password`,`name`,`created_at`,`updated_at`) VALUES (?,?,?,?,?)")).
		WithArgs("test@example.com", "password", "テストユーザー", sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	user := &domain.User{Name: "テストユーザー", Email: "test@example.com", Password: "password"}
	if err := repo.Create(user); err != nil {
		t.Errorf("ユーザーの作成に失敗しました: %v", err)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("failed to ExpectationWerMet(): %s", err)
	}
}

func TestUserRepository_GetByEmail(t *testing.T) {
	gormDB, sqlMock, err := mocks.NewMockSQL()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}

	db, err := gormDB.DB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(gormDB)

	testTime := time.Now()

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE email = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs("test@example.com", 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "name", "created_at", "updated_at"}).
			AddRow(1, "test@example.com", "password", "テストユーザー", testTime, testTime))

	user := &domain.User{}
	if err := repo.GetByEmail(user, "test@example.com"); err != nil {
		t.Errorf("ユーザーの取得に失敗しました: %v", err)
	}

	if user.Email != "test@example.com" {
		t.Errorf("期待されたユーザーが取得されませんでした: %v", user)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("failed to ExpectationWerMet(): %s", err)
	}
}

func TestUserRepository_GetByID(t *testing.T) {
	gormDB, sqlMock, err := mocks.NewMockSQL()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}

	db, err := gormDB.DB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(gormDB)

	testTime := time.Now()

	sqlMock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE `users`.`id` = ? ORDER BY `users`.`id` LIMIT ?")).
		WithArgs(1, 1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "name", "created_at", "updated_at"}).
			AddRow(1, "test@example.com", "password", "テストユーザー", testTime, testTime))

	user := &domain.User{}
	if err := repo.GetByID(user, 1); err != nil {
		t.Errorf("ユーザーの取得に失敗しました: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("期待されたユーザーが取得されませんでした: %v", user)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("failed to ExpectationWerMet(): %s", err)
	}
}

func TestUserRepository_Update(t *testing.T) {
	gormDB, sqlMock, err := mocks.NewMockSQL()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}

	db, err := gormDB.DB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(gormDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta("UPDATE `users` SET `email`=?,`password`=?,`name`=?,`created_at`=?,`updated_at`=? WHERE `id` = ?")).
		WithArgs("test@example.com", "password", "テストユーザー", sqlmock.AnyArg(), sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	user := &domain.User{ID: 1, Name: "テストユーザー", Email: "test@example.com", Password: "password"}
	if err := repo.Update(user); err != nil {
		t.Errorf("ユーザーの更新に失敗しました: %v", err)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("failed to ExpectationWerMet(): %s", err)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	gormDB, sqlMock, err := mocks.NewMockSQL()
	if err != nil {
		t.Fatalf("sqlmockの作成に失敗しました: %s", err)
	}

	db, err := gormDB.DB()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(gormDB)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(regexp.QuoteMeta("DELETE FROM `users` WHERE  `users`.`id` = ?")).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	sqlMock.ExpectCommit()

	user := &domain.User{ID: 1}
	if err := repo.Delete(user); err != nil {
		t.Errorf("ユーザーの削除に失敗しました: %v", err)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("failed to ExpectationWerMet(): %s", err)
	}
}
