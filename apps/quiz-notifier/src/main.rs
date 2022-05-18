mod rest;
mod driver;
mod error;
mod schema;

use std::env;
use log::{debug, error, log_enabled, info, Level};

#[macro_use]
extern crate serde;
#[macro_use]
extern crate diesel;
#[macro_use]
extern crate derive_getters;
#[macro_use]
extern crate derive_new;
#[cfg(test)]
extern crate mock_it;

fn main() {
    if env::var("RUST_LOG").ok().is_none() {
        env::set_var("RUST_LOG", "info");
    }
    env_logger::init();

    if let Err(e) = rest::build() {

    }
}