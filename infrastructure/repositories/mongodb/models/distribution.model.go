package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// FundDistributionModel struct
type FundDistributionModel struct {
	ID            *primitive.ObjectID         `bson:"_id,omitempty"`
	IsActive      bool                        `bson:"isActive,omitempty"`
	CreatedAt     int64                       `bson:"createdAt,omitempty"`
	ModifiedAt    int64                       `bson:"modifiedAt,omitempty"`
	Schema        string                      `bson:"schema,omitempty"`
	PortID        string                      `bson:"portId,omitempty"`
	Ticker        string                      `bson:"ticker,omitempty"`
	Distributions []*DistributionDetailsModel `bson:"distributionHistories,omitempty"`
}

// DistributionDetailsModel struct
type DistributionDetailsModel struct {
	Type               string  `bson:"type,omitempty"`
	DistributionAmount float64 `bson:"distributionAmount,omitempty"`
	ExDividendDate     string  `bson:"exDividendDate,omitempty"`
	RecordDate         string  `bson:"recordDate,omitempty"`
	PayableDate        string  `bson:"payableDate,omitempty"`
	DistDesc           string  `bson:"distDesc,omitempty"`
	DistCode           string  `bson:"distCode,omitempty"`
}
