package rabbit

import "github.com/streadway/amqp"

// Publish -
func (conn Conn) Publish(routingKey string, data []byte) error {
	return conn.Channel.Publish(
		// exchange - yours may be different
		"test.events",
		routingKey,
		// mandatory - we don't care if there I no queue
		false,
		// immediate - we don't care if there is no consumer on the queue
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}
