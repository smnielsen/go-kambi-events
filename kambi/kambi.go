package kambi

import (
	"fmt"
)

type Event struct {
	homeName string
	awayName string
}

type RetVal struct {
	sport  string
	events []Event
}

func fetchEvents(sport string) (RetVal, error) {
	fmt.Printf("Fetching events for %s", sport)

	events := []Event{}

	returns := RetVal{
		sport:  sport,
		events: events,
	}

	return returns
}
