package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/IBM/sarama"
)

type Order struct {
	CustomerName string `json:"customer_name"`
	CoffeType    string `json:"coffee_type"`
}

func main() {
	http.HandleFunc("/order", placeOrder)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	return sarama.NewSyncProducer(brokers, config)
}

func PushOrderToQueue(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}

	// Create connection
	producer, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}

	defer producer.Close()

	// Create a new message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// Send message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
	return nil
}

func placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 1. Parser request body into order
	order := new(Order)

	err := json.NewDecoder(r.Body).Decode(order)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 2. Convert body into bytes

	orderInBytes, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3. Send the bytes to kafka
	err = PushOrderToQueue("coffee_orders", orderInBytes)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4. Respond back to the user
	response := map[string]any{
		"success": true,
		"msg":     "Order for " + order.CustomerName + " is placed",
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "error placing order", http.StatusInternalServerError)
		return
	}
}
