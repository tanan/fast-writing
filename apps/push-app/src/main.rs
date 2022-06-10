use diesel::prelude::*;
use push_app::establish_connection;
use push_app::models::{User, WritingContents};
use push_app::schema::user as user_schema;
use push_app::schema::writing_contents as wc_schema;
#[macro_use] extern crate diesel;

mod schema;

fn main() {
    let connection = establish_connection();
    let users = user_schema::dsl::user
        .load::<User>(&connection)
        .expect("Error loading users");
    for user in users {
        println!("{:?}", user)
    }

    let contents = wc_schema::dsl::writing_contents
        .load::<WritingContents>(&connection)
        .expect("Error loading users");

    for c in contents {
        println!("{:?}", c)
    }
}
