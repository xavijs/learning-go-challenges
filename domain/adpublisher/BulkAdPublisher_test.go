package adpublisher

import (
	"github.com/stretchr/testify/assert"
	"learning-go-challenges/fixtures"
	"testing"
	"time"
)

func TestPublishAnAdInBulk(t *testing.T) {
	adPublishers := []AdPublisher{NewMilanunciosAdPublisher(), NewLeboncoinAdPublisher(), NewWallapopAdPublisher(), NewMilanunciosAdPublisher(), NewLeboncoinAdPublisher(), NewWallapopAdPublisher()}
	const maxElapsedTimeSeconds = 2
	bulkAdPublisher := BulkAdPublisher{adPublishers: adPublishers}

	startTime := time.Now()
	bulkAdPublisher.Execute(fixtures.FirstAd)
	elapsedTime := time.Since(startTime)

	assert.Less(t, elapsedTime, maxElapsedTimeSeconds*time.Second, "Execution took more than 3 seconds")
}
