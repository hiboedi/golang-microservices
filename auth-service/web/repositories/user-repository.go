package repositories

import (
	"auth-service/helpers"
	"auth-service/web/models"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	DeleteUser(ctx context.Context, db *gorm.DB, user models.User) error
	Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error)
	FindByEmail(ctx context.Context, db *gorm.DB, user models.User) (models.User, error)
	FindById(ctx context.Context, db *gorm.DB, user models.User) (models.User, error)
	FindAll(ctx context.Context, db *gorm.DB) ([]models.User, error)
}

type userRepositoryImpl struct {
}

func NewUserRepo() UserRepository {
	return &userRepositoryImpl{}
}

func (r *userRepositoryImpl) Create(ctx context.Context, db *gorm.DB, user models.User) (models.User, error) {
	err := db.WithContext(ctx).Create(&user).Error
	helpers.PanicIfError(err)

	return user, nil
}

func (r *userRepositoryImpl) FindAll(ctx context.Context, db *gorm.DB) ([]models.User, error) {
	var users []models.User

	err := db.WithContext(ctx).Model(&models.User{}).Find(&users).Error
	helpers.PanicIfError(err)

	return users, nil
}

func (r *userRepositoryImpl) DeleteUser(ctx context.Context, db *gorm.DB, user models.User) error {
	err := db.WithContext(ctx).Model(&models.User{}).Delete(&user).Error
	helpers.PanicIfError(err)

	return nil
}

func (r *userRepositoryImpl) FindById(ctx context.Context, db *gorm.DB, user models.User) (models.User, error) {
	err := db.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.ID).Take(&user).Error
	helpers.PanicIfError(err)

	return user, nil
}

func (r *userRepositoryImpl) FindByEmail(ctx context.Context, db *gorm.DB, user models.User) (models.User, error) {
	err := db.WithContext(ctx).Model(&models.User{}).Where("email = ?", user.Email).Take(&user).Error
	helpers.PanicIfError(err)

	return user, nil
}
