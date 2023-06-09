package activities

import (
	"time"
)

type ActivitiesEntities struct {
	ID        uint
	Title     string
	Email     string
	Createdat time.Time
	Updatedat time.Time
}

type ActivitiesService interface {
	FormData(newActivity ActivitiesEntities) (data ActivitiesEntities, row int, err error)
	GetActivity() ([]ActivitiesEntities, error)
	GetId(id int) (data ActivitiesEntities, row int, err error)
	Updata(id int, datup ActivitiesEntities) (ActivitiesEntities, error)
	Delete(id int) error
}

type ActivitiesData interface {
	FormData(newActivity ActivitiesEntities) (data ActivitiesEntities, row int, err error)
	GetActivity() ([]ActivitiesEntities, error)
	GetId(id int) (data ActivitiesEntities, row int, err error)
	Updata(id int, datup ActivitiesEntities) (ActivitiesEntities, error)
	Delete(id int) error
	UniqueData(insert ActivitiesEntities) (row int, err error)
}
