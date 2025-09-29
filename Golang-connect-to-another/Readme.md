# Go ecosystem

## Prometheus
https://blog.devops.dev/monitor-go-app-with-prometheus-grafana-guide-650c50860667

## Message Queue System/Platform
- NSQ 
- NATS 

NSQ is an open source realtime distributed messaging platform which is a successor from simplequeue.

NSQ is an open source real-time distributed messaging platform that handles millions of messages every day. It delivers reliable messages with no fault tolerance and high availability.

NSQ and Kafka are both message queuing services. NSQ is a simpler to configure and more easy to deploy message queue platform while Kafka ensures strict guarantees and reliability with no data loss.

https://medium0.com/@philipfeng/modern-open-source-messaging-apache-kafka-rabbitmq-nats-pulsar-and-nsq-ca3bf7422db5

https://www.quora.com/What-is-the-difference-between-Apache-Kafka-and-NATS

https://blog.containerize.com/nsq-vs-kafka-what-are-the-key-differences/

## Cache
Trong [Golang-connect-to-another: Cache](https://github.com/mtchuyen/Golang-Tips/edit/master/Golang-connect-to-another/Cache.md) đã nói tới các tiêu chí để đánh giá 1 Local Cache. Link dưới cũng trình bày tiêu chí về cache và ưu nhược điểm của các Local-Cache phổ biến hiện nay (bằng Golang):
- https://hackernoon.com/in-memory-caching-in-golang
- https://dgraph.io/blog/post/caching-in-go/

***Requirements for the cache***: 5 tiêu chí để đánh giá 1 thư viện (giải thuật) cache
1. Concurrent.
2. Memory-bounded (limit to configurable max memory usage).
3. Scale well as the number of cores and goroutines increase.
4. Scale well under non-random key access distribution (e.g. Zipf).
5. High cache hit ratio

Trong ***tiêu chí số 4***, DGRAPH sử dụng phân phối Zipflian, theo định luật Zipf (Zipf's Law)
> Typical access patterns follow a Zipfian distribution. The most frequently accessed keys are accessed exponentially more times than others.

***Ví dụ về 4 tiêu chí***:
- Use Go map with sync.Mutex: `Fails on requirements 2,4.`
- Use lock striping with Go maps: `Fails on requirements 2,4.`
- LRU cache: `Fails on requirement 3-4.`
- Striped LRU cache: `Would fail on requirement 4`.

Đọc thêm về `HashMap and ConcurrentHashMap in Golang`

***Popular Cache Implementations & Performance***
- BigCache
- FreeCache
- GroupCache

Trong [In-Memory Caching in Golang](https://hackernoon.com/in-memory-caching-in-golang) nêu quan điểm cá nhân về ưu/nhược điểm của từng loại cache

### Striping Cache - sharding Cache
- splits keys using a fingerprint into many smaller Go map shards protected by mutex locks
- it would only be an incremental improvement and WOULD not SCALE well.
- Implementing a Sieve cache in Go: https://medium.com/@mehul25/implementing-a-sieve-cache-in-go-9652cfa99278


### Chain cache:

https://github.com/n1ord/chaincache

https://github.com/eko/gocache

## Redis

- [Scalable event streaming with Redis and Golang](https://ably.com/blog/event-streaming-with-redis-and-golang)
- [Redis Bitmaps: Storing State in Small Places](https://www.improving.com/thoughts/redis-bitmaps-storing-state-in-small-places/)
## RabbitMQ

[Work queue with Go and RabbitMQ](https://medium.com/@masnun/work-queue-with-go-and-rabbitmq-b8c295cde861)

[Message Queues in Golang Via RabbitMQ](https://medium.com/@agiratech/message-queues-in-golang-via-rabbitmq-a3be7e426ad4)


## NATS, NAT-IO

[Building messaging in Go network clients](https://www.oreilly.com/ideas/building-messaging-in-go-network-clients)

## Websocket

## Database

[Understanding Go and Databases at Scale: Connection Pooling](https://koho.dev/understanding-go-and-databases-at-scale-connection-pooling-f301e56fa73)


### MysQL

### MariaDB

### Postgre
https://viblo.asia/p/kien-truc-cua-postgresql-E375zAB1lGW

https://viblo.asia/p/8-diem-so-sanh-giua-mysql-va-postgresql-de-chon-lua-cai-nao-phu-hop-hon-OeVKB4NElkW

http://sqladvice.com/postgresql-la-gi-so-sanh-mysql-va-postgresql/

### MongoDB

[Make yourself a Go web server with MongoDb](https://hackernoon.com/make-yourself-a-go-web-server-with-mongodb-go-on-go-on-go-on-48f394f24e)

[how to write a URL Shortener Golang Web Service with Mongodb?](http://www.minaandrawos.com/2015/09/05/link-shortener-golang-web-service-tutorial-mongodb/)

### NonSQL

### Couchbase 
[Developing a Go with Couchbase Application](https://blog.couchbase.com/create-continuous-deployment-pipeline-golang-jenkins/)

### Kafka:
[1] [Visualizing Kafka](https://timothystepro.medium.com/visualizing-kafka-20bc384803e7)

[2] [How to parallelise Kafka consumers](https://medium.com/@jhansireddy007/how-to-parallelise-kafka-consumers-59c8b0bbc37a)

[3] [How to Consume Kafka Efficiently in Golang?](https://medium.com/swlh/how-to-consume-kafka-efficiently-in-golang-264f7fe2155b)

### DuckDB:

https://betterprogramming.pub/duckdb-whats-the-hype-about-5d46aaa73196

## Microservices
[Microservices in Golang](https://ewanvalentine.io/microservices-in-golang-part-1/)

[Golang Microservices Challenge: Designing an extensible, easy-to-use, and testable HTTP client](https://medium.com/augury-research-and-development/golang-microservices-challenge-designing-an-extensible-easy-to-use-and-testable-http-client-faf43e7e5d45)

[Golang, Microservices and Twirp](https://itnext.io/golang-microservices-and-twirp-5ef495278ddf)

[Building Go Microservices Focused on Testability](https://medium.com/@rholcombe30/building-go-microservices-focused-on-testability-d6164751275d)

