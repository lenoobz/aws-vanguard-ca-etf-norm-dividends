package dividends

import (
	"context"

	"github.com/hthl85/aws-vanguard-etf-ca-norm-dividends/entities"
)

// Reader interface
type Reader interface{}

// Writer interface
type Writer interface {
	InsertAssetDividend(ctx context.Context, dividend *entities.AssetDividend) error
}

// Repo interface
type Repo interface {
	Reader
	Writer
}
