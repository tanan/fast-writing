use diesel::prelude::*;
use diesel::r2d2::{ConnectionManager, Pool};
use diesel::{Connection, MysqlConnection};
use std::rc::Rc;
use crate::error::Error;
use lazy_static::lazy_static;
use crate::schema::push_notifications::dsl::push_notifications as table_push_notifications;

lazy_static! {
    pub static ref WRITING_DB_POOL: Pool<ConnectionManager<MysqlConnection>> = Pool::builder()
        .max_size(
            std::env::var("WRITING_DB_CONNECTION_POOL_MAX_SIZE")
                .unwrap_or("10".to_string())
                .parse::<u32>()
                .unwrap()
        )
        .build(ConnectionManager::<MysqlConnection>::new(
            std::env::var("WRITING_DB_URL")
                .unwrap_or("mysql://root:password@127.0.0.1/fast_writing".to_string())
        ))
        .expect("Failed to create pool.");
}

pub struct Transaction(pub(self) Rc<r2d2::PooledConnection<ConnectionManager<MysqlConnection>>>);

pub trait WritingDbTransactional {
    type Tx;

    fn transaction<F, R>(&self, func: F) -> Result<R, Error>
    where
        F: FnOnce(Self::Tx) -> Result<R, Error>;
}

pub trait WritingDbFindBy<Tx, Q, R> {
    fn find_by(&self, tx: &Tx, query: Q) -> Result<Vec<R>, Error>;
}

// pub trait WritingDbFindAll<Tx, Q> {
//     fn find_all(%self, tx: &Tx) -> Result<Vec<R>, Error>;
// }

pub struct MySqlWritingDb;

pub mod entity {
    use super::*;

    #[derive(Identifiable, Queryable, Eq, PartialEq, Clone)]
    #[cfg_attr(test, derive(Debug))]
    #[table_name = "push_notifications"]
    pub struct Notification {
        pub id: i32,
        pub user_id: String,
    }
}

impl WritingDbTransactional for MySqlWritingDb {
    type Tx = transaction;

    fn transaction<F, R>(&self, func: F) -> Result<R, Error>
    where
        F: FnOnce(Self::Tx) -> Result<R, Error>
    {
        let conn = Rc::new(WRITING_DB_POOL.get()?);
        conn.transaction(|| func(Transaction(conn.clone())))
    }
}

impl WritingDbFindBy<Transaction, i32, entity::Notification> for MySqlWritingDb {
    fn find_by(
        &self,
        tx: &Transaction,
        query: i32
    ) -> Result<Vec<crate::writing_db::entity::Notification>, Error> {
        let entities = table_push_notifications
            .filter()
            .select()
            .load::<entity::Notification>(&*tx.0)?;

        Ok(entities);
    }
}