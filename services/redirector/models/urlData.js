// urlData.js

class URLData {
    constructor(url, shortTag) {
        this.url = url;
        this.shortTag = shortTag;
    }

    // Optional: add methods to handle or manipulate URL data
    getFullURL() {
        return `${this.url} - ${this.shortTag}`;
    }
}

module.exports = URLData;
