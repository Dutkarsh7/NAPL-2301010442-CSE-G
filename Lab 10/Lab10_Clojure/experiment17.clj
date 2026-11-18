;; ============================================================
;; Experiment 17: Clojure Core Immutable Data Structures
;; B.Tech CSE – Clojure Lab (Lab 10)
;;
;; Concepts demonstrated:
;;   - List    : ordered, linked, ideal for sequential access
;;   - Vector  : ordered, indexed, ideal for random access
;;   - Map     : key-value pairs, ideal for lookup tables
;;   - Set     : unique elements, ideal for membership tests
;;
;; KEY IDEA: Every "modification" in Clojure returns a NEW
;; collection. The original is never changed (immutability).
;; ============================================================

;; ----------------------------------------------------------------
;; 1. LIST
;;    A list is a singly-linked sequence.
;;    Defined with  '(...)  or  (list ...)
;;    - first  → returns the first element
;;    - rest   → returns everything after the first element
;;    - conj   → adds an element to the FRONT of a list
;; ----------------------------------------------------------------
(println "============================")
(println " SECTION 1: LIST")
(println "============================")

;; Create the original list
(def original-list '(10 20 30 40 50))
(println "Original list     :" original-list)

;; first – get the head element
(println "first             :" (first original-list))

;; rest – get everything after the first element
(println "rest              :" (rest original-list))

;; conj – add 99 to the list (goes to the FRONT for lists)
(def list-after-conj (conj original-list 99))
(println "After conj 99     :" list-after-conj)

;; Prove immutability: original list is still unchanged
(println "Original unchanged:" original-list)

(println)

;; ----------------------------------------------------------------
;; 2. VECTOR
;;    A vector is an indexed, random-access collection.
;;    Defined with  [...]  or  (vector ...)
;;    - conj  → adds an element to the END of a vector
;;    - get   → retrieves an element by its 0-based index
;;    - count → returns the number of elements
;; ----------------------------------------------------------------
(println "============================")
(println " SECTION 2: VECTOR")
(println "============================")

;; Create the original vector
(def original-vec [1 2 3 4 5])
(println "Original vector   :" original-vec)

;; get – retrieve element at index 2 (zero-based)
(println "get index 2       :" (get original-vec 2))

;; count – how many elements?
(println "count             :" (count original-vec))

;; conj – add 6 to the END of the vector
(def vec-after-conj (conj original-vec 6))
(println "After conj 6      :" vec-after-conj)

;; Prove immutability: original vector is still unchanged
(println "Original unchanged:" original-vec)

(println)

;; ----------------------------------------------------------------
;; 3. MAP
;;    A map stores key-value pairs (like a dictionary).
;;    Defined with  {:key val ...}
;;    - assoc  → returns a new map with an added/updated pair
;;    - dissoc → returns a new map with a key removed
;;    - get    → retrieves the value for a key
;; ----------------------------------------------------------------
(println "============================")
(println " SECTION 3: MAP")
(println "============================")

;; Create the original map
(def original-map {:name "Alice" :age 21 :course "CSE"})
(println "Original map      :" original-map)

;; get – look up a value by key
(println "get :name         :" (get original-map :name))

;; assoc – add a new key :grade, or update an existing one
(def map-after-assoc (assoc original-map :grade "A"))
(println "After assoc :grade:" map-after-assoc)

;; dissoc – remove the :age key
(def map-after-dissoc (dissoc original-map :age))
(println "After dissoc :age :" map-after-dissoc)

;; Prove immutability: original map is still unchanged
(println "Original unchanged:" original-map)

(println)

;; ----------------------------------------------------------------
;; 4. SET
;;    A set is an unordered collection of UNIQUE elements.
;;    Defined with  #{...}
;;    - conj      → adds an element (duplicates are silently ignored)
;;    - disj      → removes an element
;;    - contains? → checks whether an element exists
;; ----------------------------------------------------------------
(println "============================")
(println " SECTION 4: SET")
(println "============================")

;; Create the original set
(def original-set #{1 2 3 4 5})
(println "Original set      :" original-set)

;; contains? – check membership
(println "contains? 3       :" (contains? original-set 3))
(println "contains? 9       :" (contains? original-set 9))

;; conj – add element 6
(def set-after-conj (conj original-set 6))
(println "After conj 6      :" set-after-conj)

;; conj with duplicate – adding 3 again has no effect
(def set-dup-conj (conj original-set 3))
(println "conj duplicate 3  :" set-dup-conj)

;; disj – remove element 2
(def set-after-disj (disj original-set 2))
(println "After disj 2      :" set-after-disj)

;; Prove immutability: original set is still unchanged
(println "Original unchanged:" original-set)

(println)
(println "============================================")
(println " All four immutable data structures shown!")
(println "============================================")
