mod foo;
mod controllers;
// use controllers::user_controller;
mod my;
mod rust_practice;
mod yours;
mod misc;

fn main() {
    // my::function();

    // function();

    // my::indirect_access();

    my::nested::function();
    yours::nested::function();
    rust_practice::rust_collections::vec_example::loop_vec();
    foo::hello();
    let user_info = controllers::user_controller::get_user(1);
    println!("{}", user_info);
    println!("factorial is {} ", misc::factorial::facto(4));
    println!("factorial is {} ", misc::factorial::factorial_recursive(4)); 
}