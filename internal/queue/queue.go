package queue

import (
	"github.com/skrip42/grabbitLayer/internal/rabbitHelper"
)

type Queue struct {
    Name        string
    InputQueue  *rabbitHelper.Exchange
    OutputQueue *rabbitHelper.Exchange
}

func New(name string) (*Queue, error) {
    var err error
    queue := Queue{}
    queue.Name = name
    queue.InputQueue, err = rabbitHelper.GetExchange("input_" + name)
    if err != nil {
        return nil, err
    }

    queue.OutputQueue, err = rabbitHelper.GetExchange("output_" + name)
    if err != nil {
        return nil, err
    }
    return &queue, nil
}
