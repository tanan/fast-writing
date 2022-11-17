use crate::error::Error;

pub mod writing_db;

impl From<r2d2::Error> for Error {
    fn from(error: r2d2::Error) -> Error {
        Error::GetMysqlConnection {
            error: format!("{}", error).to_string(),
        }
    }
}

impl From<diesel::result::Error> for crate::error::Error {
    fn from(error: diesel::result::Error) -> Error {
        let message = format!("{}", error).to_string();

        match error {
            diesel::result::Error::DatabaseError(kind, error) => {
                let message = format!("{:?}", error).to_string();

                match kind {
                    diesel::result::DatabaseErrorKind::UniqueViolation => {
                        Error::UniqueViolation { message }
                    }
                    diesel::result::DatabaseErrorKind::ForeignKeyViolation => {
                        Error::ForeignKeyViolation { message }
                    }
                    _ => Error::Db { message },
                }
            }
            diesel::result::Error::NotFound => Error::NotFound { message },
            _ => Error::Db { message },
        }
    }
}
