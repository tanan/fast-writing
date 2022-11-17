use crate::domain::notification::Notifications;
use crate::error::Error;
use crate::port::notification::NotificationPort;

pub struct NotificationGetAllUseCase<T: NotificationPort> {
    pub notification_port: T,
}

impl<T: NotificationPort> NotificationGetAllUseCase<T> {
    pub fn execute(&self) -> Result<Notifications, Error> {
        self.notification_port
            .transaction(|tx | self.notification_port.get_notifications(tx))
    }
}