package importer

import (
	"pepa_pig/redis"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestCheckInReview(t *testing.T) {

}

func TestGetKeysInReview(t *testing.T) {
	t.Run("should return list of keys in review", func(t *testing.T) {
		redis.Connect()
		ctx := context.Background()
		actual := getKeysInReview(ctx)
		assert.NotEqual(t, len(actual), 0)
	})
}
