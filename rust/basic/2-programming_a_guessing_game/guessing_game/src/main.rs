// bring the io input/output library into scope
use rand::Rng;
use std::{cmp::Ordering, io};

fn main() {
    // `println!` is a macro that prints a string to the screen
    println!("Guess the number!");

    let secret_number = rand::thread_rng().gen_range(1..=100);

    // println!("The secret number is: {secret_number}");

    loop {
        println!("Please input your guess:");

        // NOTE:
        // - `mut` is added to make a variable mutable.
        // - String::new() returns a new instance of a `String`.
        // - The `::` syntax in the `::new` indicates that `new` is an associated function on the `String` type.
        let mut guess = String::new();

        // NOTE:
        // - `stdin` returns an instance that represents a handle to the standard input for terminal
        // - `read_line` returns a `Result` value, which is an enum that can be in one of multiple possible states
        //      (variants). Result's variants are `Ok` and `Err`. Result has an `expect` method that can be called.
        //      If `Err`, `expect` will cause the program to crash and display the message that is passed as an argument
        //      to `expect`. If `Ok`, `expect` will take the return value that `Ok` is holding and return just that
        //      value so it can be used.
        io::stdin()
            .read_line(&mut guess)
            .expect("Failed to read line");

        // `parse` converts a string to another type. In this case, we use it to convert from a string to a number (u32).
        let guess: u32 = match guess.trim().parse() {
            Ok(num) => num,
            Err(_) => continue, // _ is a catchcall value that match all `Err` values
        };

        println!("You guessed: {guess}");

        // - `cmp`` method compares two values and can be called on anything that can be compared. In this case, `cmp`
        //      compares `guess` to `secret_numer`.
        match guess.cmp(&secret_number) {
            Ordering::Less => println!("Too small!"),
            Ordering::Greater => println!("Too big!"),
            Ordering::Equal => {
                println!("You win!");
                break;
            }
        }
    }
}
