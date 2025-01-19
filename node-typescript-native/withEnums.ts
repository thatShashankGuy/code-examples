
enum ReplacementType {
    Fruits,
    Animals,
}
interface SearchReplaceOperationWithEnum {
    text: string;
    type: ReplacementType;
}

function searchAndReplaceWithEnum(operation: SearchReplaceOperationWithEnum): string {
    const { text, type } = operation;

    switch (type) {
        case ReplacementType.Fruits:
            return text.replace(/🍎|🍊/g, (match) => (match === "🍎" ? "🍌" : "🍍"));

        case ReplacementType.Animals:
            return text.replace(/🐵/g, "🦍");

        default:
            return text;
    }
}

const fruitOperation: SearchReplaceOperationWithEnum = {
    text: "comparing 🍎 with 🍊 and 🍎 again",
    type: ReplacementType.Fruits,
};

const fruitReplacedText = searchAndReplaceWithEnum(fruitOperation);
console.log("Original Text (Fruits):", fruitOperation.text);
console.log("Replaced Text (Fruits):", fruitReplacedText);

const animalOperation: SearchReplaceOperationWithEnum = {
    text: "🐵 see 🐵 do 🐵 say",
    type: ReplacementType.Animals,
};

const animalReplacedText = searchAndReplaceWithEnum(animalOperation);
console.log("Original Text (Animals):", animalOperation.text);
console.log("Replaced Text (Animals):", animalReplacedText);
