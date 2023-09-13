; gradeCalculate รับคะแนนและคำนวณออกมาเป็นเกรด โดยใช้ if
; ในภาษา clojure จะไม่มี if, else-if แต่จะใช้เป็น if ไปเรื่อยๆ และถ้าเป็นเงื่อนไขสุดท้ายจะเขียนเฉพาะคำสั่ง เช่น "F"
(defn gradeCalculate [score]
  (if (>= score 80) ; เงื่อนไขแรก
    "A" ; if true, grade = A, false ไปเงื่อนไขต่อไป
    (if (and (>= score 70) (< score 80)) ; เงื่อนไขที่ 2
      "B" ; if true, grade = B, false ไปเงื่อนไขสุดท้าย
      "F"))) ; คล้ายกับ else ในภาษาอื่นๆ ในกรณีที่ไม่เข้าเงื่อนไขใดๆ grade = F

; showGrade แสดงผลเกรดออกมาเป็นข้อความ โดยใช้ switch case
(defn showGrade [grade] ; รับ grade ที่ส่งมาจาก function gradeCalculate
  (let [message (case grade ; ใช้ switch case เพื่อเช็คเกรดว่าตรงกับเคสใด และทำการแสดงผลตามเกรด
                  "A" "Your grade is A"
                  "B" "Your grade is B"
                  "F" "Your grade is F"
                  "Invalid grade")]
    (println message)))

; เรียกใช้งาน function ใน main
(defn main []
  (let [score (gradeCalculate 72)] ;ใส่คะแนน =  72
    (showGrade score)))

; run function main
(main)
