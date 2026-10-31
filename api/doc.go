// Package api provides a client for working with the Discord REST API.
//
// It includes implemented helpers (SendMessage, SendMessageWithLimit) and a
// low-level request primitive that takes a method, URI, and body.
//
//   - Uses Discord API v10 by default.
//   - DoDiscordRequest accepts an endpoint URI starting with "/",
//     e.g. "/guilds/{guild_id}/members/{user_id}".
//   - DoDiscordLimitRequest respects Discord rate limits (per-route buckets
//     advertised via X-RateLimit-* headers) and accepts a context.Context.
//   - HTTP transport is based on github.com/valyala/fasthttp.
//   - A future update is planned to migrate the limiter to a dedicated
//     rate-limiter implementation.
//
// API methods are available via ctx.Context, which is passed in the command
// signature.
package api
