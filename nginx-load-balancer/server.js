const http = require('http');
const os = require('os');

const hostname = os.hostname();
const PORT = process.env.PORT || "3000";

http.createServer((req, res) => {
    let joke = `THIS IS SERIOUS BUSINESS`;
    switch (PORT) {
        case "3000":
            joke =  `Why don't skeletons fight each other? Because they don't have the guts.`
            break;
        case "8080":
            joke =  `What do you call fake spaghetti? An impasta!`
            break;
        case "1313":
            joke = `Why did the scarecrow win an award?Because he was outstanding in his field!`
            break;
        case "1517":
            joke = `Why don’t eggs tell jokes? Because they might crack up!`
            break;
        default:
            break;
    }    

    console.log(JSON.stringify({ port : PORT , joke }))

    res.writeHead(200, { 'Content-Type': 'application/json' });

    res.end(JSON.stringify({ port : PORT , joke }));
}).listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}/`);
});
