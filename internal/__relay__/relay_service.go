package relay

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
	"gorm.io/gorm"
)

func handle(db *gorm.DB) {
	messages := []MessageOrm{}
	db.Where("published = ?", false).Find(&messages).Order("created_at DESC").Limit(100)
}

func HandleCron(db *gorm.DB) error {
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
