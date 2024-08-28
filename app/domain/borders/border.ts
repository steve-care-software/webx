import {Corner} from '../surfaces/corners/corner'
import {Color} from './../colors/color'


const dotted = 0;
const dashed = 1;
const solid = 2;
const double = 3;
const groove = 4;
const ridge = 5;
const inset = 6;
const outset = 7;
const none = 8;
const hidden = 9;

export class Border {
    private style: number;
    private color: Color;
    private size: number;
    constructor(style: number, color: Color, size: number) {
        this.style = style;
        this.color = color;
        this.size = size;
    };

    // Style returns the style flag
    public Style(): number {
        return this.style;
    }

    // Color returns the color
    public Color(): Color {
        return this.color;
    }

    // Size returns the size
    public Size(): number {
        return this.size;
    }
};