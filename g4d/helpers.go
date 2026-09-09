package g4d

import (
	"context"
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"go.uber.org/zap"
)

func NewBot(token string, handler PanicHandler, gateway *gateway.Receiver, logger *zap.Logger, client *api.DiscordClient) *Bot {
	return &Bot{
		Token:        token,
		Logger:       logger,
		Gateway:      gateway,
		PanicHandler: handler,
		Client:       client,
	}
}
func (b *Bot) SetPrefix(pref string) *Bot {
	b.Prefix = pref
	return b
}

func WithDispQuantity(qnt int) _const.Quantity {
	return _const.Quantity(qnt)
}
func WithSemaphoreLimit(limit int) _const.SemaphoreLimit {
	return _const.SemaphoreLimit(limit)
}

func (b *Bot) MemReport(ctx context.Context, timeSleep time.Duration) {
	logger := b.Logger.Named("mem-reporter")
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				var m runtime.MemStats
				runtime.ReadMemStats(&m)
				logger.Info("stats",
					zap.Float32("alloc_mb", float32(m.Alloc)/1024/1024),
					zap.Float32("stack_mb", float32(m.StackInuse)/1024/1024),
					zap.Float32("sys_mb", float32(m.Sys)/1024/1024),
					zap.Float32("cache_sys_mb", float32(m.MCacheSys)/1024/1024),
					zap.Uint64("heap_objects", m.HeapObjects),
					zap.Int("goroutines", runtime.NumGoroutine()),
					zap.Float32("heap_idle_mb", float32(m.HeapIdle)/1024/1024),
					zap.Float32("heap_released_mb", float32(m.HeapReleased)/1024/1024),
					zap.Uint32("gc_count", m.NumGC),
					zap.Uint64("total_alloc", m.TotalAlloc),
				)
				time.Sleep(timeSleep)
			}
		}
	}(ctx)
}

func (b *Bot) StartProfiling(ctx context.Context) {
	f, err := os.Create("./cpu.prof")
	if err != nil {
		b.Logger.Fatal("Failed to create CPU profile", zap.Error(err))
		return
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		b.Logger.Error("Failed to start CPU profile", zap.Error(err))
		return
	}

	go func() {
		<-ctx.Done()
		pprof.StopCPUProfile()
		if err := f.Close(); err != nil {
			b.Logger.Error("Failed to close profile file", zap.Error(err))
		}
		b.Logger.Info("CPU profile saved to cpu.prof")
	}()

	b.Logger.Info("CPU profiling started")
}
