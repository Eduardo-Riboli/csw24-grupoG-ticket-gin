package repositories

import (
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{DB: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
    var users []models.User
    if err := r.DB.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func (r *UserRepository) GetByID(id uint) (models.User, error) {
    var user models.User
    if err := r.DB.First(&user, id).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
    if err := r.DB.Create(&user).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}

func (r *UserRepository) Update(user models.User) (models.User, error) {
    if err := r.DB.Save(&user).Error; err != nil {
        return models.User{}, err
    }
    return user, nil
}

func (r *UserRepository) Delete(id uint) error {
    if err := r.DB.Delete(&models.User{}, id).Error; err != nil {
        return err
    }
    return nil
}