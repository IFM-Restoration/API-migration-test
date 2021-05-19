package models

import "time"

// yellow tables + 1 red table with fourth and fifth dependencies
type JobNote struct {
	ID              uint
	JobId           []Job
	JobNoteTypeId   []JobNeedType
	JobNeedNoteId   uint // supposed to be foreign key
	Notes           string
	InternalOnly    bool
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type StateJobHistory struct {
	ID                uint
	JobStateId        []JobState
	StateExitReasonId []StateExitReason
	Entered           time.Time
	Exited            time.Time
	CreatedDatetime   time.Time
	CreatedBy         string
	LastModifiedBy    string
	IsDeleted         bool
}

type JobScheduleActivity struct {
	ID              uint
	JobScheduleId   uint // needs to be a table
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}
