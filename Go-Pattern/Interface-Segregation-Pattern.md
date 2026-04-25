# Interface Segregation Pattern

---

## 1. ISP là gì?

**Interface Segregation Principle** (trong SOLID):

> “Client không nên phụ thuộc vào những interface mà nó không sử dụng.”

👉 Hiểu đơn giản:

* ❌ Không tạo interface “to, nhiều chức năng”
* ✅ Tách thành nhiều interface nhỏ, đúng nhu cầu từng client

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

👉 Tách nhỏ:

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

### Áp dụng:

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

## 4. Tư duy cốt lõi (quan trọng nhất)

Bạn cần nhớ 3 ý này:

### 1. Interface = contract cho CLIENT, không phải cho IMPLEMENTATION

* Đừng nghĩ: “object này làm được gì?”
* Hãy nghĩ: “client cần dùng gì?”



### 2. Interface càng nhỏ càng tốt (idiomatic Go)

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

---

## 8. Checklist để review code

Khi nhìn interface, hỏi:

* Có method nào không dùng không?
* Có implement nào bị `panic` hoặc `empty`?
* Có thể tách nhỏ không?
* Client thực sự cần gì?

---

# 9. Bài tập áp dụng (24–72h)

## 🎯 Bài 1 (cực quan trọng)

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

## 🎯 Bài 2 (gần thực tế backend)

Thiết kế interface cho:

* Payment Service (VNPay, Momo, Stripe)

👉 Gợi ý:

* Không tạo 1 interface “PaymentGod”
* Tách theo hành vi:

  * Pay
  * Refund
  * QueryStatus

---

# 10. Tóm tắt để nhớ nhanh

👉 1 câu:

> “Chỉ expose đúng những gì client cần — không hơn.”

👉 3 dấu hiệu sai:

* Interface quá to
* Implement phải fake method
* Client bị phụ thuộc dư thừa

👉 3 cách fix:

* Tách nhỏ interface
* Dùng composition
* Inject đúng interface

