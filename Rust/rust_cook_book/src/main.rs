mod generate_random_numbers;
use generate_random_numbers::chapter::lesson::heading;
fn main() {
    println!("Hello, world!");
    generate_random_numbers::my_public_function();
    generate_random_numbers::grn();
    generate_random_numbers::r::print_statement();
    generate_random_numbers::r::my_public_function();
    generate_random_numbers::outer_module::inner_module::my_public_function();
    generate_random_numbers::my_module::my_public_function();
    generate_random_numbers::chapter::lesson::heading::illustration();
    heading::illustration();
}

