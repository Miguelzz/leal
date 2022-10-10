package points_evt

import (
	"context"
	"encoding/json"
	"log"
	"loyalty/internal/core/domain"

	"github.com/segmentio/kafka-go"
)

type pointsEvent struct {
}

func New() *pointsEvent {
	return &pointsEvent{}
}

func (evn *pointsEvent) Redeem(points domain.Points) (int, error) {

	data, err := json.Marshal(points)
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9091", "points", 0)
	res, err := conn.WriteMessages(kafka.Message{Value: data})

	if err != nil {
		log.Fatal("failed:", err)
		return -1, nil
	}

	return res, nil
}

func (evn *pointsEvent) Buy(points domain.Points) (int, error) {

	data, err := json.Marshal(points)
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9091", "points", 0)
	res, err := conn.WriteMessages(kafka.Message{Value: data})

	if err != nil {
		log.Fatal("failed:", err)
		return -1, nil
	}

	return res, nil
}
