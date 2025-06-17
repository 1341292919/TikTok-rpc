package cache

import (
	"math/rand"
	"testing"
)

func buildTestModelVideoId(t *testing.T) int64 {
	t.Helper()
	return int64(rand.Uint32())
}
func buildTestModelUserId(t *testing.T) int64 {
	t.Helper()
	return int64(10000)
}
func buildTestModelCommentId(t *testing.T) int64 {
	t.Helper()
	return int64(rand.Uint32())
}
