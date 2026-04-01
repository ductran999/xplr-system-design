### 04. Microkernel Architecture

**The "Plug-in" Strategy for Systems**

---

#### Core Concept

- **The Kernel (Core):** Minimal logic to run the system. Manages lifecycles and registration.
- **The Plug-ins:** Independent modules for specific features (Stripe, PayPal, etc.).
- **The Contract:** A strict **Interface** that plugins must follow to be "pluggable."

```mermaid
graph TD
    subgraph "Core System (The Kernel)"
        Registry[Plugin Registry]
        Lifecycle[Lifecycle Manager]
        Engine[Execution Engine]
    end

    subgraph "Plug-in Modules"
        P1[Stripe Payment Plugin]
        P2[PayPal Payment Plugin]
        P3[Crypto Payment Plugin]
    end

    %% Connections
    P1 -.->|Registers| Registry
    P2 -.->|Registers| Registry
    P3 -.->|Registers| Registry
    
    Engine -->|Calls via Interface| P1
    Engine -->|Calls via Interface| P2
    Engine -->|Calls via Interface| P3
```

---

#### Design Pattern Relation

- **Macro-level:** It is essentially the **Strategy** and **Factory** patterns scaled up to govern the **entire application structure**.

---

#### Key Features

- ✅ **Extensible:** Add features without touching or recompiling the Core.
- ✅ **Isolated:** If a plugin crashes, the Core ideally stays alive.
- ✅ **Custom:** Users pick and choose which "extensions" to install.
- ❌ **Complex:** Extremely hard to design a "perfect" Interface that never needs to change.

---

#### Real-world Examples

- **IDEs:** VS Code, IntelliJ, Eclipse.
- **Browsers:** Chrome/Firefox Extensions.
- **OS:** Linux/Windows Kernels (Drivers).

---
