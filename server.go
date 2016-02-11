package main

import (
	"os"

	"github.com/jessemillar/rytsar-server/accessors"
	"github.com/jessemillar/rytsar-server/controllers"
	"github.com/zenazn/goji"
)

func main() {
	// Construct the dsn used for the database
	dsn := os.Getenv("RYTSAR_DB_USER") + ":" + os.Getenv("RYTSAR_DB_PASS") + "@tcp(" + os.Getenv("RYTSAR_DB_HOST") + ":" + os.Getenv("RYTSAR_DB_PORT") + ")/" + os.Getenv("RYTSAR_DB_NAME")

	// Construct a new AccessorGroup and connects it to the database
	ag := new(accessors.AccessorGroup)
	ag.ConnectToDB("mysql", dsn)

	// Constructs a new ControllerGroup and gives it the AccessorGroup
	cg := new(controllers.ControllerGroup)
	cg.Accessors = ag

	goji.Get("/health", cg.Health) // Service health
	goji.Get("/database/:latitude/:longitude/:radius", cg.DumpDatabase)
	goji.Get("/count/:latitude/:longitude/:radius", cg.CountNearbyEnemies)
	goji.Delete("/delete/:id", cg.DeleteEnemy)
	goji.Serve()
}
