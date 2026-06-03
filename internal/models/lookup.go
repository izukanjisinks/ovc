package models

import "github.com/google/uuid"

type OVCCategory struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Requisite struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	DefaultPrice float64   `json:"default_price"`
}

// ChildRequisite is a requisite as assigned to a child (includes quantity/checked/price)
type ChildRequisite struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Quantity     int       `json:"quantity"`
	Checked      bool      `json:"checked"`
	PricePerItem float64   `json:"price_per_item"`
}

type Sponsor struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type NameRequest struct {
	Name string `json:"name"`
}

type RequisiteRequest struct {
	Name         string  `json:"name"`
	DefaultPrice float64 `json:"default_price"`
}

func (r NameRequest) Validate() string {
	if r.Name == "" {
		return "name is required"
	}
	return ""
}

func (r RequisiteRequest) Validate() string {
	if r.Name == "" {
		return "name is required"
	}
	return ""
}

// Relationship set requests
type SetCategoriesRequest struct {
	CategoryIDs []uuid.UUID `json:"category_ids"`
}

type SetRequisitesRequest struct {
	Requisites []RequisiteItem `json:"requisites"`
}

type RequisiteItem struct {
	RequisiteID  uuid.UUID `json:"requisite_id"`
	Quantity     int       `json:"quantity"`
	Checked      bool      `json:"checked"`
	PricePerItem float64   `json:"price_per_item"`
}

type SetSponsorsRequest struct {
	SponsorIDs []uuid.UUID `json:"sponsor_ids"`
}
