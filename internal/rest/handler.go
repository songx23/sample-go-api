package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type greeting struct {
	Content string `json:"content"`
}

type crazy struct {
	Answer string `json:"answer"`
}

func handleGreeting(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		name = "World"
	}

	c.JSON(http.StatusOK, greeting{
		Content: fmt.Sprintf("Hello, %s!", name),
	})
}

func handleGoCrazy(c *gin.Context) {
	n := c.Query("number")
	r := c.Query("repeat")
	number, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, crazy{})
		return
	}
	repeat, err := strconv.Atoi(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, crazy{})
		return
	}
	if err != nil {
		return
	}
	const workers = 1000
	jobs := make(chan int64, workers)
	results := make(chan int64, repeat)
	defer close(jobs)
	defer close(results)
	for w := 0; w < workers; w++ {
		go worker(w, jobs, results)
	}
	i := 0
	for i < repeat {
		i++
		jobs <- int64(number)
	}

	for a := 1; a <= repeat; a++ {
		<-results
	}

	c.JSON(http.StatusOK, crazy{Answer: "Calculation finished."})
}

func worker(id int, jobs <-chan int64, results chan<- int64) {
	for j := range jobs {
		f := factorial(j)
		results <- f
	}
}

func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
