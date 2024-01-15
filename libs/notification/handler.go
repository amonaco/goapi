package notification

import (
	"log"

	"github.com/amonaco/goapi/libs/worker"
)

func Handler(task *worker.Task) {
	switch task.Name {

	case "user_create":
		log.Printf("[handler] user_create notification received %v, %v\n", task.Name, task.Fields)
		userCreate(task.Fields)
	case "user_forgot_password":
		log.Printf("[handler] forgot password notification received %v, %v\n", task.Name, task.Fields)
		userForgotPassword(task.Fields)
	default:
		log.Printf("[handler] unknown task received %v, %v\n", task.Name, task.Fields)
		return
	}
}
