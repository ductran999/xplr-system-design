# 03. Pipeline Architecture (Pipe and Filter)

**Pipeline Architecture** decomposes a complex task into a series of smaller, independent processing steps called **Filters**, connected by **Pipes**.

## 1. Core Concepts

- **Filter:** A self-contained processing unit. It receives input, transforms it, and produces output. It doesn't know about other filters.
- **Pipe:** The communication channel that passes data from one filter to the next. In Go, we use **Channels**.
- **Source:** The start of the pipeline (input data).
- **Sink:** The end of the pipeline (final result).

## 2. Why Use Pipeline Architecture?

### ✅ Advantages

- **Modularity:** Each filter is easy to write, test, and maintain.
- **Reusability:** You can rearrange or swap filters like Lego blocks.
- **Parallelism:** Each stage can run on a different CPU core (Goroutines), making data processing extremely fast.
- **Separation of Concerns:** Each step only cares about its specific transformation logic.

### ❌ Disadvantages

- **Data Overhead:** Passing data between stages can add slight overhead.
- **Complexity:** Error handling across multiple stages (channels) can be tricky.
- **Fixed Flow:** Harder to implement complex conditional branching (if/else) inside the pipe.

---

## 3. Real-world Examples

- **Compilers:** Code -> Lexer -> Parser -> Optimizer -> Machine Code.
- **Image/Video Processing:** Raw Video -> Resize -> Filter -> Watermark -> Compress.
- **ETL Systems:** Extract -> Transform -> Load (Big Data).
- **UNIX Shell:** `cat file.txt | grep "error" | sort | uniq`.
