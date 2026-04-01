# 05. Clean Architecture (The Onion Architecture)

**Clean Architecture** is a software design philosophy that separates the elements of a design into ring-level groups. The main rule is **The Dependency Rule**: source code dependencies can only point **inward**, towards the center of the onion.

## 1. The Four Layers (The Circles)

1.  **Entities (Core):** Business objects of the application. They encapsulate the most general and high-level rules. They are least likely to change when something external changes.
2.  **Use Cases:** Contains application-specific business rules. It orchestrates the flow of data to and from the entities.
3.  **Interface Adapters:** Converts data from the format most convenient for use cases/entities to the format most convenient for external agencies (Database, Web). This is where **Repositories** and **Controllers** live.
4.  **Frameworks & Drivers (External):** The outermost layer, composed of tools like Databases, Web Frameworks, Devices, etc.

---

## 2. The Dependency Rule

- **Inward only:** Code in an inner circle cannot know anything about code in an outer circle.
- **Abstraction:** To talk to an outer layer, the inner layer defines an **Interface**, and the outer layer implements it (Dependency Inversion).

```mermaid
graph BT
    %% Inner Layer: Entities
    subgraph Entities [Layer 1: Enterprise Business Rules]
        E[Entities / Models]
    end

    %% Layer 2: Use Cases
    subgraph UseCases [Layer 2: Application Business Rules]
        UC[Use Cases / Interactors]
    end

    %% Layer 3: Interface Adapters
    subgraph Adapters [Layer 3: Interface Adapters]
        CO[Controllers / Handlers]
        RE[Repositories / Gateways]
        PR[Presenters]
    end

    %% Layer 4: External Tools
    subgraph Infrastructure [Layer 4: Frameworks & Drivers]
        DB[(Database: SQL / NoSQL)]
        WEB[Web Frameworks: Gin / Echo]
        EXT[External APIs / Devices]
    end

    %% Dependency Rule: Inward Only
    Infrastructure -.-> Adapters
    Adapters -.-> UseCases
    UseCases -.-> Entities

    %% Interaction Flow
    WEB --> CO
    CO --> UC
    UC --> E
    UC --> RE
    RE --> DB

    %% Dark Mode Styling
    style Entities fill:#2d1b4d,stroke:#9d70ff,stroke-width:2px,color:#e0e0e0
    style UseCases fill:#1a237e,stroke:#5c6bc0,stroke-width:2px,color:#e0e0e0
    style Adapters fill:#004d40,stroke:#26a69a,stroke-width:2px,color:#e0e0e0
    style Infrastructure fill:#212121,stroke:#757575,stroke-width:2px,color:#e0e0e0
    
    %% Node Styling
    style E fill:#4527a0,color:#fff
    style UC fill:#283593,color:#fff
    style CO fill:#00695c,color:#fff
    style RE fill:#00695c,color:#fff
    style DB fill:#424242,color:#fff
    style WEB fill:#424242,color:#fff
```

---

## 3. Why Use Clean Architecture?

### ✅ Advantages

- **Independent of Frameworks:** You don't have to rely on a specific library.
- **Testable:** Business rules can be tested without the UI, Database, or any other external element.
- **Independent of UI:** The UI can change easily (e.g., from Web to Mobile) without changing the business logic.
- **Independent of Database:** You can swap SQL for MongoDB or any other storage easily.

### ❌ Disadvantages

- **High Complexity:** Requires many layers, interfaces, and data mappers (DTOs).
- **Boilerplate:** Even simple features require a lot of files.
- **Overkill:** Not recommended for simple CRUD applications.

---

## 4. Real-world Example

- Banking systems where core transaction logic must be 100% bug-free and independent of whether the customer uses an ATM, a Web App, or a Mobile App.
