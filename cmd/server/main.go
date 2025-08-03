package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/ricardolindner/go-expert-stress-test/internal/stresstest"
	"github.com/schollz/progressbar/v3"
)

var (
	url         string
	totalReqs   int
	concurrency int
)

func init() {
	flag.StringVar(&url, "url", "", "URL do serviço a ser testado")
	flag.IntVar(&totalReqs, "requests", 1, "Número total de requests")
	flag.IntVar(&concurrency, "concurrency", 1, "Número de chamadas simultâneas")
}

func main() {
	flag.Parse()

	if url == "" {
		fmt.Println("Erro: A URL é um parâmetro obrigatório.")
		flag.Usage()
		os.Exit(1)
	}
	if totalReqs <= 0 || concurrency <= 0 {
		fmt.Println("Erro: O número de requests e a concorrência devem ser maiores que zero.")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("\nIniciando teste de carga para a URL: %s\n", url)
	fmt.Printf("Total de requests: %d, Concorrência: %d\n", totalReqs, concurrency)
	fmt.Println("----------------------------------------")

	resultsChan := make(chan stresstest.TestResult, totalReqs)
	var wg sync.WaitGroup

	bar := progressbar.NewOptions(totalReqs,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(50),
		progressbar.OptionSetDescription("[cyan]Testando...[reset]"),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	startTime := time.Now()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			stresstest.Worker(context.Background(), workerID, url, totalReqs/concurrency, resultsChan, bar)
		}(i)
	}

	remainingReqs := totalReqs % concurrency
	if remainingReqs > 0 {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			stresstest.Worker(context.Background(), workerID, url, remainingReqs, resultsChan, bar)
		}(concurrency)
	}

	wg.Wait()
	close(resultsChan)

	duration := time.Since(startTime)
	fmt.Println("\nTeste concluído.")
	fmt.Println("----------------------------------------")

	var allResults []stresstest.TestResult
	for result := range resultsChan {
		allResults = append(allResults, result)
	}

	stresstest.GenerateReport(allResults, duration, totalReqs)
}
