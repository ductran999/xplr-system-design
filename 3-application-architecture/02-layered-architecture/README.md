# 02. Layered Architecture (N-Tier Architecture)

**Layered Architecture** is the most common software architecture pattern. It organizes code into horizontal layers, each staying within its own responsibility (Separation of Concerns).

## 1. The Standard Layers (The 3-Tier Pattern)

1.  **Presentation Layer (Controller/Handler):** Responsible for handling external requests (HTTP, gRPC, CLI). It validates input and calls the Service layer.
2.  **Business Logic Layer (Service/Domain):** The "brain" of the application. It contains all business rules and orchestrates the flow. It should be independent of HTTP or Database details.
3.  **Data Access Layer (Repository/Persistence):** Responsible for communicating with the Database (SQL, NoSQL, File).

---

## 2. Strict vs. Relaxed Layering

- **Strict Layering:** A layer can only call the layer directly below it. (Controller -> Service -> Repository).
- **Relaxed Layering:** A layer can call any layer below it. (Controller -> Repository).
- _Best Practice:_ Always aim for **Strict Layering** to ensure high maintainability.

---

## 3. The "Pain" of Traditional Layered Architecture

In traditional implementations, layers are **tightly coupled**.

- Example: `UserService` has a hard dependency on `MySQLRepository`.
- **Problem:** If you want to change the database or write Unit Tests, you have to rewrite the Service code.

**Solution:** Use **Dependency Inversion (DI)** and **Interfaces**. This is the gateway to **Clean Architecture**.

---
