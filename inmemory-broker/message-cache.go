package inmemorybroker

import "sync"


const (
	LenCleanUp = 400
)

type queueMap map[uint64]*WrappedMessage

// memory cache сообщений чата
type messageCache struct {
	mu        sync.Mutex
	queue     queueMap
	queueSize int
	Buffer    chan interface{}
}

func newMessageCache(cap int) *messageCache {
	return &messageCache{
		queue:     make(queueMap, cap),
		queueSize: cap, // для алокации памяти под новую map queue при отчистке старой
		Buffer:    make(chan interface{}, cap),
	}
}

func (mc *messageCache) getBuffer() <-chan interface{} {
	return mc.Buffer
}

// создаём новую map queue с сообщениями со статусом NotDef и Ready
func (mc *messageCache) cleanupDeque() {
	if len(mc.queue) < LenCleanUp {
		return
	}

	mc.mu.Lock()
	defer mc.mu.Unlock()

	newQueue := make(queueMap, mc.queueSize)
	for k, v := range mc.queue {
		if v.Status < Cancelled { // Сохраняем только неотправленные и неопределенные
			newQueue[k] = v
		}
	}
	mc.queue = newQueue
}

func (mc *messageCache) set(m *WrappedMessage) {
	mc.mu.Lock()
	// if len(mc.queue) > mc.queueSize {
	// 	log.Printf("WARNING: large deque size: %d", len(mc.queue))
	// }
	mc.queue[m.ID] = m
	mc.mu.Unlock()
}

func (mc *messageCache) processDequeBatch() {
	mc.mu.Lock()
	for _, v := range mc.queue {
		if v.Status == Ready {
			v.Status = Sended
			select {
			case mc.Buffer <- v.Message:
				// успех
			default:
				v.Status = Ready
			}
		}
	}
	mc.mu.Unlock()
}
