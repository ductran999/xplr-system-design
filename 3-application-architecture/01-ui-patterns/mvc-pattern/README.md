# 01. MVC (Model-View-Controller) Pattern

The **Model-View-Controller (MVC)** pattern is a foundational architectural pattern for building user interfaces. It's especially popular in traditional server-side rendered web applications (e.g., Ruby on Rails, Django, ASP.NET MVC).

## 1. The Three Components

1.  **Model:** Represents the data and business logic of the application. It's the "brain" and the "source of truth."
2.  **View:** The UI layer. Its sole job is to display data provided by the Controller. In classic MVC, the View is often a template (HTML, ERB, Blade).
3.  **Controller:** The "traffic cop." It receives user input, interacts with the Model to update state, and then selects a View to render and pass the data to.

## 2. The Flow of MVC

A typical request flows like this:

1.  User clicks a button, sending a request to the server.
2.  The `Controller` receives the request.
3.  The `Controller` calls the `Model` to fetch or update data.
4.  The `Model` returns the data to the `Controller`.
5.  The `Controller` chooses a `View` and passes the data to it.
6.  The `View` renders the final HTML and sends it back to the user.

---

## 3. MVC vs. Layered Architecture

While they seem similar, there's a key philosophical difference:

| Feature             | MVC Pattern                                   | Layered Architecture (3-Tier)          |
| :------------------ | :-------------------------------------------- | :------------------------------------- |
| **Focus**           | User Interface Interaction                    | **Code Organization & Dependencies**   |
| **"Brain"**         | Model                                         | **Service Layer**                      |
| **Controller Role** | A router and data-passer                      | A thin entry point                     |
| **Data Flow**       | Triangular (Controller talks to Model & View) | Linear (Controller -> Service -> Repo) |
| **Use Case**        | Server-Side Web Apps                          | Any application (Backend, CLI, etc.)   |

**In short:** MVC is about how you handle **user interaction**, while Layered is about how you structure **backend dependencies**. They can even be used together!

---
