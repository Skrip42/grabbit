package rabbitHelper

import "github.com/streadway/amqp"

type Exchange struct {
    Name    string
    Channel *amqp.Channel
    Consume <-chan amqp.Delivery
}

func (ex *Exchange) Send(message string, corelation string) error {
    err := ex.Channel.Publish(
        ex.Name,
        "",
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            CorrelationId: corelation,
            Body: []byte(message),
        },
    )
    if err != nil {
        return err
    }
    return nil
}
