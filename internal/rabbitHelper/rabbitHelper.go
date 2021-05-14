package rabbitHelper

import "github.com/streadway/amqp"

var isConnected = false
var connection *amqp.Connection
var isChannel = false
var channel *amqp.Channel

//singleton connection
func getConnection() (*amqp.Connection, error) {
    if !isConnected {
        c, err := amqp.Dial("amqp://grabbit_local:123@10.1.19.242:5672/")
        if err != nil {
            // panic(err)
            return nil, err
        }
        connection = c
    }
    return connection, nil
}

//singleton channel
func getChannel() (*amqp.Channel, error) {
    if !isChannel {
        cn, err := getConnection()
        if err != nil {
            // panic(err)
            return nil, err
        }
        ch, err := cn.Channel()
        if err != nil {
            // panic(err)
            return nil, err
        }
        channel = ch
    }
    return channel, nil
}

func GetExchange(name string) (*Exchange, error) {
    var err error
    ex := Exchange{}
    ex.Name = name
    ex.Channel, err = getChannel()
    if err != nil {
            // panic(err)
        return nil, err
    }
    // ex.Channel = channel
    err = ex.Channel.ExchangeDeclare(
        name,
        "fanout",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
            // panic(err)
        return nil, err
    }
    q, err := ex.Channel.QueueDeclare(
        name + "__grabbit",
        false,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
            // panic(err)
        return nil, err
    }
    err = ex.Channel.QueueBind(
        q.Name,
        "",
        name,
        false,
        nil,
    )
    if err != nil {
            // panic(err)
        return nil, err
    }
    msgs, err := ex.Channel.Consume(
        q.Name,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
            // panic(err)
        return nil, err
    }
    ex.Consume = msgs
    return &ex, nil
}
