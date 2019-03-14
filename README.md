# mqlogger

Basic logging tool for monitoring an MQTT topic subscription. Adds timestamp data to received payload before writing to file. 

### Download

```go get``` then build or install in the normal way

### Usage

```
./mqlogger -h

  -file string
        file to log to (default "log.txt")
  -host string
        host to connect to (default "localhost")
  -port string
        port to use (default "1883")
  -topic string
        topic to subscribe to (default "test")
```