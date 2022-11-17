use std::slice::Iter;

#[derive(Clone)]
#[cfg_attr(test, derive(Eq, PartialEq, Debug))]
pub struct NotificationId(pub i32);

#[derive(Clone)]
#[cfg_attr(test, derive(Eq, PartialEq, Debug))]
pub struct UserId(pub String);

impl<T: Into<String>> From<T> for UserId {
    fn from(string: T) -> Self { Self(string.into()) }
}

pub struct Notification {
    id: NotificationId,
    user_id: UserId,
}

#[derive(Clone)]
#[cfg_attr(test, derive(Eq, PartialEq, Debug))]
pub struct Notifications(Vec<Notification>);

impl Notifications {
    pub fn iter(&self) -> Iter<Notification> { self.0.iter() }
}

impl From<Vec<Notification>> for Notifications {
    fn from(vec: Vec<Notification>) -> Self {
        Notifications(vec)
    }
}