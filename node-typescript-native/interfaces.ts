export default interface SearchReplaceOperation {
    text: string;
    searchValue: string | RegExp;
    replaceValue: string;
}
