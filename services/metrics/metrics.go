package metrics

import (
	amqp "github.com/kaellybot/kaelly-amqp"
	"github.com/kaellybot/kaelly-metrics/repositories/metrics"
	"github.com/rs/zerolog/log"
)

func New(broker amqp.MessageBroker, repository metrics.Repository) *Impl {
	return &Impl{
		broker:     broker,
		repository: repository,
	}
}

func GetBinding() amqp.Binding {
	return amqp.Binding{
		Exchange:   amqp.ExchangeRequest,
		RoutingKey: routingkey,
		Queue:      queueName,
	}
}

func (service *Impl) Consume() {
	log.Info().Msgf("Consuming command requests...")
	service.broker.Consume(queueName, service.consume)
}

func (service *Impl) consume(ctx amqp.Context, message *amqp.RabbitMQMessage) {
	service.repository.Write(message, ctx.CorrelationID, ctx.ReplyTo, ctx.Timestamp)
}
