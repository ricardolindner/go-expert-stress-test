package stresstest

import (
	"context"
	"net/http"
	"time"

	"github.com/schollz/progressbar/v3"
)

type TestResult struct {
	StatusCode int
	Duration   time.Duration
}

func Worker(ctx context.Context, id int, url string, numRequests int, results chan<- TestResult, bar *progressbar.ProgressBar) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	for i := 0; i < numRequests; i++ {
		startTime := time.Now()
		resp, err := client.Get(url)
		duration := time.Since(startTime)

		statusCode := 0
		if err == nil {
			statusCode = resp.StatusCode
			resp.Body.Close()
		}

		results <- TestResult{
			StatusCode: statusCode,
			Duration:   duration,
		}

		bar.Add(1)
	}
}
