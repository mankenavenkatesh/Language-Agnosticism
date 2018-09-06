package main;
import "fmt";

func main() {
	var a int = 21;
	var b int = 10;
	var c int;

	c = a + b;
	fmt.Printf("Sum of 21 and 10 is: %d\n", c);
	c = a - b;
	fmt.Printf("Differnce of 21 and 10 is : %d\n", c);
	c = a * b;
	fmt.Printf("Multiplication of 21 and 10 is : %d\n", c);
	c = a / b;
	fmt.Printf("Division of 21 and 10 is : %d\n", c);
	c = a % b;
	fmt.Printf("Remainder when 21 divided by 10 is : %d\n", c);
	a++;
	fmt.Printf("Increment a by 1 using increment operator:%d\n", a);
	a--;
	fmt.Printf("Decrement a by 1 using decrement operator:%d\n", a);
	
}