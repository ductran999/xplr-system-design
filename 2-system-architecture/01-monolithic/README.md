# 01. Monolithic Architecture

Monolithic is the most traditional software architecture. All application functionalities (User Management, Product Management, Payment, Email Sending, etc.) are packed into **ONE single codebase, running in ONE single process, and typically sharing ONE database.**

## 1. Communication in a Monolith

- You don’t need REST APIs, gRPC, or WebSocket for modules to communicate.
- **How it works:** Direct function calls in memory (Function Call / Method Invocation).
- **Example:** The `Checkout()` function directly calls `DeductInventory()` and `SendEmail()`. The speed is virtually instant (0 milliseconds) because there is no network boundary involved.

## 2. Advantages (Why do 99% of startups begin with a Monolith?)

- **Easy to develop:** Open one project and you have the entire codebase. Debugging and navigating between files is very straightforward.
- **Easy to test:** Run a single command to test the entire flow.
- **Easy to deploy:** Just ship one `.exe` file (Go) or a code directory (Node/PHP) to the server. No need for Docker or Kubernetes complexity.
- **High performance (with low traffic):** No network latency between modules.

## 3. Disadvantages (As the company grows, Monolith becomes a burden)

Imagine a system like Shopee built as a monolith:

- **Hard to scale (no fine-grained scaling):**  
  During big sales (e.g., 11.11), the Payment module gets overloaded. You want to scale only Payment? You can’t. You must replicate the **ENTIRE** monolith across multiple servers, wasting resources on unnecessary modules (like Avatar updates).

- **Single point of failure:**  
  A new developer accidentally introduces a memory leak in the Email module. BOOM! The entire system (browsing, checkout, everything) crashes because all modules run in the same process.

- **Team collaboration conflicts:**  
  100 developers working on the same repository. Merging code and resolving conflicts becomes a nightmare. When bugs appear in production, it’s hard to identify which team caused them.

- **Technology lock-in:**  
  The project is written in Java. Now the AI team wants to use Python for a recommendation feature? Not possible, since everything must use the same tech stack.

## 4. High-Level Design: Monolithic Architecture

```text
            +----------------------+
            |   Client / Browser   |
            +----------+-----------+
                       |
                       v
            +----------------------+
            |   Monolithic App     |
            |   (Single Process)   |
            +----------+-----------+
                       |
        ---------------------------------------
        |          |          |              |
        v          v          v              v
   [User Module] [Order] [Inventory]   [Payment]
                   |          |              |
                   v          v              v
               [Notification / Email Module]
                       |
                       v
                +--------------+
                |   Database   |
                | (Shared DB)  |
                +--------------+

            ---> Response returned to Client
```