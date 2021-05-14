package executor

import (
	"encoding/json"
	"log"

	"github.com/skrip42/grabbitLayer/internal/message"
	"github.com/skrip42/grabbitLayer/internal/queue"
	"github.com/skrip42/grabbitLayer/internal/queueDispetcher"
	"github.com/skrip42/grabbitLayer/internal/request"
)

type Executor struct {
    Queues map[string]*queue.Queue
    Messages map[string]*message.Message
    QueueDispetcher *queueDispetcher.QueueDispecther
}

func New() *Executor {
    ex := Executor{}
    ex.Queues = make(map[string]*queue.Queue)
    ex.Messages = make(map[string]*message.Message)
    ex.QueueDispetcher = queueDispetcher.New()
    return &ex;
}

func (ex *Executor) createQueueWorker(name string) error {
    var err error
    ex.Queues[name], err = ex.QueueDispetcher.GetQueue(name)
    if err != nil {
        return err
    }

    go func (n string) {
        for d := range ex.Queues[n].InputQueue.Consume {
            correlation := d.CorrelationId
            msg := ex.Messages[correlation]
            response, headers, err := msg.Request.Send()
            if  err != nil {
                continue
            }
            respReq := request.Request{
                Url: msg.Callback,
                Headers: headers,
                Method: "GET",
                Body: response,
            }
            msg.Response = respReq
            encResp, err := json.Marshal(respReq)
            if err != nil {
                continue
            }
            ex.Queues[n].OutputQueue.Send(string(encResp), correlation)
            msg.Status = 2
        }
        log.Println("close input " + name + " processor")
    } (name)

    go func (n string) {
        for d := range ex.Queues[n].OutputQueue.Consume {
            correlation := d.CorrelationId
            msg := ex.Messages[correlation]
            msg.Response.Send()
            msg.Status = 3
            delete(ex.Messages, n)
        }
        log.Println("close output " + name + " processor")
    } (name)
    return nil
}

func (ex *Executor) Execute(msg *message.Message) error {
    var err error
    _, ok := ex.Queues[msg.Queue]
    if !ok {
        err = ex.createQueueWorker(msg.Queue)
        if err != nil {
            return err
        }
    }
    ex.Messages[msg.Corelation] = msg
    marshaledRequest, err := json.Marshal(msg.Request)
    if err != nil {
        return err
    }
    ex.Queues[msg.Queue].InputQueue.Send(string(marshaledRequest), msg.Corelation)
    msg.Status = 1
    return nil
}
