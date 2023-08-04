$('#newpost').on('submit', createPost);
//$('.like-post').on('click', likePost);
$(document).on('click', '.like-post', likePost);
$(document).on('click', '.unlike-post', unlikePost);

function createPost(event) {

    event.preventDefault();

    $.ajax({
        url: "/posts",
        method: "POST",
        data: {
            title: $('#title').val(),
            content: $('#content').val()
        }
    }).done(function () {
        alert("Publicado com sucesso!");
        window.location = "/home";
    }).fail(function () {
        alert("Erro ao publicar!");
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
        alert("Erro ao curtir!")
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
        alert("Erro ao descurtir!")
    }).always(function () {
        heart.prop('disabled', false);
    });
}
