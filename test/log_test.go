package test

import (
	logs "github.com/sirupsen/logrus"
	"sync"
	"testing"
	"time"
)

type People struct {
	Name string
	Age  int
}

func TestFile(t *testing.T) {
	var log = logs.New()
	logs.SetFormatter(&logs.JSONFormatter{})

	log.WithFields(logs.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		log := logs.WithFields(logs.Fields{
			"log_id": "qwer",
		})
		log.Info("with log id")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(5 * time.Second)
		logs.Info("with no")
	}()
	wg.Wait()
	t.Logf("end")
}
