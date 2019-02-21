package main

import (
	"fmt"
	"os"
	"os/signal"
	//"time"

	//"github.com/robfig/cron"
	cron "gopkg.in/robfig/cron.v2"

	"github.com/ess/cronenberg"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: cronenberg /path/to/jobs/file")
		os.Exit(1)
	}

	logger := cronenberg.NewLogger()
	jobs := cronenberg.LoadJobs(os.Args[1], logger)
	logger.Info("cronenberg", "Initializing")

	errs := make([]error, 0)
	c := cron.New()

	for _, j := range jobs {
		logger.Info("cronenberg", "Scheduling job "+j.Name)
		if _, err := c.AddJob("0 "+j.When, j); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		err := fmt.Sprintf("Could not schedule some jobs, aborting: %s", errs)
		logger.Error("cronenberg", err)
		os.Exit(1)
	}

	if len(c.Entries()) == 0 {
		logger.Error("cronenberg", "No jobs were loaded, aborting")
		os.Exit(1)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	c.Start()

	<-sigs
	c.Stop()
	logger.Info("cronenberg", "Terminating")
}
