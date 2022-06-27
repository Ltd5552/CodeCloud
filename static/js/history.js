var userID = location.search.match(new RegExp("userID=(.*)&"))[1];
var codeID = location.search.match(new RegExp("codeID=(.*)"))[1];
alert(userID)
alert(codeID)
// $(function () {
// /*
    $.get("/user/findName", {}, function (data) {
        //{uid:1,name:'李四'}
        var msg = "欢迎回来，" + data.username;
        $("#span_username").html(msg);
    });

    $.get("/code/detail", {'codeID': codeID}, function (data) {
        //{result:text, code:text, type:String, runTime:time}
        $("#input1").html(data.code);
        $("#resultInfo").html(data.result);
        $("#type").html("&nbsp;&nbsp;代码的语言为："+data.type+"；运行的时间为："+data.runTime+"&nbsp;&nbsp;")
    });
// */
    $("#edit").click(function () {
        // location.href = "../../views/runner.html";
        location.href = "http://localhost:63342/CodeCloud/views/history.html?userID="+userID+"&codeID="+codeID;
    });

    $("#delete").click(function (){
        $.delete("/code/detail", {'codeID':codeID},);
    });
// });