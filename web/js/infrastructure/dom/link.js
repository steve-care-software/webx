class DomLink {
    constructor(
        parentDom,
        identifier,
        onClickEvent,
    ) {
        this.parentDom = parentDom;
        this.identifier = identifier;
        if (onClickEvent != null) {
            this.clickListener = this.createClickListener(onClickEvent);
        }

        this.instance = null;
        this.styleInstance = null;
    };

    createClickListener(onClickEvent) {
        return function(evt) {
            evt.preventDefault(); // Prevent the default action (navigating to the href)
            onClickEvent(evt);
        };
    };

    Init() {
        this.styleInstance = document.createElement('style');
        this.styleInstance.textContent = "text/css";
        this.parentDom.appendChild(this.styleInstance);

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

    UpdateStyles(styles) {
        var normalStr = styles.GetNormal().String();
        var hoverStr = styles.GetHover().String();

        var style = "";
        if (normalStr != "") {
            style = `#${this.identifier} { ${normalStr} }`;
        }

        if (hoverStr != "") {
            style += `#${this.identifier}:hover { ${hoverStr} }`;
        }

        if (style == "") {
            return;
        }

        this.styleInstance.innerHTML = style;
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