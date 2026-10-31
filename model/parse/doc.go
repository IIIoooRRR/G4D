/*
This package intentionally trades some "magic" for predictability.

Design goals:

 1. One parse per event, not per command.
    Every command must NOT re-unmarshal the raw event. The parser
    unmarshals once, stores the result in a per-processor cache slot,
    and commands read it via GetEvent[T]. This keeps CPU and GC pressure
    proportional to the number of *events*, not the number of *commands*.

 2. No reflection in the hot path.
    Earlier versions used reflect.New(t) on every event. That was
    replaced with per-type factories:

    _const.EventMessageCreate: func() any { return &schema.GetMessage{} }

    The factory allocates the concrete type directly. No reflect.New,
    no reflect.TypeOf, no interface boxing on the hot path.

    Reflection is *not* used here at all. If you want to add a type
    that is not registered, use types.Add / types.Change *before*
    Bot.Run. The registry is read-only at runtime.

 3. Two build modes: safe and unsafe.
    The package ships two implementations selected by build tag:

    //go:build !unsafe   -> EventEntry.Data is any, GetEvent uses a type assertion
    //go:build unsafe    -> EventEntry.Data is unsafe.Pointer, GetEvent uses a direct cast

    The safe build is the default. It is fully memory-safe and works
    with the race detector. The unsafe build removes one interface
    indirection per GetEvent call; it is opt-in via `-tags unsafe` and
    is intended for users who have audited the code and accept the
    trade-off.

    The unsafe build does not change the public API. Both builds expose
    the same GetEvent[T] signature; only the internal representation
    differs.

 4. Isolation between commands.
    Each command runs in its own goroutine, gated by a semaphore
    (limiter). A slow command does not  block other commands for the
    same event, and it does block the processor from picking up
    the next event - except for the explicit wg.Wait() below.

 5. Event pooling and explicit lifetime.
    The processor waits for all commands of an event to finish
    (wg.Wait) before returning the event to the pool. This is a
    deliberate trade-off: throughput is slightly lower, but the event
    buffer is reused, which keeps allocations flat under load.

    If you need maximum throughput and do not want the pool, fork the
    processor or implement your own ProcessorType. The library does
    not hide this choice from you.

 6. No mutexes in the parse path.
    Each processor owns its cache slot (indexed by seq). There is no
    shared mutable state between processors, so no locks are needed.
    The only synchronization is the WaitGroup that ties a command's
    lifetime to the event's lifetime.

 7. Adding new event types.
    If a Discord event type is missing from the default registry, use
    the helpers in parse/types:

    types.Add(eventType, func() any { return &MyStruct{} })   // add, no-op if present
    types.Change(eventType, func() any { return &MyStruct{} }) // replace existing

    Both must be called before Bot.Run. After Run, the registry is
    treated as immutable.

 8. What this package does NOT do.
    It does not retry, rate-limit, or deduplicate events. It does not
    guess your command routing. It does not hide the WebSocket or the
    HTTP client. If you want a router, build one on top of
    CommandTemplate - the library gives you the primitives, not the
    policy.
*/
package parse
