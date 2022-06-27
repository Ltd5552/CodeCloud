var userID = location.search.match(new RegExp("userID=(.*)&"))[1];
var codeID = location.search.match(new RegExp("codeID=(.*)"))[1];
// alert(userID)
// alert(codeID)
$(function () {

    $.get("http://localhost/user/findName", {'userID': userID}, function (data) {
        //{uid:1,name:'李四'}
        var msg = "欢迎回来，" + data.UserName;
        $("#span_username").html(msg);
    });
    // });
        if(codeID!="none")
        {
            $.get("/code/runAgain", {'codeID': codeID}, function (data) {
                $("#input1").html(data.Code);
            });
        }
        $.get("/code/history", {'userID': userID}, function (data) {
            var his = ''
            for(var i=1;i<=data.CodeID.length;i++){
                var item = "<el-menu-item class='history_item' index='1-" +i+ "1' codeID='"+ data.codeID[i-1]+"'>" +
                    "<a href='http://localhost/CodeCloud/views/history.html?userID="+userID+"&codeID="+data.codeID[i-1]+"'"+data.runTime[i-1]+"</a></el-menu-item>";
                his += item;
            }
            $("#history_group").html(his)
        });

    $('#runCode').click(function () {
        var code = $("#input1").val();
        var type = $("#langSelect").val();
        var date = new Date();
        var runTime = date.getTime();
        $.post("/code/run", {'code': code, 'type': type, 'userID': userID, 'runTime': runTime}, function (data) {
            $('#resultInfo').html(data.result)
        })
    })

});