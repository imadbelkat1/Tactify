package services

import (
	"context"
	"fmt"
	"sync"

	fpl_api "github.com/imadbelkat1/fpl-service/internal/api"
	"github.com/imadbelkat1/fpl-service/internal/models"
)

type FixturesApiService struct {
	Client *fpl_api.FplApiClient
}

func (s *FixturesApiService) UpdateFixtures() error {
	var fixtures models.Fixtures

	ctx := context.Background()

	if err := s.Client.GetAndUnmarshal(ctx, fixturesEndpoint, &fixtures); err != nil {
		return err
	}

	if err := publishFixtures(ctx, fixtures); err != nil {
		return fmt.Errorf("update fixtures: %w", err)
	}

	return nil
}

func publishFixtures(ctx context.Context, Fixtures models.Fixtures) error {

	toDelete := []string{"stats"}

	const deleteWorkers = 10
	const publishWorkers = 10

	jobs := make(chan models.Fixture, len(Fixtures))
	fixturesChan := make(chan ProcessedModel, len(Fixtures))
	results := make(chan error, len(Fixtures)*2) // delete + publish

	// delete stage
	var deleteWg sync.WaitGroup
	for i := 0; i < deleteWorkers; i++ {
		deleteWg.Add(1)
		go func() {
			defer deleteWg.Done()
			for element := range jobs {
				processed, err := processDelete(element, toDelete)
				if err != nil {
					results <- err
					continue
				}
				fixturesChan <- ProcessedModel{ID: element.ID, Data: processed}
			}
		}()
	}

	// close fixturesChan after delete stage finishes
	go func() {
		deleteWg.Wait()
		close(fixturesChan)
	}()

	// publish stage
	var publishWg sync.WaitGroup
	for i := 0; i < publishWorkers; i++ {
		publishWg.Add(1)
		go func() {
			defer publishWg.Done()
			for element := range fixturesChan {
				key := []byte(fmt.Sprintf("%d", element.ID))
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
	for _, element := range Fixtures {
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
