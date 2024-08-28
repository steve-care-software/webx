export class Dimension {
    private width?: number;
    private height?: number;
    constructor(width?: number, height?: number) {
        this.width = width;
        this.height = height;
    };

    // HasWidth returns true if there is a width, false otherwise
    public HasWidth(): boolean {
        return this.width != undefined
    }

    // Width returns the width, if any
    public Width(): number {
        let result: number | undefined = this.width;
        if (result == undefined) {
            throw new Error(`the width is undefined`);
        }

        return result as number
    }

    // Height returns the height
    public Height(): number {
        let result: number | undefined = this.height;
        if (result == undefined) {
            throw new Error(`the height is undefined`);
        }

        return result as number
    }
};