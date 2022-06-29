var userID = location.search.match(new RegExp("userID=(.*)"))[1];
var codeID = location.search.match(new RegExp("codeID=(.*)&"))[1];
$(function () {
    $.get("/user/findName", {'userID':userID}, function (data) {
        //{uid:1,name:'李四'}
        var msg = "欢迎回来，" + data.UserName;
        $("#span_username").html(msg);
    });

    $.post("/code/detail", {'codeID': codeID}, function (data) {
        //{result:text, code:text, type:String, runTime:time}
        $("#input1").html(data.Code);
        $("#resultInfo").html(data.Result);
        $("#type").html("&nbsp;&nbsp;代码的语言为："+data.Type+"；运行的时间为："+data.RunTime+"&nbsp;&nbsp;")
    });
    $("#edit").click(function () {
        location.href = "http://localhost:8088/code/run?userID="+userID+"&codeID="+codeID;
    });
    $("#delete").click(function (){
        $.ajax({
            url: "http://localhost:8088/code/detail?codeID="+codeID+"",
            type: 'DELETE',
            // data:{'codeID':codeID},
            success: function(data) {
                var flag = data.Flag
                if (flag){
                    alert("删除成功")
                    location.href = "http://localhost:8088/code/run?userID="+userID+"&codeID=";
                }
                else {
                    alert("删除失败，请稍后尝试")
                }
            }
        });

    });
});