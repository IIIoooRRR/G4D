package g4d

import (
	"context"
	"os"
	"runtime"
	"runtime/metrics"
	"runtime/pprof"
	"time"

	"github.com/IIIoooRRR/G4D/api"
	"github.com/IIIoooRRR/G4D/gateway"
	"github.com/IIIoooRRR/G4D/model/_const"
	"go.uber.org/zap"
)

func NewBot(token string, handler PanicHandler, gateway *gateway.Receiver, logger *zap.Logger, client *api.DiscordClient) *Bot {
	bot := &Bot{
		Token:        token,
		Logger:       logger,
		Gateway:      gateway,
		PanicHandler: handler,
		Client:       client,
	}
	bot.processorPool = map[_const.ProcessorType]func(*Bot, int, chan struct{}){
		_const.StaticEventProcessor:  (*Bot).staticEventProcessor,
		_const.DynamicEventProcessor: (*Bot).dynamicEventProcessor,
	}
	return bot
}

func (b *Bot) SetPrefix(pref string) *Bot {
	b.Prefix = pref
	return b
}

func WithProcessorsQuantity(qnt int) _const.Quantity {
	return _const.Quantity(qnt)
}
func WithSemaphoreLimit(limit int) _const.SemaphoreLimit {
	return _const.SemaphoreLimit(limit)
}

func (b *Bot) StartProfiling(ctx context.Context) {
	f, err := os.Create("./cpu.prof")
	if err != nil {
		b.Logger.Error("Failed to create CPU profile", zap.Error(err))
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

func (b *Bot) MemReport(ctx context.Context, timeSleep time.Duration) {
	if b.isMemReportActive.Swap(true) {
		return
	}

	logger := b.Logger.Named("mem-reporter")

	const (
		allocName    = "/memory/classes/heap/objects:bytes"
		stackName    = "/memory/classes/heap/stacks:bytes"
		sysName      = "/memory/classes/total:bytes"
		mcacheName   = "/memory/classes/metadata/mcache:bytes"
		objectsName  = "/memory/classes/heap/objects:objects"
		gcName       = "/gc/cycles:total"
		heapIdleName = "/memory/classes/heap/free:bytes"
	)

	samples := []metrics.Sample{
		{Name: allocName},
		{Name: stackName},
		{Name: heapIdleName},
	}

	go func(ctx context.Context) {
		defer b.isMemReportActive.Store(false)

		ticker := time.NewTicker(timeSleep)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				logger.Info("mem reporter stopped")
				return
			case <-ticker.C:
				metrics.Read(samples)

				var alloc, stack, heapIdle uint64

				if samples[0].Value.Kind() == metrics.KindUint64 {
					alloc = samples[0].Value.Uint64()
				}
				if samples[1].Value.Kind() == metrics.KindUint64 {
					stack = samples[1].Value.Uint64()
				}
				if samples[2].Value.Kind() == metrics.KindUint64 {
					heapIdle = samples[2].Value.Uint64()
				}

				logger.Info("stats",
					zap.Float32("alloc_mb", float32(alloc)/1024/1024),
					zap.Float32("stack_mb", float32(stack)/1024/1024),
					zap.Int("goroutines", runtime.NumGoroutine()),
					zap.Float32("heap_idle_mb", float32(heapIdle)/1024/1024),
				)
			}
		}
	}(ctx)
}
