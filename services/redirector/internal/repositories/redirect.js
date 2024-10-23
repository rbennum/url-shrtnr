const { db } = require('../../utils/db/driver');

async function findOriginalUrl(shortCode) {
    const result = await db('link_mappers')
        .select('url')
        .where({ short_tag: shortCode })
        .first();
    
    return result ? result.url : null;
}

module.exports = {
    findOriginalUrl,
};
