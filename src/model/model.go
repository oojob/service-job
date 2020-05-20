package model

import (
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// AddressModel address
type AddressModel struct {
	Country    string `bson:"country,omitempty"`
	Locality   string `bson:"locality,omitempty"`
	Region     string `bson:"region,omitempty"`
	PostalCode int64  `bson:"postal_code,omitempty"`
	Street     string `bson:"street,omitempty"`
}

// MetadataModel metadata
type MetadataModel struct {
	CreatedAt     timestamp.Timestamp `bson:"created_at,omitempty"`
	UpdatedAt     timestamp.Timestamp `bson:"updated_at,omitempty"`
	PublishedDate timestamp.Timestamp `bson:"published_date,omitempty"`
	EndDate       timestamp.Timestamp `bson:"end_date,omitempty"`
	LastActive    timestamp.Timestamp `bson:"last_active,omitempty"`
}

// Range range
type Range struct {
	Min int64 `bson:"min,omitempty"`
	Max int64 `bson:"max,omitempty"`
}

// Time time
type Time struct {
	ValidFrom    timestamp.Timestamp `bson:"valid_from:omitempty"`
	ValidThrough timestamp.Timestamp `bson:"valid_through:omitempty"`
}

// ApplicantModel applicants
type ApplicantModel struct {
	Applications []string `bson:"applications,omitempty"`
	Shortlisted  []string `bson:"shortlisted,omitempty"`
	Onhold       []string `bson:"onhold,omitempty"`
	Rejected     []string `bson:"rejected,omitempty"`
}

// IdentifierModel identifier
type IdentifierModel struct {
	Identifier                string `bson:"identifier,omitempty"`
	Name                      string `bson:"name,omitempty"`
	AlternateName             string `bson:"alternate_name,omitempty"`
	Type                      string `bson:"type,omitempty"`
	AdditionalType            string `bson:"additional_type,omitempty"`
	Description               string `bson:"description,omitempty"`
	DisambiguatingDescription string `bson:"disambiguating_description,omitempty"`
	Headline                  string `bson:"headline,omitempty"`
	Slogan                    string `bson:"slogan,omitempty"`
}

// Job model schema
type Job struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Address       AddressModel       `bson:"address,omitempty"`
	Metadata      MetadataModel      `bson:"metadata,omitempty"`
	Employment    string             `bson:"employment,omitempty"`
	Salary        Range              `bson:"salary,omitempty"`
	Experience    Range              `bson:"experience,omitempty"`
	Time          Time               `bson:"time,omitempty"`
	Skills        []string           `bson:"skills,omitempty"`
	WorkingHours  string             `bson:"working_hours,omitempty"`
	Status        string             `bson:"status,omitempty"`
	Qualification []string           `bson:"qualification,omitempty"`
	Type          string             `bson:"type,omitempty"`
	Applicant     ApplicantModel     `bson:"applicants,omitempty"`
	Identity      IdentifierModel    `bson:"identity,omitempty"`
}
