const fs = require('fs');
const quotes = require('./quotes.json');
const outputFilePath = './cli/quotes.go';

function* convertQuoteToGoCode(quote) {
    const entries = {
        'Author': quote.author,
        'SourceTitle': quote.source,
        'SourceURL': quote.link,
        'Quote': quote.quote,
    };

    yield '{\n';

    for (const key in entries) {
        yield '   ' + key + ': ' + JSON.stringify(entries[key]) + ',\n';
    }

    yield '},\n';
}

fs.writeFileSync(outputFilePath, '');
const writeStream = fs.createWriteStream(outputFilePath);
writeStream.write('package main\n\nvar quotes = []Quote{\n');

for (const quote of quotes) {
    for (const line of convertQuoteToGoCode(quote)) {
        writeStream.write(line);
    }
}

writeStream.write('}');