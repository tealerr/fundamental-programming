;loop/recur
(loop [x 1]
  (when (<= x 5) ; เช็คเงื่อนไข  x <= 5 หรือไม่
    (println "Count is "x)  ; print ค่า x
    (recur (inc x)))) ; เพิ่มค่า x ทีละ 1
(println "")

; while loop คล้ายกับภาษาอื่น จะทำงานในขณะที่เงื่อนไขเป็นจริง
(let [x (atom 3)] ; กำหนดให้ค่าเริ่มต้นของ x = 3
  (while (<= @x 9) ; กำหนดเงื่อนไข
    (println "while count is" @x) ; แสดงผล x 
    (swap! x + 2))) ; เพิ่มค่า x ทีละ 2
(println "")

;doseq loop ใช้เพื่อทำงานกับสมาชิกในลิสต์หรือคอลเล็คชันต่าง ๆ
(doseq [n [0 1 2 "A" "T"]] ; n เก็บค่าที่ออกมาจาก array
  (println "deseq output is "n)) ; d
(println "")

;Dotimes ใช้เพื่อทำงานซ้ำตามจำนวนรอบที่กำหนดไว้
(dotimes [n 5] ; วนซ้ำจำนวน 5 รอบ โดยค่าเริ่มต้น = 0
  (println "dotimes count is " (inc n))) ;ใช้ inc เพื่อให้ค่าเริ่มต้นเป็น 1