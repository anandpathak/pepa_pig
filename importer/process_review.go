package importer

import (
	"context"
	"fmt"

	"pepa_pig/redis"
	"pepa_pig/utils"
)

func CheckInReview() {
	var ctx = context.Background()

	redis.Connect()

	keys := getKeysInReview(ctx)

	for _, key := range keys {
		fmt.Printf("%s %s", key, redis.Instance.Get(ctx, key).Val())
	}
}

func getKeysInReview(ctx context.Context) []string {
	keys := redis.Instance.Keys(ctx, fmt.Sprintf("%s*", utils.REDIS_REVIW_KEY_PREFIX))
	return keys.Val()
}
