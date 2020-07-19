package importer

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/go-redis/redis/v8"
)

// InsertRedis insert the csv data into redis
func InsertRedis(filepath string) {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer file.Close()
	scanner := csv.NewReader(file)
	for {
		record, err := scanner.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("error in line: %s", err.Error())
			continue
		}
		space := regexp.MustCompile(`\s+`)
		productName := space.ReplaceAllString(strings.TrimSpace(record[0]), "_")
		country := space.ReplaceAllString(strings.TrimSpace(record[1]), "_")
		err = rdb.Set(ctx, productName, country, 0).Err()
		if err != nil {
			fmt.Print(err)
		}
	}
}
