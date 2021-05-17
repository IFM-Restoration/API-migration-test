package main

import (
	"time"
)

// green tables with second level dependencies
type WorkOrderDetail struct {
	ID                        uint
	WorkOrderId               []WorkOrder
	InteriorExteriorId        []InteriorExterior
	PriorityId                []Priority
	IsWarranty                bool
	WarrantyClientWorkOrderId string
	Description               string
	DoNotExceedAmount         uint
	IsOverDoNotExceed         bool
	IsHomeVacant              bool
	CreatedDatetime           time.Time
	CreatedBy                 string
	LastModifiedBy            string
	IsDeleted                 bool
}

type WorkOrderState struct {
	ID              uint
	WorkOrderId     []WorkOrder
	StateId         []State
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type WorkOrderLocation struct {
	ID                uint
	WorkOrderId       []WorkOrder
	MarketId          []Market
	Address1          string
	Address2          string
	City              string
	State             string
	Zip               string
	AccessInformation string
	CreatedDatetime   time.Time
	CreatedBy         string
	LastModifiedBy    string
	IsDeleted         bool
}

type Estimate struct {
	EstimateId                     uint // do we need to add foreign key?
	StatusId                       uint
	Version                        string
	WorkOrderNumber                string
	Total                          uint
	Created                        time.Time
	Modified                       time.Time
	Expired                        time.Time
	Notes                          string
	QbId                           string
	QbSynctoken                    string
	Author                         string
	QbParentRef                    string
	IsDeleted                      uint
	IsChangeOrder                  uint
	ServiceMix                     string
	Duration                       string
	AllowableSpendRange            string
	IsRejected                     uint
	DynamicExperienceWorkOrderUuid string
	DynamicExperienceCustomerWo    string
}

type Job struct {
	ID                  uint
	WorkOrderId         []WorkOrder
	CustomerId          []Customer
	MajorTradeId        []MajorTrade
	Uuid                string
	Name                string
	IsWarranty          bool
	OrderOfJobExecution uint
	Description         string
	CustomerAccepted    bool
	JobContactSignature string
	IsJobCanceled       bool
	ProjectManagerId    string
	PendingSignOffDate  time.Time
	CompletionDate      time.Time
	ExpectedJobDuration uint
	QcCompleteDate      time.Time
	IsSelfScheduleSent  bool
	CreatedDatetime     time.Time
	CreatedBy           string
	LastModifiedBy      string
	IsDeleted           bool
}

type Customer struct {
	ID               uint
	CustomerStatusId []CustomerStatus
	MarketId         []Market
	FirstName        string
	LastName         string
	Company          string
	Phone            string
	Email            string
	PrimaryTradeName string
	WorkRadius       uint
	CustomerScore    uint
	Street           string
	City             string
	State            string
	Zip              string
	CreatedDatetime  time.Time
	CreatedBy        string
	LastModifiedBy   string
	IsDeleted        bool
}

type CustomerTrait struct {
	ID              uint
	CustomerId      []Customer
	TraitId         []Trait
	CreatedDatetime time.Time
	CreatedBy       string
	LastModifiedBy  string
	IsDeleted       bool
}

type ServiceDetail struct {
	ID                uint
	ClientGeographyId []ClientGeography
	ServiceCategoryId []ServiceCategory
	UnitId            uint
	Detail            string
	Rate              uint
	CreatedDatetime   time.Time
	CreatedBy         string
	LastModifiedBy    string
	IsDeleted         bool
}
