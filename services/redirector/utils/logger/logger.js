// logger.js
const winston = require('winston');

// Define log format
const logFormat = winston.format.combine(
    winston.format.timestamp({ format: 'YYYY-MM-DD HH:mm:ss' }),
    winston.format.json() // Logs in JSON format
);

// Initialize logger
const logger = winston.createLogger({
    level: process.env.LOG_LEVEL || 'info', // Set log level from environment variable or default to 'info'
    format: logFormat,
    transports: [
        new winston.transports.Console({
        format: process.env.NODE_ENV === 'development'
            ? winston.format.combine(
                winston.format.colorize(), // Makes console output colorful in development
                winston.format.simple()
            )
            : logFormat
        }),
        new winston.transports.File({ filename: 'logs/error.log', level: 'error' }), // Separate file for error logs
        new winston.transports.File({ filename: 'logs/combined.log' }) // General log file
    ]
});

// If not in production, log to the `console` with the colorized simple format
if (process.env.NODE_ENV !== 'production') {
    logger.add(
        new winston.transports.Console({
            format: winston.format.combine(
                winston.format.colorize(),
                winston.format.simple()
            )
        })
    );
}

module.exports = logger;
