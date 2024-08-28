export class Overflow {
    private isOneLine: boolean;
    private isMultiLine: boolean;
    private isMultiLineReverse: boolean;
    constructor(isOneLine: boolean, isMultiLine: boolean, isMultiLineReverse: boolean) {
        this.isOneLine = isOneLine;
        this.isMultiLine = isMultiLine;
        this.isMultiLineReverse = isMultiLineReverse;
    }

    // IsOneLine returns true if oneLine, false otherwise
    public IsOneLine(): boolean {
        return this.isOneLine;
    }

     // IsMultiLine returns true if multi-line, false otherwise
     public IsMultiLine(): boolean {
        return this.isMultiLine;
    }

    // IsMultiLineReverse returns true if multi-line reverse, false otherwise
    public IsMultiLineReverse(): boolean {
        return this.isMultiLineReverse;
    }
};