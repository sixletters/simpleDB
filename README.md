author: Harris Maung
co-author: Dion Ng

This is a simple distributed key-value store based on a Btree Data structure, using Raft as the consensus algorithm. During my time at Goldman, I worked on the data serialization of a K-V store, and was introduced to B Trees and B+ Trees. At Risingwave, I worked mainly with Etcd, and as we all know Etcd uses Raft as the consensus algorithm. This is my simple implementation of the two combined for learning purposes.