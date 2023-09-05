package Cache

import (
	"github.com/patrickmn/go-cache"
	"simpledouyin/Response"
	"sync"
	"time"
)

var (
	C          = cache.New(10*time.Hour, 20*time.Second) // 创建一个go-cache实例
	VideoList  []*Response.VideoList
	VideoMutex sync.Mutex
)
