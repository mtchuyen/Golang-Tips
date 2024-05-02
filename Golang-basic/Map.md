# Map
***Map*** có các loại:
- ...
- TreeMap
- HashMap
- EnumMap
- SynchronizedMap

### HashMap
Dùng giải thích từ ngôn ngữ Java trong link: https://www.geeksforgeeks.org/difference-hashmap-concurrenthashmap/

[***HashMap***](https://www.geeksforgeeks.org/java-util-hashmap-in-java-with-examples/):
> It provides the basic implementation of the Map interface

HashMap in Java stores the data in (Key, Value) pairs, and you can access them by an index of another type (e.g. an Integer).
> One object is used as a key (index) to another object (value). If you try to insert the duplicate key in HashMap, it will REPLACE the element of the corresponding key.
>
> it is unsynchronized.

***Một số loại HashMap***
- IdentityHashMap
- LinkedHashMap
- WeakHashMap


#### HashMap vs HashSet

#### HashMap vs ConcurrentHashMap
***HashMap*** is NOT Thread-safe whereas ***ConcurrentHashMap*** is Thread-safe in nature.
> HashMap performance is relatively high because it is non-synchronized in nature and any number of threads can perform simultaneously.
> But ConcurrentHashMap performance is low sometimes because sometimes Threads are required to wait on ConcurrentHashMap.


> In HashMap, NULL values are allowed for key and values, whereas in ConcurrentHashMap null value is NOT allowed for key and value,

#### ConcurrentHashMap vs SynchronizedHashMap
ConcurrentHashMap and SynchronizedHashMap both are thread-safe and can be used in multithreaded and concurrent

[Link](https://www.geeksforgeeks.org/difference-between-concurrenthashmap-and-synchronizedhashmap/)
