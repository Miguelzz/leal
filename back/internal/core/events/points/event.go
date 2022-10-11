package points_evt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"loyalty/internal/core/domain"
	"loyalty/internal/repositories/mocks"

	"github.com/segmentio/kafka-go"
)

type pointsEvent struct {
}

func New() *pointsEvent {

	go func() {

		fmt.Println("points <-")
		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9091", "points", 0)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		for {
			m, err := conn.ReadMessage(10e3)
			if err != nil {
				break
			}

			p := domain.Points{}
			json.Unmarshal(m.Value, &p)
			mocks.MockReadPoints[p.ID] = p
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close reader:", err)
		}
	}()
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
	fmt.Println("points ->")
	data, err := json.Marshal(points)
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9091", "points", 0)
	res, err := conn.WriteMessages(kafka.Message{Value: data})

	if err != nil {
		log.Fatal("failed:", err)
		return -1, nil
	}

	return res, nil
}
