export class Color {
    private red: number;
    private blue: number;
    private green: number;
    private alpha: number;
    constructor(
        red: number,
        blue: number,
        green: number,
        alpha: number,
    ) {
        this.red = red;
        this.blue = blue;
        this.green = green;
        this.alpha = alpha;
    }

    // Red returns the red number
    public Red(): number {
        return this.red
    }

    // Blue returns the blue number
    public Blue(): number {
        return this.blue
    }

    // Green returns the green number
    public Green(): number {
        return this.green
    }

    // Alpha returns the alpha number
    public Alpha(): number {
        return this.alpha
    }
};