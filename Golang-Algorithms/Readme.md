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

### Fibonacci series with Golang and Dynamic Programming
https://medium.com/@deepalipawade/fibonacci-series-with-golang-and-dynamic-programming-75bd60c7fd21


## Data Structures and Algorithms

### Linked Lists
Linked Lists là một cấu trúc dữ liệu dạng collection, nó là tập họp của nhiều object được gọi là node, và các node này được nối với nhau thông qua một liên kết được gọi là `link`.

`Linked Lists` cũng tương tự như Array, nhưng ở một số ngôn ngữ khi ta làm việc với Array, thì ta thường phải đối mặt với một số vấn đề về độ dài của Array. `Linked Lists` sẽ giúp ta giải quyết các vấn đề đó.

#### Real World Example
Vậy Linked Lists sẽ được sử dụng ở đâu trong thực tế và nó giúp ta giải quyết những vấn đề gì?

Linked Lists được sử dụng cho khá nhiều ứng dụng trong khoa học máy tính (Computer Science), ví dụ:

- Dynamic memory allocation : We use linked list of free blocks.
- Maintaining directory of names.
- Performing arithmetic operations on long integers.
- Manipulation of polynomials by storing constants in the node of linked list.
- Representing sparse matrices.


Còn trong thực tế ta có thể sử dụng Linked Lists cho các ứng dụng như:
- Image viewer.
- Previous and next page in web browser.
- Music Player.

Các ứng dụng trên thì ta không thể dùng Linked Lists bình thường được, mà ta cần dùng ***Doubly Linked Lists***, các bạn có thể đọc để biết rõ hơn.



#### Ref:

https://viblo.asia/p/golang-data-structures-and-algorithms-linked-lists-0gdJz07vVz5



