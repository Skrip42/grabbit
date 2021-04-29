package main

import (
	"log"
	// "github.com/streadway/amqp"
)

type QueuePair struct {
    Input  chan string
    Output chan string
    Empty  chan bool
}

type QueueRequest struct {
    RequestSchema RequestSchema
    Queue         string
}

type QueueDispecther struct {
    Input chan QueueRequest
    QueuePairs map[string]QueuePair
}

func NewQueueDispecther() *QueueDispecther {
    qd := new(QueueDispecther)
    qd.Init()
    // qd.Input = make(chan QueueRequest)
    return qd
}


func (qd *QueueDispecther) Init() {
    qd.Input = make(chan QueueRequest)
    go func () {
        for request := range qd.Input {
            log.Print(request.RequestSchema)
            log.Print(request.Queue)
            // go sendRequestFromRequestSchema(request.RequestSchema)
        }
    } ()
}
//
// func connectToRabbit() *amqp.Connection {
//     connection, err := amqp.Dial("amqp://test:123@10.1.19.242:5672/")
//     if err != nil {
//         panic(err)
//     }
//     return connection
// }
//
// func queueDispecther() chan QueueRequest {
//     connection := connectToRabbit()
//     defer connection.Close()
//
//     queueStorage := make(map[string]QueuePair)
//
//     queueRequestInput := make(chan QueueRequest, 100)
//
//     go func () {
//         for request := range queueRequestInput {
//             log.Print(request.Queue)
//             // go sendRequestFromRequestSchema(request.RequestSchema)
//         }
//     } ()
//
//     return queueRequestInput
// }
