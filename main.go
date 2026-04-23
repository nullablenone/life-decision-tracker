package main

import (
	"log"

	"github.com/nullablenone/life-decision-tracker/config"
	"github.com/nullablenone/life-decision-tracker/domain/decision"
	"github.com/nullablenone/life-decision-tracker/routes"
)

func main() {
	db, err := config.ConnectPostgre()
	if err != nil {
		log.Fatal(err)
	}

	if err = db.AutoMigrate(decision.DecisionCategory{}); err != nil {
		log.Fatal("failed to create table decision catagories")
	}

	if err = decision.SeedDecisions(db); err != nil {
		log.Fatal("failed to seed decision categories")
	}

	decisionRepo := decision.NewRepository(db)
	decisionService := decision.NewService(decisionRepo)
	decisionHandler := decision.NewHandler(decisionService)

	router := routes.SetRoutes(decisionHandler)
	if err = router.Run(":3333"); err != nil {
		log.Fatalf("failed to start HTTP server on :3333: %v", err)
	}

}
