package algorithms

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {

	// Token bucket that allows max 5 events every 5 seconds
	bucket := NewTokenBucket(5, 1000*time.Millisecond)

	assert.Equal(t, int64(5), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, false, bucket.TakeN(6), "Not enough tokens")
	assert.Equal(t, int64(5), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, true, bucket.Take(), "Not enough tokens")
	assert.Equal(t, int64(4), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, true, bucket.Take(), "Not enough tokens")
	assert.Equal(t, int64(3), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, true, bucket.Take(), "Not enough tokens")
	assert.Equal(t, int64(2), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, true, bucket.Take(), "Not enough tokens")
	assert.Equal(t, int64(1), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, true, bucket.Take(), "Not enough tokens")
	assert.Equal(t, int64(0), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, false, bucket.Take(), "Can take token")
	assert.Equal(t, int64(0), bucket.AvailableTokens(), "Invalid available tokens count")

	time.Sleep(800 * time.Millisecond)

	assert.Equal(t, false, bucket.Take(), "Can take token")
	assert.Equal(t, int64(0), bucket.AvailableTokens(), "Invalid available tokens count")

	time.Sleep(200 * time.Millisecond)

	assert.Equal(t, int64(1), bucket.AvailableTokens(), "Invalid available tokens count")
	assert.Equal(t, true, bucket.Take(), "Not enough tokens")
	assert.Equal(t, int64(0), bucket.AvailableTokens(), "Invalid available tokens count")
}
