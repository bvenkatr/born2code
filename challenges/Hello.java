public class Hello {
   String name;
   int age;
   Hello(String name, int age){
      this.name = name;
      this.age = age;
   }
   public static void main(String[] args) {
      Hello obj = new Hello("Venkat", 27);
      System.out.println(obj.name);
      System.out.println(obj.age);
   }
}
