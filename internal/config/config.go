// package config

// import "os"

// // KafkaConfig holds configuration for Kafka
// type KafkaConfig struct {
//     BootstrapServers string
//     SecurityProtocol string
//     SaslMechanism    string
//     SaslJaasConfig   string
//     Topic            string
//     GroupID          string
// }

// // LoadKafkaConfig loads the Kafka configuration from environment variables
// func LoadKafkaConfig() KafkaConfig {
//     return KafkaConfig{
//         BootstrapServers: getEnv("KAFKA_BOOTSTRAP_SERVERS", "b-2-public.cmndev351.ezayi5.c3.kafka.ap-south-1.amazonaws.com:9196,b-1-public.cmndev351.ezayi5.c3.kafka.ap-south-1.amazonaws.com:9196"),
//         SecurityProtocol: getEnv("KAFKA_SECURITY_PROTOCOL", "SASL_SSL"),
//         SaslMechanism:    getEnv("KAFKA_SASL_MECHANISM", "SCRAM-SHA-512"),
//         // SaslJaasConfig:   getEnv("KAFKA_SASL_JAAS_CONFIG", "org.apache.kafka.common.security.scram.ScramLoginModule required username='devkafka' password='DevAtACLPublicKafka';"), // Set this in your environment
//         Topic:            getEnv("KAFKA_TOPIC", "test2"),
//         GroupID:          getEnv("KAFKA_CONSUMER_GROUP_ID", "PKD_ARTICLE-GROUP"),
//     }
// }

// func getEnv(key, fallback string) string {
//     if value, ok := os.LookupEnv(key); ok {
//         return value
//     }
//     return fallback
// }



package config

import "os"

// KafkaConfig holds configuration for Kafka
type KafkaConfig struct {
    BootstrapServers string
    SecurityProtocol string
    SaslMechanism    string
    SaslJaasConfig   string
    Topic            string
    GroupID          string
    User             string
    Pass             string
}

// LoadKafkaConfig loads the Kafka configuration from environment variables
func LoadKafkaConfig() KafkaConfig {
    return KafkaConfig{
        BootstrapServers: getEnv("KAFKA_BOOTSTRAP_SERVERS", "localhost:9092"), // Externalize this value
        SecurityProtocol: getEnv("KAFKA_SECURITY_PROTOCOL", "SASL_SSL"),
        SaslMechanism:    getEnv("KAFKA_SASL_MECHANISM", "SCRAM-SHA-512"),
        Topic:            getEnv("KAFKA_TOPIC", "test2"),
        GroupID:          getEnv("KAFKA_CONSUMER_GROUP_ID", "PKD_ARTICLE-GROUP"),
        User:         getEnv("KAFKA_USERNAME", ""),
        Pass:         getEnv("KAFKA_PASSWORD", ""),
    }
}

func getEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

