use thiserror::Error;

#[derive(Error, Debug, Clone, Eq, PartialEq)]
pub enum Error {
    #[error("Failed to execute a query with the error: {message:?]")]
    Db { message: String },
    #[error("Failed to get mysql connection from connection pool with the error: {error:?}")]
    GetMysqlConnection { error: String },
    #[error("Unique violation: {message:?}")]
    NotFound { message: String },
    #[error("Bad Request: {message:?}")]
    BadRequest { message: String },
}