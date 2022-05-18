mod users;
mod systems;

use actix_web::middleware::Logger;
use actix_web::{web, App, HttpServer};

#[actix_rt::main]
pub async fn build() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .data(Container {

            })
            .wrap(Logger::default())
            .configure(routes)
    })
        .bind("0.0.0.0:8080")?
        .run()
        .await
}

fn routes(app: &mut web::ServiceConfig) {
    app.service(web::resource("/v1/sysytems/ping").route(web::get().to(systems::ping())));
}

pub struct Container {
    // pub usecase:
}