# File I/O vs Network I/O in Goroutines

## Why They Behave Differently

The key difference is that **Network I/O uses Go's Netpoller**, while **File I/O typically does not**.

---

## File I/O

Examples:

```go
file.Read(...)
file.Write(...)
os.WriteFile(...)
```

### What Happens?

1. Goroutine `G` performs a file syscall.
2. OS thread `M` enters the kernel.
3. If the syscall blocks, `M` is blocked.
4. Go runtime detaches `P` from the blocked `M`.
5. Another `M` receives the `P` and continues running goroutines.
6. When the syscall completes, the kernel wakes the blocked `M`.

```text
G1 -> File Write
      |
      v
M1 blocked in kernel
P1 detached
P1 -> M2
```

**Result:** The OS thread (`M`) is blocked.

---

## Network I/O

Examples:

```go
conn.Read(...)
conn.Write(...)
http.Get(...)
```

### What Happens?

1. Socket is configured as non-blocking.
2. If data is unavailable, the syscall returns immediately.
3. Goroutine `G` is parked.
4. `M` and `P` continue running other goroutines.
5. The socket is registered with the Netpoller.
6. When the socket becomes ready, the Netpoller wakes the goroutine.

```text
G1 -> conn.Read()
      |
      v
No data available

G1 parked
M1 + P1 continue working

Socket becomes ready
      |
      v
Netpoller wakes G1
```

**Result:** Only the goroutine is blocked; the OS thread keeps working.

---

## What is Netpoller?

The Netpoller is Go's network event system built on top of:

* Linux → `epoll`
* macOS/BSD → `kqueue`
* Windows → `IOCP`

Instead of blocking an OS thread for every network connection, the runtime asks the OS:

> "Notify me when this socket is ready."

When the socket becomes readable or writable, the Netpoller marks the waiting goroutine as runnable again.

```text
Socket Not Ready
       |
       v
Goroutine Parked
       |
       v
Netpoller Waits
       |
       v
Socket Ready
       |
       v
Goroutine Resumed
```

---

## Summary

| Aspect                      | File I/O     | Network I/O  |
|-----------------------------|--------------|--------------|
| Uses Netpoller              | No           | Yes          |
| Blocks Goroutine            | Yes          | Yes          |
| Blocks OS Thread (M)        | Usually Yes  | Usually No   |
| P Detached from M           | Yes          | No           |
| Scales to Many Connections  | Limited      | Excellent    |

### Mental Model

```text
File I/O
G blocks -> M blocks -> P detached

Network I/O
G blocks -> Netpoller waits
M continues -> P continues
```
