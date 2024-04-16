package utils

import "time"

type TokenBucket struct {
	maxToken int
	token    int
	addRate  int
}

func NewTokenBucket(maxToken int, addRate int) *TokenBucket {
	return &TokenBucket{
		maxToken: maxToken,
		token:    maxToken,
		addRate:  addRate,
	}
}

func (tb *TokenBucket) AddToken() {
	if tb.token < tb.maxToken {
		tb.token += tb.addRate
	}
	tb.token += tb.addRate
	if tb.token > tb.maxToken {
		tb.token = tb.maxToken
	}
}

func (tb *TokenBucket) StartAddToken(gap time.Duration) {
	go func() {
		for {
			tb.AddToken()
			time.Sleep(gap)
		}
	}()
}

func (tb *TokenBucket) GetToken() bool {
	if tb.token > 0 {
		tb.token--
		return true
	}
	return false
}

func (tb *TokenBucket) GetTokenCount() int {
	return tb.token
}

func (tb *TokenBucket) GetTokenWithTimeout(timeout time.Duration) bool {
	select {
	case <-time.After(timeout):
		return false
	default:
		if tb.GetToken() {
			return true
		}
		return false
	}
}
