# Chapter 9 Exercises

Exercise 9.1: Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program. The result should indicate whether the transaction succeeded or failed due to insufficient funds. The message sent to the monitor goroutine must contain both the amount to withdraw and a new channel over which the monitor goroutine can send the boolean result back to Withdraw.

Exercise 9.2: Rewrite the PopCount example from Section 2.6.2 so that it initializes the lookup table using sync.Once the first time it is needed. (Realistically, the cost of synchronization would be prohibitive for a small and highly optimized function like PopCount.)

```go
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
    for i := range pc {
        pc[i] = pc[i/2] + byte(i&1)
    }
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
    return int(pc[byte(x>>(0*8))] +
        pc[byte(x>>(1*8))] +
        pc[byte(x>>(2*8))] +
        pc[byte(x>>(3*8))] +
        pc[byte(x>>(4*8))] +
        pc[byte(x>>(5*8))] +
        pc[byte(x>>(6*8))] +
        pc[byte(x>>(7*8))])
}
```

Exercise 9.3: Extend the Func type and the (*Memo).Get method so that callers may provide an optional done channel through which they can cancel the operation (§8.9). The results of a cancelled Func call should not be cached.

Exercise 9.4: Construct a pipeline that connects an arbitrary number of goroutines with channels. What is the maximum number of pipeline stages you can create without running out of memory? How long does a value take to transit the entire pipeline?

Exercise 9.5: Write a program with two goroutines that send messages back and forth over two unbuffered channels in ping-pong fashion. How many communications per second can the program sustain?

Exercise 9.6: Measure how the performance of a compute-bound parallel program (see Exercise 8.5) varies with GOMAXPROCS. What is the optimal value on your computer? How many CPUs does your computer have?

