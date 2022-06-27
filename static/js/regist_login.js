function checkUserName(){
    // alert('用户名校验')
    // 获取用户名
    var username = $("#username").val();
    // 定义正则表达式
    var reg_username = /^\w{8,16}$/;
    // 判断
    var flag = reg_username.test(username);
    if(flag){
        // 校验通过， 用户名是合法的
        $("#username").css('border', '')
    }
    else{
        // 校验失败
        $("#username").css('border', '1px solid red')
    }
    return flag;
}
// 2. 密码：单词字符，8-20位
function checkPassword(){
    // alert('Password校验')
    // 获取密码
    var password = $("#password").val();
    // 定义正则表达式
    var reg_password = /^\w{8,16}$/;
    // 判断
    var flag = reg_password.test(password);
    if(flag){
        // 校验通过， 密码是合法的
        $("#password").css('border', '')
    }
    else{
        // 校验失败
        $("#password").css('border', '1px solid red')
    }
    return flag;
}
$(function () {
    //1.给登录按钮绑定单击事件
    $("#btn_sub").click(function () {
        var myflag = true;

        if($(".regin-type").attr('va')==0) {
            if(!(checkPassword() && checkUserName()))
                myflag = false;
        } else {
            myflag = true;
        }

        if(myflag){
           // 2.发送ajax请求，提交表单数据
           //alert($("#loginForm").serialize()+'&type='+$(".regin-type").attr('va'))
            $.post("/user/login",$("#loginForm").serialize()+'&type='+$(".regin-type").attr('va'),function (data) {
                //data : {flag:false,msg:'',userID:''}
                // alert(data.UserID)
                if(data.Flag){
               // if(true){
                    //登录成功，跳转到执行代码界面
                    location.href="http://localhost/code/run?userID="+data.UserID+"&codeID=none";

                }else{
                    // //登录失败，给出提示信息
                    $("#errorMsg").removeProp("hidden");

                    $('#errorMsg').html(data.Msg);
                }
             });
        }
    });
    $(".regin-type").click(function () {
        if($(this).attr('va')==1){
            // $.delay(1)
            $(this).html('已有账号？现在登录！')
            $(this).attr('va',0)
            $('#btn_sub').html('注册')
            $('#username').prop('placeholder','8-16位字符')
            $('#password').prop('placeholder','8-16位字符')
            $("#username").css('border', '')
            $("#password").css('border', '')
        } else {
            // $.delay(1)
            $(this).html('没有账号？现在注册！')
            $(this).attr('va',1)
            $('#btn_sub').html('登录')
            $('#username').prop('placeholder','用户名')
            $('#password').prop('placeholder','密码')
            $("#username").css('border', '')
            $("#password").css('border', '')
        }
    });
    $("#username").blur(checkUserName||$(".regin-type").attr('va'))
    $("#password").blur(checkPassword||$(".regin-type").attr('va'))
});