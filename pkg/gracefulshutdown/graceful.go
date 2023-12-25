package gracefulshutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type GracefulShutdownCallbackFunc func(context.Context) error

type GracefulShutdownI interface {
	ExitProgram()
	CallbackRegister(string, GracefulShutdownCallbackFunc)
}

type gracefulShutdown struct {
	sigChan     chan os.Signal
	stopChan    chan struct{}
	callbacks   map[string]GracefulShutdownCallbackFunc
	proccessing int
	mu          sync.RWMutex
}

func NewWithContextTimeout(ctx context.Context, duration time.Duration) (GracefulShutdownI, context.Context) {
	g := gracefulShutdown{
		sigChan:   make(chan os.Signal),
		stopChan:  make(chan struct{}),
		callbacks: make(map[string]GracefulShutdownCallbackFunc),
		mu:        sync.RWMutex{},
	}

	signal.Notify(g.sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), duration)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("context cancel -> graceful shutdown")
				cancel()
				g.shutdown(ctxTimeout)
				cancelTimeout()

				os.Exit(0)
			case sig := <-g.sigChan:
				log.Printf("recieved signal %v -> graceful shutdown", sig)
				cancel()
				g.shutdown(ctxTimeout)
				cancelTimeout()

				os.Exit(0)
			}
		}
	}()

	return &g, ctx
}

func New() GracefulShutdownI {
	g := gracefulShutdown{
		sigChan:   make(chan os.Signal),
		stopChan:  make(chan struct{}),
		callbacks: make(map[string]GracefulShutdownCallbackFunc),
		mu:        sync.RWMutex{},
	}

	signal.Notify(g.sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		sig := <-g.sigChan
		log.Printf("recieved signal %v -> graceful shutdown", sig)

		for {
			time.Sleep(100 * time.Millisecond)
			if g.proccessing == 0 {
				g.shutdown(context.Background())
				os.Exit(0)
			}
		}
	}()

	return &g
}

func (g *gracefulShutdown) shutdown(ctx context.Context) {
	// shutdown callback
	if len(g.callbacks) > 0 {
		var wg sync.WaitGroup

		for key, callback := range g.callbacks {
			wg.Add(1)

			innerCallback := callback
			innerKey := key

			go func(ctx context.Context) {
				defer wg.Done()
				if err := innerCallback(ctx); err != nil {
					log.Printf("clean up fail [%s] %s\n", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully\n", innerKey)
			}(ctx)
		}

		wg.Wait()
	}
}

func (g *gracefulShutdown) CallbackRegister(name string, cb GracefulShutdownCallbackFunc) {
	if _, ok := g.callbacks[name]; !ok {
		g.callbacks[name] = cb
	}
}

func (g *gracefulShutdown) ExitProgram() {
	g.sigChan <- syscall.SIGTERM
}

func (g *gracefulShutdown) IncreaseProcessing() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.proccessing++
}

func (g *gracefulShutdown) DecreaseProcessing() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.proccessing--
}
