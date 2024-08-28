class Link {
    constructor(anchor, onClickEvent) {
        this.anchor = anchor;
        this.onClickEvent = onClickEvent;
    };

    Anchor() {
        return this.anchor;
    };

    HasOnClickEvent() {
        return (this.onClickEvent != null);
    };

    OnClickEvent() {
        return this.onClickEvent;
    };
};