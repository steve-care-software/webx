
// main:
document.addEventListener('DOMContentLoaded', function () {

    var stylesBuilder = new StylesBuilder(
        new StyleBuilder(),
    );
    
    var styles = stylesBuilder.Create().WithNormalValues({
        "color": "#ffffff",
        "background-color": "#df3079",
        "padding": "20px",
        "margin": "5px",
        "text-decoration": "none",
    }).WithHoverValues({
        "color": "#000000",
        "background-color": "#ffffff",
        "text-decoration": "underline",
    }).Now();

    const parent = document.getElementById("root");
    var domLink = new DomLink(
        parent,
        "my_link",
    );

    var link = new Link(
        "Save", 
        styles,
        function(evt) {
            console.log("save clicked!" + Math.random());
        },
    );

    var appLink = new ApplicationLink(
        domLink,
        link,
    );

    var updatedStyles = stylesBuilder.Create().WithExisting(styles).WithNormalValues({
        "background-color": "#ff0000",
        "text-decoration": "underline",
    }).Now();

    appLink.Update(new Link(
        "Updated", 
        updatedStyles,
        function(evt) {
            console.log("update clicked!" + Math.random());
        },
    ));

} );

