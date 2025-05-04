const fs = require('fs');
const path = require('path');
const [,, langArg, ...fileArgs] = process.argv;

const language = langArg || 'en';
const fileList = fileArgs.length > 0 ? fileArgs : ['quotes']; // fallback to quotes.json

// const quotes = require('./quotes-data/es/quotes.json');
const outputFilePath = './cli/quotes.go';
fs.writeFileSync(outputFilePath, '');
const writeStream = fs.createWriteStream(outputFilePath);
writeStream.write('package main\n\nvar quotes = []Quote{\n');

function* convertQuoteToGoCode(quote) {
    const entries = {
        'Author': quote.author,
        'AuthorURL': quote.authorlink,
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

for (const file of fileList) {
    const filePath = path.join(__dirname, 'quotes-data', language, `${file}.json`);
    if (!fs.existsSync(filePath)) {
        console.error(`File not found: ${filePath}`);
        continue;
    }

    const quotes = JSON.parse(fs.readFileSync(filePath, 'utf8'));

    for (const quote of quotes) {
        for (const line of convertQuoteToGoCode(quote)) {
        writeStream.write(line);
        }
    }
}

writeStream.write('}');
writeStream.end();