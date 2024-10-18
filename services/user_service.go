package services

import (
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/user"
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
)

type UserService struct {
    Repository *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{Repository: repo}
}

func (userService *UserService) GetAllUsers() ([]entities.UserResponse, error) {
    users, err := userService.Repository.GetAll()
    if err != nil {
        return nil, err
    }

    var userResponses []entities.UserResponse
    for _, user := range users {
        userResponses = append(userResponses, entities.UserResponse{
            ID:       user.ID,
            TenantID: user.TenantID,
            Name:     user.Name,
            Email:    user.Email,
        })
    }

    return userResponses, nil
}

func (userService *UserService) GetUserByID(id uint) (entities.UserResponse, error) {
    user, err := userService.Repository.GetByID(id)
    if err != nil {
        return entities.UserResponse{}, err
    }

    return entities.UserResponse{
        ID:       user.ID,
        TenantID: user.TenantID,
        Name:     user.Name,
        Email:    user.Email,
    }, nil
}

func (userService *UserService) CreateUser(userRequest entities.UserCrRequest) (entities.UserResponse, error) {
    user := models.User{
        TenantID: userRequest.TenantID,
        Name:     userRequest.Name,
        Email:    userRequest.Email,
    }

    createdUser, err := userService.Repository.Create(user)
    if err != nil {
        return entities.UserResponse{}, err
    }

    return entities.UserResponse{
        ID:       createdUser.ID,
        TenantID: createdUser.TenantID,
        Name:     createdUser.Name,
        Email:    createdUser.Email,
    }, nil
}

func (userService *UserService) UpdateUser(id uint, userRequest entities.UserUpRequest) (entities.UserResponse, error) {
    user, err := userService.Repository.GetByID(id)
    if err != nil {
        return entities.UserResponse{}, err
    }

    if userRequest.Name != "" {
        user.Name = userRequest.Name
    }
    if userRequest.Email != "" {
        user.Email = userRequest.Email
    }

    updatedUser, err := userService.Repository.Update(user)
    if err != nil {
        return entities.UserResponse{}, err
    }

    return entities.UserResponse{
        ID:       updatedUser.ID,
        TenantID: updatedUser.TenantID,
        Name:     updatedUser.Name,
        Email:    updatedUser.Email,
    }, nil
}

func (userService *UserService) DeleteUser(id uint) error {
    return userService.Repository.Delete(id)
}