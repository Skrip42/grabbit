package queueDispetcher

import (

	"github.com/skrip42/grabbitLayer/internal/queue"
)

type QueueDispecther struct {
    Queues map[string]*queue.Queue
}

func New() *QueueDispecther {
    qd := QueueDispecther{}
    qd.Queues = make(map[string]*queue.Queue)
    return &qd
}

func (qd *QueueDispecther) GetQueue(name string) (*queue.Queue, error) {
    _, ok := qd.Queues[name]
    if !ok {
        q, err := queue.New(name)
        if err != nil {
            return nil, err
        }
        qd.Queues[name] = q
    }
    return qd.Queues[name], nil
}
