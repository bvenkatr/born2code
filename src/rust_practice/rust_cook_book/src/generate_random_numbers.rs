use rand::Rng;
// In the provided code snippet, although the use rand::Rng; statement brings the Rng trait into scope, the code doesn't directly reference the Rng trait in its explicit form (Rng is not used explicitly in the code). However, the usage of Rng is indirectly implied and utilized.

// Here's the explanation:

// use rand::Rng;: This statement imports the Rng trait from the rand crate into the current scope. Importing the Rng trait allows the methods defined in the Rng trait to be used on types that implement the Rng trait.

// let mut rng = rand::thread_rng();: This line creates a random number generator (rng) using the thread_rng() function provided by the rand crate. The thread_rng() function returns a type that implements the Rng trait behind the scenes (even though Rng isn't explicitly mentioned in the code).

// println!("Random f64: {}", rng.gen::<f64>());: This line demonstrates the usage of the random number generator rng by calling its gen method with the type parameter <f64>. The gen method is part of the Rng trait, although it's not explicitly stated in the code. This method generates a random f64 value using the underlying functionality provided by the Rng trait.

// So, while the code doesn't directly reference Rng explicitly in the main logic, the code utilizes the functionality of the Rng trait indirectly through the thread_rng() function and the gen method provided by types that implement the Rng trait. The Rng trait serves as an abstraction behind the scenes to enable random number generation in a concise and usable manner.


pub fn grn() {
    let mut rng = rand::thread_rng();
    println!("Random f64: {}", rng.gen::<f64>());
}

// declare a module
pub fn my_public_function() {
    println!("I am a public function in my_mod.rs");
}

fn my_function(){
    println!("Hi, you came inside the root function using super");
}

// declare a module
pub mod r {
    pub fn print_statement(){
      println!("Hi, this a function of module r");
    }

    fn my_private_function(){
        println!("Hi, I'm a private function within the module");
      }
      pub fn my_public_function(){
        //! also works without writing self i.e.
        //! my_private_function();
        println!("Hi,I'm a public function within the module");
        println!("I'll invoke private function within the module");
        self::my_private_function(); 
        
      }
}

// declare a module
pub mod outer_module {
    // function within outer module
    fn my_private_function() {
        println!("Hi, I got into the private function of outer module");
    }
    // declare a nested module
    pub mod inner_module {
        // function within nested module
        pub fn my_public_function() {
        println!("Hi, I got into the public function of inner module");
        println!("I'll invoke private function of outer module");
        super::my_private_function();
        }
    }
}

  // declare a module
pub mod my_module {
    // function within outer module
    pub fn my_public_function() {
      println!("Invoke root function");
      super::my_function();
    }
}


pub mod chapter {
    pub mod lesson { // mod level 1
        pub mod heading { // mod level 2
            pub fn illustration() {  // mod level 3
              println!("Hi, I'm a 3rd level nested function");
            }
        }
    }
}
