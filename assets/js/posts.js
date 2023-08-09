//$('#newpost').on('submit', createPost);
//$('.like-post').on('click', likePost);
$(document).on('click', '.insert-post', insertPost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);
$(document).on('click', '.edit-post', updatePost);
$(document).on('click', '.delete-post', deletePost);
//$('#update-post').on('click', updatePost);

function insertPost(event) {

    event.preventDefault();

    $.ajax({
        url: "/post",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function () {
        Swal.fire("Concluído!", "Publicado com sucesso!", "success");
        window.location = "/home";
    }).fail(function () {
        Swal.fire("Erro!", "Erro ao publicar!", "error");
    })
}

function likePost(event) {
    event.preventDefault();

    const heart = $(event.target);
    const postId = heart.closest('div').data('post-id');

    heart.prop('disabled', true);
    //console.log(postId);

    $.ajax({
        url: `/post/${postId}/like`,
        method: "POST"
    }).done(function () {
        const likes = heart.next('span');
        const likesConv = parseInt(likes.text());

        likes.text(likesConv + 1);

        heart.addClass('unlike-post');
        heart.addClass('text-danger');
        heart.removeClass('like-post');

    }).fail(function () {
        Swal.fire("Erro!", "Erro ao curtir!", "error");
    }).always(function () {
        heart.prop('disabled', false);
    });
}

function unlikePost(event) {
    event.preventDefault();

    const heart = $(event.target);
    const postId = heart.closest('div').data('post-id');

    heart.prop('disabled', true);

    $.ajax({
        url: `/post/${postId}/unlike`,
        method: "POST"
    }).done(function () {
        const likes = heart.next('span');
        const likesConv = parseInt(likes.text());

        likes.text(likesConv - 1);

        heart.removeClass('unlike-post');
        heart.removeClass('text-danger');
        heart.addClass('like-post');

    }).fail(function () {
        Swal.fire("Erro!", "Erro ao descurtir!", "error");
    }).always(function () {
        heart.prop('disabled', false);
    });
}

function updatePost(event) {
    //event.preventDefault();
    $(this).prop('disabled', true);

    const postId = $(this).data('post-id');
    //console.log(postId);

    $.ajax({
        url: `/post/${postId}`,
        method: "PUT",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function () {
        Swal.fire("Concluído!", "Atualizado com sucesso!", "success")
            .then(function () { window.location = "/home"; })
    }).fail(function () {
        Swal.fire("Falhou!", "Erro ao Atualizar!", "error")
            .then(function () { window.location = "/home"; })
    }).always(function () {
        $('#edit-post').prop('disabled', false);
    });
}

function deletePost(event) {
    event.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Deseja mesmo excluir esta publicação?",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function (confirmation) {
        if (!confirmation.value) return;

        const trash = $(event.target);
        const post = trash.closest('div')
        const postId = post.data('post-id');

        trash.prop('disabled', true);

        $.ajax({
            url: `/post/${postId}`,
            method: "DELETE",
        }).done(function () {
            post.fadeOut("slow", function () {
                $(this).remove();
            })
        }).fail(function () {
            Swal.fire("Erro!", "Erro ao excluir!", "error");
        });
    })
}