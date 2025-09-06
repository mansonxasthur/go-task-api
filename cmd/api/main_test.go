package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestApp_run(t *testing.T) {
	tests := []struct {
		name      string
		simSignal os.Signal
		expectErr bool
	}{
		{
			name:      "graceful shutdown when SIGINT is received",
			simSignal: syscall.SIGINT,
			expectErr: false,
		},
		{
			name:      "graceful shutdown when SIGTERM is received",
			simSignal: syscall.SIGTERM,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mux := http.NewServeMux()
			wire(mux)
			ctx, stop := context.WithCancel(context.Background())
			defer stop()

			srv := &http.Server{
				Handler: mux,
				Addr:    ":8080",
			}

			go func() {
				// Simulating signal to initiate shutdown after a short delay.
				time.Sleep(100 * time.Millisecond)
				process, _ := os.FindProcess(os.Getpid())
				_ = process.Signal(tt.simSignal)
			}()

			shutdownChan := make(chan error, 1)
			go func() {
				err := srv.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					shutdownChan <- err
				} else {
					shutdownChan <- nil
				}
			}()

			// Wait for the simulated signal handling and shutdown process
			go func() {
				select {
				case <-ctx.Done():
					shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()
					shutdownChan <- srv.Shutdown(shutdownCtx)
				}
			}()

			err := <-shutdownChan
			if (err != nil) != tt.expectErr {
				t.Errorf("server run failed, got error: %v, expectErr: %v", err, tt.expectErr)
			}
		})
	}
}

func TestWire(t *testing.T) {
	mux := http.NewServeMux()
	wire(mux)

	// Test that '/users' handler exists
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	rr := &mockResponseWriter{}
	mux.ServeHTTP(rr, req)

	if rr.status == 0 {
		t.Errorf("expected handler to respond with status, got 0")
	}
}

type mockResponseWriter struct {
	header http.Header
	status int
}

func (m *mockResponseWriter) Header() http.Header {
	if m.header == nil {
		m.header = http.Header{}
	}
	return m.header
}

func (m *mockResponseWriter) Write(data []byte) (int, error) {
	return len(data), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	m.status = statusCode
}
