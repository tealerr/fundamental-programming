(defn gradeCalculate [score]
  (cond
    (>= score 80) "A"
    (and (>= score 70) (< score 80)) "B"
    :else "F"))

(gradeCalculate 75)