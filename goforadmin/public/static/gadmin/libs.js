toastr.options = {
    "autoDismiss": false,
    "positionClass": "toast-top-full-width",
    "type": "error",
    "timeOut": "5000",
    "extendedTimeOut": "2000",
    "allowHtml": false,
    "closeButton": false,
    "tapToDismiss": false,
    "progressBar": true,
    "newestOnTop": false,
    "maxOpened": 0,
    "preventDuplicates": false,
    "preventOpenDuplicates": false
};

$.notify.addStyle("bootstrap", {
    html: "<div>\n<span data-notify-text></span>\n</div>",
    classes: {
        base: {
            "padding": "4px 8px 4px 4px",
            "text-shadow": "0 1px 0 rgba(255, 255, 255, 0.5)",
            "background-color": "#fcf8e3",
            "border": "1px solid #fbeed5",
            "border-radius": "4px",
            "white-space": "nowrap",
            "background-repeat": "no-repeat",
        },
        error: {
            "color": "#ffffff",
            "background-color": "rgba(232, 86, 86, 0.85)",
            "border-color": "rgba(232, 86, 86, 0.85)",
        }
    }
});

$.notify.addStyle("login", {
    html: "<div>\n<span data-notify-text></span>\n</div>",
    classes: {
        base: {
            "padding": "4px 8px 4px 4px",
            "text-shadow": "0 1px 0 rgba(255, 255, 255, 0.5)",
            "background-color": "#fcf8e3",
            "border": "1px solid #fbeed5",
            "border-radius": "4px",
            "white-space": "nowrap",
            "background-repeat": "no-repeat",
        },
        error: {
            "color": "#000000",
            "background-color": "rgba(124, 144, 130, 0.85)",
            "border-color": "rgba(124, 144, 130, 0.85)",
        }
    }
});

$(document).ready(function () {
    // 调整内容区块的高度
    $(window).on('resize', function () {
        $('.al-main').css("min-height", $(window).height() + "px");
        if ($(window).width() < 1194) {
            if (!$('body>main').hasClass('menu-collapsed')) {
                $('body>main').addClass('menu-collapsed')
            }
        }
    });
    $(window).resize();
    // 输入框验证的逻辑
    app_bind_form_validate('blur');
    // 初始化选择框
    $('select').selectpicker();
    // 初始化上传框
    app_upload();
    // 初始化照片弹出框
    $(document).on('click', '.pop-img',function(){
        app_modal({
            type: "info",
            message: '<a href="' + $(this).attr('src') + '" target="_blank"> <img style="width:90%" src="' +
                        $(this).attr('src') + '" /></a>'
        });
    });
    // 初始化日期选择按钮
    $('.datetime').datetimepicker({language:'zh-CN'});
    $('.daterange').each(function(){
        var $target = $(this).find('input');
        var format = $target.attr('date-format') || "YYYY-MM-DD";
        app_daterange($target[0], {format: format});
        $(this).find('i').click(function(){
            $target.data('daterangepicker').show();
        });
    });

    // 添加按钮的逻辑
    $('.btn-add').click(function () {
        $('#add').toggle();
    });
    // 初始化搜索框
    $('#searchInput').keydown(function (e) {
        if (e.keyCode == 13) {
            app_toast_error("未找到匹配的搜索结果!");
            var self = this;
            setTimeout(function () {
                $(self).val('')
            }, 3000);
        }
    });
    // 初始化菜单的效果
    restore_menu();
    $('.collapse-menu-link').click(function () {
        if ($('body>main').hasClass('menu-collapsed')) {
            $('body>main').removeClass('menu-collapsed')
        } else {
            $('body>main').addClass('menu-collapsed')
        }
    });
    $('.al-sidebar').mousemove(function (e) {
        $('.sidebar-hover-elem').css("top", e.clientY - $('.page-top').height() - 21);
    }).mouseleave(function () {
        $('.sidebar-hover-elem').css("top", -150);
    });
    function save_menu() {
        sessionStorage.setItem('menu', $('.al-sidebar').html())
    }

    function restore_menu() {
        var menu = sessionStorage.getItem('menu')
        if (menu && menu.length > 0 && (menu.split('<li').length - 1 == $('.al-sidebar li').length)) {
            $('.al-sidebar').html(menu);
        }
        $('.al-sidebar').css("min-height", $(window).height() + "px").css('visibility', 'visible');
    }

    $('.al-sidebar-list>li').click(function () {
        var $self = $(this);
        if ($self.hasClass('with-sub-menu')) {
            var $subMenu = $self.find('.al-sidebar-sublist');
            if (!$subMenu.attr('data-height')) {
                $subMenu.attr('data-height', $subMenu.find("li").length * 29);
                $subMenu.height(0).removeClass('hidden');
            }
            if ($self.hasClass('ba-sidebar-item-expanded')) {
                $self.removeClass('ba-sidebar-item-expanded');
                $subMenu.animate({height: "0", overflow: "hidden"}, "slow", 'swing', 'save_menu');
            } else {
                $self.addClass('ba-sidebar-item-expanded');
                $subMenu.animate({height: $subMenu.attr('data-height'), overflow: "hidden"}, "slow", 'swing', 'save_menu')
            }
        } else {
            $('.al-sidebar-list li').removeClass('selected');
            $(this).addClass('selected');
            save_menu();
        }
    });
    $('.al-sidebar-list>li li').click(function () {
        $('.al-sidebar-list li').removeClass('selected');
        $(this).addClass('selected');
        save_menu();
    });
    // 添加.form选择器的ajax统一提交的逻辑
    $('.form').find('button.btn-submit').click(function () {
        var self = this;
        var $form = $(self).parents(".form");
        var err = app_form_validate_error($form[0]);
        if (err) {
            return app_toast_error(err);
        }


        app_ajax({
            currentTarget: self,
            url: $form.attr('action'),
            method: $form.attr('method') || "POST",
            data: $form.serialize(),
            success: function (data) {
                if (app_parse_response(data)) {
                    app_toast_success(data.msg);
                    setTimeout(function () {
                        window.location.reload();
                    }, 1000)
                }
            }
        });

        return false;
    });
    // 全选按钮的逻辑
    $(document).on('click', '.table thead tr:eq(0) th:eq(0) input', function () {
        var checked = $(this).prop('checked');
        $(this).parents('.table').find('tbody tr').each(function () {
            $(this).find('td:eq(0) input').prop('checked', checked);
        });
    });
    // 全选提交按钮的逻辑
    $('.btn-submit-checked').click(function () {
        var self = this;
        var data = "";
        var selector = $(self).attr('data-selector') || ".table";
        $(selector).find('tbody tr').each(function () {
            var $c = $(this).find('td:eq(0) input');
            if ($c.prop('checked')) {
                data += "&" + $c.attr('name') + "=" + $c.val();
            }
        });

        if (data == "") {
            return app_toast_error("请选择");
        }

        var ajaxFunc = function () {
            app_ajax({
                currentTarget: self,
                url: $(self).attr('data-url'),
                method: $(self).attr('data-method') || "POST",
                data: data.substr(1),
                success: function (data) {
                    if (app_parse_response(data)) {
                        app_toast_success(data.msg);
                        setTimeout(function () {
                            window.location.reload();
                        }, 1000)
                    }
                }
            });
        };

        app_modal({
            type: "danger",
            message: "确定删除吗?",
            cancel: true,
            ok: ajaxFunc,
        });

        return false;
    });
    // 选择输入框的动态请求逻辑
    $(document).on('keyup', '.bs-searchbox input', function () {
        var self = this;
        var val = $(self).val() || "";
        val = val.replace(/^[\s\uFEFF\xA0]+|[\s\uFEFF\xA0]+$/g, '');
        var $select = $(self).parent().parent('.dropdown-menu').next('select[data-url!=""]');
        if ($select.length > 0) {
            if (val.length == 0) {
                $select.html('');
                $select.selectpicker('refresh');
            } else {
                app_ajax({
                    currentTarget: self,
                    url: $select.attr('data-url'),
                    method: $select.attr('data-method') || "POST",
                    data: "keyword=" + val,
                    success: function (data) {
                        if (app_parse_response(data)) {
                            var list = data.data;
                            var html = "";
                            var i;
                            if (list && list.length > 0) {
                                for (i in list) {
                                    html += '<option value="' + list[i][0] + '">' + list[i][1] + '</option>';
                                }
                            }
                            $select.html(html);
                            $select.selectpicker('refresh');
                        }
                    }
                });
            }
        }
    });
});

function app_ajax(options) {
    var $target = $(options.currentTarget);
    var disableClick = $target.attr('data-disable-click');
    if (disableClick) {
        return;
    }
    $target.attr('data-disable-click', true);

    options.error = options.error || function () {
            app_toast_error("网络错误, 请重试!")
        };
    options.timeout = options.timeout || 30000;
    var cancelDisableClickFunc = function () {
        $target.removeAttr('data-disable-click');
    };

    var complete = options.complete || function () {
        };
    options.complete = function () {
        cancelDisableClickFunc();
        complete();
    };

    $.ajax(options);
}

function app_ajax_with_notice(target, title, url) {
    var ajaxFunc = function () {
        app_ajax({
            currentTarget: target,
            url: url,
            method: "POST",
            data: "",
            success: function (data) {
                if (app_parse_response(data)) {
                    app_toast_success(data.msg);
                    $(target).toggleClass('selected')
                }
            }
        });
    };

    app_modal({
        type: "danger",
        message: title,
        cancel: true,
        ok: ajaxFunc
    });

    return false;
}

function app_parse_response(data) {
    if (typeof data.ret !== "undefined" && data.ret == 0) {
        return true;
    }

    if (typeof data.msg !== "undefined") {
        app_toast_error(data.msg);
    } else {
        app_toast_error("服务器繁忙, 请重试!")
    }

    return false;
}

function app_toast_error(msg) {
    toastr.error(msg);

    return false
}

function app_toast_success(msg) {
    toastr.success(msg);

    return false
}

function app_json_decode(str) {
    str = $('<div/>').html(str).text();
    try {
        return JSON.parse(str)
    } catch (e) {
        console.log(e);
    }
    return null
}

function app_bind_form_validate(event) {
    $(document).on(event, '[data-regexp!=""]', function () {
        var regexp = $(this).attr('data-regexp');
        var msg = $(this).attr('data-msg');
        var options = {};
        try {
            options = JSON.parse($(this).attr('data-options'));
        } catch (e) {
        }
        var val = $(this).val();
        if (typeof msg == "string" && msg.length > 0 && typeof regexp == "string" && regexp.length > 0) {
            var regObj = new RegExp(regexp);
            if (!regObj.test(val)) {
                options["position"] = options["position"] || "right";
                $(this).notify(msg, options);
            }
        }
        return false
    });
}

function app_form_validate_error(selector) {
    var ret = null;
    $(selector).find('[data-regexp!=""]').each(function () {
        var must = $(this).attr('data-must') || "";
        var regexp = $(this).attr('data-regexp');
        var msg = $(this).attr('data-msg');
        var val = $(this).val() || "";
        if ((must == "" || must == "1")
            && typeof msg == "string"
            && msg.length > 0
            && typeof regexp == "string"
            && regexp.length > 0) {
            var regObj = new RegExp(regexp);
            if (!regObj.test(val)) {
                ret = msg;
                return false;
            }
        }
    });

    return ret;
}

function app_data_table(options) {
    var $target = $(options.currentTarget);
    options.processing = options.processing || true;
    options.serverSide = options.serverSide || true;
    options.autoWidth = options.autoWidth || true;

    options.ajax = {
        url: $target.attr('data-url'),
        method: $target.attr('data-method') || "POST",
        data: function (d) {
            d.total = $target.attr('data-total');
        }
    };

    if (!options.columnDefs) {
        options.columnDefs = [];
    }

    if ($target.find("thead th:eq(0) input").length > 0) {
        options.columnDefs.push({"width": "20px", "targets": 0});
    }

    var conf = $target.attr('data-dtconf');
    if (conf && conf.length > 0) {
        conf = app_json_decode(conf);
        if (conf) {
            var i, v;
            var lengthMenu = [[], []];
            for (i in conf.limits) {
                v = conf.limits[i];
                lengthMenu[0].push(v);
                lengthMenu[1].push(v < 0 ? "所有" : v);
            }
            options.lengthMenu = lengthMenu;

            var order = [];
            var orderDisable = [];
            var orderAbleMap = [];
            var size = $target.find('thead tr:eq(0) th').length;
            for (i in conf.columns) {
                v = conf.columns[i];
                if (v.order != "") {
                    order.push([v.index, v.order])
                }
                if (v.sortable) {
                    orderAbleMap[v.index] = true;
                }
            }
            options.order = order;
            for (i = 0; i < size; i++) {
                if (true !== orderAbleMap[i]) {
                    orderDisable.push(i);
                }
            }
            options.columnDefs.push({"orderable": false, "targets": orderDisable});
        }
    }

    $target.on('xhr.dt', function (e, options, json) {
        if (typeof json.recordsTotal != "undefined") {
            $target.attr('data-total', json.recordsTotal);
        }
    });

    options.language = {
        "decimal": "",
        "emptyTable": "没有记录",
        "info": "当前展示 _START_ - _END_ 条, 共 _TOTAL_ 条记录",
        "infoEmpty": "总计0条记录",
        "infoFiltered": "",
        "infoPostFix": "",
        "thousands": ",",
        "lengthMenu": "每页 _MENU_ 条",
        "loadingRecords": "加载中...",
        "processing": "加载中...",
        "search": "搜索:",
        "zeroRecords": "没有符合要求的记录",
        "paginate": {
            "first": "首页",
            "last": "末页",
            "next": "下一页",
            "previous": "上一页"
        },
        "aria": {
            "sortAscending": ": 增序排列",
            "sortDescending": ": 倒序排序"
        }
    };

    $target.DataTable(options);
}

function app_submit_button(title, url, ajax, highlight) {

    var icon = "ion-edit";
    var notice = '确定' + title + '吗?';

    if (title == '删除') {
        icon = 'ion-trash-b';
    } else if (title == '置顶') {
        icon = 'ion-arrow-up-c';
    } else if (title == '置尾') {
        icon = 'ion-arrow-down-c';
    }else if (title == '点赞') {
        icon = 'ion-heart'
    } else if (title == '收藏') {
        icon = 'ion-star';
    } else if (title == '推荐') {
        icon = 'ion-flag';
    }

    return '<span style="display: inline-block;padding: 0 5px">' +
        '<a ' + (ajax ? 'onclick="return app_ajax_with_notice(this,\'' + notice + '\',\'' + url + '\');"' : 'href="' + url + '"') +
        ' title="' +
        title +
        '"' +
        (highlight ? ' class="selected"' : '') +
        '>' +
        '<i class="' + icon + '"></i></span></a>';
}

function app_modal(options) {
    options = options || {};
    var type = options.type || "success";
    var message = options.message || "set options.message first!";
    var cancel = options.cancel || false;
    var ok = options.ok || false;

    var modelClass = "gadmin-modal";
    var icon = "ion-checkmark";
    var title = type.charAt(0).toUpperCase() + type.slice(1);
    var cancelFunc = '$(this).parents(\'.' + modelClass + '\').remove();';
    var top = $(window).height() / 2 - 50;

    if (type == "danger") {
        icon = "ion-flame";
    } else if (type == "info") {
        icon = "ion-information-circled";
    } else if (type == "warning") {
        icon = "ion-android-warning";
    }


    var html =
        '<div class="' + modelClass + '">' +
        '<div role="dialog" class="modal fade in" style="z-index: 1050; display: block;background:#000000;opacity:0.5">' +
        '</div>' +
        '<div class="modal-dialog" style="position:fixed;left:50%; top:50%;z-index:9999">' +
        '<div class="modal-content">' +
        '<div class="modal-header bg-' + type + '">' +
        '<i class="' + icon + ' modal-icon"></i><span> ' + title + '</span>' +
        '</div>' +
        '<div class="modal-body text-center">' + message + '</div>' +
        '<div class="modal-footer">' +
        (cancel ? '<button type="button" onclick="' + cancelFunc + '" class="btn btn-' + type + '">取消</button>' : '') +
        '<button type="button" ' +
        (ok ? '' : 'onclick="' + cancelFunc + '" ') +
        'class="btn modal-ok btn-' + type + '">确定</button>' +
        '</div>' +
        '</div>' +
        '</div>' +
        '</div>';

    $('.' + modelClass).find('.modal-ok').unbind('click');
    $('.' + modelClass).remove();
    $("body").append(html);
    var $dialog = $('.' + modelClass).find('.modal-dialog');
    $dialog.css({'margin-left': -($dialog.width() / 2) + "px", 'margin-top': -($dialog.height() / 2) + "px"});
    if (ok) {
        $('.' + modelClass).find('.modal-ok').click(function () {
            $(this).parents('.' + modelClass).remove();
            ok();
        });
    }
}

function app_upload(options) {
    app_upload.index = app_upload.index || 0;
    options = options || {};
    options.selector = ".gupload";
    var callback = options.callback || function(data, target) {
            if (!data || data.length == 0) {
                app_toast_error('上传失败!');
            } else{
                var urls = [];
                for (var i in data) {
                    urls.push(data[i].url);
                }

                var ret = urls.join(",");
                var inputId = $(target).attr('data-target');
                $('#' + inputId).val(ret);
                app_toast_success('上传成功!');
            }
        };

    $(options.selector).each(function () {
        var $self = $(this);
        var domain = $self.attr('data-domain');
        var uptokenUrl = $self.attr('data-uptoken-url') || "/resource/token";
        var flashSwfUrl = $self.attr('data-flash-swf-url') || "/static/bower_components/plupload/js/Moxie.swf";
        var multiSelection = !$self.attr('data-multi') ||
            ($self.attr('data-multi') != "0" && $self.attr('data-multi') != "false");
        var maxFileSize = $self.attr('data-max-file-size') || "100mb";
        var id = 'gupload_' + (app_upload.index++);

        var html = '<button type="button" style="float: left;" class="btn btn-primary" id="' +
            id +
            '">浏览</button><div class="gupload-info" style="float:left;width: 80%;margin-left: 10px">';
        $self.html(html);
        var results = [];

        var getProcessHtml = function (id, name) {

            return '<div id="' + id + '">' +
                '<div style="display: inline-block;width: 60px;margin: 5px;word-break: break-all">' +
                name +
                '   </div>' +
                '   <div class="progress" style="display: inline-block; margin-bottom: 0px; width: 50%">' +
                '   <div class="progress-bar progress-bar-info" role="progressbar" aria-valuenow="20" aria-valuemin="0" aria-valuemax="100" style="width: 0%">' +
                '   0% Complete' +
                '   </div>' +
                '   </div>' +
                '   </div>';
        };

        // https://developer.qiniu.com/kodo/sdk/javascript
        Qiniu.uploader({
            runtimes: 'html5,flash,html4',
            browse_button: id,
            uptoken_url: uptokenUrl,
            get_new_uptoken: false,
            //unique_names: true,
            //save_key: true,
            domain: domain,
            max_file_size: maxFileSize,
            flash_swf_url: flashSwfUrl,
            max_retries: 3,
            dragdrop: false,
            drop_element: null,
            chunk_size: '4mb',
            auto_start: true,
            multi_selection: multiSelection,
            init: {
                'FilesAdded': function (up, files) {
                    var $div = $self.find('.gupload-info');
                    $div.html('');
                    plupload.each(files, function (file) {
                        $div.append(getProcessHtml(file.id, file.name));
                    });
                },
                'BeforeUpload': function (up, file) {
                },
                'UploadProgress': function (up, file) {
                    var $progressBar = $('#' + file.id).find('.progress-bar');
                    $progressBar.css("width", file.percent + "%")
                        .html(file.percent + "% Complete");
                },
                'FileUploaded': function (up, file, info) {
                    var fileName = file.name;
                    var res = app_json_decode(info);
                    var domain = up.getOption('domain');
                    var sourceLink = 'http://' + domain + "/" + res.key;
                    var ext = fileName.substr(fileName.lastIndexOf('.') + 1).toLowerCase();
                    if (ext == 'png' || ext == 'jpg' || ext == 'jpeg' || ext == 'bmp'
                        || ext == 'gif' || ext == 'ico') {
                        var html = '<img class="pop-img" style="width: 60px; height: 60px;" src="' + sourceLink + '"/>';
                        $('#' + file.id).find('>div:eq(0)')
                            .html(html);
                    }
                    results.push({
                        key: res.key,
                        url: sourceLink,
                    });
                },
                'Error': function (up, err, errTip) {
                    app_toast_error(errTip);
                },
                'UploadComplete': function () {
                    callback(results, $self[0]);
                }
            }
        });
    });
}

function app_daterange(target, options, callback) {
    options = options || {};
    callback = callback || function(){};
    var format = options.format || "YYYY-MM-DD";
    $(target).daterangepicker({
        "showDropdowns": true,
        "autoApply": true,
        ranges: {
            '今天': [moment(), moment()],
            '昨天': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
            '最近七天': [moment().subtract(6, 'days'), moment()],
            '最近30天': [moment().subtract(29, 'days'), moment()],
            '本月': [moment().startOf('month'), moment().endOf('month')],
            '上个月': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
        },
        locale : {
            applyLabel : '确定',
            cancelLabel : '取消',
            fromLabel : '起始时间',
            toLabel : '结束时间',
            customRangeLabel : '自定义',
            daysOfWeek : [ '日', '一', '二', '三', '四', '五', '六' ],
            monthNames : [ '一月', '二月', '三月', '四月', '五月', '六月',
                '七月', '八月', '九月', '十月', '十一月', '十二月' ],
            firstDay : 1,
            format: format,
            separator : ' ~ '
        }
    }, callback);
}