// Copy  and paste this file into the browser
// console at  https://worrydream.com/quotes/
// and paste results to the quotes.json file.

let data = [];

for (const node of document.querySelectorAll(".entry")) {
    const id = node.id;
    const author = node.attributes["data-author"].value;
    const quote = document.querySelector(`#${id} .quote`)?.innerText ?? "";
    const source = document.querySelector(`#${id} .source .title`);
    const title = source?.innerText.replace(/^: /, "") ?? "";
    const link = source?.href ?? "";

    data.push({
        author: author,
        source: title,
        link: link,
        quote: quote
    });
};

JSON.stringify(data)