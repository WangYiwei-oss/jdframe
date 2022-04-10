package configs

import (
	"log"
	"sync"
	"time"
)

var doOnceSnow sync.Once

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type SnowflakeGenerator struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

var SnowGenerator *SnowflakeGenerator

func init() {
	doOnceSnow.Do(func() {
		SnowGenerator = NewSnowflakeGenerator(1)
	})
}

func NewSnowflakeGenerator(workerId int64) *SnowflakeGenerator {
	if workerId < 0 || workerId > workerMax {
		log.Fatalln("workerId error, must >0")
		return nil
	}
	return &SnowflakeGenerator{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}
}

func (s *SnowflakeGenerator) GetId() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if s.timestamp == now {
		s.number++
		if s.number > numberMax {
			for now <= s.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.number = 0
		s.timestamp = now
	}
	return int64((now-startTime)<<timeShift | (s.workerId << workerShift) | (s.number))
}
