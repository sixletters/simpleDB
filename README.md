## SimpleDB
author: Harris Maung
co-author: Dion Ng

This is a simple distributed key-value store based on a Btree Data structure, using Raft as the consensus algorithm. During my time at Goldman, I worked on the data serialization of a K-V store, and was introduced to B Trees and B+ Trees. At Risingwave, I worked mainly with Etcd, and as we all know Etcd uses Raft as the consensus algorithm. This is my simple implementation of the two combined for learning purposes.

## Build
``` make build```

outputs a binary in the output/bin directory.

## Usage
```./cmd --host $host --port $port```
+ host: host of the server
+ port: port that server should listen on

## Run by make file
``` make run``` 

## Test
``` make test```