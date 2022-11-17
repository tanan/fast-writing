table! {
    push_notifications (id) {
        id -> Unsigned<Bigint>,
        user_id -> Varchar,
        created_at -> Nullable<Datetime>,
        updated_at -> Nullable<Datetime>,
    }
}

table! {
    user (id) {
        id -> Varchar,
        name -> Varchar,
        email -> Varchar,
        last_updated -> Nullable<Datetime>,
    }
}

table! {
    writing_contents (id) {
        id -> Bigint,
        user_id -> Varchar,
        title -> Varchar,
        description -> Varchar,
        scope -> Varchar,
        tags -> Nullable<Varchar>,
        last_updated -> Nullable<Datetime>,
        quiz -> Nullable<Json>,
    }
}

joinable!(push_notifications -> user (user_id));

allow_tables_to_appear_in_same_query!(
    push_notifications,
    user,
    writing_contents,
);
