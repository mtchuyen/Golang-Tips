# Interface Segregation Pattern

Dave Cheney talks: **Small interfaces are beautiful**

---

## 1. ISP là gì?

**Interface Segregation Principle** (trong SOLID):

> “Client không nên phụ thuộc vào những interface mà nó không sử dụng.”

👉 Hiểu đơn giản:

* ❌ Không tạo interface “to, nhiều chức năng”
* ✅ Tách thành nhiều interface nhỏ, đúng nhu cầu từng client

**Interface Segregation Pattern (ISP) (Nguyên lý Phân tách Interface)** là một nguyên lý trong bộ **SOLID** (5 nguyên lý thiết kế phần mềm hướng đối tượng), nói rằng: c**lient (code sử dụng interface) không nên bị ép phụ thuộc vào những method nó không dùng**. 

Nói đơn giản: **đừng tạo interface quá to**, hãy chia thành nhiều interface nhỏ, đúng nhu cầu từng nơi dùng. 

Trong Golang, nguyên lý này đặc biệt quan trọng vì Go xem interface là **hành vi (behavior)** chứ không phải “hệ phân cấp kiểu dữ liệu” như nhiều ngôn ngữ khác. Làm tốt ISP giúp code dễ test, ít coupling (phụ thuộc chặt), dễ thay đổi, và tránh những “God Interface” (interface khổng lồ làm mọi thứ).

### Ví dụ: thuê nhân viên văn phòng.

Một công ty viết hợp đồng thuê nhân viên:
- Biết pha cà phê
- Biết kế toán
- Biết lập trình
- Biết bảo vệ
- Biết sửa điều hòa

Mọi nhân viên đều phải làm tất cả.

**Kết quả? Kế toán phải học sửa điều hòa.**

**Đó là vi phạm ISP.**

Công ty tốt hơn chia vai trò:
- Accountant
- Programmer
- Security

Ai làm việc nào ký đúng hợp đồng đó.

Interface cũng vậy. **Đừng bắt cá phải leo cây.**

### Ví dụ

#### Anti-pattern (God Interface)
```
                +---------------------+
Client A ------>| Big Interface        |
Client B ------>| Read                 |
Client C ------>| Write                |
                | Delete               |
                | Backup               |
                | Monitor              |
                +---------------------+

```

Clients bị ép phụ thuộc method không dùng

#### ISP đúng

```
         +---------+
ClientA->| Reader  |
         +---------+

         +---------+
ClientB->| Writer  |
         +---------+

         +----------------+
ClientC->| Reader + Writer |
         +----------------+
```

Dependency đúng nhu cầu.

---

## 2. Vấn đề nếu KHÔNG dùng ISP

**Ví dụ sai (anti-pattern)**:

```go
type Worker interface {
    Work()
    Eat()
    Sleep()
}
```

Giờ bạn có:

```go
type Robot struct{}

func (r Robot) Work() {}
func (r Robot) Eat() {
    panic("robot does not eat")
}
func (r Robot) Sleep() {
    panic("robot does not sleep")
}
```

**👉 Vấn đề:**

* Robot bị **ép implement những thứ không cần**
* Code smell: `panic`, `empty function`
* Vi phạm ISP

---

## 3. Cách đúng: Segregate Interface

**Quy tắc "Need, not Future"**
- Thiết kế interface cho nhu cầu hiện tại, không cho tưởng tượng tương lai.

**Khi nào thì tác nhỏ:**
- Có client nào dùng method này chưa?
- Nếu bỏ method này interface có gọn hơn không?

Nếu có → tách.

**👉 Tách nhỏ:**

```go
type Worker interface {
    Work()
}

type Eater interface {
    Eat()
}

type Sleeper interface {
    Sleep()
}
```

**Áp dụng:**

```go
type Human struct{}

func (h Human) Work()  {}
func (h Human) Eat()   {}
func (h Human) Sleep() {}

type Robot struct{}

func (r Robot) Work() {}
```

**👉 Kết quả:**

* Human dùng 3 interface
* Robot chỉ cần `Worker`
* Không còn “function thừa”

---

## 4. Tư duy cốt lõi

Bạn cần nhớ 3 ý này:

### 1. Interface = contract cho CLIENT, không phải cho IMPLEMENTATION

* Đừng nghĩ: “object này làm được gì?”
* Hãy nghĩ: “client cần dùng gì?”



### 2. Interface càng nhỏ càng tốt (idiomatic Go)

**Interface nhỏ, tập trung một vai trò**
- Mỗi interface chỉ nên mô tả một nhóm hành vi nhỏ, có mục đích rõ ràng. Nếu interface có quá nhiều method, thường nó đang gánh nhiều trách nhiệm.

Go có pattern nổi tiếng:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

**👉 Chỉ 1 method nhưng cực kỳ mạnh**



### 🔹 3. Composition > God Interface

Thay vì:

```go
type FileManager interface {
    Read()
    Write()
    Delete()
    Compress()
}
```

👉 Hãy:

```go
type Reader interface { Read() }
type Writer interface { Write() }
type Deleter interface { Delete() }
type Compressor interface { Compress() }
```

---

## 5. Case thực tế

### 🎯 Scenario: hệ thống upload file

**❌ Sai (thiết kế tệ)**:

```go
type Storage interface {
    Upload(file []byte)
    Download(id string)
    Delete(id string)
}
```

**👉 Vấn đề:**

* Có service chỉ cần upload
* Nhưng vẫn phải implement download + delete


**Đúng (áp dụng ISP)**

```go
type Uploader interface {
    Upload(file []byte)
}

type Downloader interface {
    Download(id string)
}

type Deleter interface {
    Delete(id string)
}
```

---

**💡 Service chỉ cần Upload**

```go
type UploadService struct {
    uploader Uploader
}
```

👉 Không phụ thuộc vào Delete/Download

---

## 6. Pattern quan trọng trong Go

### ⭐ Accept Interface, Return Struct

```go
func Process(r io.Reader) {
    // dùng Reader thôi
}
```

👉 Đây chính là ISP trong thực tế:

* Function chỉ nhận đúng thứ nó cần

---

## 7. Khi nào nên áp dụng ISP?

### ✅ Nên dùng khi:

* Interface có > 3 method
* Có implement bị “ép dùng method không cần”
* Có nhiều loại client khác nhau

---

### ❌ Không cần dùng khi:

* Interface nhỏ sẵn rồi
* Code đơn giản (over-engineering là sai)


### Lỗi thường gặp: Over-segregation (tách quá mức)

**Sai**: Khi logic quá đơn giản
```
type Opener interface { Open() }
type Closer interface { Close() }
type Flusher interface { Flush() }
```
---

## 8. Checklist để review code

Khi nhìn interface, hỏi:

* Có method nào không dùng không?
* Có implement nào bị `panic` hoặc `empty`?
* Có thể tách nhỏ không?
* Client thực sự cần gì?

---

## 9. Bài tập áp dụng (24–72h)

### Bài 1 (cực quan trọng)

Refactor code này:

```go
type Notification interface {
    SendEmail(msg string)
    SendSMS(msg string)
    SendPush(msg string)
}
```

👉 Yêu cầu:

* Tách theo ISP
* Viết 2 implementation:

  * EmailService
  * SMSService

---

### Bài 2 (gần thực tế backend)

Thiết kế interface cho:

* Payment Service (VNPay, Momo, Stripe)

👉 Gợi ý:

* Không tạo 1 interface “PaymentGod”
* Tách theo hành vi:

  * Pay
  * Refund
  * QueryStatus

---

## 10. Tóm tắt

**Interface Segregation** nói rằng **đừng ép code phụ thuộc vào method nó không cần**. Trong Go, điều này mạnh hơn vì **interface mô tả hành vi** chứ không phải kiểu dữ liệu. Mẫu thiết kế đúng là: interface nhỏ, định nghĩa ở nơi sử dụng, dễ mock, ghép lại bằng composition khi cần. Dấu hiệu vi phạm thường là “God Interface”, mock phức tạp, hoặc implement phải viết method rỗng. Hãy nhớ quy tắc **Need, not Future** (thiết kế cho nhu cầu thật, không cho giả định).

👉 1 câu: **Chỉ expose đúng những gì client cần — không hơn.**

👉 3 dấu hiệu sai:

* Interface quá to
* Implement phải fake method
* Client bị phụ thuộc dư thừa

👉 3 cách fix:

* Tách nhỏ interface
* Dùng composition
* Inject đúng interface

