Go không có sẵn đầy đủ các cấu trúc dữ liệu như trong Java/C++ (List, Set, Stack...), nhưng có nhiều công cụ built-in (array, slice, map, channel) để xây dựng các cấu trúc đó một cách hiệu quả.


| Loại cấu trúc                 | Mục đích chính                              | Go hỗ trợ trực tiếp? | Gợi ý cài đặt                             |
| ----------------------------- | ------------------------------------------- | -------------------- | ----------------------------------------- |
| **Array**                     | Lưu danh sách phần tử có kích thước cố định | ✅ Có sẵn             | `[N]T`                                    |
| **Slice**                     | Danh sách động, mở rộng được                | ✅ Có sẵn             | `[]T`                                     |
| **Map (HashMap)**             | Ánh xạ key → value                          | ✅ Có sẵn             | `map[K]V`                                 |
| **Struct**                    | Gom nhóm các trường dữ liệu khác nhau       | ✅ Có sẵn             | `struct {}`                               |
| **Queue**                     | Cấu trúc hàng đợi (FIFO)                    | ❌ Không sẵn          | Tự cài bằng `slice`                       |
| **Stack**                     | Cấu trúc ngăn xếp (LIFO)                    | ❌ Không sẵn          | Tự cài bằng `slice`                       |
| **Set**                       | Tập hợp các phần tử duy nhất                | ❌ Không sẵn          | Dùng `map[T]struct{}`                     |
| **LinkedList / Tree / Graph** | Cấu trúc phức tạp                           | ❌ Không sẵn          | Cài thủ công hoặc dùng thư viện bên ngoài |

# Ref 
- https://github.com/emirpasic/gods
- https://github.com/deckarep/golang-set
- https://github.com/idsulik/go-collections
- https://github.com/Workiva/go-datastructures
- https://groups.google.com/g/golang-nuts/c/ROeTHEbwUgw?pli=1
- https://blog.stackademic.com/why-go-has-only-3-data-structures-and-why-thats-a-good-thing-850b8e59ef70
- https://medium.com/@rodrigobme/why-go-has-only-3-data-structures-4fa607fb847b
