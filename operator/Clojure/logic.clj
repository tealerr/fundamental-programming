(def x 10)
(def y 5)

(def str1 "Hello")
(def str2 "World")

(def lst1 [1 2 3])
(def lst2 [])

(println "Number logical")
(println "x > y and x < 20 is " (and (> x y) (< x 20)))
(println "x < y or y < 3 is " (or (< x y) (< y 3)))
(println "\n")

(println "List logical")
(println "lst1 is not empty and lst2 is not empty: " (and (seq lst1) (seq lst2)))
(println "At least one of lst1 or lst2 is not empty: " (or (seq lst1) (seq lst2)))
(println "\n")

(println "Not (!) logical operator")
(println "10 < 3 is" (not (< x 3)))
(println (not nil))
(println (not true))
(println "\n")
