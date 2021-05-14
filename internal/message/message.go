package message

import (
	"math/rand"

	"github.com/skrip42/grabbitLayer/internal/request"
)

type Message struct {
    Status     int
    Request    request.Request
    Response   request.Request
    Callback   string
    Corelation string
    Queue      string
}

func randInt(min int, max int) int {
        return min + rand.Intn(max-min)
}

func randomString(l int) string {
        bytes := make([]byte, l)
        for i := 0; i < l; i++ {
                bytes[i] = byte(randInt(65, 90))
        }
        return string(bytes)
}

func New(
    request request.Request,
    callback string,
    queue string,
) *Message {
    m := Message{}
    m.Status = 0
    m.Request = request
    m.Callback = callback
    m.Queue = queue
    m.Corelation = randomString(32)
    return &m
}
