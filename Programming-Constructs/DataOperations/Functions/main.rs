fn main() {
    println!("Hello !world");
    another_function();
    another_function1(10, 4);
}

fn another_function() {
    println!("Another function.");
}

fn another_function1(x: i32, y: i32) {
    println!("The value of x is: {} and y is:{}", x, y);
}

