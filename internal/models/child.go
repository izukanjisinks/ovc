package models

import (
	"time"

	"github.com/google/uuid"
)

type Child struct {
	ID                 uuid.UUID  `json:"id"`
	PupilID            string     `json:"pupil_id"`
	FirstName          string     `json:"first_name"`
	LastName           string     `json:"last_name"`
	Address            string     `json:"address"`
	ClassName          string     `json:"class_name"`
	ImageURL           *string    `json:"image_url"`
	GuardianFirstName  string     `json:"guardian_first_name"`
	GuardianLastName   string     `json:"guardian_last_name"`
	GuardianAddress    string     `json:"guardian_address"`
	GuardianPhone      string     `json:"guardian_phone"`
	CreatedBy          *uuid.UUID `json:"created_by"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`

	// Populated on detail fetch
	Categories []OVCCategory     `json:"categories,omitempty"`
	Requisites []ChildRequisite  `json:"requisites,omitempty"`
	Sponsors   []Sponsor         `json:"sponsors,omitempty"`
}

type CreateChildRequest struct {
	PupilID           string  `json:"pupil_id"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Address           string  `json:"address"`
	ClassName         string  `json:"class_name"`
	ImageURL          *string `json:"image_url"`
	GuardianFirstName string  `json:"guardian_first_name"`
	GuardianLastName  string  `json:"guardian_last_name"`
	GuardianAddress   string  `json:"guardian_address"`
	GuardianPhone     string  `json:"guardian_phone"`

	CategoryIDs []uuid.UUID    `json:"category_ids"`
	SponsorIDs  []uuid.UUID    `json:"sponsor_ids"`
	Requisites  []RequisiteItem `json:"requisites"`
}

type UpdateChildRequest struct {
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Address           string  `json:"address"`
	ClassName         string  `json:"class_name"`
	ImageURL          *string `json:"image_url"`
	GuardianFirstName string  `json:"guardian_first_name"`
	GuardianLastName  string  `json:"guardian_last_name"`
	GuardianAddress   string  `json:"guardian_address"`
	GuardianPhone     string  `json:"guardian_phone"`

	CategoryIDs []uuid.UUID     `json:"category_ids"`
	SponsorIDs  []uuid.UUID     `json:"sponsor_ids"`
	Requisites  []RequisiteItem `json:"requisites"`
}

func (r CreateChildRequest) Validate() string {
	switch {
	case r.PupilID == "":
		return "pupil_id is required"
	case r.FirstName == "":
		return "first_name is required"
	case r.LastName == "":
		return "last_name is required"
	case r.Address == "":
		return "address is required"
	case r.ClassName == "":
		return "class_name is required"
	case r.GuardianFirstName == "":
		return "guardian_first_name is required"
	case r.GuardianLastName == "":
		return "guardian_last_name is required"
	case r.GuardianAddress == "":
		return "guardian_address is required"
	case r.GuardianPhone == "":
		return "guardian_phone is required"
	}
	return ""
}

func (r UpdateChildRequest) Validate() string {
	switch {
	case r.FirstName == "":
		return "first_name is required"
	case r.LastName == "":
		return "last_name is required"
	case r.Address == "":
		return "address is required"
	case r.ClassName == "":
		return "class_name is required"
	case r.GuardianFirstName == "":
		return "guardian_first_name is required"
	case r.GuardianLastName == "":
		return "guardian_last_name is required"
	case r.GuardianAddress == "":
		return "guardian_address is required"
	case r.GuardianPhone == "":
		return "guardian_phone is required"
	}
	return ""
}
