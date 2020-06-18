# redisserver

## Running & Testing
### Build + Test
Run test & build with `make`
```
[vagrant@gonew redisserver]# make
# Run go test with cache disabled
CGO_ENABLED=0 go test -cover ./... -count=1
?       github.com/nvaatstra/redisserver/cmd/server     [no test files]
ok      github.com/nvaatstra/redisserver/pkg/commands   0.002s  coverage: 90.5% of statements
?       github.com/nvaatstra/redisserver/pkg/datatypes  [no test files]
ok      github.com/nvaatstra/redisserver/pkg/server     2.004s  coverage: 76.5% of statements
CGO_ENABLED=0 go build -o bin/redis_server ./cmd/server/
```

### Build + Test & Run
Run test & build + start server with `make run`
```
[vagrant@gonew redisserver]# make run
# Run go test with cache disabled
CGO_ENABLED=0 go test -cover ./... -count=1
?       github.com/nvaatstra/redisserver/cmd/server     [no test files]
ok      github.com/nvaatstra/redisserver/pkg/commands   0.002s  coverage: 90.5% of statements
?       github.com/nvaatstra/redisserver/pkg/datatypes  [no test files]
ok      github.com/nvaatstra/redisserver/pkg/server     2.005s  coverage: 76.5% of statements
CGO_ENABLED=0 go build -o bin/redis_server ./cmd/server/
bin/redis_server -addr=127.0.0.1:1234
```