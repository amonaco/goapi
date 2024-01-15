package worker

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"

	"github.com/amonaco/goapi/libs/database"
)

// Task message and channel
type Task struct {
	Name   string
	Fields map[string]string
}

// The struct for the worker
type Worker struct {
	channels [](chan *Task)
	quit     chan bool
	Max      int
	Handler  func(work *Task)
}

type Notification struct {
	ID        uint32    `json:"id"`
	Data      string    `json:"data"`
	CreatedAt time.Time `json:"created_at" pg:"default:now()"`
}

// Package namespace variables
var worker *Worker
var max = 4

// Adds a field to a task
func (task *Task) AddField(key string, value string) {
	task.Fields[key] = value
}

// Creates a new worker
func Init(handler func(task *Task)) {
	w := &Worker{
		Max:     max,
		Handler: handler,
	}
	w.channels = make([](chan *Task), max)
	w.quit = make(chan bool)

	for i := 0; i < max; i++ {
		log.Printf("[worker][%d] starting up\n", i)
		w.channels[i] = make(chan *Task)
		go w.runner(w.channels[i], i)
	}
	worker = w
}

// Initializes a new task
func (worker *Worker) NewTask(name string) *Task {
	task := &Task{Name: name}
	task.Fields = make(map[string]string)
	return task
}

// Stop a worker (may be never used)
func (worker *Worker) Stop() {
	worker.quit <- true
}

// Pushes a message to a random worker
func (worker *Worker) Push(task *Task) {
	id := getRandom()
	worker.channels[id] <- task

	// Model for notifications
	notification := &Notification{}

	db := database.Get()
	data, err := json.Marshal(task)
	if err != nil {
		log.Printf("[worker] cannot marshal task into json\n")
	}
	notification.Data = string(data)

	// Persist notifications in database
	_, err = db.Model(notification).Insert()
	if err != nil {
		log.Printf("[worker] issue persisting notification\n")
	}
}

// Handles messages and runs the worker
func (worker *Worker) runner(c chan *Task, id int) {
	for {
		select {
		case <-worker.quit:
			return
		case work := <-c:
			log.Printf("[worker][%d] received task\n", id)
			worker.Handler(work)
		}
	}
}

// Gets a random number between 0 and max workers
func getRandom() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

// Returns the worker pointer
func Get() *Worker {
	return worker
}
