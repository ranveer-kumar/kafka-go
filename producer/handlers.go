package producer

import (
	"fmt"
	"net/http"
)

// SendMessageHandler is an HTTP handler function for sending messages to Kafka.
func SendMessageHandler(prod *Producer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		message := r.URL.Query().Get("message")
		topic := r.URL.Query().Get("topic")

		if message == "" || topic == "" {
			http.Error(w, "Missing 'message' or 'topic' query parameter", http.StatusBadRequest)
			return
		}

		err := prod.SendMessage(topic, message)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to send message: %v", err), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Message sent to topic '%s': %s\n", topic, message)
	}
}
