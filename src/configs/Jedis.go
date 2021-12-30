package configs

import (
	"context"
	"fmt"
	"github.com/WangYiwei-oss/jdframe/src/configinjector"
	"github.com/WangYiwei-oss/jdframe/src/configparser"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

func init() {
	configinjector.RegisterBean("REDIS", NewJedis)
}

var doOnceRedis sync.Once

//用于迭代MGET对象
type Iterator struct {
	data  []interface{}
	index int
}

func (i *Iterator) HasNext() bool {
	if i.data == nil || len(i.data) == 0 {
		return false
	}
	return i.index < len(i.data)
}

func (i *Iterator) Next() interface{} {
	ret := i.data[i.index]
	i.index++
	return ret
}

type Jedis struct {
	*redis.Client
	ctx context.Context
}

func NewIterator(data []interface{}) *Iterator {
	return &Iterator{data: data}
}

var redisClient *redis.Client
var bctx context.Context

func NewJedis() *Jedis {
	doOnceRedis.Do(func() {
		config := configparser.GlobalSettings["REDIS_CONFIG"].(map[string]interface{})
		redisClient = redis.NewClient(&redis.Options{
			Network:  "tcp",
			Addr:     fmt.Sprintf("%s:%s", config["IP"].(string), config["PORT"].(string)),
			Password: config["PASSWORD"].(string),
			DB:       0, //数据库
			//连接池容量及闲置连接数量，连接池用来保护后端的数据库
			PoolSize:     15,
			MinIdleConns: 10, //好比最小连接数
			//超时
			DialTimeout:  5 * time.Second, //连接建立超时事件
			ReadTimeout:  3 * time.Second, //读超时
			WriteTimeout: 3 * time.Second,
			PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长
			//闲置连接检查包括IdleTimeout MaxConnAge
			IdleCheckFrequency: 60 * time.Second, //闲置连接检查周期，默认为一分钟。-1表示不做检查，只在客户端获取连接时对闲置连接进行处理
			IdleTimeout:        5 * time.Second,  //闲置超时
			MaxConnAge:         0 * time.Second,  //连接存活时长，即从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接

			//命令执行失败时的重试策略
			MaxRetries:      0,                      //命令执行失败时，最多重试多少次
			MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，-1表示取消间隔
			MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，-1表示取消间隔
		})
		pong, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalln("连接redis失败", err)
		}
		log.Println("redis连接成功", pong)
		bctx = context.Background()
	})
	return &Jedis{
		Client: redisClient,
		ctx:    bctx,
	}
}

//TODO 错误处理需要继续封装
func (j *Jedis) Get(key string) (string, error) {
	return j.Client.Get(j.ctx, key).Result()
}

func (j *Jedis) MGet(keys ...string) ([]interface{}, error) {
	return j.Client.MGet(j.ctx, keys...).Result()
}

//func (j *Jedis)Set(key string){
//	//j.Client.Set()
//}
