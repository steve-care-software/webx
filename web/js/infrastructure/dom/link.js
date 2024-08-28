class DomLink {
    constructor(
        parentDom,
        identifier,
        normalStyle,
        hoverStyle,
        onClickEvent,
    ) {
        this.parentDom = parentDom;
        this.identifier = identifier;
        this.normalStyle = normalStyle;
        this.hoverStyle = hoverStyle;
        if (onClickEvent != null) {
            this.clickListener = this.createClickListener(onClickEvent);
        }

        this.instance = null;
    };

    createClickListener(onClickEvent) {
        return function(evt) {
            evt.preventDefault(); // Prevent the default action (navigating to the href)
            onClickEvent(evt);
        };
    };

    Init() {
        var style = `#${this.identifier} { ${this.normalStyle} }`;
        if (this.hoverStyle != null) {
            style += `#${this.identifier}:hover { ${this.hoverStyle} }`;
        }

        const styleIns = document.createElement('style');
        styleIns.textContent = "text/css";
        styleIns.innerHTML = style;
        this.parentDom.appendChild(styleIns);

        const aIns = document.createElement('a');
        aIns.setAttribute("id", this.identifier);
        aIns.setAttribute("href", "#");
        this.parentDom.appendChild(aIns);

        // store the link:
        this.instance = aIns;
        return this;
    };

    SetAnchor(anchor) {
        this.instance.innerText = anchor;
        return this;
    };

    SetOnClickEvent(onClickEvent) {
        // if there is currently a click listener, remove it:
        if (this.clickListener != null) {
            this.instance.removeEventListener('click',  this.clickListener);
        }

        // add the new click event:
        this.clickListener = this.createClickListener(onClickEvent);
        this.instance.addEventListener('click', this.clickListener);
        return this;
    };
};