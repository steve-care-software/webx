export class Alignment {
    private isStart: boolean;
    private isEnd: boolean;
    private isCenter: boolean;
    constructor(isStart: boolean, isEnd: boolean, isCenter: boolean) {
        this.isStart = isStart;
        this.isEnd = isEnd;
        this.isCenter = isCenter;
    }

    // IsStart returns true if at start, false otherwise
    public IsStart(): boolean {
        return this.isStart;
    }

    // IsEnd returns true if at end, false otherwise
    public IsEnd(): boolean {
        return this.isEnd;
    }

    // IsCenter returns true if at center, false otherwise
    public IsCenter(): boolean {
        return this.isCenter;
    }
};