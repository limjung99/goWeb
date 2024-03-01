$(function() {
    if (!window.EventSource){
        alert("Noe Event Source!")
        return
    };

    var $chatlog = $(`#chat-log`);
    var $chatmsg = $(`#chat-msg`);

    var isBlank = (string) => {
        return string == null || string.trim() === "";
    };
    var username;
    while(isBlank(username)){
        username = prompt("What's ur name?");
        if(!isBlank(username)) {
            $(`#user-name`).html('<b>'+username+'</b>');
        }
    }

    $(`#input-form`).on('submit',(e)=>{
        $.post('/messages',{
            msg : $chatmsg.val(),
            name : username
        });
        $chatmsg.val("");
        $chatmsg.focus();
        return false;
    })

    var addMessage = (data) => {
        var text = "";
        if(!isBlank(data.name)) {
            text = '<string>' + data.name + ':</string>';
        }
        text += data.msg;
        $chatlog.prepend('<div><span>'+text+'</span></div>');

    }

    var es = new EventSource('/stream');
    es.onopen = (e) => {
        $.post('/users' , {
            name : username
        })
    }
    es.onmessage = (e) => {
        var msg = JSON.parse(e.data)
        addMessage(msg)
    }
})