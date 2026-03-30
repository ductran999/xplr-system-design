# MVVM Pattern (Model-View-ViewModel)

**MVVM** is a modern UI architectural pattern that leverages **Data Binding** to separate the development of the graphical user interface from the business logic or back-end logic.

## 1. The Three Components

1.  **Model:** Represents the data and business logic (Entities, Repositories).
2.  **View:** The UI layer (HTML, XML, SwiftUI). It is "passive," meaning it doesn't contain logic but "observes" the ViewModel.
3.  **ViewModel:** The "State of the View." It exposes data and commands that the View can bind to. It acts as a converter, transforming Model data into a format the View can easily display.

## 2. The "Magic" of Data Binding

The most critical part of MVVM is the **Binder**.

- **One-way Binding:** When the ViewModel changes, the View updates automatically.
- **Two-way Binding:** When the user types in a text field (View), the variable in the ViewModel updates instantly, and vice-versa.

---

## 3. Comparison: MVC vs. MVP vs. MVVM

| Feature          | MVC                 | MVP                       | MVVM                        |
| :--------------- | :------------------ | :------------------------ | :-------------------------- |
| **View's Role**  | Knows the Model     | Total isolation (Passive) | **Binds to ViewModel**      |
| **Mediator**     | Controller (Router) | Presenter (Controller)    | **ViewModel (State)**       |
| **Coordination** | Manual              | Manual (via Interface)    | **Automatic (via Binding)** |
| **Testing**      | Hard (UI coupled)   | Great (Interface-based)   | **Excellent (Pure Logic)**  |

## 4. Why Use MVVM?

- **Developer-Designer Workflow:** Designers can work on the View (HTML/CSS) while developers work on the ViewModel without breaking each other's code.
- **Reusability:** A single ViewModel can often be used for multiple Views.
- **Maintainability:** Logic is separated from UI code, making it easier to test and change.

---
