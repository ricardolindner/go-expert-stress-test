package stresstest

import (
	"fmt"
	"net/http"
	"sort"
	"time"
)

func GenerateReport(results []TestResult, totalDuration time.Duration, totalRequests int) {
	statusCodeCounts := make(map[int]int)
	status200Count := 0
	totalLatency := time.Duration(0)

	for _, result := range results {
		statusCodeCounts[result.StatusCode]++
		totalLatency += result.Duration
		if result.StatusCode == http.StatusOK {
			status200Count++
		}
	}

	fmt.Println("\n### Relatório do Teste de Carga ###")
	fmt.Printf("Tempo total gasto: %s\n", totalDuration.Truncate(time.Millisecond))
	fmt.Printf("Total de requests: %d\n", totalRequests)
	fmt.Println("\n----------------------------------------")
	fmt.Printf("\nRequisições retornadas com status HTTP 200:\n")
	fmt.Printf("  - Requisições com status 200: %d\n", status200Count)

	var otherStatusCodes []int
	connectionErrorCount := statusCodeCounts[0]
	delete(statusCodeCounts, 0)
	for code := range statusCodeCounts {
		if code != http.StatusOK {
			otherStatusCodes = append(otherStatusCodes, code)
		}
	}
	sort.Ints(otherStatusCodes)

	if connectionErrorCount > 0 {
		fmt.Printf("\nRequisições com URL inválida ou servidor offline:\n")
		fmt.Printf("  - Requisições com erro: %d\n", connectionErrorCount)
	}

	if len(otherStatusCodes) > 0 {
		fmt.Println("\nDistribuição dos demais status HTTP:")
		for _, code := range otherStatusCodes {
			fmt.Printf("  - Requisições com status %d: %d\n", code, statusCodeCounts[code])
		}
	} else if connectionErrorCount == 0 {
		fmt.Println("Não foram registrados status diferente de sucesso")
	}
	fmt.Println("\n----------------------------------------")

	if totalRequests > 0 {
		avgLatency := totalLatency / time.Duration(totalRequests)
		fmt.Printf("\nMédia de latência por request: %s\n", avgLatency.Truncate(time.Microsecond))
	}
}
