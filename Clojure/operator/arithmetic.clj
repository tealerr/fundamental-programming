;Number
(println "Number arithmetic")
(def num1 10)
(def num2 5)

(println (str num1 " + " num2 " = " (+ num1 num2)))
(println (str num1 " - " num2 " = " (- num1 num2)))
(println (str num1 " * " num2 " = " (* num1 num2)))
(println (str num1 " / " num2 " = " (/ num1 num2)))
(println (str num1 " % " num2 " = " (mod num1 num2)))
(println "\n")

;String


;Collections (Vectors):
(println "Vectors arithmetic")
(def nums [1 2 3 4 5])

(println (reduce + nums)) ; Sum of elements
(println (reduce * nums)) ; Product of elements
(println "\n")


;Collections (Maps):
(println "Maps arithmetic")
(def person {:age 30})

(println (+ (:age person) 5)) ; Increase age by 5
(println "\n")
