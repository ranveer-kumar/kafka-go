// kafka_client.go

package kafka

import (
	"context"
	"crypto/tls"
	"net"
	"strings"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/pkg/sasl/scram"

	"kafka-go/internal/config"
)

// CreateKafkaClient creates a Kafka client with the specified configuration.
func NewClient(cfg config.KafkaConfig) (*kgo.Client, error) {
	tlsConfig := &tls.Config{}

	dialerFunc := func(ctx context.Context, network, address string) (net.Conn, error) {
		return tls.Dial(network, address, tlsConfig)
	}

	opts := []kgo.Opt{
		kgo.SeedBrokers(strings.Split(cfg.BootstrapServers, ",")...),
		kgo.SASL(scram.Auth{
			User: cfg.User,
			Pass: cfg.Pass,
		}.AsSha512Mechanism()),
		kgo.Dialer(dialerFunc),
		kgo.ConsumeTopics(cfg.Topic),
	}

	client, err := kgo.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return client, nil
}
