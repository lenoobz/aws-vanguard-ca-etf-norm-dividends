package distributions

import (
	"context"

	"github.com/lenoobz/aws-vanguard-etf-ca-norm-dividends/entities"
)

///////////////////////////////////////////////////////////
// Fund Repository Interface
///////////////////////////////////////////////////////////

// Reader interface
type Reader interface {
	FindFundDistributions(context.Context) ([]*entities.FundDistribution, error)
}

// Writer interface
type Writer interface{}

// Repo interface
type Repo interface {
	Reader
	Writer
}
