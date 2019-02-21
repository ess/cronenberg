package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/ess/cronenberg/pkg/cronenberg/cron"
	"github.com/ess/cronenberg/pkg/cronenberg/fs"
	"github.com/ess/cronenberg/pkg/cronenberg/logger"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: cronenberg /path/to/jobs/file")
		os.Exit(1)
	}

	log := logger.New()
	jobs := fs.NewJobService(os.Args[1], log)
	log.Info("cronenberg", "Initializing")

	errs := make([]error, 0)
	c := cron.New()

	for _, j := range jobs.All() {
		log.Info("cronenberg", "Scheduling job "+j.Name)
		if err := c.Manage(cron.NewRunner(j, log)); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		err := fmt.Sprintf("Could not schedule some jobs, aborting: %s", errs)
		log.Error("cronenberg", err)
		os.Exit(1)
	}

	if c.Count() == 0 {
		log.Error("cronenberg", "No jobs were loaded, aborting")
		os.Exit(1)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	c.Start()

	<-sigs
	c.Stop()
	log.Info("cronenberg", "Terminating")
}
