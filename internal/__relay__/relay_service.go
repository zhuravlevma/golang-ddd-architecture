package relay

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	amqp "github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

func handle(db *gorm.DB, amqpChannel *amqp.Channel) {
	messages := []MessageOrm{}
	db.Where("published = ?", false).Find(&messages).Order("created_at DESC").Limit(100)

	var ids []int
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for _, message := range messages {
		ids = append(ids, message.ID)
		mapB, _ := json.Marshal(message.Payload)
		amqpChannel.PublishWithContext(ctx,
			"",     // exchange
		message.MessageType, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(mapB),
		})
	}

	if len(messages) > 0 {
		db.Delete(&MessageOrm{}, ids)
	}
}

func HandleCron(db *gorm.DB, amqpChannel *amqp.Channel) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return nil
	}
	j, err := s.NewJob(
		gocron.DurationJob(
			5*time.Second,
		),
		gocron.NewTask(
			handle,
			db,
			amqpChannel,
		),
	)
	if err != nil {
		return err
	}
	// each job has a unique id
	fmt.Println(j.ID())

	// start the scheduler
	s.Start()
	return nil
}
