# 07. CQRS (Command Query Responsibility Segregation)

**CQRS** is an architectural pattern that separates the models for updating data (Commands) from the models for reading data (Queries). 

## 1. The Core Concept

*   **Command:** An operation that changes the state of the system (Create, Update, Delete). Commands should not return data (except perhaps a confirmation ID).
*   **Query:** An operation that reads the state of the system. Queries should be side-effect-free (they do not change any data).

```mermaid
graph TD
    subgraph Client [User Interface]
        UI[Web / Mobile]
    end

    subgraph CQRS_Logic [Application Layer]
        direction TB
        subgraph Commands [Command Side - WRITE]
            CH[Command Handlers]
            WM[Write Model / Domain]
        end
        
        subgraph Queries [Query Side - READ]
            QH[Query Handlers]
            RM[Read Model / DTOs]
        end
    end

    subgraph Storage [Data Layer]
        DB[(Database)]
    end

    %% Flows
    UI -->|1. Execute Command| CH
    CH --> WM
    WM -->|Update| DB

    UI -->|2. Request Data| QH
    QH --> RM
    DB -->|Fetch| RM
    RM -->|Return Result| UI

    %% Dark Mode Styling
    style CQRS_Logic fill:#1a237e,stroke:#5c6bc0,stroke-width:2px,color:#e0e0e0
    style Commands fill:#2d1b4d,stroke:#9d70ff,color:#fff
    style Queries fill:#004d40,stroke:#26a69a,color:#fff
    style Storage fill:#212121,stroke:#757575,color:#fff
```

## 2. Why Use CQRS?

### ✅ Advantages
*   **Independent Scaling:** You can scale the Read side (which is usually 90% of traffic) independently from the Write side.
*   **Optimized Data Models:** The Read model can be highly optimized for UI display (denormalized), while the Write model is optimized for business logic (normalized).
*   **Security:** Easier to manage permissions (who can change data vs. who can see it).

### ❌ Disadvantages
*   **Complexity:** You have to manage two separate models and potentially two different databases.
*   **Eventual Consistency:** If you use separate databases for Read and Write, there will be a delay (milliseconds) before the Read side sees the update.

---

## 3. Levels of CQRS
1.  **Code-level CQRS:** Separate classes/functions for Commands and Queries but sharing the same DB.
2.  **Database-level CQRS:** Separate databases (e.g., SQL for Write, Elasticsearch for Read).

---