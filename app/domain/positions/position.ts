export class Position {
    private top: number;
    private left: number;
    constructor(top: number, left: number) {
        this.top = top;
        this.left = left;
    };

    // Top returns the top, if any
    public Top(): number {
       return this.top;
    }

    // Left returns the left
    public Left(): number {
        return this.left;
    }
};