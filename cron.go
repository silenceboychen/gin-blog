package main

import (
	"gin-blog/models"
	"gin-blog/pkg/logging"
	"github.com/robfig/cron"
	"time"
)

//https://segmentfault.com/a/1190000014666453
func main() {
	logging.Info("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		logging.Info("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		logging.Info("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}
