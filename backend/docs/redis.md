# Redis

## Overview

Redis is used in this project for:

- **Caching**: Reduce database load for frequently accessed data
- **Session Store**: Manage user sessions (if JWT blacklisting/token invalidation is needed)
- **Rate Limiting**: Track API request rates per user/IP
- **Message Queue**: Background job processing (email, notifications, etc.)
- **Distributed Locking**: Prevent race conditions in concurrent operations

## Connection

| Env Variable      | Default       | Description          |
|-------------------|---------------|----------------------|
| `REDIS_HOST`      | `localhost`   | Redis server host    |
| `REDIS_PORT`      | `6379`        | Redis server port    |
| `REDIS_PASSWORD`  | (empty)       | Redis auth password  |
| `REDIS_DB`        | `0`           | Redis database index |

Connection is initialized in `pkg/database/redis.go`.

## Key Naming Convention

All Redis keys follow this format:

```
{service}:{entity}:{identifier}:{field}
```

Examples:
- `destiny:session:{userID}` — user session data
- `destiny:cache:user:{userID}` — cached user profile
- `destiny:ratelimit:api:{userID}` — rate limit counters
- `destiny:queue:email:{jobID}` — email job data
- `destiny:lock:task:{taskID}` — distributed lock for task updates

All keys use the `destiny:` prefix to namespace within the shared Redis instance.

## Expiration / TTL

| Data Type         | Default TTL |
|-------------------|-------------|
| Session data      | 24h         |
| Cache entries     | 5m          |
| Rate limit window | 1m          |
| Locks             | 10s         |

## Usage Patterns

### Caching
```go
// Get from cache
data, err := redis.Get(ctx, key).Bytes()
if err == redis.Nil {
    // Miss — load from DB, then set cache
    data = loadFromDB()
    redis.Set(ctx, key, data, 5*time.Minute)
}
```

### Rate Limiting
```go
// Increment counter with expiry
count, err := redis.Incr(ctx, "destiny:ratelimit:api:"+userID).Result()
if count == 1 {
    redis.Expire(ctx, key, 1*time.Minute)
}
if count > 100 {
    // Reject request
}
```

### Distributed Locking
```go
ok, err := redis.SetNX(ctx, lockKey, processID, 10*time.Second).Result()
if ok {
    defer redis.Del(ctx, lockKey)
    // Do work
}
```

## Implementation Files

- `pkg/database/redis.go` — Redis client initialization and configuration
- `internal/repository/` — Data access layer using Redis for caching
- `internal/middleware/` — Rate limiting middleware using Redis