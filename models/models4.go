package models

import "time"

// orange tables with third level dependencies
type WorkOrderMajorTrade struct {
	ID                uint
	MajorTradeId      []MajorTrade
	WorkOrderDetailId []WorkOrderDetail
	CreatedDatetime   time.Time
	CreatedBy         string
	LastModifiedBy    string
	IsDeleted         bool
}

type StateWorkOrderHistory struct {
	ID                uint
	WorkOrderStateId  []WorkOrderState
	StateExitReasonId []StateExitReason
	Entered           time.Time
	Exited            time.Time
	CreatedDatetime   time.Time
	CreatedBy         string
	LastModifiedBy    string
	IsDeleted         bool
}

type ExceptionTransaction struct {
	ID                    uint
	WorkOrderId           []WorkOrder
	JobId                 []Job // not null
	ExceptionReasonCodeId []ExceptionReasonCode
	Description           string
	CreatedDatetime       time.Time
	CreatedBy             string
	LastModifiedBy        string
	IsDeleted             bool
}

type EstimateNeedsRevisionReason struct {
	ID                    uint
	NeedsRevisionReasonId []NeedsRevisionReason
	EstimateId            []Estimate
	Notes                 string
	CreatedDatetime       time.Time
	CreatedBy             string
	LastModifiedBy        string
	IsDeleted             bool
}

type LineItem struct {
	ID              uint
	JobId           []Job
	UnitId          []Unit
	Order           uint
	ServiceDetailId []ServiceDetail
	Description     string
	Quantity        uint
	Rate            uint
	Amount          uint
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Pclsm struct {
	ID                      uint
	JobId                   []Job
	MaterialId              []Material
	Problem                 string
	Cause                   string
	Location                string
	IsApproved              bool
	NeedSpecialPartsOrdered bool
	CreatedDatetime         time.Time
	CreatedBy               string
	LastModifiedBy          string
	IsDeleted               bool
}

type JobNeed struct {
	ID                       uint
	JobId                    []Job
	JobNeedTypeId            []JobNeedType
	OrderNumber              string
	OrderDate                time.Time
	EstimatedTimeArrivalDate time.Time
	CompletionDate           time.Time
	Notes                    string
	CreatedDatetime          time.Time
	CreatedBy                string
	LastModifiedBy           string
	IsDeleted                bool
}

type JobState struct {
	ID              uint
	JobId           []Job
	StateId         []State
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type JobTrait struct {
	ID              uint
	TraitId         []Trait
	JobId           []Job
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type Invoice struct {
	ID              uint
	EstimateId      []Estimate
	PoNumber        string
	TestData        bool
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type ResidentContactAttempt struct {
	ID              uint
	JobId           []Job
	ResidentId      []Resident
	AttemptDate     time.Time
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type ResidentContactInformation struct {
	ID              uint
	ResidentId      []Resident
	Phone           uint
	AltPhone        uint
	Email           string
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}
