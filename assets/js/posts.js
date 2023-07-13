$('#newpost').on('submit', createPost);

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
