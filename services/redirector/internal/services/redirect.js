const redirectRepository = require('../repositories/redirect');

async function getOriginalUrl(shortCode) {
    return await redirectRepository.findOriginalUrl(shortCode);
}

module.exports = {
    getOriginalUrl,
};
