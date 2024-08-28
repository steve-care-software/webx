class StyleBuilder {
    constructor() {
        this.values = null;
        this.separator = ":";
        this.eol = ";";
    };

    Create() {
        return new StyleBuilder();
    };

    WithValues(values) {
        this.values = values
        return this
    };

    WithSeparator(separator) {
        this.separator = separator
        return this
    };

    WithEOL(eol) {
        this.eol = eol
        return this
    };

    Now() {
        return new Style(this.values, this.separator, this.eol);
    };
};


class Style {
    constructor(values, separator, eol) {
        this.values = values;
        this.separator = separator;
        this.eol = eol;
    };

    Values() {
        return this.values;
    };

    String() {
        var output = "";
        for (const key in this.values) {
            output += key + this.separator + this.values[key] + this.eol;
        }

        return output;
    };
};

class StylesBuilder {
    constructor(
        styleBuilder,
    ) {
        this.styleBuilder = styleBuilder;
        this.existing = null;
        this.normalValues = null;
        this.hoverValues = null;
        this.separator = ":";
        this.eol = ";";
    };

    Create() {
        return new StylesBuilder(this.styleBuilder);
    };

    WithExisting(styles) {
        this.existing = styles
        return this
    };

    WithNormalValues(values) {
        this.normalValues = values
        return this
    };

    WithHoverValues(values) {
        this.hoverValues = values
        return this
    };

    WithSeparator(separator) {
        this.separator = separator
        return this
    };

    WithEOL(eol) {
        this.eol = eol
        return this
    };

    Now() {

        // verify that separator and eol are valid:

        var normalValues = (this.normalValues != null) ? this.normalValues : {};
        var hoverValues = (this.hoverValues != null) ? this.hoverValues : {};

        var normalBuilder = this.styleBuilder.Create().
                                                WithValues(this.normalValues).
                                                WithSeparator(this.separator).
                                                WithEOL(this.eol);

        var hoverBuilder = this.styleBuilder.Create().
                                            WithValues(this.hoverValues).
                                            WithSeparator(this.separator).
                                            WithEOL(this.eol);

        if (this.existing != null) {
            var normal = {...this.existing.GetNormal().Values(), ...normalValues};
            normalBuilder.WithValues(normal)

            var hover = {...this.existing.GetHover().Values(), ...hoverValues};
            hoverBuilder.WithValues(hover)
        };
        
        var normal = normalBuilder.Now();
        var hover = hoverBuilder.Now();
        return new Styles(normal, hover);
    };
};

class Styles {
    constructor(normal, hover) {
        this.normal = normal;
        this.hover = hover;
    };

    GetNormal() {
        return this.normal;
    };

    GetHover() {
        return this.hover;
    };
};

class Link {
    constructor(
        anchor, 
        styles,
        onClickEvent,
    ) {
        this.anchor = anchor;
        this.styles = styles;
        this.onClickEvent = onClickEvent;
    };

    Anchor() {
        return this.anchor;
    };

    HasStyles() {
        return this.styles != null
    };

    GetStyles() {
        return this.styles
    };

    HasOnClickEvent() {
        return (this.onClickEvent != null);
    };

    OnClickEvent() {
        return this.onClickEvent;
    };
};