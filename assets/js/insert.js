$('#forminsuser').on('submit', InsertUser);
$('#btn-cancel').on('submit', OpenLoginPage);

function InsertUser(event) {
    event.preventDefault();
    console.log("Javascript!!!")

    if ($('#pass').val() != $('#confpass').val()) {
        alert('As senhas estão diferentes!')
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
    });
}

function OpenLoginPage(event) {
    event.preventDefault();
    console.log("Não voltou!!!")

    window.location.href("/");
}
