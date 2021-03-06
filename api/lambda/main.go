package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	logger "github.com/lenoobz/aws-lambda-logger"
	"github.com/lenoobz/aws-vanguard-etf-ca-norm-dividends/config"
	"github.com/lenoobz/aws-vanguard-etf-ca-norm-dividends/infrastructure/repositories/mongodb/repos"
	"github.com/lenoobz/aws-vanguard-etf-ca-norm-dividends/usecase/distributions"
	"github.com/lenoobz/aws-vanguard-etf-ca-norm-dividends/usecase/dividends"
)

func main() {
	appConf := config.AppConf

	// create new logger
	zap, err := logger.NewZapLogger()
	if err != nil {
		log.Fatal("create app logger failed")
	}
	defer zap.Close()

	// create new repository
	distributionRepo, err := repos.NewDistributionMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create distribution mongo repo failed")
	}
	defer distributionRepo.Close()

	// create new repository
	dividendRepo, err := repos.NewDividendMongo(nil, zap, &appConf.Mongo)
	if err != nil {
		log.Fatal("create dividend mongo repo failed")
	}
	defer dividendRepo.Close()

	// create new service
	distributionService := distributions.NewService(distributionRepo, zap)
	dividendService := dividends.NewService(dividendRepo, *distributionService, zap)

	lambda.Start(dividendService.PopulateFundDividends)
}
