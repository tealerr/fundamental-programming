public class loop {
    public static void main(String args[]){
        forLoop();
        forEachLoop();
        whileLoop();
        doWhileLoop();
        breakAndCon();
        forLength();
    }

    //ทำงานในข
    public static void forLoop(){
        for (int i = 1; i < 5; i++) {
            System.out.println("Count at "+i);
        }
            System.out.println("");    
    }

    public static void forEachLoop(){
        String[] country = {"Japan", "United States", "Thailand"};

        // วนค่าทั้งหมดใน country array และเก็บไว้ที่ตัวแปร i
        for (String i : country) {
            System.out.println("My country is "+i); //แสดงค่าของสมาชิกใน array
          }
          System.out.println("");    
    }

    //ทำงานในขณะที่เงื่อนไขเป็นจริง
    public static void whileLoop(){
        int i = 5;
        while (i>=0 ){
            System.out.println("while loop countdown "+i);
            i--;
        }
            System.out.println("");    
    }

    public static void doWhileLoop(){
        int i = 1; 
        do{
            System.out.println(i);  //print i ไปเรื่อยๆ
            i++;
        }
        // ในขณะที่ i <5 (เงื่อนไขยังเป็นจริง)
        // ถ้า  i = 6 จะหยุดการทำงาน
        while( i <= 5);
        System.out.println("");    

    }

    public static void breakAndCon(){
        //นับ 1 ถึง 10
        System.out.println("Count 1 to 10 !");

        //ใช้ for loop  
        for (int i = 1; i<= 10; i++){
            if (i == 7 ){ // ให้เบรคเมื่อนับถึงครั้งที่ 7
            System.out.println("I'm use break");
            break; //ใช้เพื่อให้ออกจาก loop ทันทีโดยไม่สนใจเงื่อนไขของ loop
        }
            System.out.println(i); // print ไปเรื่อยๆจนครบเงื่อนไข
        }
        System.out.println("");    

        // 
        System.out.println("Count 1 to 5 !");
        int i = 1;

        //ใช้ while loop  
        while (i <= 5) {
            if (i == 3) { //ข้ามการนับครั้งที่ 3
                System.out.println("I'm using continue");
                i++;
                continue; //ใช้เพื่อข้ามการทำงานบางส่วนของลูป
            }
            System.out.println(i);
            i++;
        }
        System.out.println("");    
    }

    public static void forLength(){
        //สร้าง array ของยี่ห้อรถยนต์
        String[] cars = { "BMW", "Benz", "Ferrari", "Volvo"};

        System.out.println("Use for loop with my cars array");    
        for(int i = 0; i < cars.length; i++){
            System.out.println("i have "+cars[i]); //print สมาชิกออกมาทีละตัว 
        }
    }
}

