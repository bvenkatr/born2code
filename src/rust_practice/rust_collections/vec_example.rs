pub fn loop_vec() {
    let mut v = Vec::new();
    for i in 1..10 {
        v.push(i);
    }
    println!("length is = {:?}", v.len());
}
