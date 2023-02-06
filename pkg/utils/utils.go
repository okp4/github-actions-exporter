package utils

import (
	"log"
	"time"

	"github.com/spendesk/github-actions-exporter/pkg/config"
)

var (
	nbOfCallBetweenReset = 20
)

func WaitEvery15Calls(counter *int) {
	if config.Github.SecondsBetween15calls <= 0 {
		return
	}
	if *counter <= 0 || *counter > nbOfCallBetweenReset {
		*counter = 1
	}

	if *counter%nbOfCallBetweenReset == 0 {
		*counter = 1
		log.Printf("Wait %ds to try not to reach github API rate limit", config.Github.SecondsBetween15calls)
		time.Sleep(time.Duration(config.Github.SecondsBetween15calls) * time.Second)
	} else {
		*counter++
	}
}
