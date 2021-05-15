package entities

import (
	"context"

	logger "github.com/hthl85/aws-lambda-logger"
	"github.com/hthl85/aws-vanguard-etf-ca-norm-dividends/utils/datetime"
)

// FundDistribution struct
type FundDistribution struct {
	PortID                string                 `json:"portId,omitempty"`
	Ticker                string                 `json:"ticker,omitempty"`
	DistributionHistories []*DistributionDetails `json:"distributionHistories,omitempty"`
}

// DistributionDetails struct
type DistributionDetails struct {
	Type               string  `json:"type,omitempty"`
	DistributionAmount float64 `json:"distributionAmount,omitempty"`
	ExDividendDate     string  `json:"exDividendDate,omitempty"`
	RecordDate         string  `json:"recordDate,omitempty"`
	PayableDate        string  `json:"payableDate,omitempty"`
	DistDesc           string  `json:"distDesc,omitempty"`
	DistCode           string  `json:"distCode,omitempty"`
}

// MapFundDistributionToAssetDividend map fund distribution to asset dividend
func (f *FundDistribution) MapFundDistributionToAssetDividend(ctx context.Context, log logger.ContextLog) *AssetDividend {
	assetDividend := &AssetDividend{
		Ticker:    f.Ticker,
		Dividends: make(map[int64]*DividendDetails),
	}

	for _, distribution := range f.DistributionHistories {
		exDividendDate, err := datetime.GetStarDateFromString(distribution.ExDividendDate)
		if err != nil {
			log.Error(ctx, "parse exDividendDate failed", "error", err)
		}

		recordDate, err := datetime.GetStarDateFromString(distribution.RecordDate)
		if err != nil {
			log.Error(ctx, "parse recordDate failed", "error", err)
		}

		payableDate, err := datetime.GetStarDateFromString(distribution.PayableDate)
		if err != nil {
			log.Error(ctx, "parse payableDate failed", "error", err)
		}

		dividendDetails := &DividendDetails{
			DistDesc:       distribution.DistDesc,
			DistCode:       distribution.DistCode,
			Amount:         distribution.DistributionAmount,
			ExDividendDate: exDividendDate,
			RecordDate:     recordDate,
			PayableDate:    payableDate,
		}

		if payableDate != nil {
			dividendTime := payableDate.Unix()
			assetDividend.Dividends[dividendTime] = dividendDetails
		}
	}

	return assetDividend
}
