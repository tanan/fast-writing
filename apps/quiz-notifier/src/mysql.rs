use mysql::*;
use mysql::prelude::*;

#[derive(Debug, PartialEq, Eq)]
struct User {
    name: String,
    email: String,
}

fn create() {
  // 
}

fn query() {
  let url = "mysql://root:password@127.0.0.1:3306/fast_writing";
  let pool = Pool::new(url)?;
  let mut conn = pool.get_conn()?;
  let users = conn
      .query_map(
          "SELECT name, email from user",
          |(name, email)| {
              User { name, email }
          },
      )?;
  for row in users {
    Ok(v) => {
      
    }
  }
}
