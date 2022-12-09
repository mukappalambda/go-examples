package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)
	proteinLogger := log.WithFields(log.Fields{
		"event":    "Drink Protein Shake",
		"duration": "2 min",
		"place":    "Room",
	})

	warmupLogger := log.WithFields(log.Fields{
		"event":    "Warm-Up",
		"duration": "10 min",
		"place":    "Playground",
	})

	HIITLogger := log.WithFields(log.Fields{
		"event":    "HIIT Workout",
		"duration": "40 min",
		"place":    "Workout place",
	})
	MarathonLogger := log.WithFields(log.Fields{
		"event":    "Run a marathon",
		"duration": "3 hours",
		"place":    "Riverside",
	})

	proteinLogger.Info("morning")
	warmupLogger.Info("morning")
	HIITLogger.Warn("morning")
	MarathonLogger.Fatal("afternoon")

}
