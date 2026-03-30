# MVP Pattern (Model-View-Presenter)

**MVP** is a derivation of the MVC pattern. It is commonly used to separate data modeling (Model) and user interface (View) by introducing a third component called the **Presenter**.

## 1. The Three Components

1.  **Model:** Represents the data and business logic.
2.  **View:** A "Passive View." It only displays data and sends user actions to the Presenter. It **does not know** about the Model. It usually communicates with the Presenter through an **Interface**.
3.  **Presenter:** The "Heart" of the pattern. It acts as the intermediary.
    - It receives user actions from the View.
    - It fetches data from the Model.
    - It formats the data and **manually updates** the View.

## 2. Key Differences from MVC and MVVM

- **vs. MVC:** In MVC, the Controller chooses the View. In MVP, the Presenter is **coupled** to a specific View (usually 1 Presenter per 1 View).
- **vs. MVVM:** In MVVM, data updates are **automatic** (Data Binding). In MVP, the Presenter must **explicitly** call methods on the View (e.g., `view.showUser(name)`).

---

## 3. Why Use MVP?

- **Total Isolation:** View and Model are 100% separated.
- **Testability:** Since the Presenter communicates with the View via an Interface, you can easily "mock" the View to unit test all the UI logic.
- **Clean Code:** Business logic is moved out of the UI components (Activities, Forms, etc.).

---
