
// main:
document.addEventListener('DOMContentLoaded', function () {
    const parent = document.getElementById("root");
    var domLink = new DomLink(
        parent,
        "my_link",
        `
            color: #ffffff;
            background-color: #df3079;
            padding: 20px;
            margin: 5px;
            text-decoration: none;
        `,
        `
            color: #000000;
            background-color: #ffffff;
            text-decoration: underline;
        `,
        
    );

    var link = new Link(
        "Save", 
        function(evt) {
            console.log("save clicked!" + Math.random());
        },
    );

    var appLink = new ApplicationLink(
        domLink,
        link,
    );

    appLink.Update(new Link(
        "Updated", 
        function(evt) {
            console.log("update clicked!" + Math.random());
        },
    ));

} );

