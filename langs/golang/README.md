# What's Go

The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.


## Concurrency with Goroutines
A goroutine is a lightweight, independently executing function managed by the Go runtime, used to achieve concurrency. It runs concurrently with other functions, allowing a program to perform multiple tasks at the same time without the overhead of traditional operating system threads. 
To start a goroutine, you simply place the go keyword before a function call, like **go myFunction()**

### Samples
- [Buffered Channels](./concurrent_programming/buffer_channel_example/)
- [for loops with channels.](./concurrent_programming/for_channels_example/)
- [Mutex with Channels.](./concurrent_programming/mutex_sample/)
- [Multile Producers  with Multiple Consumers example.](./concurrent_programming/ordersapp_multipleprod_multiplecon/)
- [Multiple Producers with Single consumer pattern example.](./concurrent_programming/ordersapp_multipleprod_singlecon/)
- [Single Producer with Multiple consumer pattern example.](./concurrent_programming/ordersapp_singleprod_multiplecon/)
- [Single Producer with Single consumer pattern example.](./concurrent_programming/ordersapp_singleprod_singlecon/) 
