
### 1. **Why `package main` and not `package forum`?**

In Go, the entry point of an executable program **must be in `package main`** and contain a `func main()`.

* **`package main`** signals this is an **executable program**.
* Other package names (like `forum`) are usually for **libraries or reusable code**, not for running directly.

If you name your package `forum` but try to run it, Go won’t find a proper entry point.

---

### 2. **What does it mean that `ListenAndServe` is blocking?**

“Blocking” means when you call:

```go
http.ListenAndServe(":8080", nil)
```

your program **waits right here indefinitely**, running the web server **and doesn’t continue to the next lines** until:

* The server stops (e.g., error or shutdown)
* The program crashes

Think of it like calling a function that “goes into an infinite loop” to keep accepting web requests. So any code after `ListenAndServe` won’t run *unless* the server ends.

---

### 3. **What does `log.Fatal(err)` do?**

* `log.Fatal(err)` does **two things**:

  1. Prints the error message to the console.
  2. Immediately **exits the program with a non-zero status** (indicating failure).

Why use it here?
If `ListenAndServe` fails (e.g., port 8080 is already in use), the error is returned — `log.Fatal` will print the error and stop the program, instead of silently ignoring it.

---

### 4. **What *exactly* is `ListenAndServe` doing?**

It’s a convenience function in Go’s `net/http` package that:

* Creates an HTTP server listening on the specified port (`:8080` here).
* Waits for incoming HTTP requests.
* For each request, dispatches it to the registered handlers (routes).
* Runs **forever** unless an error occurs.

Essentially, it’s your **web server’s main loop**. When you call it, you’re saying:

> “Start listening on port 8080, and handle HTTP requests using these routes.”

---

### To summarize:

| Concept           | What it means                              | Why it matters                             |
| ----------------- | ------------------------------------------ | ------------------------------------------ |
| `package main`    | Defines the executable’s entry point       | Go runs `main()` only in this package      |
| `ListenAndServe`  | Starts the web server, waits for requests  | Blocks execution until stopped/error       |
| Blocking function | Stops further code from running until done | Log after it won’t show unless server ends |
| `log.Fatal(err)`  | Logs error and exits immediately           | Good for fatal server startup errors       |

1) what is an http server
2) what is an http request
3) what is ListenAndServe