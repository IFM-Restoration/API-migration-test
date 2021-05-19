package models

import "time"

// purple tables without foreign key

type Client struct {
	ID              uint
	Name            string
	ProcessRule     string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Market struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type MajorTrade struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type InteriorExterior struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Priority struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type WorkOrderAcceptanceStatus struct {
	ID               uint
	AcceptanceStatus string
	CreatedDatetime  time.Time
	CreatedBy        string
	LastModifiedBy   string
	IsDeleted        bool
}

type ExceptionReasonCode struct {
	ID              uint
	Reason          string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type JobNeedType struct {
	ID              uint
	Type            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type JobNoteType struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type JobRescheduleReason struct {
	ID              uint
	Reason          string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type JobScheduleActivityStatus struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type TraitType struct {
	ID              uint
	TraitType       string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type State struct {
	ID              uint
	Name            string
	ForWorkOrder    bool
	ForJob          bool
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type StateExitReason struct {
	ID              uint
	Reason          string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type DelayReason struct {
	ID              uint
	Reason          string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type AttachmentType struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type AttachmentCategory struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type AttachmentExtension struct {
	ID              uint
	Name            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type NeedsRevisionReason struct {
	ID              uint
	Reason          string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type ServiceCategory struct {
	ID              uint
	Category        string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type CustomerStatus struct {
	ID              uint
	Status          string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type LaborType struct {
	ID              uint
	Type            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Material struct {
	ID              uint
	Material        string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Unit struct {
	ID              uint
	Unit            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type QbAccess struct {
	ID           uint
	AccessToken  string
	RefreshToken string
	RealmId      string
	Expired      time.Time
}

type Permissions struct {
	ID              uint
	Permission      string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Roles struct {
	ID              uint
	Role            string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Users struct {
	ID              uint
	UserName        string
	Uuid            string
	Email           string
	IsIfmInternal   bool
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}
