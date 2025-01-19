
import SearchReplaceOperation  from "./interfaces.ts";

function searchAndReplace(operation: SearchReplaceOperation): string {
    const { text, searchValue, replaceValue } = operation;
    return text.replace(searchValue, replaceValue);
}

const originalPhrase: string = "comparing🍎 with 🍊";
const search: string = "🍊";
const replacement: string = "🍌";

const operation: SearchReplaceOperation = {
    text: originalPhrase,
    searchValue: search,
    replaceValue: replacement,
};

const newPhrase = searchAndReplace(operation);
console.log("Original Text:", originalPhrase);
console.log("Replaced Text:", newPhrase);

// Using regex
const anotherPhrase = "🐵 see 🐵 do";
const regexOperation: SearchReplaceOperation = {
    text: anotherPhrase,
    searchValue: /🐵/g,
    replaceValue: "🦍",
};

const regexPhrase = searchAndReplace(regexOperation);
console.log("Original Text:", anotherPhrase);
console.log("Regex Replaced Text:", regexPhrase);
