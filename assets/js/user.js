$('#stop-follow').on('click', stopFollow);
$('#follow').on('click', follow);
$('#btn-save').on('click', save);
$('#btn-cancel').on('click', cancel);
$('#update-password').on('click', updatePass);
$('#delete-user').on('click', deleteUser);

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
            })
    }).fail(function () {
        Swal.fire("Erro!", "Falha ao alterar usuário!", "error");
    });
}

function cancel(event) {
    event.preventDefault();
    window.location = "/profile";
}

function updatePass(event) {
    event.preventDefault();

    if ($('new-password').val() != $('confirm-password').val()) {
        Swal.fire("Atenção!", "As senhas digitadas não são iguais!", "warning");
        return;
    }

    $.ajax({
        url: "/update-pass",
        method: 'POST',
        data: {
            current: $('#current-password').val(),
            new: $('#new-password').val()
        }
    }).done(function () {
        Swal.fire("Concluído!", "Senha alterada com sucesso!", "success")
            .then(function () {
                window.location = "/profile";
            })
    }).fail(function () {
        Swal.fire("Erro!", "Falha ao alterar a senha!", "error");
    });
}

function deleteUser() {

    Swal.fire({
        title: "Atenção!!!",
        text: "Deseja mesma excluir sua conta de usuário??? Você não terá mais acesso ao seus conteúdos!!!!!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function (confirm) {
        if (confirm.value) {
            $.ajax({
                url: "/delete-user",
                method: 'DELETE'
            }).done(function () {
                Swal.fire("Concluído!", "Conta excluída com sucesso!", "success")
                    .then(function () {
                        window.location = "/logout";
                    })
            }).fail(function () {
                Swal.fire("Erro!", "Falha ao excluir a conta!", "error");
            });
        }
    });
}
