export class Direction {
    private isRow: boolean;
    private isReverse: boolean;
    constructor(isRow: boolean, isReverse: boolean) {
        this.isRow = isRow;
        this.isReverse = isReverse;
    }

    // IsRow returns true if row, false otherwise
    public IsRow(): boolean {
        return this.isReverse;
    }

     // IsReverse returns true if reverse, false otherwise
     public IsReverse(): boolean {
        return this.isReverse;
    }
};