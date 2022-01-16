package repository

import (
	"go-gorm-test/domain/models"

	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	UserGet() (*[]models.User, error)
	UserPost(Email string, PasswordDigest string) (*models.User, error)
	// UserPut(ID int) (*models.User, error)
	// UserDelete(ID int) (*models.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db}
}

func (ur *userRepositoryImpl) UserGet() (*[]models.User, error) {
	findUsers := []models.User{}
	err := ur.db.Find(&findUsers).Error
	if err != nil {
		return nil, err
	}
	return &findUsers, nil
}

func (ur *userRepositoryImpl) UserPost(Email string, PasswordDigest string) (*models.User, error) {
	insertUser := models.User{}
	insertUser.Email = Email
	insertUser.PasswordDigest = PasswordDigest
	err := ur.db.Create(&insertUser).Error
	if err != nil {
		return nil, err
	}
	return &insertUser, nil
}

// func (ur *userRepositoryImpl) UserPut(ID int) (*models.User, error) {
// 	user := models.User{ID: 1}
// 	user.Email = null.StringFrom("update@example.com")
// 	user.Update(context.Background(), db.DB, boil.Infer())

// 	return &user, nil
// }

// func (ur *userRepositoryImpl) UserDelete(ID int) (*models.User, error) {
// 	user := models.User{ID: ID}
// 	user.Delete(context.Background(), db.DB)

// 	return &user, nil
// }
