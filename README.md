# Go CLI Todo Manager

A modular, persistent Command Line Interface (CLI) application built in Go. This project emphasizes clean code architecture, data persistence, and the use of modern Go features.

## 📂 Project Structure
- **`main.go`**: Entry point and program flow.
- **`todo.go`**: Core `Todo` logic and list methods (Add, Toggle, Print).
- **`storage.go`**: Generic JSON engine for saving and loading data.
- **`command.go`**: CLI flag parsing and input handling.
- **`todos.json`**: Persistent storage for your tasks.

---

## 🧠 Concepts Learned

### 1. Pointer Logic & Nullability
* **Pointer Receivers**: Used `(todos *Todos)` to ensure methods modify the original list in memory rather than a temporary copy.
* **Nullable Fields**: Implemented `*time.Time` for `CompletedAt` to allow for `nil` values, representing tasks that aren't yet finished.


### 2. Modern Generics (`[T any]`)
* **Reusable Logic**: Developed a `Storage[T]` struct that handles File I/O for any data type using type parameters.
* **Type Safety**: Leveraged Go's type system to create a flexible yet safe storage engine.


### 3. Custom Type Systems
* **Type Aliasing**: Used `type Todos []Todo` to encapsulate list logic and attach custom methods directly to slices.
* **Organization**: Grouped data and behavior together to keep the codebase clean and modular.

### 4. Data Persistence (JSON)
* **Serialization**: Used `json.MarshalIndent` to transform Go memory structures into human-readable JSON files.
* **File Permissions**: Managed local disk access using `os` package with standard `0644` permissions.


### 5. Slice Mechanics
* **Dynamic Slices**: Managed task lists using `append()` for growth and specific indexing for targeted updates.
* **Memory Efficiency**: Learned how slices point to underlying arrays to manage data without excessive copying.

---

## 🛠️ Usage
1. **Setup**: Run `go mod tidy` to sync dependencies.
2. **Execute**: Run `go run .`
3. **Flags**: Use flags like `-add`, `-list`, or `-toggle` to manage your tasks.
