package main

import "time"

// blue tables with at least one foreign key
type WorkOrder struct {
	ID                          uint
	ClientId                    []Client
	ResidentId                  []Resident
	WorkOrderAcceptanceStatusId []WorkOrderAcceptanceStatus
	ClientContactApprovedById   uint
	Name                        string
	Uuid                        string
	DueDate                     time.Time
	IsCanceled                  bool
	CreatedDatetime             time.Time
	CreatedBy                   string
	LastModifiedDatetime        time.Time
	LastModifiedBy              string
	IsDeleted                   bool
}

type Resident struct {
	ID              uint
	ClientId        []Client
	FirstName       string
	LastName        string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}
