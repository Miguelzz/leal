package points_repo

import (
	"context"
	"encoding/json"
	"log"
	"loyalty/internal/core/domain"
	"loyalty/internal/repositories/mocks"

	"github.com/segmentio/kafka-go"
)

type pointsTable struct {
}

func New() *pointsTable {

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "points", 0)
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

	return &pointsTable{}
}

func (repo *pointsTable) Get(id int) (domain.Points, error) {
	return mocks.MockReadPoints[id], nil
}
