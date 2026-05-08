package main

import (
	"log"

	"github.com/nullablenone/life-decision-tracker/config"
	"github.com/nullablenone/life-decision-tracker/domain/activity"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
	"github.com/nullablenone/life-decision-tracker/domain/home"
	"github.com/nullablenone/life-decision-tracker/routes"
)

func main() {

	env, err := config.NewEnv()
	if err != nil {
		log.Fatal("failed to load .env file: %w", err)
	}

	db, err := config.ConnectPostgre(env)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(decision.DecisionCategory{}, &activity.Board{}, activity.Activity{}); err != nil {
		log.Fatal("failed to migrate tables")
	}

	if err = decision.SeedDecisions(db); err != nil {
		log.Fatal("failed to seed decision categories")
	}

	decisionRepo := decision.NewRepository(db)
	decisionService := decision.NewService(decisionRepo)
	decisionHandler := decision.NewHandler(decisionService)

	activityRepo := activity.NewRepository(db)
	activityService := activity.NewService(activityRepo)
	activityHandler := activity.NewHandler(activityService, decisionService)

	homeHandler := home.NewHome(activityService, decisionService)

	router := routes.SetRoutes(decisionHandler, activityHandler, homeHandler)

	if err = router.Run(":777"); err != nil {
		log.Fatalf("failed to start HTTP server on :777: %v", err)
	}

}
