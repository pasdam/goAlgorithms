package algorithms

import (
	"sync"
	"time"
)

// TokenBucket implement the token bucket algorithm.
// The update interval and the size are not customizable after the object is created,
// but they are specified in the factory method NewTokenBucket.
type TokenBucket struct {
	// number of tokens available
	availableTokens int64

	// interval indicates after how much time generate new tokens
	interval time.Duration

	// last is the last time the limiter's tokens field was updated
	lastUpdateTimestamp time.Time

	// mutex controll tokens assigments
	mutex sync.Mutex

	// maximum number of tokens available
	size int64
}

// NewTokenBucket is the factory method that returns a new TokenBucket,
// configured with the specified parameters:
// - size: maximum number of tokens in the bucket;
// - interval: indicates after how much time new tokens are generated
func NewTokenBucket(size int64, interval time.Duration) *TokenBucket {
	return &TokenBucket{
		availableTokens:     size,
		interval:            interval,
		lastUpdateTimestamp: time.Now(),
		size:                size,
	}
}

// AvailableTokens return the number of tokens currently available
func (bucket *TokenBucket) AvailableTokens() int64 {
	bucket.mutex.Lock()
	defer bucket.mutex.Unlock() // save the function and execute it at exit (after return)

	// evaluate actual number of available tokens
	tokenToAdd := int64(time.Since(bucket.lastUpdateTimestamp) / bucket.interval)
	if tokenToAdd > 0 {
		bucket.availableTokens += tokenToAdd
		bucket.lastUpdateTimestamp = time.Now()
	}

	return bucket.availableTokens
}

// Interval returns the interval in ms in which limit the number of tokens
func (bucket *TokenBucket) Interval() time.Duration {
	return bucket.interval
}

// Size returns the maximum number of tokens available
func (bucket *TokenBucket) Size() int64 {
	return bucket.size
}

// TakeN take N tokens from the bucket, if they are available, or none otherwise.
// It returns true if the N tokens were available, false otherwise
func (bucket *TokenBucket) TakeN(n int64) bool {
	bucket.mutex.Lock()
	defer bucket.mutex.Unlock() // save the function and execute it at exit (after return)

	// evaluate actual number of available tokens
	tokenToAdd := int64(time.Since(bucket.lastUpdateTimestamp) / bucket.interval)
	if tokenToAdd > 0 {
		bucket.availableTokens += tokenToAdd
		bucket.lastUpdateTimestamp = time.Now()
	}

	if bucket.availableTokens >= n {
		bucket.availableTokens -= n

		return true
	}

	return false
}

// Take 1 token from the bucket, same as TakeN(1).
// It returns true if the token was available, false otherwise
func (bucket *TokenBucket) Take() bool {
	return bucket.TakeN(1)
}
