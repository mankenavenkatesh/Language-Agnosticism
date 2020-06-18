fn main() {
    // fixed type in array
    let a = [1,2,3,4,5];
    println!("a= {},{},{},{},{}", a[0],a[1],a[2],a[3],a[4]);

    let months = ["January", "February", "March", "April", "May", "June", "July",
              "August", "September", "October", "November", "December"];
    println!("months= {},{},{},{},{}", months[0],months[1],months[2],months[3],months[4]);

    // array type and size
    let b: [i32; 5] = [11, 12, 13, 14, 15];
    println!("b= {},{},{},{},{}", b[0],b[1],b[2],b[3],b[4]);

    // same as let a = [3, 3, 3, 3, 3];
    let c = [3; 5];
    println!("c= {},{},{},{},{}", c[0],c[1],c[2],c[3],c[4]);
    
}
