# Learn Golang by Building a Redis Clone

Redis-compatible database engine.

---

## Overview

Practicing Go engineering principles, including encapsulation, data manipulation, thread safety, and concurrency primitives like channels and contexts.

---

## Key Concepts Covered

* **Custom Data Structures:** Wrapping Go primitives into thread-safe, domain-specific stores.
* **Encapsulation:** Hiding internal map states while exposing clean CRUD APIs (`Get`, `Set`, `Delete`).
* **Key Sorting:** Implementing custom key extraction and alphabetical ordering.
* **Concurrency & Synchronization:** Safely handling concurrent reads/writes and background processes.
* **Context & Deadlines:** Implementing timeouts and cancellation using Go's `context` package.

---

## Repository Structure

```text
.
├── store/          # Key-value data storage logic & operations
├── server/         # TCP/HTTP protocol handling and server loop
├── concurrency/    # Context management, routines, and channel handlers
└── main.go         # Core entry point
```

---

## Quick Start

### Prerequisites
* **Go** version `1.18` or higher installed.

### Installation & Execution

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/go-redis.git
   cd go-redis
   ```

2. Run the application:
   ```bash
   go run main.go
   ```
