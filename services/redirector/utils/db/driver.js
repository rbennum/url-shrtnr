const knex = require('knex');
const logger = require('../logger/logger');

const db = knex({
    client: 'mysql2',
    connection: {
        host: process.env.DB_HOST,
        user: process.env.DB_USER,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_NAME,
        port: parseInt(process.env.DB_PORT)
    },
    migrations: {
        directory: './migrations'
    }
});

async function migrateDatabase() {
    try {
        logger.info('Running database migrations...');
        await db.migrate.latest();
        logger.info('Database migration completed.');
    } catch (error) {
        logger.error('Database migration failed:', error);
        process.exit(1);
    }
}

module.exports = {
    db,
    migrateDatabase
};