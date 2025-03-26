package xcache

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/somphongph/lib-golang-packages/xlogger"
)

type Redis struct {
	Host         string
	Port         int
	Password     string
	InstanceName string
}

type Servicer interface {
	Set(key string, value any, expiration time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Client() *redis.Client
	Health() HealthStats
}

type service struct {
	client *redis.Client
}

type HealthStats struct {
	Code      string
	Message   string
	Error     error
	PoolStats *redis.PoolStats
}

var (
	instanceName string
)

func NewRedis(sect *Redis) (*service, error) {
	instanceName = sect.InstanceName

	addr := fmt.Sprintf("%s:%d", sect.Host, sect.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: sect.Password,
	})

	if _, err := client.Ping().Result(); err != nil {
		xlogger.Errorf("Failed to ping Redis: %v", err)

		return nil, err
	}

	xlogger.Infof("Redis initialized")

	return &service{client}, nil
}

func (s *service) Set(key string, value any, expiration time.Duration) error {
	k := fmt.Sprintf("%s::%s", instanceName, key)
	return s.client.Set(k, value, expiration).Err()
}

func (r *service) Get(key string) (string, error) {
	k := fmt.Sprintf("%s::%s", instanceName, key)
	return r.client.Get(k).Result()
}

func (s *service) Delete(key string) error {
	k := fmt.Sprintf("%s::%s", instanceName, key)
	return s.client.Del(k).Err()
}

func (s *service) Client() *redis.Client {
	return s.client
}

func (s *service) Health() HealthStats {

	stats := HealthStats{}

	err := s.client.Ping().Err()
	if err != nil {
		stats.Code = "down"
		stats.Error = fmt.Errorf("redis down: %v", err)
		xlogger.Errorf("Redis health check failed: %v", err)

		return stats
	}

	stats.Code = "up"
	stats.Message = "Redis is healthy"
	stats.PoolStats = s.client.PoolStats()

	// Evaluate stats to provide a health message
	if stats.PoolStats.TotalConns > 100 {
		stats.Message = "Redis is experiencing high connection usage."
	}

	return stats
}
