fn main() {
    println!("Hello !world");
    another_function();
    another_function1(10, 4);
    let s = sum(1,2);
    println!("sum of 1 and 2 is {}",s)
}

fn another_function() {
    println!("Another function.");
}

fn another_function1(x: i32, y: i32) {
    println!("The value of x is: {} and y is:{}", x, y);
}

fn sum(a: i32, b: i32) -> i32 {
    a+b
}