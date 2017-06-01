<?php

if(!isset($_GET['secret']) && $_GET['secret'] != 'jI10L39OS5)'){
    header("HTTP/1.1 404 Not Found");
    header("Status: 404 Not Found");
    exit;
}
//$scheme = isset( $_SERVER['HTTPS'] ) && 'on' == $_SERVER['HTTPS'] ? 'wss' : 'ws';
$scheme = 'ws';
$ws = $scheme.'://'.$_SERVER['SERVER_ADDR'].':9505';
?>
<!DOCTYPE HTML>
<html>
<head>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <title>服务器负载监控</title>

    <script type="text/javascript" src="js/jquery-2.1.4.js"></script>
    <script src="js/highcharts.js"></script>
    <script src="js/exporting.js"></script>
    <style type="text/css">${demo.css}</style>
    <style type="text/css">
        *{margin: 0; padding: 0;}
        body{
            background-color: black;
        }
        .cpu,.memory{
            float: left;
            border: 1px solid #000000;
        }
        rect{
            fill: black;
        }
    </style>
    <script>
        var globalData = {
            api_cpu : [],
            api_memory : [],
            mysql_cpu : [],
            mysql_memory : [],
            mongo_cpu : [],
            mongo_memory : [],
            chat_cpu : [],
            chat_memory : [],
            api_request: []
        };
        var processIntervalHandle = null;
        var wsServer = '<?=$ws?>';
        var websocket = new WebSocket(wsServer);
        websocket.onopen = function (evt) {

        };

        websocket.onclose = function (evt) {

        };

        websocket.onmessage = function (evt) {
            var data = null;
            try{
                data = JSON.parse(evt.data);
            }catch(e){
                return;
            }
            for(var machine in data){
                for(var type in data[machine]){
                    var k = machine + '_' + type;
                    for(var i in data[machine][type]){
                        var d = {
                            x : data[machine][type][i].x * 1000,
                            y : data[machine][type][i].y
                        }
                        globalData[k].push(d);
                    }
                }
            }
            call_high_charts('#container0', 'api_cpu','API服务器CPU使用率','CPU(%)','CPU使用率');
            call_high_charts('#container00','api_memory','API服务器内存占用情况(4G)','Memory(MB)','内存占用');
            call_high_charts('#container1','mysql_cpu','MYSQL服务器CPU使用率','CPU(%)','CPU使用率');
            call_high_charts('#container01','mysql_memory','MYSQL服务器内存占用情况(8G)','Memory(MB)','内存占用');
            call_high_charts('#container2','mongo_cpu','MONGO服务器CPU使用率','CPU(%)','CPU使用率');
            call_high_charts('#container02','mongo_memory','MONGO服务器内存占用情况(4G)','Memory(MB)','内存占用');
            call_high_charts('#container3','chat_cpu','聊天服务器CPU使用率','CPU(%)','CPU使用率');
            call_high_charts('#container03','chat_memory','聊天服务器内存占用情况(1G)','Memory(MB)','内存占用');

            call_high_charts('#container4','api_request','帖子列表API请求时间','Time(ms)','请求时间');
            if(processIntervalHandle != null){
                $('.process').remove();
                clearInterval(processIntervalHandle);
                processIntervalHandle = null;
            }
        };

        Highcharts.setOptions({
            global: {
                useUTC: false
            },
            colors:['#00e100']
        });

        var init = [];
        function call_high_charts(id,type,title,yName,noticeName){
            if(typeof init[id] != 'undefined'){
                return;
            }
            init[id] = true;
            $(id).highcharts({
                chart: {
                    type: 'spline',
                    plotBackgroundColor: '#000500',
                    animation: Highcharts.svg,
                    marginRight: 0,
                    events: {
                        load: function () {
                            var series = this.series[0];
                            setInterval(function () {
                                var d = globalData[type].shift()
                                if(d){
                                    if(type.indexOf('memory') >= 0){
                                        d.y /= 1024;
                                    }
                                    series.addPoint([d.x, d.y], true, true);
                                }
                            }, 1000);
                        }
                    }
                },
                title: {
                    text: title,
                    style: { "color": "#00e100"}
                },
                xAxis: {
                    type: 'datetime',
                    tickPixelInterval: 150
                },
                yAxis: {
                    title: {
                        text: yName,
                        style: {
                            color: '#00e100'
                        }
                    },
                    plotLines: [{
                        value: 0,
                        width: 1,
                        color: '#00e100'
                    }]
                },
                tooltip: {
                    formatter: function () {
                        return '<b>' + this.series.name + '</b><br/>' +
                            Highcharts.dateFormat('%Y-%m-%d %H:%M:%S', this.x) + '<br/>' +
                            Highcharts.numberFormat(this.y, 2);
                    }
                },
                legend: {
                    enabled: false
                },
                exporting: {
                    enabled: false
                },
                series: [{
                    name: noticeName,
                    lineWidth: 0.8,
                    data: (function () {

                        var data =  globalData[type];
                        globalData[type] = [];
                        return data;
                    }())
                }]
            });
        }

        websocket.onerror = function (evt, e) {
            console.log('Error occured: ' , evt, e);
        };
    </script>
    <script type="text/javascript">
        $(function () {
            $(document).ready(function () {

            });
        });
    </script>
</head>
<body>
<hr class="process" style="width: 0; color: #00e100" />

<div id="container0" class="cpu e" style="min-width: 310px; margin: 0 auto"></div>
<div id="container1" class="cpu e" style="min-width: 310px; margin: 0 auto"></div>
<div id="container2" class="cpu e" style="min-width: 310px; margin: 0 auto"></div>
<div id="container3" class="cpu e" style="min-width: 310px; margin: 0 auto"></div>

<div id="container00" class="memory e" style="min-width: 310px; margin: 0 auto"></div>
<div id="container01" class="memory e" style="min-width: 310px; margin: 0 auto"></div>
<div id="container02" class="memory e" style="min-width: 310px; margin: 0 auto"></div>
<div id="container03" class="memory e" style="min-width: 310px; margin: 0 auto"></div>

<div id="container4" style="min-width: 310px; margin: 0"></div>
<script>
    var w = $(window).width();
    var h = $(window).height();
    var mainW = (w - 30) / 4;
    var mainH = (h - 30) / 2;
    $(".cpu").css("width",mainW);
    $(".cpu").css("height",mainH);
    $(".memory").css("width",mainW);
    $(".memory").css("height",mainH);
    $("#container4").css("height",mainH);
    $(".e").click(function(){
        if($(this).hasClass("change")){
            $(this).css("width",mainW);
            $(".e").removeClass("change");
        }else{
            $(".e").removeClass("change");
            $(this).addClass("change");
            $(this).css("width",w).siblings().css("width",mainW);
            $("#container4").css("width",w);
        }
        for(var i in Highcharts.charts){
            Highcharts.charts[i].reflow();
        }
    });
    processIntervalHandle = setInterval(function(){
        var processWidth = $('.process').width() + 10;
        if(processWidth > w - 50){
            processWidth = 10;
        }
        $('.process').width(processWidth);
    },400);
</script>
</body>
</html>
