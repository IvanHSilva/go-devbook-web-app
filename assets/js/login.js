$('#login-form').on('submit', Login);

function Login(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        dataType: "text",
        data: {
            email: $('#email').val(),
            pass: $('#pass').val(),
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        alert("Usuário ou senha inválidos!")
    });
}
