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
        Swal.fire("Erro!", "Usuário ou senha inválidos!!", "error");
    });
}
