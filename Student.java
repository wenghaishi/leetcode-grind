//write a class student that contains attributes: name, age and courses. Write methods for getters and setters.

import java.util.ArrayList;

public class Student {
  public String name;
  public int age;
  public ArrayList<String> courses;

  public Student(String name, int age) {
    this.name = name;
    this.age = age;
    this.courses = new ArrayList<String>();
  }

  public String getName() {
    return name;
  }

  public int getAge() {
    return age;
  }

  public ArrayList<String> getCourses() {
    return courses;
  }
}
