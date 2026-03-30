# VIPER Architecture (View-Interactor-Presenter-Entity-Router)

**VIPER** is a back-port of **Clean Architecture** to iOS/Android development. It is the most modular pattern, dividing each screen (module) into five distinct layers.

## 1. The Five Components

1.  **V (View):** Passive UI. It displays what the Presenter tells it and sends user actions to the Presenter.
2.  **I (Interactor):** Contains the **Business Logic**. It fetches data (API/DB) and performs calculations. It is completely independent of the UI.
3.  **P (Presenter):** The "Coordinator." It receives actions from the View, requests data from the Interactor, applies UI logic (formatting), and tells the Router when to navigate.
4.  **E (Entity):** Simple Data Objects (Models).
5.  **R (Router):** Also called the **Wireframe**. It contains navigation logic (e.g., "Open the Detail Screen"). This is the unique part of VIPER.

---

## 2. Why Use VIPER?

### ✅ Advantages

- **Ultimate Separation of Concerns:** Each file has one and only one responsibility.
- **Perfect for Large Teams:** Multiple developers can work on the same screen (one on logic, one on UI, one on navigation) without merge conflicts.
- **100% Testable:** Every component can be tested in total isolation.

### ❌ Disadvantages

- **Extreme Overkill:** Too many files for small/medium features.
- **Steep Learning Curve:** New developers struggle to follow the data flow.
- **Boilerplate:** Creating a new screen requires creating at least 5 files and 5 interfaces.

---
