use crate::domain::notification::{Notification, NotificationId, UserId};
use crate::domain::notification::Notifications;
use crate::port::notification::NotificationPort;
use crate::driver::writing_db::{entity, WritingDbFindBy, WritingDbTransactional};
use crate::error::Error;

#[cfg_attr(test, derive(Debug, Clone, Eq, PartialEq))]
pub struct  NotificationGatewayTransaction<T>(pub(self) T);

pub struct NotificationGateway<T> {
    pub writing_db: T,
}

impl<T> NotificationPort for NotificationGateway<T>
where
    T: WritingDbTransactional
        + WritingDbFindBy<<T as WritingDbTransactional>::Tx, i32, entity::Notification>
{
    type Transaction = NotificationGatewayTransaction<T::Tx>;

    fn transaction<F, R>(&self, func: F) -> Result<R, Error> where F: FnOnce(&Self::Transaction) -> Result<R, Error> {
        self.writing_db.transaction(|tx | {
            let tx = NotificationGatewayTransaction(tx);
            func(&tx)
        })
    }

    fn get_notifications(&self, tx: &Self::Transaction) -> Result<Notifications, Error> {
        self.writing_db.find_by(&tx.0).map(|values | {
            values
                .into_iter()
                .map(|entity| {
                    Notification::new(
                        NotificationId(entity.id),
                        UserId(entity.user_id),
                    )
                })
        })
    }
}