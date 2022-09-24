use std::env;
use rand::Rng;
use diesel::prelude::*;
use serde_json::{json, Value};
use push_app::establish_connection;
use push_app::models::{User, WritingContents};
use push_app::schema::user as user_schema;
use push_app::schema::writing_contents as wc_schema;
#[macro_use] extern crate diesel;

mod schema;
#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {

    let connection = establish_connection();
    let users = user_schema::dsl::user
        .filter(user_schema::email.eq("toshifumi.anan@gmail.com"))
        .limit(1)
        .load::<User>(&connection)
        .expect("Error loading users");

    for user in users {
        let contents = find_contents_by_user_id(user, &connection);

        let num = rand::thread_rng().gen_range(0..contents.len());

        let attachment = build_message(format!("https://front-cmfot4feta-an.a.run.app/contents/{}", &contents[num].id));

        let body = json!({
            "channel": "tanan-notice1",
            "attachments": vec![attachment],
        });

        let slack_url = env::var("SLACK_URL").expect("slack url must be set.");
        let client = reqwest::Client::new();
        let res = client.post(slack_url).body(body.to_string()).send().await?;
        println!("{}", &res.text().await?);
    }

    Ok(())


}

fn find_contents_by_user_id(user: User, connection: &MysqlConnection) -> Vec<WritingContents> {
    return wc_schema::dsl::writing_contents
        .filter(wc_schema::user_id.eq(user.id))
        .load::<WritingContents>(connection)
        .expect("Error loading users");
}

fn build_message(contents_url: String) -> Value {
    let text = format!("今日のクイズ：{}", contents_url);
    return json!({
        "text": text,
        "color": "good",
    })
}
