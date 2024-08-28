class ApplicationLink {
    constructor(infra, link) {
        this.infra = infra.Init();
        this.Update(link);
    };

    Update(updatedLink) {
        if  ((this.link == undefined) || (this.link.Anchor() != updatedLink.Anchor())) {
            this.infra.SetAnchor(updatedLink.Anchor());
        }

        if (updatedLink.HasOnClickEvent()) {
            var  onClickEvent = updatedLink.OnClickEvent()
            this.infra.SetOnClickEvent(onClickEvent);
        };

        if (updatedLink.HasStyles()) {
            var styles = updatedLink.GetStyles();
            this.infra.UpdateStyles(styles);
        }
        
        this.link = updatedLink;
        return this;
    };
};