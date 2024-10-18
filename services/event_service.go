package services

import (
    entities "github.com/grupoG/csw24-grupoG-ticket-gin/entities/event"
    "github.com/grupoG/csw24-grupoG-ticket-gin/models"
    "github.com/grupoG/csw24-grupoG-ticket-gin/repositories"
)

type EventService struct {
    Repository *repositories.EventRepository
}

func NewEventService(repo *repositories.EventRepository) *EventService {
    return &EventService{Repository: repo}
}

func (eventService *EventService) GetAllEvents() ([]entities.EventResponse, error) {
    events, err := eventService.Repository.GetAll()
    if err != nil {
        return nil, err
    }

    var eventResponses []entities.EventResponse
    for _, event := range events {
        eventResponses = append(eventResponses, entities.EventResponse{
            ID:       event.ID,
            TenantID: event.TenantID,
            Name:     event.Name,
            Location: event.Location,
            Date:     event.Date,
        })
    }

    return eventResponses, nil
}

func (eventService *EventService) GetEventByID(id uint) (entities.EventResponse, error) {
    event, err := eventService.Repository.GetByID(id)
    if err != nil {
        return entities.EventResponse{}, err
    }

    return entities.EventResponse{
        ID:       event.ID,
        TenantID: event.TenantID,
        Name:     event.Name,
        Location: event.Location,
        Date:     event.Date,
    }, nil
}

func (eventService *EventService) CreateEvent(eventRequest entities.EventCrRequest) (entities.EventResponse, error) {
    event := models.Event{
        TenantID: eventRequest.TenantID,
        Name:     eventRequest.Name,
        Location: eventRequest.Location,
        Date:     eventRequest.Date,
    }

    createdEvent, err := eventService.Repository.Create(event)
    if err != nil {
        return entities.EventResponse{}, err
    }

    return entities.EventResponse{
        ID:       createdEvent.ID,
        TenantID: createdEvent.TenantID,
        Name:     createdEvent.Name,
        Location: createdEvent.Location,
        Date:     createdEvent.Date,
    }, nil
}

func (eventService *EventService) UpdateEvent(id uint, eventRequest entities.EventUpRequest) (entities.EventResponse, error) {
    event, err := eventService.Repository.GetByID(id)
    if err != nil {
        return entities.EventResponse{}, err
    }

    if eventRequest.Name != "" {
        event.Name = eventRequest.Name
    }
    if eventRequest.Location != "" {
        event.Location = eventRequest.Location
    }
    if !eventRequest.Date.IsZero() {
        event.Date = eventRequest.Date
    }

    updatedEvent, err := eventService.Repository.Update(event)
    if err != nil {
        return entities.EventResponse{}, err
    }

    return entities.EventResponse{
        ID:       updatedEvent.ID,
        TenantID: updatedEvent.TenantID,
        Name:     updatedEvent.Name,
        Location: updatedEvent.Location,
        Date:     updatedEvent.Date,
    }, nil
}

func (eventService *EventService) DeleteEvent(id uint) error {
    return eventService.Repository.Delete(id)
}