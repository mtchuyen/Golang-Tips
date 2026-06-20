# Golang with protocol

## UDP vs TCP Protocol

[TCP/IP Networking](https://appliedgo.net/networking/)

[Implementing UDP vs TCP in Golang](http://www.minaandrawos.com/2016/05/14/udp-vs-tcp-in-golang/)

## Modbus Protocol
[How to build a Modbus driver in Go?](http://www.minaandrawos.com/2014/11/26/how-to-write-a-modbus-driver/)

## Remote Procedure Call (RPC)

### Remote Procedure Call (RPC) briefly explained with Golang

https://mipsmonsta.medium.com/remote-procedure-call-rpc-briefly-explained-958842f6d786


## Protocol buffers (Protobuf)

See: https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-with-protocol/grpc.md

### What is Protobuf and Why You Should Use It
https://betterprogramming.pub/what-is-protobuf-and-why-you-should-use-it-14d52646f2a7

- [Salient Points for Protocol Buffers](https://mipsmonsta.medium.com/salient-points-for-protocol-buffers-f2a759ceb48)

### Go gRPC vs REST
https://helios04.medium.com/go-grpc-vs-rest-2c6a2428edbd

## [The Benchmark](https://medium.com/@maahisoft20/forget-json-these-4-data-formats-made-my-apis-5x-faster-a43a7b3935d6)

```
Format        | Size     | Serialize  | Deserialize
------------- | -------- | ---------- | -----------
JSON          | 1.8 MB   | 142 ms     | 98 ms
MessagePack   | 1.1 MB   | 61 ms      | 44 ms
Protobuf      | 680 KB   | 38 ms      | 29 ms
Avro          | 590 KB   | 35 ms      | 31 ms
FlatBuffers   | 720 KB   | 28 ms      | ~2 ms *
```
