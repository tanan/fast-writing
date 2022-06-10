#[derive(Debug, Queryable)]
pub struct User {
    pub id: String,
    pub name: String,
    pub email: String,
}

#[derive(Debug, Queryable)]
pub struct WritingContents {
    pub id: i32,
    pub user_id: String,
    pub title: String,
    pub description: String,
}