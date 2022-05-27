package montiors

func AlertTasks(c chan string) {
	c <- "send string back to tasks"
}
