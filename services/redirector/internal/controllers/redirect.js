const logger = require('../../utils/logger/logger');
const redirectService = require('../services/redirect');

async function handleRedirect(_, res, parsedUrl) {
    try {
        const shortCode = parsedUrl.pathname.slice(1);
        logger.info(`Short URL: ${shortCode}`);

        if (shortCode) {
            const originalUrl = await redirectService.getOriginalUrl(shortCode);
            logger.info(`Original URL: ${originalUrl}`);
            
            if (originalUrl) {
                res.writeHead(302, { Location: originalUrl });
                res.end();
            } else {
                res.writeHead(404, { 'Content-Type': 'text/plain' });
                res.end('Short URL not found');
            }
        } else {
            res.writeHead(404, { 'Content-Type': 'text/plain' });
            res.end('Short URL not found');
        }
    } catch (error) {
        logger.error('Error handling request:', error);
        res.writeHead(500, { 'Content-Type': 'text/plain' });
        res.end('Internal Server Error');
    }
}

module.exports = {
    handleRedirect,
};
