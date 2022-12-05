# Mutex

## Tổng quan
Để bảo vệ dữ liệu trong critical resource (tức là handle data/resource race), người ta có thể sử dụng semaphore hoặc lock.

### several ways to implement locks.
1. Semaphore
2. Semaphores and locks.
3. Spin lock
4. Pessimistic locking and optimistic locking

### 1. Semaphore
Chúng ta sẽ xem lại cơ chế hoạt động của semaphore trong Linux kernel
- Semaphore là một cấu trúc, vừa dùng để đồng bộ tài nguyên, vừa dùng để đồng bộ hoạt động.

Semaphore gồm 2 thành phần chính: biến `count` và hàng đợi `wait_list`. 
- Biến ***count*** giúp kiểm soát số lượng thread còn lại được phép truy cập vào critical resource. 
- Còn hàng đợi ***wait_list*** chứa danh sách các thread đang phải chờ đợi trước khi có thể truy cập critical resource.

Linux kernel sử dụng cấu trúc semaphore để biểu diễn một semaphore.
```C
struct semaphore {
    /*
     * Do cấu trúc semaphore cũng bị nhiều thread truy cập đồng thời,
     * nên semaphore cũng được xem là một critical resource. Biến @lock là
     * một spinlock bảo vệ @count và @wait_list trong cấu trúc semaphore.
     */
    raw_spinlock_t		lock;

    /*
     * Biến count vừa thể hiện trạng thái của semaphore, vừa thể hiện
     * trạng thái của critical resource.
     *   > 0: semaphore đang ở trạng thái AVAILABLE,
     *        còn critical resource đang ở trạng thái READY.
     *        count cũng thể hiện còn bao nhiêu thread nữa được phép
     *        sử dụng critical resource.
     *   = 0: semaphore đang ở trạng thái UNAVAILABLE,
     *        còn critical resource đang ở trạng thái BUSY.
     */
    unsigned int		count;

    // wait_list là danh sách các thread đang chờ đợi để có được semaphore
    struct list_head	wait_list;
};

```

### 2. Lock
Trong cơ chế `Lock` có thể phân làm `spin lock` và `mutex lock`

`Lock` cũng có thể phân làm: Optimistic locking và Pessimistic locking
- Optimistic locking: dữ liệu có thể được ***access*** bởi nhiều process, tuy nhiên nếu các process đồng thời ***cập nhật*** dữ liệu thì sẽ xảy ra conflict, chỉ một process thành công và các process khác sẽ không thực hiện được.
- Pessimistic locking: chỉ process đầu tiên truy cập đến đỗi tượng sẽ có thể cập nhật nó. Các process khác sẽ không được cập nhật và thậm chí là ***KHÔNG thể đọc*** dữ liệu.


### CAS theory
CAS stands for ***Compare And Swap***, which is a well-known lock-free algorithm.

`CAS theory is a kind of spinlock.`

#### bài toán Mutual Exclusion
Khi lập trình, ta cần đảm bảo ***chỉ duy nhất 1 thread*** được truy cập vào ***shared resource*** tại một thời điểm. Từ đó sẽ giải quyết được ***data race***.
- Đoạn code dùng để ***read/write*** [[shared resource]] gọi là ***critical section***. 

***mutual exclusion*** - nó giúp chúng ta implement [[critical section]]. Dịch ra là ***loại trừ lẫn nhau***: Nó mang ý nghĩa nếu A thì thôi B, nếu B thì thôi A, chỉ 1 và duy nhất 1, ai nhanh hơn thì được.

***mutex***: viết tắt của ***Mut***ual ***ex***cluision.

***mutual exclusion*** hay còn gọi là ***Mutex*** có các đặc điểm:
- Là cách để xử lý data race.
- Chỉ có duy nhất một thread được sở hữu key tại một thời điểm.
- Giới hạn quyền truy cập đến critical section.


#### Hoạt động của CAS

CAS hoạt động dựa trên toán hạng (`operands`):
```
- Vị trí của bộ nhớ sẽ hoạt động (M).
- Giá trị của biến hiện tại (A).
- Giá trị mới của biến sau khi thay đổi (B).
```

Các bước thực hiện CAS:
```
- Kiểm tra giá trị hiện tại của biến tại (M).
- So sánh giá trị đó với (A):
- Nếu giá trị đó khác (A), không làm gì, kết thúc CAS.
- Nếu giá trị là (A), thực hiện việc thay đổi vùng nhớ sang giá trị (B).
```

Khác biệt sẽ nằm ở đây, khi nhiều thread cùng thực hiện ***CAS*** trên một biến, chỉ có duy nhất 1 thread truy cập thành công và thay đổi giá trị. Các thread còn lại không bị block, chúng vẫn thực hiện CAS nhưng không có gì thay đổi vì giá trị mới đã ***khác*** (A).

Do đó, các thread không bị block và không xảy ra context switch. 

Có một thuật ngữ khác để diễn tả cơ chế lock này là ***lock-free***.


## Ref
- [1: Deep Understanding of Golang Mutex](https://levelup.gitconnected.com/deep-understanding-of-golang-mutex-9964b02c56e9)
- [2: Mutex Examples in Go](https://levelup.gitconnected.com/mutex-examples-in-go-ad7c440461a4)
- [3: Tìm hiểu về khái niệm Mechanical Sympathy trong phần mềm](https://batnamv.medium.com/t%C3%ACm-hi%E1%BB%83u-v%E1%BB%81-kh%C3%A1i-ni%E1%BB%87m-mechanical-sympathy-v%C3%A0-b%E1%BB%99-th%C6%B0-vi%E1%BB%87n-lmax-disruptor-4d553dc7fa55)
- [4: Go Concurrency dành cho Java Developers: medium](https://batnamv.medium.com/go-concurrency-d%C3%A0nh-cho-java-developers-c7709f1f8752)
- [5: Go Concurrency dành cho Java Developers: techmaster](https://techmaster.vn/posts/36313/go-concurrency-danh-cho-java-developers)
