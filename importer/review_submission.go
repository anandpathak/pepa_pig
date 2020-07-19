package importer

import (
	"context"
	"fmt"
	"os"
	"strings"

	"pepa_pig/redis"
	"pepa_pig/utils"
)

func CheckInReview() {
	var ctx = context.Background()
	redis.Connect()

	fmt.Println("use y or n to select or reject the update")
	keys := getKeysInReview(ctx)

	fmt.Printf("Total Key In review: %d", len(keys))
	for _, key := range keys {
		k := strings.Replace(key, utils.REDIS_REVIW_KEY_PREFIX, "", 1)
		val := redis.Instance.Get(ctx, key).Val()
		fmt.Printf("%s | %s", k, val)
		if isReviewAccepted() {
			redis.Instance.SetNX(ctx, k, val, 0)
		}
		redis.Instance.Del(ctx, key)
	}
}

func isReviewAccepted() bool {

	var b []byte = make([]byte, 1)
	var YES string = "y"
	var NO string = "n"
	for {
		os.Stdin.Read(b)
		if string(b) == YES {
			return true
		} else if string(b) == NO {
			return false
		}

		fmt.Print("can only accept y or n")
		return isReviewAccepted()
	}
}

func getKeysInReview(ctx context.Context) []string {
	keys := redis.Instance.Keys(ctx, fmt.Sprintf("%s*", utils.REDIS_REVIW_KEY_PREFIX))
	return keys.Val()
}
