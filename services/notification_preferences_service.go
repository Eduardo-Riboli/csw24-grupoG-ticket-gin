package services

import (
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/notification_preferences"
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
)

type NotificationPreferencesService struct {
    Repository *repositories.NotificationPreferencesRepository
}

func NewNotificationPreferencesService(repo *repositories.NotificationPreferencesRepository) *NotificationPreferencesService {
    return &NotificationPreferencesService{Repository: repo}
}

func (service *NotificationPreferencesService) GetAllPreferences() ([]entities.NotificationPreferencesResponse, error) {
    preferences, err := service.Repository.GetAll()
    if err != nil {
        return nil, err
    }

    var responses []entities.NotificationPreferencesResponse
    for _, preference := range preferences {
        responses = append(responses, entities.NotificationPreferencesResponse{
            ID:            preference.ID,
            UserID:        preference.UserID,
            ReceiveEmails: preference.ReceiveEmails,
        })
    }

    return responses, nil
}

func (service *NotificationPreferencesService) GetPreferenceByID(id uint) (entities.NotificationPreferencesResponse, error) {
    preference, err := service.Repository.GetByID(id)
    if err != nil {
        return entities.NotificationPreferencesResponse{}, err
    }

    return entities.NotificationPreferencesResponse{
        ID:            preference.ID,
        UserID:        preference.UserID,
        ReceiveEmails: preference.ReceiveEmails,
    }, nil
}

func (service *NotificationPreferencesService) CreatePreference(request entities.NotificationPreferencesCrRequest) (entities.NotificationPreferencesResponse, error) {
    preference := models.NotificationPreferences{
        UserID:        request.UserID,
        ReceiveEmails: request.ReceiveEmails,
    }

    createdPreference, err := service.Repository.Create(preference)
    if err != nil {
        return entities.NotificationPreferencesResponse{}, err
    }

    return entities.NotificationPreferencesResponse{
        ID:            createdPreference.ID,
        UserID:        createdPreference.UserID,
        ReceiveEmails: createdPreference.ReceiveEmails,
    }, nil
}

func (service *NotificationPreferencesService) UpdatePreference(id uint, request entities.NotificationPreferencesUpRequest) (entities.NotificationPreferencesResponse, error) {
    preference, err := service.Repository.GetByID(id)
    if err != nil {
        return entities.NotificationPreferencesResponse{}, err
    }

    preference.ReceiveEmails = request.ReceiveEmails

    updatedPreference, err := service.Repository.Update(preference)
    if err != nil {
        return entities.NotificationPreferencesResponse{}, err
    }

    return entities.NotificationPreferencesResponse{
        ID:            updatedPreference.ID,
        UserID:        updatedPreference.UserID,
        ReceiveEmails: updatedPreference.ReceiveEmails,
    }, nil
}

func (service *NotificationPreferencesService) DeletePreference(id uint) error {
    return service.Repository.Delete(id)
}