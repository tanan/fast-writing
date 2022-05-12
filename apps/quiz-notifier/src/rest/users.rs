use diesel::prelude::*;
use diesel::r2d2::{ConnectionManager, Pool};
use diesel::{Connection, Insertable, MysqlConnection};
use lazy_static::lazy_static;
use std::rc::Rc;

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
                .unwrap_or("mysql://root:password@127.0.0.1/fast_writing".to_string())
        ))
        .expect("Failed to create pool.");
}

use crate::error::Error;

pub struct Transaction(pub(self) Rc<r2d2::PooledConnection<ConnectionManager<MysqlConnection>>>);

pub trait UserDbTransactional {
    type Tx;

    fn transaaction<F, R>(&self, func: F) -> Result<R, Error>
    where
        F: FnOnce(Self::Tx) -> Result<R, Error>;
}

pub trait UserDbFindBy<Tx, Q, R> {
    fn find_by(&self, tx: &Tx, query: Q) -> Result<Vec<R>, Error>;
}

pub trait UserDbFindAll<Tx, R> {
    fn find_all(&self, tx: &Tx) -> Result<Vec<R>, Error>;
}