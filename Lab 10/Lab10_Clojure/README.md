# Lab 10 – Clojure Programming Practicals

**Subject:** Programming Paradigms / Functional Programming  
**Language:** Clojure 1.12  
**Java Runtime:** Java 25 (LTS)

---

## Contents

| File | Experiment | Marks |
|------|-----------|-------|
| `experiment17.clj` | Core Immutable Data Structures | 5 |
| `experiment18.clj` | Sequence Transformations | 5 |

---

## Experiment 17 – Clojure Core Immutable Data Structures

### Objective
Demonstrate all four core immutable data structures in Clojure — **List**, **Vector**, **Map**, and **Set** — and show that every "modification" produces a new collection while leaving the original unchanged.

### Concept Overview

| Structure | Literal Syntax | Characteristic |
|-----------|---------------|----------------|
| List | `'(1 2 3)` | Linked, `conj` adds to the front |
| Vector | `[1 2 3]` | Indexed, `conj` adds to the back |
| Map | `{:k v}` | Key-value pairs, O(log n) lookup |
| Set | `#{1 2 3}` | Unique elements, O(log n) membership |

### Operations Demonstrated

**List**
- `first` – returns the head element
- `rest` – returns the tail (everything after the first)
- `conj` – adds an element (at the front for lists)

**Vector**
- `conj` – adds an element at the end
- `get` – retrieves an element by zero-based index
- `count` – returns the number of elements

**Map**
- `get` – retrieves the value for a key
- `assoc` – returns a new map with an added/updated key-value pair
- `dissoc` – returns a new map with a key removed

**Set**
- `contains?` – checks whether an element is a member
- `conj` – adds an element (duplicates silently ignored)
- `disj` – removes an element

### Execution

```bash
clojure -M experiment17.clj
```

### Sample Output

```
============================
 SECTION 1: LIST
============================
Original list     : (10 20 30 40 50)
first             : 10
rest              : (20 30 40 50)
After conj 99     : (99 10 20 30 40 50)
Original unchanged: (10 20 30 40 50)
============================
 SECTION 2: VECTOR
============================
Original vector   : [1 2 3 4 5]
get index 2       : 3
count             : 5
After conj 6      : [1 2 3 4 5 6]
Original unchanged: [1 2 3 4 5]
============================
 SECTION 3: MAP
============================
Original map      : {:name Alice, :age 21, :course CSE}
get :name         : Alice
After assoc :grade: {:name Alice, :age 21, :course CSE, :grade A}
After dissoc :age : {:name Alice, :course CSE}
Original unchanged: {:name Alice, :age 21, :course CSE}
============================
 SECTION 4: SET
============================
Original set      : #{1 4 3 2 5}
contains? 3       : true
contains? 9       : false
After conj 6      : #{1 4 6 3 2 5}
conj duplicate 3  : #{1 4 3 2 5}
After disj 2      : #{1 4 3 5}
Original unchanged: #{1 4 3 2 5}
============================================
 All four immutable data structures shown!
============================================
```

---

## Experiment 18 – Sequence Transformations

### Objective
Demonstrate the core higher-order sequence functions in Clojure — **map**, **filter**, **reduce**, **take**, **drop** — and the **`->>`** threading macro for chaining operations. Each function is applied to two different collections.

### Collections Used
- `numbers` – `[1 2 3 4 5 6 7 8 9 10]`
- `scores`  – `[85 42 91 60 73 55 88 37]` (student exam scores)

### Concept Overview

| Function | Purpose | Returns |
|----------|---------|---------|
| `map` | Transform every element | Lazy sequence |
| `filter` | Keep elements matching a predicate | Lazy sequence |
| `reduce` | Fold a sequence into one value | Single value |
| `take n` | First N elements | Lazy sequence |
| `drop n` | All elements after the first N | Lazy sequence |
| `->>` | Thread value as last arg through forms | Whatever the last form returns |

### Operations Demonstrated

1. **map** – Square every number; add a 10-point bonus to every score.
2. **filter** – Keep only even numbers; keep only passing scores (≥ 60).
3. **reduce** – Sum all numbers; find the maximum score; count passing scores.
4. **take / drop** – Take first 4 numbers, drop first 4; take top-3 scores.
5. **`->>`** – Pipeline: `filter odd? → map square → take 3` on numbers;  
   `filter ≥60 → map +5 → sort descending` on scores.

### Execution

```bash
clojure -M experiment18.clj
```

### Sample Output

```
============================================
 EXPERIMENT 18: SEQUENCE TRANSFORMATIONS
============================================
----------------------------
 1. MAP
----------------------------
Original numbers      : [1 2 3 4 5 6 7 8 9 10]
Squared (map)         : [1 4 9 16 25 36 49 64 81 100]
Original scores       : [85 42 91 60 73 55 88 37]
Scores + 10 bonus     : [95 52 101 70 83 65 98 47]
----------------------------
 2. FILTER
----------------------------
Original numbers      : [1 2 3 4 5 6 7 8 9 10]
Even numbers (filter) : [2 4 6 8 10]
Original scores       : [85 42 91 60 73 55 88 37]
Passing scores (>=60) : [85 91 60 73 88]
----------------------------
 3. REDUCE
----------------------------
Original numbers      : [1 2 3 4 5 6 7 8 9 10]
Sum (reduce +)        : 55
Original scores       : [85 42 91 60 73 55 88 37]
Maximum score (reduce): 91
Passing score count   : 5
----------------------------
 4. TAKE and DROP
----------------------------
Original numbers      : [1 2 3 4 5 6 7 8 9 10]
take 4                : [1 2 3 4]
drop 4                : [5 6 7 8 9 10]
Original scores       : [85 42 91 60 73 55 88 37]
Top 3 scores          : [91 88 85]
----------------------------
 5. THREADING MACRO ->>
----------------------------
Original numbers      : [1 2 3 4 5 6 7 8 9 10]
Pipeline (odd→sq→take3): [1 9 25]
Original scores       : [85 42 91 60 73 55 88 37]
Pipeline (pass→+5→sort): [96 93 90 78 65]
============================================
 All sequence transformations demonstrated!
============================================
```

---

## Prerequisites

- [Clojure CLI](https://clojure.org/guides/install_clojure) (tested with 1.12.6)
- Java 11 or later (tested with Java 25 LTS)

Verify your installation:

```bash
clojure --version
java -version
```

---

## Running Both Programs

From inside the `Lab10_Clojure` folder:

```bash
clojure -M experiment17.clj
clojure -M experiment18.clj
```
