# 06. gRPC (gRPC Remote Procedure Call)

In a Microservices architecture, your system can be split into dozens of services (Order Service, User Service, Payment Service). These services constantly need to call each other to exchange data.

If you use **REST API (JSON)** for inter-service communication: it’s quite SLOW! Because JSON is text-based, the computer must spend time parsing it into objects, plus the overhead of HTTP/1.1.

**gRPC was developed by Google to solve this problem.**

## 1. The Nature of gRPC

- **RPC (Remote Procedure Call):** Call a function that lives on another server as if it were in your local code.
- **HTTP/2:** gRPC requires HTTP/2 (faster, supports multiple parallel requests over a single connection - multiplexing).
- **Protobuf (Protocol Buffers):** Instead of sending data as JSON strings (like REST), gRPC serializes data into **binary format**. This reduces payload size by 3–10x, and machines can read it instantly without heavy parsing.

## 2. How It Works (Code Generation Mechanism)

The most unique aspect of gRPC is that you DON’T manually write communication logic.

1. You define a `.proto` file (function names, input/output) — like a contract.
2. You use a compiler tool provided by Google.
3. The tool automatically generates thousands of lines of code for both Server and Client. You just call the function and you're done! (Even if the Server is written in Go and the Client in Java, the tool generates compatible code).

## 3. Powerful Comparison: REST vs gRPC

| Criteria             | REST API (01)                | gRPC (06)                                         |
| :------------------- | :--------------------------- | :------------------------------------------------ |
| **Data Format**      | JSON (Text - Human-readable) | Protobuf (Binary - Machine-readable)              |
| **Speed / Size**     | Slower / Heavier             | Extremely fast / Lightweight                      |
| **Network Protocol** | HTTP/1.1                     | HTTP/2                                            |
| **Communication**    | One-way (Request-Response)   | Supports bidirectional streaming (like WebSocket) |
| **Best Use Case**    | Frontend → Backend           | **Microservice → Microservice**                   |
