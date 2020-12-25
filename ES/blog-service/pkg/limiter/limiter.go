package limiter

import (
	"github.com/gin-gonic/gin"
	"time"
)
import "github.com/juju/ratelimit"

type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LImiterBucketRule) LimiterIface
}
type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}
type LImiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64
}
