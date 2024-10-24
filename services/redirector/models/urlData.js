class URLData {
    constructor(url, shortTag) {
        this.url = url;
        this.shortTag = shortTag;
    }

    toJSON() {
        return {
            url: this.url,
            short_tag: this.shortTag
        }
    }
}

module.exports = URLData;
