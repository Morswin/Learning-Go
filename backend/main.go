package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/* For anybody who will read this before I manage to split this into multiple files.
 * Yes, this will get reorganized.
 * No need to tell me. ;3
 */

// Idk if this is the intended way to use enums, this is what I found on the internet. Hopefully it will work with Gin
type TaskStatus int

const (
	taskActive TaskStatus = iota
	taskFinished
	taskCanceled
)

/* Yeah, I definitelly need to start reorganizing it.
 * I'm nowhere near anything working and it starts to become a mess.
 */

var taskNextId uint = 0

type Task struct {
	ID          uint
	title       string
	description string
	status      TaskStatus
}

func TaskGetNextId() uint {
	var ret = taskNextId
	taskNextId++
	return ret
}

var Tasks []Task

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func init() {
	Tasks = append(Tasks, Task{
		ID:          TaskGetNextId(),
		title:       "",
		description: "",
		status:      taskActive,
	})
}

func main() {
	r := gin.Default()

	r.GET("/", helloWorld)

	r.Run()
}
