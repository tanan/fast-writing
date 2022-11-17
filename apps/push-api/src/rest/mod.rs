mod systems;

use actix_web::{web, App, HttpServer};
use actix_web::middleware::Logger;
use crate::usecase::notification_get_all::NotificationGetAllUseCase;

#[actix_web::main]
pub async fn build() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .app_data(Container {
                notification_get_all_usecase: NotificationGetAllUseCase {
                    notification_port:
                }
            })
            .wrap(Logger::default())
            .configure(routes)
    })
    .bind(("127.0.0.1", 10001))?
    .run()
    .await
}

fn routes(app: &mut web::ServiceConfig) {
    app.route("v1/systems/ping",web::get().to(systems::ping));
}

pub struct Container {
    pub notification_get_all_usecase: NotificationGetAllUseCase<>,
}