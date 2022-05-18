use diesel::prelude::*;
use diesel::{Connection, MysqlConnection};
use diesel::r2d2::ConnectionManager;
use lazy_static::lazy_static;
use r2d2::Pool;
use std::rc::Rc;
use actix_http::RequestHeadType::Rc;
use crate::schema::user;
use crate::schema::user::dsl::user as table_user;

lazy_static! {
    pub static ref USER_DB_POOL: Pool<ConnectionManager<MysqlConnection>> = Pool::builder()
        .max_size(
            std::env::var("USER_DB_CONNECTION_POOL_MAX_SIZE")
                .unwrap_or("10".to_string())
                .parse::<u32>()
                .unwrap()
        )
        .build(ConnectionManager::<MysqlConnection>::new(
            std::env::var("USER_DB_URL")
                .unwrap_or("mysql://developer:developer@localhost/user".to_string())
        ))
        .expect("Failed to create pool.");
}

use crate::error::Error;

pub struct Transaction(pub(self) Rc<r2d2::PooledConnection<ConnectionManager<MysqlConnection>>>);

pub trait UserDbTransactional {
    type Tx;

    fn transaction<F, R>(&self, func: F) -> Result<R, Error>
        where
            F: FnOnce(Self::Tx) -> Result<R, Error>;
}

pub trait UserDbSave<Tx, E, R> {
    fn save(&self, tx: &Tx, entity: E) -> Result<R, Error>;
}

pub trait UserDbFindBy<Tx, Q, R> {
    fn find_by(&self, tx: &Tx, query: Q) -> Result<Vec<R>, Error>;
}

pub trait UserDbFindAll<Tx, R> {
    fn find_all(&self, tx: &Tx) -> Result<Vec<R>, Error>;
}

pub trait UserDbDelete<Tx, E, Q> {
    fn delete(&self, tx: &Tx, query: Q) -> Result<(), Error>;
}

pub struct MySqlUserDb;

pub mod entity {
    use super::*;

    #[derive(Insertable, Queryable)]
    #[cfg_attr(test, derive(Clone, Eq, PartialEq, Debug))]
    #[table_name = "user"]
    pub struct User {
        pub id: String,
        pub name: String,
        pub email: String,
    }
}

impl UserDbTransactional for MySqlUserDb {
    type Tx = Transaction;

    fn transaction<F, R>(&self, func: F) -> Result<R, Error>
        where F: FnOnce(Self::Tx) -> Result<R, Error>
    {
        let conn = Rc::new(USER_DB_POOL.get()?);
        conn.transaction(|| func(Transaction(conn.clone())))
    }
}

impl UserDbFindBy<Transaction, String, entity::User> for MySqlUserDb {
    fn find_by(&self, tx: &Transaction, query: String) -> Result<Vec<User>, Error> {
        let entities =
    }
}