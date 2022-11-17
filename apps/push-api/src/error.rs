use thiserror::Error;

#[derive(Error, Debug)]
pub enum Error {
    #[error("Failed to execute a query with the error: {message:?}")]
    Db { message: String },
    #[error("Failed to get mysql connection from connection pool with the error: {error:?}")]
    GetMysqlConnection { error: String },
    #[error("Unique violation: {message:?}")]
    UniqueViolation { message: String },
    #[error("Foreign key violation: {message:?}")]
    ForeignKeyViolation { message: String },
    #[error("Not Found: {message:?}")]
    NotFound { message: String },
    #[error("Bad Request: {message:?}")]
    BadRequest { message: String },
}