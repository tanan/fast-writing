use crate::error::Error;
use crate::domain::notification::Notifications;

pub trait NotificationPort {
    type Transaction;

    fn transaction<F, R>(&self, func: F) -> Result<R, Error>
    where
        F: FnOnce(&Self::Transaction) -> Result<R, Error>;

    fn get_notifications(&self, transaction: &Self::Transaction) -> Result<Notifications, Error>;

}