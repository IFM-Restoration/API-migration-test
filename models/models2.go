package models

import "time"

// blue tables with at least one foreign key
type WorkOrder struct {
	ID                          uint
	ClientId                    []Client
	ResidentId                  []Resident
	WorkOrderAcceptanceStatusId []WorkOrderAcceptanceStatus
	ClientContactApprovedById   []Client // not sure about foreign key!!!!
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

type RolePermission struct {
	ID              uint
	RoleId          []Roles
	PermissionID    []Permissions
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type UserClientRole struct {
	ID              uint
	UserId          []Users
	RoleId          []Roles
	ClientId        []Client
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type ClientContact struct {
	ID              uint
	ClientId        []Client
	Name            string
	Email           string
	Phone           string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Trait struct {
	ID              uint
	TraitTypeId     []TraitType
	Trait           string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type ClientGeography struct {
	ID                  uint
	ClientId            []Client
	MarketId            []Market
	BoundaryCoordinates string
	CreatedDatetime     time.Time
	CreatedBy           string
	LastModifiedBy      string
	IsDeleted           bool
}
