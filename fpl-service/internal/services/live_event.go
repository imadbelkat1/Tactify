package services

import (
	"context"
	"fmt"
	"strconv"
	"sync"

	fpl_api "github.com/imadbelkat1/fpl-service/internal/api"
	"github.com/imadbelkat1/fpl-service/internal/models"
)

type LiveEventApiService struct {
	Client *fpl_api.FplApiClient
}

func (s *LiveEventApiService) UpdateLiveEvent(eventID string) error {
	var liveEvent models.LiveEvent
	ctx := context.Background()

	endpoint := fmt.Sprintf(liveEventEndpoint, eventID)

	if err := s.Client.GetAndUnmarshal(ctx, endpoint, &liveEvent); err != nil {
		return fmt.Errorf("failed to fetch live event data: %v", err)
	}

	if err := publishLiveEvent(ctx, liveEvent, eventID); err != nil {
		return fmt.Errorf("failed to publish live event data: %v", err)
	}

	return nil
}

func publishLiveEvent(ctx context.Context, liveEvent models.LiveEvent, eventID string) error {
	gameweek, err := strconv.Atoi(eventID)
	if err != nil {
		return fmt.Errorf("invalid event ID: %v", err)
	}

	toDelete := []string{"explain", "modified"}

	const deleteWorkers = 10
	const publishWorkers = 10

	jobs := make(chan models.LiveElement, len(liveEvent.Elements))
	elements := make(chan ProcessedModel, len(liveEvent.Elements))
	results := make(chan error, len(liveEvent.Elements)*2) // delete + publish

	// delete stage
	var deleteWg sync.WaitGroup
	for i := 0; i < deleteWorkers; i++ {
		deleteWg.Add(1)
		go func() {
			defer deleteWg.Done()
			for element := range jobs {
				element.Gameweek = gameweek
				processed, err := processDelete(element, toDelete)
				if err != nil {
					results <- err
					continue
				}
				elements <- ProcessedModel{ID: element.ID, Data: processed}
			}
		}()
	}

	// close elements after delete stage finishes
	go func() {
		deleteWg.Wait()
		close(elements)
	}()

	// publish stage
	var publishWg sync.WaitGroup
	for i := 0; i < publishWorkers; i++ {
		publishWg.Add(1)
		go func() {
			defer publishWg.Done()
			for element := range elements {
				key := []byte(fmt.Sprintf("%d-%d", gameweek, element.ID))
				err := Publish(ctx, liveEventProducer, liveEventTopic, key, element.Data)
				results <- err
			}
		}()
	}

	// close results after publish stage finishes
	go func() {
		publishWg.Wait()
		close(results)
	}()

	// feed jobs
	for _, element := range liveEvent.Elements {
		jobs <- element
	}
	close(jobs)

	// check results
	for err := range results {
		if err != nil {
			return err
		}
	}

	return nil
}
