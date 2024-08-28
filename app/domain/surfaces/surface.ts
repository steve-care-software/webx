import {Dimension} from '../dimensions/dimension'
import {Position} from '../positions/position'
import {Color} from '../colors/color'
import {Border} from '../borders/border'

export class Surface {
    private dimension?: Dimension;
    private position?: Position;
    private color?: Color
    private border?: Border;
    constructor(
        dimension?: Dimension, 
        position?: Position, 
        color?: Color, 
        border?: Border,
    ) {
        this.dimension = dimension;
        this.position = position;
        this.color = color;
        this.border = border;
    };

    // HasDimension returns true if there is a dimension, false otherwise
    public HasDimension(): boolean {
        return this.dimension != undefined;
    }

    // Dimension returns the dimension, if any
    public Dimension(): Dimension {
        let result: Dimension | undefined = this.dimension;
        if (result == undefined) {
            throw new Error(`the dimension is undefined`);
        }

        return result as Dimension
    }

    // HasPosition returns true if there is a position, false otherwise
    public HasPosition(): boolean {
        return this.position != undefined;
    }

    // Position returns the position, if any
    public Position(): Position {
        let result: Position | undefined = this.position;
        if (result == undefined) {
            throw new Error(`the position is undefined`);
        }

        return result as Position
    }

    // HasColor returns true if there is a color, false otherwise
    public HasColor(): boolean {
        return this.color != undefined;
    }

    // Color returns the color, if any
    public Color(): Color {
        let result: Color | undefined = this.color;
        if (result == undefined) {
            throw new Error(`the color is undefined`);
        }

        return result as Color
    }

    // HasBorder returns true if there is a border, false otherwise
    public HasBorder(): boolean {
        return this.border != undefined;
    }

    // Border returns the border, if any
    public Border(): Border {
        let result: Border | undefined = this.border;
        if (result == undefined) {
            throw new Error(`the border is undefined`);
        }

        return result as Border
    }
};