package crack

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

var ctx = context.Background()

func rediscon(ctx context.Context, cancel context.CancelFunc, ip, user, passwd string, port, timeout int, ch <-chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		<-ch
	}()
	select {
	case <-ctx.Done():
		return
	default:
	}

	client := redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%d", ip, port),
		Username:        user,
		Password:        passwd,
		DB:              0,
		DialTimeout:     time.Duration(timeout) * time.Second,
		MinRetryBackoff: time.Duration(timeout) * time.Second,
		ReadTimeout:     time.Duration(timeout) * time.Second,
	})
	_, err := client.Ping(ctx).Result()
	if err == nil {
		end(ip, user, passwd, port, "Redis")
		cancel()
	}
}