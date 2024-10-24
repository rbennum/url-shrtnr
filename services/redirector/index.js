require('dotenv').config({ path: '.env.local', debug: true });
const url = require('url');
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
const { handleRedirect } = require('./internal/controllers/redirect.js');

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

    // open RabbitMQ connection
    await openRabbitMQConnection();

    // start listening to messages
    consumeMessages();

    // migrate database
    await migrateDatabase();

    // start the server
    server = http.createServer((req, res) => {
        const parsedUrl = url.parse(req.url, true);
        handleRedirect(req, res, parsedUrl);
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

rabbitEvents.on(rabbitEventName, async (data) => {
    logger.debug(data.url)
    logger.debug(data.shortTag)
    try {
        const result = await db('link_mappers').insert({
            url: data.url,
            short_tag: data.shortTag
        });
        logger.info('Data successfully saved to MariaDB:', result);
    } catch (error) {
        logger.error('Error saving data to MariaDB:', error);
    }
})
