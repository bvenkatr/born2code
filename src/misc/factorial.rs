pub fn facto(n: i32) -> i32 {
    let mut factorial = 1;
    if n == 0 {
        0
    } else if n == 1 {
        1
    } else {
        for i in 1..n+1 {
            factorial *= i;
        }
        factorial
    }
}

pub fn factorial_recursive(n: i32) -> i32 {
    if n == 0 {
        0
    } else if n == 1 {
        1
    } else {
        n * factorial_recursive(n - 1)
    }
}