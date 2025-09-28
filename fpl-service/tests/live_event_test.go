package main

import (
	"log"
	"testing"
	"time"

	fplApi "github.com/imadbelkat1/fpl-service/internal/api"
	live_event_service "github.com/imadbelkat1/fpl-service/internal/services"
)

func TestLiveEventApiService(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping real API test")
	}

	// Setup service with real client
	service := &live_event_service.LiveEventApiService{
		Client: fplApi.NewFplApiClient(),
	}

	// Wait for consumer to start
	time.Sleep(200 * time.Millisecond)

	// Test with real API
	log.Println("Calling FPL API...")
	err := service.UpdateLiveEvent("1")
	if err != nil {
		t.Fatalf("UpdateLiveEvent(Test) with API failed: %v", err)
	}

	// Wait for messages to be processed
	time.Sleep(3 * time.Second)

	t.Log("Real API test completed successfully")
}
