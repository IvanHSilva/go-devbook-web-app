$('#stop-follow').on('click', stopFollow);
$('#follow').on('click', follow);
$('#btn-save').on('click', save);
$('#btn-cancel').on('click', cancel);

function stopFollow() {

    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/user/$(userId)/unfollow`,
        method: 'POST'
    }).done(function () {
        window.location(`/user/$(userId)`,);
    }).fail(function () {
        Swal.fire("Falha!", "Erro ao deixar de seguir usuário!", "error");
        $('#stop-follow').prop('disabled', false);
    });
}

function follow() {

    const userId = $(this).data('user-id');
    $(this).prop('disabled', true);

    $.ajax({
        url: `/user/$(userId)/follow`,
        method: 'POST'
    }).done(function () {
        window.location(`/user/$(userId)`,);
    }).fail(function () {
        Swal.fire("Falha!", "Erro ao seguir usuário!", "error");
        $('#follow').prop('disabled', false);
    });
}

function save(event) {

    event.preventDefault();

    $.ajax({
        url: "/edit-user",
        method: 'PUT',
        data: {
            name: $('#name').val(),
            email: $('#email').val()
        }
    }).done(function () {
        Swal.fire("Concluído!", "Usuário alterado com sucesso!", "success")
            .then(function () {
                window.location = "/profile";
            });
    }).fail(function () {
        Swal.fire("Erro!", "Falha ao alterar usuário!", "error");
    });
}

function cancel(event) {
    event.preventDefault();
    window.location = "/profile";
}
