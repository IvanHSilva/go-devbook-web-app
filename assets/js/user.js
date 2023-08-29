$('#stop-follow').on('click', stopFollow);
$('#follow').on('click', follow);

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
