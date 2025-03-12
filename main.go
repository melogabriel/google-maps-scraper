package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/melogabriel/google-maps-scraper/runner"
	"github.com/melogabriel/google-maps-scraper/runner/databaserunner"
	"github.com/melogabriel/google-maps-scraper/runner/filerunner"
	"github.com/melogabriel/google-maps-scraper/runner/installplaywright"
	"github.com/melogabriel/google-maps-scraper/runner/lambdaaws"
	"github.com/melogabriel/google-maps-scraper/runner/webrunner"
)

var (
	logs []string
	mu   sync.Mutex
)

// Adds a message to the log buffer
func addLog(message string) {
	mu.Lock()
	logs = append(logs, message)
	mu.Unlock()
	log.Println(message) // Also prints to the Render console
}

// HTTP endpoint to return logs as JSON
func logsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"logs": %q}`, logs)
}

// Start an HTTP server for logs
func startLogServer() {
	http.HandleFunc("/logs", logsHandler)
	go func() {
		port := "8080"
		log.Println("Log server running at http://localhost:" + port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Failed to start HTTP server:", err)
		}
	}()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	runner.Banner()
	addLog("Application started")

	// Start log server
	startLogServer()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		addLog("Received termination signal, shutting down...")
		cancel()
	}()

	cfg := runner.ParseConfig()

	runnerInstance, err := runnerFactory(cfg)
	if err != nil {
		addLog("Error creating runner: " + err.Error())
		cancel()
		os.Stderr.WriteString(err.Error() + "\n")

		runner.Telemetry().Close()
		os.Exit(1)
	}

	if err := runnerInstance.Run(ctx); err != nil {
		addLog("Execution error: " + err.Error())
		os.Stderr.WriteString(err.Error() + "\n")

		_ = runnerInstance.Close(ctx)
		runner.Telemetry().Close()
		cancel()
		os.Exit(1)
	}

	addLog("Execution completed successfully!")

	_ = runnerInstance.Close(ctx)
	runner.Telemetry().Close()
	cancel()
	os.Exit(0)
}

func runnerFactory(cfg *runner.Config) (runner.Runner, error) {
	switch cfg.RunMode {
	case runner.RunModeFile:
		addLog("Running in File mode")
		return filerunner.New(cfg)
	case runner.RunModeDatabase, runner.RunModeDatabaseProduce:
		addLog("Running in Database mode")
		return databaserunner.New(cfg)
	case runner.RunModeInstallPlaywright:
		addLog("Installing Playwright")
		return installplaywright.New(cfg)
	case runner.RunModeWeb:
		addLog("Running in Web mode")
		return webrunner.New(cfg)
	case runner.RunModeAwsLambda:
		addLog("Running in AWS Lambda mode")
		return lambdaaws.New(cfg)
	case runner.RunModeAwsLambdaInvoker:
		addLog("Invoking AWS Lambda")
		return lambdaaws.NewInvoker(cfg)
	default:
		err := fmt.Errorf("%w: %d", runner.ErrInvalidRunMode, cfg.RunMode)
		addLog("Error: " + err.Error())
		return nil, err
	}
}
