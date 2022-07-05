# Algorithms in Golang

Basic: 
https://github.com/go-awesome/Data-Structures-and-Algorithms

https://github.com/kylesliu/awesome-golang-algorithm

[Dynamic Programming in Go](https://medium.com/@kayahuseyin/dynamic-programming-in-go-a87e0b3c5ae0)

## Encode & Decode
### Base64

- `ad-02-03-fc-14-ee` --> `YWQtMDItMDMtZmMtMTQtZWU=`


### Base16 to Decimal

***Encode***:

```go
ecid, erD := strconv.ParseUint(stringid, 16, 64)
```

- `stringid=fffffffffffffffa` --> `ecid=18446744073709551610`
- stringid = `regex:[a-f0-9]`, length <= `2x`
- ecid = `uint64`

***Decode***:

```go
ma, err := strconv.ParseInt(i, 10, 64)
```
- raw: `ad:02:03:fc:14:ee` --> encode: `190224168391918` --> decode: `ad0203fc14ee`.



## Hash Algorithms

[See all contents](https://github.com/mtchuyen/Golang-Tips/blob/master/Golang-Algorithms/hash.md)

***List:***
- md: md5, md6
- SHA
- RIPEMD
- WHIRLPOOL
- TIGER
- SNEFRU
- GOST
- ADLER
- CRC
- FNV
- JOAAT
- HAVAL
- murmur3
- DJB2
- SDBM
- Farm hash fingerprint32, fingerprint64
- Pearson
- Loselose


## Sorting Algorithms: 

### Quick Sort in Golang

https://agupta97.medium.com/sorting-algorithms-quick-sort-in-golang-cf1866f395c8

### Insertion Sort in Golang

https://agupta97.medium.com/sorting-algorithms-insertion-sort-in-golang-c014e8e3d082

### Merge Sort in Golang

https://medium.com/geekculture/sorting-algorithms-merge-sort-in-golang-2ae73ff07906

### Dijkstra’s algorithm in Go

https://go-recipes.dev/dijkstras-algorithm-in-go-e1129b2f5c9e

### binary search algorithm

https://medium.com/@alissakhall/binary-search-with-go-7d6291fb551e?

### Skip List 
https://lmsang1986.violet.vn/entry/show/entry_id/7018749

https://medium.com/codex/skiplist-in-go-267bb81e51cd



