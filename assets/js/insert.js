$('#forminsuser').on('submit', InsertUser);
$('#btn-cancel').on('submit', OpenLoginPage);

function InsertUser(event) {
    event.preventDefault();
    console.log("Javascript!!!")

    if ($('#pass').val() != $('#confpass').val()) {
        Swal.fire("Erro!", "As senhas estão diferentes!", "error");
        return;
    }

    const today = new Date();
    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $('#name').val(),
            email: $('#email').val(),
            pass: $('#pass').val()
            //date: today.toLocaleDateString('pt-BR')
        }
    }).done(function () {
        Swal.fire("Concluído!", "Usuário cadastrado com sucesso!", "success")
            .then(function () {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    dataType: "text",
                    data: {
                        email: $('#email').val(),
                        pass: $('#pass').val()
                    }
                }).done(function () {
                    window.location = "/home";
                }).fail(function () {
                    Swal.fire("Erro!", "Falha ao fazer login!", "error");
                })
            })
    }).fail(function () {
        Swal.fire("Erro!", "Falha ao cadastrar usuário!", "error");
    });
}

function OpenLoginPage(event) {
    event.preventDefault();
    console.log("Não voltou!!!")

    window.location.href("/");
}
