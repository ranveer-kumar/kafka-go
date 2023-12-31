package consumer

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/twmb/franz-go/pkg/kgo"

	"kafka-go/internal/config"
	kafka "kafka-go/pkg/kafka"
)

type Consumer struct {
	client     *kgo.Client
	messageCh  chan string
	numWorkers int
}

func NewConsumer(cfg config.KafkaConfig, numWorkers int) (*Consumer, error) {
	client, err := kafka.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Consumer{
		client:     client,
		messageCh:  make(chan string, numWorkers*2),
		numWorkers: numWorkers,
	}, nil
}

func ConsumeHandler(w http.ResponseWriter, r *http.Request, consumer *Consumer) {
	go consumer.StartConsuming()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Kafka consumer started.")
}

func (c *Consumer) StartConsuming() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(c.numWorkers)

	for i := 0; i < c.numWorkers; i++ {
		go func() {
			defer wg.Done()
			c.processMessages()
		}()
	}

	go func() {
		<-sigCh
		fmt.Println("Received shutdown signal. Stopping Kafka consumer...")
		cancel()
		wg.Wait()
		close(c.messageCh)
	}()

	c.consumeMessages(ctx)
}

func (c *Consumer) consumeMessages(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fetches := c.client.PollFetches(ctx)
			if errs := fetches.Errors(); len(errs) > 0 {
				panic(fmt.Sprint(errs))
			}

			iter := fetches.RecordIter()
			for !iter.Done() {
				record := iter.Next()
				c.messageCh <- string(record.Value)
			}
		}
	}
}

func (c *Consumer) processMessages() {
	for message := range c.messageCh {
		fmt.Println("Processed message:", message)
	}
}

func (c *Consumer) Close() {
	if c.client != nil {
		c.client.Close()
		close(c.messageCh)
	}
}
