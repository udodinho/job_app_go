package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Status string

const (
	Pending      Status = "pending"
	Interviewing Status = "interviewing"
	Approved     Status = "approved"
	Rejected     Status = "rejected"
)

type Job struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id" validate:"required,uuid"`
	Company   string    `json:"company" validate:"required,lte=255"`
	Position  string    `json:"position" validate:"required,lte=255"`
	Status    Status    `json:"status" gorm:"default:pending"`
	CreatedBy uuid.UUID `json:"created_by" validate:"required,uuid"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type UpdateJob struct {
	ID        uuid.UUID `gorm:"primarykey" json:"id" validate:"required,uuid"`
	Company   string    `json:"company"`
	Position  string    `json:"position"`
	Status    Status    `json:"status"`
	CreatedBy uuid.UUID `json:"created_by" validate:"required,uuid"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (jb *Job) CreateJob() (*Job, error) {
	DB.Create(&jb)
	return jb, nil
}

func GetAllJob(createdBy uuid.UUID) ([]Job, error) {
	var jobs []Job

	DB.Where("created_by", createdBy).Find(&jobs)

	return jobs, nil
}

func GetSingleJob(id, createdBy uuid.UUID) (Job, *gorm.DB, error) {
	var job Job

	db := DB.Where("id", id).Where("created_by", createdBy).First(&job)

	return job, db, nil
}

func DeleteJob(id, createdBy uuid.UUID) (Job, error) {
	var job Job

	DB.Where("id", id).Where("created_by", createdBy).Delete(job)

	return job, nil
}
