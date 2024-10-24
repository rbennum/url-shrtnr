const amqp = require('amqplib');
const logger = require('../logger/logger.js');
const URLData = require('../../models/urlData.js');
const { EventEmitter } = require('events');

let connection;
let channel;
const queue = 'url-shrtnr';
const rabbitEvents = new EventEmitter();
const rabbitEventName = 'EVENT_RABBIT';
const rabbitEventErrorName = 'EVENT_RABBIT_ERROR';

async function openConnection() {
    try {
        logger.info(`Trying to connect to ${process.env.RABBITMQ_URL}`);
        connection = await amqp.connect(process.env.RABBITMQ_URL || 'amqp://localhost');
        channel = await connection.createChannel();
        await channel.assertQueue(queue, { durable: true });
        logger.info(`Connected to RabbitMQ, queue: ${queue}`);
    } catch (error) {
        logger.error('Failed to connect to RabbitMQ:', error);
        throw error;
    }
}

async function closeConnection() {
    try {
        if (channel) {
            await channel.close();
            logger.info('RabbitMQ channel closed');
        }
        if (connection) {
            await connection.close();
            logger.info('RabbitMQ connection closed');
        }
    } catch (error) {
        logger.error('Error closing RabbitMQ connection:', error);
    }
}

async function consumeMessages() {
    if (!channel) {
        logger.error('RabbitMQ channel is not available for consuming messages');
        return;
    }
    channel.consume(queue, (message) => {
        if (message !== null) {
            try {
                const messageContent = JSON.parse(message.content.toString());
                const urlDataInstance = new URLData(
                    messageContent.url,
                    messageContent.short_tag
                );
                rabbitEvents.emit(rabbitEventName, urlDataInstance);
                channel.ack(message);
            } catch (error) {
                logger.error('Failed to process message:', error);
                rabbitEvents.emit(rabbitEventErrorName);
                channel.nack(message);
            }
        }
    });
}

// since handling error event is just to log it out,
// we're doing it here
rabbitEvents.on(rabbitEventErrorName, (error) => {
    logger.error("Something's wrong when listening to RabbitMQ:", error);
})

module.exports = {
    openConnection,
    closeConnection,
    consumeMessages,
    rabbitEventName,
    rabbitEvents
};
