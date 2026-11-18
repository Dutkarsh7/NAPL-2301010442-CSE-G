;; ============================================================
;; Experiment 18: Sequence Transformations in Clojure
;; B.Tech CSE – Clojure Lab (Lab 10)
;;
;; Concepts demonstrated:
;;   - map    : transform every element of a sequence
;;   - filter : keep elements that satisfy a predicate
;;   - reduce : accumulate elements into a single value
;;   - take   : take the first N elements
;;   - drop   : drop the first N elements, keep the rest
;;   - ->>    : threading macro to chain operations cleanly
;;
;; Two different collections are tested for each operation.
;; ============================================================

;; ----------------------------------------------------------------
;; Collections used throughout this experiment
;; ----------------------------------------------------------------
(def numbers  [1 2 3 4 5 6 7 8 9 10])   ;; a simple integer vector
(def scores   [85 42 91 60 73 55 88 37]) ;; student exam scores

(println "============================================")
(println " EXPERIMENT 18: SEQUENCE TRANSFORMATIONS")
(println "============================================")
(println)

;; ----------------------------------------------------------------
;; 1. MAP – Transform every element
;;    (map f coll) applies function f to each element and
;;    returns a new lazy sequence of the results.
;; ----------------------------------------------------------------
(println "----------------------------")
(println " 1. MAP")
(println "----------------------------")

;; Square every number in 'numbers'
(def squared (map #(* % %) numbers))
(println "Original numbers      :" numbers)
(println "Squared (map)         :" (vec squared))

;; Add a 10-point bonus to every score in 'scores'
(def boosted-scores (map #(+ % 10) scores))
(println "Original scores       :" scores)
(println "Scores + 10 bonus     :" (vec boosted-scores))

(println)

;; ----------------------------------------------------------------
;; 2. FILTER – Select elements based on a condition
;;    (filter pred coll) keeps only elements for which
;;    pred returns true.
;; ----------------------------------------------------------------
(println "----------------------------")
(println " 2. FILTER")
(println "----------------------------")

;; Keep only even numbers from 'numbers'
(def evens (filter even? numbers))
(println "Original numbers      :" numbers)
(println "Even numbers (filter) :" (vec evens))

;; Keep only passing scores (>= 60) from 'scores'
(def passing (filter #(>= % 60) scores))
(println "Original scores       :" scores)
(println "Passing scores (>=60) :" (vec passing))

(println)

;; ----------------------------------------------------------------
;; 3. REDUCE – Aggregate a sequence into a single value
;;    (reduce f coll) folds the collection left-to-right using f.
;;    (reduce f init coll) starts with an initial value.
;; ----------------------------------------------------------------
(println "----------------------------")
(println " 3. REDUCE")
(println "----------------------------")

;; Sum of all numbers
(def total (reduce + numbers))
(println "Original numbers      :" numbers)
(println "Sum (reduce +)        :" total)

;; Maximum score using reduce
;; max returns the larger of two values; reduce applies it pairwise
(def highest (reduce max scores))
(println "Original scores       :" scores)
(println "Maximum score (reduce):" highest)

;; Count of passing scores (reduce used as a counter)
(def pass-count
  (reduce (fn [acc s] (if (>= s 60) (inc acc) acc)) 0 scores))
(println "Passing score count   :" pass-count)

(println)

;; ----------------------------------------------------------------
;; 4. TAKE and DROP – Slice a sequence
;;    take  → returns the first N elements
;;    drop  → skips the first N elements, returns the rest
;; ----------------------------------------------------------------
(println "----------------------------")
(println " 4. TAKE and DROP")
(println "----------------------------")

;; Take first 4 numbers
(def first-four (take 4 numbers))
(println "Original numbers      :" numbers)
(println "take 4                :" (vec first-four))

;; Drop first 4 numbers (keep the rest)
(def after-four (drop 4 numbers))
(println "drop 4                :" (vec after-four))

;; Take top-3 scores (sort descending first, then take 3)
(def top-three (take 3 (sort > scores)))
(println "Original scores       :" scores)
(println "Top 3 scores          :" (vec top-three))

(println)

;; ----------------------------------------------------------------
;; 5. THREADING MACRO ->>
;;    ->> passes each result as the LAST argument to the next form.
;;    It makes a pipeline of transformations readable top-to-bottom
;;    instead of nesting function calls inside out.
;;
;;    Example pipeline on 'numbers':
;;      1. filter  – keep only odd numbers
;;      2. map     – square each odd number
;;      3. take    – take the first 3 results
;; ----------------------------------------------------------------
(println "----------------------------")
(println " 5. THREADING MACRO ->>")
(println "----------------------------")

;; Pipeline on 'numbers': odd → squared → first 3
(def pipeline-result-1
  (->> numbers
       (filter odd?)          ;; keep 1 3 5 7 9
       (map #(* % %))         ;; square: 1 9 25 49 81
       (take 3)))             ;; first three: 1 9 25

(println "Original numbers      :" numbers)
(println "Pipeline (odd→sq→take3):" (vec pipeline-result-1))

;; Pipeline on 'scores': keep >=60 → add 5 bonus → sort descending
(def pipeline-result-2
  (->> scores
       (filter #(>= % 60))    ;; keep passing scores
       (map #(+ % 5))         ;; add 5 bonus marks
       (sort >)))             ;; sort high to low

(println "Original scores       :" scores)
(println "Pipeline (pass→+5→sort):" (vec pipeline-result-2))

(println)
(println "============================================")
(println " All sequence transformations demonstrated!")
(println "============================================")
