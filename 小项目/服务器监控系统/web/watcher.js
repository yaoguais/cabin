(function($,_window){

    var Watcher = {
        url: null,
        listener: null,
        infoShower: null,
        commandLine: null,
        chartManager: null,
        auth: false,
        background: '#000000',
        color: '#00e100',
        getWidth: function(){ return $(_window).width() / 4;},
        getHeight: function(){return $(_window).height() / 3;},
        init: function(options){
            this.infoShower = new InfoShower();
            this.infoShower.init();
            this.commandLine = new CommandLine();
            this.commandLine.init();
            this.chartManager = new ChartManager();
            if(options && typeof options == 'object'){
                for(var attr in options){
                    this[attr] = options[attr];
                }
            }
            $('body').css({"background":this.background,"color":this.color});
            return this;
        },
        start: function(){
            this.listener = new Listener(this.url);
            this.listener.listening();
        }
    };
    function WatcherProtocol(command, code, message, data){

        this.command = command || WatcherProtocol.COMMAND_UNKNOWN;
        this.code = code || WatcherProtocol.ERROR_NOT_EXISTS;
        this.message = message || null;
        this.data = data || {};
    }
    WatcherProtocol.COMMAND_UNKNOWN = 0;
    WatcherProtocol.COMMAND_GET        = 100;
    WatcherProtocol.COMMAND_PUSH       = 101;
    WatcherProtocol.COMMAND_AUTH       = 102;
    WatcherProtocol.COMMAND_NEED_AUTH  = 103;
    WatcherProtocol.COMMAND_RE_FLOW    = 104;
    WatcherProtocol.COMMAND_FLOW       = 105;
    WatcherProtocol.COMMAND_HELP       = 199;
    WatcherProtocol.COMMAND_DISK       = 200;
    WatcherProtocol.ERROR_NOT_EXISTS   = 0;
    WatcherProtocol.firstGet           = true;
    WatcherProtocol.prototype.getAuthMessage = function(username,password){
        this.command = WatcherProtocol.COMMAND_AUTH;
        this.data.username = username;
        this.data.password = password;
        return this.serialize();
    };

    WatcherProtocol.prototype.serialize = function(){

        return JSON.stringify(this);
    };

    WatcherProtocol.prototype.unSerialize = function(raw){

        var obj = null;
        try{
            obj = JSON.parse(raw);
            this.command = obj.command;
            this.code = obj.code;
            this.message = obj.message;
            this.data = obj.data;
            this.resolveCodeMessage();
            return true;
        }catch (e){
            return false;
        }
    };
    WatcherProtocol.prototype.resolveCodeMessage = function(){

        if(this.code == WatcherProtocol.ERROR_NOT_EXISTS){
            return;
        }
        if(!this.message){
            this.message = 'unknown code, maybe you should update your client,see github.com/yaoguais/watcher';
        }
    };


    function Listener(url) {

        try{
            this.ws = new WebSocket(url);
            this.ws.onopen = this.open;
            this.ws.onmessage = this.message;
            this.ws.onclose = this.close;
            this.ws.onerror = this.error;
        }catch (e){

        }
        this.auth = Watcher.auth;
    }
    Listener.prototype.listening = function(){};
    Listener.prototype.send = function(data){
        this.ws.send(data);
        console.log(data);
    };
    Listener.prototype.open = function(evt){

    };
    Listener.prototype.message = function(evt){
        console.log(evt);
        var protocol = new WatcherProtocol();
        if(protocol.unSerialize(evt.data)){
            if(protocol.code != WatcherProtocol.ERROR_NOT_EXISTS){
                Watcher.infoShower.show(protocol.message);
            }else{
                switch(protocol.command){
                    case WatcherProtocol.COMMAND_DISK:
                        Disk.addDisk(protocol.data.id,protocol.data.info);
                        break;
                    case WatcherProtocol.COMMAND_GET:
                        if(WatcherProtocol.firstGet){
                            Watcher.listener.send((new WatcherProtocol(WatcherProtocol.COMMAND_RE_FLOW)).serialize());
                        }
                    case WatcherProtocol.COMMAND_PUSH:
                    case WatcherProtocol.COMMAND_FLOW:
                        for(var id in protocol.data){
                            var chart = new Chart();
                            chart.id  = id;
                            chart.data = protocol.data[id];
                            chart.type = protocol.data[id]['type'];
                            Watcher.chartManager.handle(chart);
                        }
                        break;
                    case WatcherProtocol.COMMAND_NEED_AUTH:
                        if(Watcher.auth){
                            if($.cookie('watcher_username') && $.cookie('watcher_password')){
                                Watcher.listener.send((new WatcherProtocol()).getAuthMessage($.cookie('watcher_username'), $.cookie('watcher_password')));
                            }else{
                                Watcher.commandLine.show();
                            }
                        }else{
                            Watcher.listener.send((new WatcherProtocol()).getAuthMessage('',''));
                        }
                        break;
                    case WatcherProtocol.COMMAND_AUTH://auth ok,send get
                        Watcher.listener.send((new WatcherProtocol(WatcherProtocol.COMMAND_GET)).serialize());
                        break;
                    case  WatcherProtocol.COMMAND_RE_FLOW:// reflow ok
                        WatcherProtocol.firstGet = false;
                        break;
                }
            }
        }else{
            Watcher.infoShower.show('server message error');
        }
    };
    Listener.prototype.close = function(evt){
        Watcher.infoShower.show('server close your connection');
    };
    Listener.prototype.error = function(evt){

        Watcher.infoShower.show('error occur, connect retry after 5 second');
        setTimeout(function(){
            Watcher.start();
        }, 5000);
    };
    function Disk(){}
    Disk.handle = null;
    Disk.addDisk = function(id,info){
        if(!Disk.handle){
            $("body").append('<div style="clear: both" id="diskContainer"></div>');
            Disk.handle = $('#diskContainer');
        }
        if($('#disk_' + id).length == 0){
            Disk.handle.append('<div style="margin: 15px" id="disk_' + id + '"></div>');
        }

        if(typeof info == "string"){
            var html = '<ul>',rows = info.split(";");
            for(var i in rows){
                if(rows[i].length != 0){
                    html += '<li>' + rows[i] + '</li>';
                }
            }
            html += '</ul>';
            $('#disk_' + id).html(html);
        }
    };

    function Chart(){
        this.id = null;
        this.type = null;
        this.data = {};
    }

    function ChartManager(){
        this.handler = null;
        this.handlers = [];
    }
    ChartManager.prototype.resize = function(){
        var w = Watcher.getWidth();
        var h = Watcher.getHeight();
        for(var i in this.handlers){
            for(var j in this.handlers[i]){
                var handle = this.handlers[i][j];
                handle.setSize(w,h);
            }
        }
    };
    ChartManager.prototype.createDefault = function(chart){

        var exists = false;
        for(var k in this.handlers){
            if(k == chart.id){
                exists = true;
            }
        }
        if(!exists){
            if(!this.handler){
                $("body").append('<div id="chartContainer"></div>');
                this.handler = $('#chartContainer');
            }
            var attr = 'chart' + chart.type.substring(0,1).toUpperCase()+chart.type.substring(1) + 'Types';
            if(typeof Watcher[attr] != "undefined" && typeof chart.data.list[0] != "undefined"){
                this.handlers[chart.id] = [];
                var data = chart.data.list[0];
                for(var i in data){
                    if(typeof Watcher[attr][i] != "undefined"){
                        var info = Watcher[attr][i];
                        var d = 'chart_' + chart.id + '_' + info.name.replace(' ','_');
                        this.handler.append('<div style="float: left" id="' + d +'"></div>');
                        var handle = new Highcharts.Chart({
                            chart: {
                                renderTo : d,
                                type: 'spline',
                                plotBackgroundColor: '#000000',
                                backgroundColor: '#000000',
                                animation: Highcharts.svg,
                                marginRight: 0,
                                events : {
                                    click: function(){
                                        if(this.chartWidth <= Watcher.getWidth()){
                                            $('#' + this.options.chart.renderTo).prependTo(Watcher.chartManager.handler);
                                            this.setSize($(_window).width(), Watcher.getHeight());
                                            this.reflow();
                                        }else{
                                            Watcher.chartManager.resize();
                                        }
                                    }
                                }
                            },
                            colors: ['#00e100'],
                            title: {
                                text: info.title,
                                style: { "color": "#00e100"}
                            },
                            xAxis: {
                                type: 'datetime',
                                tickPixelInterval: 150
                            },
                            yAxis: {
                                title: {
                                    text: info.y,
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
                                name: info.name,
                                lineWidth: 0.8,
                                data: null
                            }]
                        });
                        this.handlers[chart.id][i] = handle;
                    }
                }
            }
            this.resize();
            return true;
        }else{
            return false;
        }
    };

    ChartManager.prototype.parseToXY = function(d){
        var list = [];
        for (var i in d) {
            for (var j = 1, l = d[i].length; j < l; ++j) {
                if (typeof list[j] == 'undefined') {
                    list[j] = [];
                }
                list[j].push({x: d[i][0] * 1000, y: d[i][j]});
            }
        }
        return list;
    };

    ChartManager.prototype.handle = function(chart){
        this.createDefault(chart);
        var d = this.parseToXY(chart.data.list);
        for(var i in this.handlers[chart.id]){
            var handle = this.handlers[chart.id][i];
            var justOne = d[i].length == 1;
            for(var j in d[i]){
                handle.series[0].addPoint(d[i][j], justOne, handle.series[0].length >= 100);
            }
            if(!justOne){
                handle.setSize(handle.chartWidth, handle.chartHeight);
            }
        }
    };
    function InfoShower(){
        this.messages = [];
        this.intervalId = null;
        this.handler = null;
        return this;
    }
    InfoShower.prototype.init = function(){
        $('body').append('<div id="infoShower" style="width: 100%;height: 100px;vertical-align:middle;line-height:100px;text-align: center;background: #000000;z-index: 99999999;display: none"><span style="font-size: 32px;font-weight:700; color: #FF0000"></span></div>');
        this.handler = $('#infoShower');
        this.handler.css("opacity",0.8);
    };
    InfoShower.prototype.show = function(message){
        this.messages.push(message);
        if(null === this.intervalId){
            this.pop();
            this.intervalId = setInterval(this.pop.bind(this), 2000);
        }
    };
    InfoShower.prototype.pop = function(){
        var message = this.messages.shift();
        if(message){
            var h = $(_window).height(), w = $(_window).width();
            var span = this.handler.find('span');
            this.handler.css({width:w+'px'});
            span.text(message);
            this.handler.show();
        }else{
            this.handler.hide();
        }
    };

    function CommandLine(){

        this.handler = null;
        this.input = null;
        this.output = null;
        return this;
    }
    CommandLine.prototype.init = function(){

        $('body').append('<div id="commandLine" style="width: 600px;height: 240px;position: absolute;top: 1px;z-index: 99999998;overflow: hidden;display: none"><div><input type="text" id="commandLineInput" style="border: 0;text-indent: 4px;font-weight: 700" /></div><div id="commandLineOutput" style="height: 220px;font-size:12px;overflow-y: auto;margin-left: 4px"></div></div>');
        this.handler = $('#commandLine'), this.input = $('#commandLineInput'), this.output = $('#commandLineOutput');
        this.input.css({"background":Watcher.background,"color":Watcher.color});
        this.input.width(this.handler.width()), this.output.width(this.handler.width()-4);
        this.input.keydown(this.keydown.bind(this));
        this.output.click(this.click.bind(this));
        $('body').keydown(function(e){
            if(e.keyCode == 13){
                this.show();
            }
        }.bind(this));
        //this.show();
    };
    CommandLine.prototype.click = function(e){
        this.handler.hide();
    };
    CommandLine.prototype.show = function(){
        this.handler.css({"left":($(_window).width()-this.handler.width())/2});
        if(this.input.val() == ''){
            this.input.val('Input help to get more details, RETURN to commit');
        }
        this.input[0].focus();
        this.handler.show();
    };
    CommandLine.prototype.hide = function(){
        this.handler.hide();
    };
    CommandLine.prototype.notice = function(data){
        this.output.html(data);
    };
    CommandLine.prototype.addNotice = function(data){
        this.output.html(this.output.html() + '<br>' +data);
    };
    CommandLine.prototype.keydown = function(e){

        if(e.keyCode == 13){
            var val = this.input.val();
            if(val == ''){
                Watcher.infoShower.show('command can not be empty');
            }else{
                var protocol = this.getProtocolFromCommand(val);
                if(protocol instanceof WatcherProtocol){
                    Watcher.listener.ws.send(protocol.serialize());
                }
            }
        }
    };
    CommandLine.prototype.getProtocolFromCommand = function(line){

        var arg, args = line.split(' ');
        var protocol = new WatcherProtocol();
        while(args.length > 0){
            arg = args.shift();
            switch (arg.toLowerCase()){
                case 'auth':
                    protocol.command = WatcherProtocol.COMMAND_AUTH;
                    arg = args.shift();
                    if(arg){
                        protocol.data['username'] = arg.trim();
                        arg = args.shift();
                        if(arg){
                            protocol.data['password'] = md5(arg.trim());
                            $.cookie('watcher_username', protocol.data['username']);
                            $.cookie('watcher_password', protocol.data['password']);
                            return protocol;
                        }
                    }
                    //break
                default :
                    this.notice('usage:<br>auth username password<br>help<br>');
                    return null;
            }
        }
        return protocol;
    };


    function Queue(id, size){
        this.id = id;
        this.size = size;
    }
    function QueueManager(){
        this.queueList = {};
    }
    QueueManager.prototype.isIdExists = function(id){
        for(var k in this.queueList){
            if(k == id){
                return true;
            }
        }
        return false;
    };
    QueueManager.prototype.add = function(queue){

        if(this.isIdExists(queue.id)){
            return false;
        }
        this.queueList[queue.id] = queue;
        return true;
    };
    QueueManager.prototype.remove = function(queue){
        for(var k in this.queueList){

        }
    };
    _window.Watcher = Watcher;
})($,window);