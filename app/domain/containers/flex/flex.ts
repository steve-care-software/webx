import {Alignment} from './alignments/alignment'
import {Overflow} from './overflows/overflow'
import {Direction} from './directions/direction'

export class Flex {
    private horizontal: Alignment;
    private vertical: Alignment;
    private overflow: Overflow;
    private direction: Direction;
    constructor(horizontal: Alignment, vertical: Alignment, overflow: Overflow, direction: Direction) {
        this.horizontal = horizontal;
        this.vertical = vertical;
        this.overflow = overflow;
        this.direction = direction;
    };

    // Horizontal returns the horizontal alignment
    public Horizontal(): Alignment {
        return this.horizontal;
    }

    // Vertical returns the vertical alignment
    public Vertical(): Alignment {
        return this.vertical;
    }

     // Overflow returns the overflow
     public Overflow(): Overflow {
        return this.overflow;
    }

    // Direction returns the direction
    public Direction(): Direction {
        return this.direction;
    }
};