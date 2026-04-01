# 06. Hexagonal Architecture (Ports & Adapters)

**Hexagonal Architecture**, created by Alistair Cockburn, aims to create loosely coupled application components that can be easily connected to their software environment by means of ports and adapters.

## 1. The Core Philosophy: Inside vs. Outside

*   **The Inside (The Core):** Contains the Business Logic and Domain Models. It is agnostic to any external technology (Database, Framework, UI).
*   **The Outside:** Contains the Database, UI, External APIs, Frameworks, etc. These are considered "details" that can be swapped.

## 2. Ports and Adapters

1.  **Ports (Interfaces):** The "sockets" through which the Core communicates with the outside world.
    *   *Driving Ports (Primary):* Interfaces for the outside to call into the Core (e.g., Service interfaces).
    *   *Driven Ports (Secondary):* Interfaces for the Core to call the outside (e.g., Repository interfaces).
2.  **Adapters (Implementations):** The "plugs" that fit into the ports.
    *   *Input Adapters (Driving):* HTTP Handlers, CLI, Test Suites.
    *   *Output Adapters (Driven):* SQL Database, In-memory Cache, External Mail Service.

### 3. 🏗️ Hexagonal Architecture Diagram (Dark Mode)

```mermaid
graph LR
    subgraph Outside_Input [Adapters: Driving / Input]
        UI[Web UI / Mobile]
        CLI[Command Line]
        TEST[Test Suite]
    end

    subgraph Inside [The Core: Business Logic]
        direction TB
        subgraph Ports_In [Driving Ports]
            IP[UseCase / Service Interface]
        end
        
        subgraph Domain [Domain Models & Services]
            DS[Domain Services]
            DM[Domain Models]
        end
        
        subgraph Ports_Out [Driven Ports]
            OP[Repository / Gateway Interface]
        end
    end

    subgraph Outside_Output [Adapters: Driven / Output]
        DB[(Database: SQL / NoSQL)]
        MAIL[Email Service]
        SMS[SMS Gateway]
    end

    %% Connections
    UI & CLI & TEST --> IP
    IP --> DS
    DS --> DM
    DS --> OP
    OP --> DB & MAIL & SMS

    %% Dark Mode Styling
    style Inside fill:#1a237e,stroke:#5c6bc0,stroke-width:2px,color:#e0e0e0
    style Outside_Input fill:#212121,stroke:#757575,stroke-width:1px,color:#e0e0e0
    style Outside_Output fill:#212121,stroke:#757575,stroke-width:1px,color:#e0e0e0
    
    style Domain fill:#2d1b4d,stroke:#9d70ff,color:#fff
    style IP fill:#004d40,stroke:#26a69a,color:#fff
    style OP fill:#004d40,stroke:#26a69a,color:#fff
```
---

## 4. Why Hexagonal Architecture?

### ✅ Advantages
*   **Technology Agnostic:** You can change your database or web framework without touching a single line of business logic.
*   **High Testability:** You can test the Core by plugging in "Mock Adapters" instead of real databases or APIs.
*   **Maintainability:** Clear boundaries prevent infrastructure code from leaking into business logic.

---