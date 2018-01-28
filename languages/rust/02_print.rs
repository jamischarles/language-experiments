// from https://rustbyexample.com/hello/print.html
//
// TODO STARTING: Finish the exercise on this page

fn main() {
    // In general, the `{}` will be automatically replaced with any
    // arguments. These will be stringified.
    println!("{} days", 31); //interpolation


    // There are various optional patterns this works with. Positional
    // arguments can be used.
    println!("{0}, this is {1}. {1}, this is {0}", "Alice", "Bob");

    // As can named arguments.
    println!("{subject} {verb} {object}",
             object="the lazy dog",
             subject="the quick brown fox",
             verb="jumps over");

    // Special formatting can be specified after a `:`.
    println!("{} of {:b} people know binary, the other half doesn't", 1, 2);

    // You can right-align text with a specified width. This will output
    // "     1". 5 white spaces and a "1".
    println!("{number:>width$}", number=1, width=6);

    // You can pad numbers with extra zeroes. This will output "000001".
    println!("{number:>0width$}", number=1, width=6);

     // It will even check to make sure the correct number of arguments are
    // used.
    println!("My name is {0}, {1} {0}", "Jamis", "Bond");
    // FIXME ^ Add the missing argument: "James". FIXED


    // Create a structure which contains an `i32`. Name it `Structure`.
    #[allow(dead_code)] // NOTE TO compiler, that allows code that isn't used!!! Cool
    struct Structure(i32);


    // However, custom types such as this structure require more complicated
    // handling. This will not work.
    // println!("This struct `{}` won't print...", Structure(3));
    // FIXME ^ Comment out this line. FIXED
}
