# nats-exploration

```bash

https://docs.nats.io/running-a-nats-service/clients

brew install nats-server

nats-server
[62402] 2024/10/07 13:20:02.338024 [INF] Starting nats-server
[62402] 2024/10/07 13:20:02.338576 [INF]   Version:  2.10.21
[62402] 2024/10/07 13:20:02.338580 [INF]   Git:      [not set]
[62402] 2024/10/07 13:20:02.338583 [INF]   Name:     NCK2IP3ZI23YIDSMQYKWJI2F3T5PILY6TEFAIIQ2AC7RRVMQ4ZYHVFEM
[62402] 2024/10/07 13:20:02.338586 [INF]   ID:       NCK2IP3ZI23YIDSMQYKWJI2F3T5PILY6TEFAIIQ2AC7RRVMQ4ZYHVFEM
[62402] 2024/10/07 13:20:02.340066 [INF] Listening for client connections on 0.0.0.0:4222
[62402] 2024/10/07 13:20:02.340314 [INF] Server is ready

brew tap nats-io/nats-tools
brew install nats-io/nats-tools/nats

# Subscribe
nats subscribe ">" -s nats://0.0.0.0:4222

# Publish, subscribe should receive hello world
nats pub hello world -s nats://0.0.0.0:4222
```

