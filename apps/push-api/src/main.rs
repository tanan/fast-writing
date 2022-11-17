mod error;
mod rest;
mod usecase;
mod port;
mod domain;
mod gateway;
mod schema;
mod driver;

use std::env;

fn main() {
    env::set_var("RUST_BACKTRACE", "1");
    if env::var("RUST_LOG").ok().is_none() {
        env::set_var("RUST_LOG", "info");
    }
    env_logger::init();

    if let Err(e) = rest::build() {
        eprintln!("Application startup failed: {:?}!", e);

        std::process::exit(1);
    }
}
