fn main() {
    // auto types
    let tup = (500, 4.5, 1);
    let (x,y,z) = tup;
    println!("tup = {},{},{}", x,y,z);

    // static typing
    let tup1:(f64, i32, char) = (3.2, 1, 'c');
    let (a, b, c) = tup1;
    println!("tup1 = {},{},{}", a,b,c);
    
    // indexing
    let tup2=(100, 200, 300);
    println!("tup2={},{},{}",tup2.0, tup2.1,tup2.2);
}