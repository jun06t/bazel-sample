package dice

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var cli *redis.Client

const (
	historyKey = "dice_history"
)

func PlayDice() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	n := r.Intn(6)
	return n + 1
}

func SaveNumber(ctx context.Context, num int) (string, error) {
	timeKey := strconv.FormatInt(time.Now().Unix(), 10)
	_, err := cli.HSet(ctx, historyKey, timeKey, num).Result()
	if err != nil {
		return "", err
	}
	return timeKey, nil
}

func GetHistory(ctx context.Context) (map[string]string, error) {
	m, err := cli.HGetAll(ctx, historyKey).Result()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func newClient(ctx context.Context, addr string) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
