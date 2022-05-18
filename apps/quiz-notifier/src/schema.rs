use diesel::table;

table! {
    user (id) {
        id -> VarChar,
        name -> VarChar,
        email -> VarChar,
    }
}