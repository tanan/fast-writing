table! {
    user (id) {
        id -> VarChar,
        name -> VarChar,
        email -> VarChar,
    }
}

table! {
    writing_contents (id) {
        id -> Integer,
        user_id -> VarChar,
        title -> VarChar,
        description -> VarChar,
        // last_updated -> Timestamp,
    }
}
