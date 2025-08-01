# Go_Assignment

⚙️ Key Constructs & How They Work

make(chan func(), 10)
Creates a buffered channel that can hold up to 10 anonymous functions of type func().
This allows sending up to 10 functions into the channel without blocking the sender.
Use Case: Task queues — workers can pull tasks from the channel and execute them concurrently.
for i := 0; i < 4; i++ { go func() { ... } }
Starts 4 goroutines—lightweight threads managed by Go.
Each goroutine waits in a loop and pulls a function from the channel to execute it.
Use Case: Worker pool — multiple workers handle tasks from a shared source.

🔄 Why 4 Iterations?

This loop is spawning four workers, creating a basic concurrency model where multiple goroutines can process tasks in parallel. It’s a common approach to avoid overwhelming system resources while allowing parallelism. You can adjust the number depending on your workload or CPU cores.

📦 Why a Buffered Channel of Size 10?

The buffer lets you queue up to 10 function tasks without blocking the main goroutine. This way, your program can start putting tasks into the channel even if not all workers are immediately ready. It’s a way to smooth out work distribution and avoid backpressure early on.

❌ Why "HERE1" Might Not Be Printed

Several possibilities could be at play here:
Main goroutine exits early:
fmt.Println("Hello") runs right after the task is sent to the channel.
If the main goroutine ends before any worker pulls and runs f(), "HERE1" may never be printed.
Race condition:
Your workers depend on being scheduled quickly enough to consume the channel before main exits.

✅ How to Fix That

View code_fix.go
