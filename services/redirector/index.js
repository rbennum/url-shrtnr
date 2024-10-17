require('dotenv').config({ path: '.env.local', debug: true });
const { migrateDatabase, db } = require('./utils/db/driver.js');
const logger = require('./utils/logger/logger.js');
const { 
    openConnection: openRabbitMQConnection,
    closeConnection: closeRabbitMQConnection,
    rabbitEventName,
    consumeMessages,
    rabbitEvents
} = require('./utils/rabbitmq/driver.js');
const http = require('http');

let server;
const PORT = process.env.SERVER_PORT;

async function gracefulShutdown(signal) {
    logger.info(`Received ${signal}. Shutting down gracefully...`);

    // Close HTTP server to stop accepting new requests
    if (server) {
        server.close(() => {
            logger.info('Closed HTTP server');
        });
    }

    // Close database connection
    db.destroy(() => {
        logger.info('Closing database');
    });

    // Close rabbitmq connection
    closeRabbitMQConnection();

    process.exit(0);
}

async function startApp() {
    logger.info('Initiating app')

    // fetch env vars
    // run dotenv from the outside system, like:
    // node -r dotenv/config index.js dotenv_config_path=.env.local \
    // dotenv_config_debug=true

    // open RabbitMQ connection
    await openRabbitMQConnection();

    // start listening to messages
    consumeMessages();

    // migrate database
    await migrateDatabase();

    // start the server
    server = http.createServer((req, res) => {
        if (req.url === '/') {
            res.writeHead(200, { 'Content-Type': 'text/plain' });
            res.end('Hello, world!');
        } else {
            res.writeHead(404, { 'Content-Type': 'text/plain' });
            res.end('Page not found');
        }
    });
    server.listen(PORT, () => {
        logger.info(`Server listening on port ${PORT}`);
    })
}

// start the server
startApp().catch((error) => {
    logger.error('Error starting server:', error);
    process.exit(1);
});

// listening to termination signals
['SIGINT', 'SIGTERM'].forEach(signal => {
    process.on(signal, () => gracefulShutdown(signal));
});

rabbitEvents.on(rabbitEventName, (data) => {
    // TODO: Process fetched messages here
})
