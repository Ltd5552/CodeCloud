var userID = location.search.match(new RegExp("userID=(.*)&"))[1];
var codeID = location.search.match(new RegExp("codeID=(.*)"))[1];
// alert(userID)
// alert(codeID)
$(function () {

    $.get("http://localhost:8088/user/findName", {'userID': userID}, function (data) {
        //{uid:1,name:'李四'}
        var msg = "欢迎回来，" + data.UserName;
        $("#span_username").html(msg);
    });
    // });
    if(codeID!="none")
    {
        $.get("http://localhost:8088/code/runAgain", {'codeID': codeID}, function (data) {
            $("#input1").html(data.Code);
        });
    }

    $.get("http://localhost:8088/code/history", {'userID': userID}, function (data) {
        var his = '';
        for(var i = 0;i < data.Lists.length;i++){
            var one = data.Lists[i];
            var item = '<el-menu-item class="history_item" index="1-'+ i+1 +'" codeID='+one.CodeID+'> <a href="http://localhost:8088/code/detail?codeID='+one.CodeID+'&userID='+userID+'">'+one.RunTime+'</a></el-menu-item><br>';
            his += item;
        }
        $("#history_group").html(his)
    });

    $('#runCode').click(function () {
        var code = $("#input1").val();
        var type = $("#langSelect").val();

        //添加0
        function getNow(s) {
            return s < 10 ? '0' + s: s;
        }
        //获取时间戳
        var myDate = new Date();
        var year=myDate.getFullYear();        //获取当前年
        var month=myDate.getMonth()+1;   //获取当前月
        var date=myDate.getDate();            //获取当前日

        var h=myDate.getHours();              //获取当前小时数(0-23)
        var m=myDate.getMinutes();          //获取当前分钟数(0-59)
        var s=myDate.getSeconds();

        var runTime=year+'-'+getNow(month)+"-"+getNow(date)+" "+getNow(h)+':'+getNow(m)+":"+getNow(s);
        //请求
        $.post("/code/run", {'code': code, 'type': type, 'userID': userID, 'runTime': runTime}, function (data) {
            $('#resultInfo').html(data.output+data.errors)
        })
    })

});