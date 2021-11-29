// Package consts @Description  TODO
// @Author  	 jiangyang
// @Created  	 2021/11/28 1:56 下午
package consts

const VerificationCodeTpl=`<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title></title>
    <style>
        body {
            margin: 0;
            font-family: "lucida Grande", Verdana, "Microsoft YaHei";
            font-size: 12px;
            background: #fff;
            color: #555555;
        }

        a {
            color: #4c83f5;
            text-decoration: none;
        }

        div {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }

        .mail {
            width: 600px;
            margin: 0 auto;
            margin-top: 15px;
            background: #fff;
            box-shadow: 0 0 25px rgb(234 234 234);
            -webkit-box-shadow: 0 0 25px rgb(234 234 234);
        }

        .mail .top {
            width: 100%%;
            height: 20px;
            background-color: #4c83f5;
        }

        .mail .header {
            margin: 0 38px;
        }

        .mail .header .logo {
            height: 86px;
            line-height: 86px;
            position: relative;
        }

        .mail .header .logo img {
            width: 100px;
            position: absolute;
            top: 0px;
            left: -17px;
        }

        .mail .header .hello p {
            line-height: 25px;
        }

        .mail .line {
            width: 100%%;
            height: 1px;
            border-top: solid #e6e6e6 1px;
        }

        .mail .footer {
            margin: 0 38px;
            padding-bottom: 8px;
            border-top: 1px solid #e6e6e6;
        }
    </style>
</head>

<body>
    <div class="mail" style="font-family: Microsoft YaHei;
    font-size: 12px;
    background: #fff;
    color: #555555;
    box-sizing: border-box;
    width: 600px;
    margin: 0 auto;
    margin-top: 15px;
    background: #fff;
    box-shadow: 0 0 25px rgb(234 234 234);
    -webkit-box-shadow: 0 0 25px rgb(234 234 234);
    border: 1px solid #f5f5f5">
        <div class="top" style="width: 100%%;
        height: 20px;
        background-color: #4c83f5;"></div>
        <div class="header" style="margin: 0 38px;">
            <div class="logo" style="height: 86px;
            line-height: 86px;
            position: relative;">
                <a href="https://www.shixiseng.com/" target="_blank" style="color: #4c83f5;
                text-decoration: none;">
                    <img src="https://tva1.sinaimg.cn/large/0060QYYfly1gwutrb3rfwj30rs0rs3z0.jpg" alt="" style="width: 100px;
                    position: absolute;
                    top: 0px;
                    left: -17px;">
                </a>
            </div>
            <div class="hello" style="margin-top: 15px;">
                <p style="line-height: 25px;">您好：<br />    <span style="margin-left: 20px;">您的验证码为: {{code}}</span></p>
                
            </div>
        </div>
        <div class="footer" style="text-align: right;margin: 0 38px;
        padding-bottom: 8px;
        border-top: 1px solid #e6e6e6;">
            <p> xxx</p>
        </div>
    </div>
</body>

</html>`

