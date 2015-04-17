define(['loglevel', 'i18n', 'can', 'settings', 'app/partial/navbar'],
    function (log, i18n, can, settings, Navbar) {
    /**
     * A set for general purpose utilities.
     * @author dorajistyle
     * @module util
     */
    var SPACE_REGEX = /^\s*|\s\s+\n*|\s+$/g,
        INVALID_USER_NAME_REGEX = /[^a-zA-Z0-9\-.]+/gi,
        NON_DIGIT = /\D/g;
    var messages = new Array();
    return {
        /**
         * @param {string} name of param.
         * @returns {string} return param value.
         * @example <caption>
         * var lang =util.getParam('lang');
         * </caption>
         */
        getParam: function (name) {
            var value = new RegExp('[\\?&]' + name + '=([^&#]*)').exec(window.location.href);
            return value != null ? value[1] : undefined;
        },

        /**
         * Check user's browser that less than IE8.
         * @returns {*|jQuery}
         */
        isLessThanIE8: function () {
            return $('html').hasClass('lt-ie8');
        },

        /**
         * stringify JSON if browser greater than IE7.
         * @param obj
         * @returns {*}
         */
        stringify: function (obj) {
            if(this.isLessThanIE8()) return i18n.t('util.stringify.notSupport');
            return JSON.stringify(obj);
        },
         /* Get encoded Current URL
         * @returns {string}
         */
        getCurrentURL: function() {
            return window.location.href;
        },

        /**
        * Get Current host
        * @returns {string}
        */
        getCurrentHost: function() {
          var host = window.location.protocol+'//'+window.location.host;
          return host;
        },
        /**
        * Get static root from Settings
        * @returns {string}
        */
        getStaticPath: function() {
            return settings.staticPath;
        },

        /**
         * Get id data from event target.
         * @param ev
         * @returns {*}
         */
        getId: function (ev) {
          return this.getData(ev, 'id');
        },
        /**
         * Get data from event target.
         * @param ev
         * @param name
         * @returns {*|jQuery}
         */
        getData: function(ev, name) {
          return $(ev.target).data(name);
        },
        /**
        * Init array.
        * @param arr
        */
        initArray: function (arr) {
          if(arr == undefined) {
            arr = {};
          }
        },
        /**
         * Clear array.
         * @param arr
         */
        clearArray: function (arr) {
          arr.splice(0, arr.length);
        },
        /**
         * Set canjs observe array data.
         * @param arr
         * @param data
         */
        setArray: function(arr, data) {
          arr.attr(new can.List(data).attr());
            //old versions set Array
//          arr.attr(data.attr());
        },
        /**
         * Set canjs observe object data.
         * @param obj
         * @param data
         */
        setObject: function(obj, data) {
          if(can.isEmptyObject(obj)){
            obj = data;
          } else {
            obj.attr(data.attr());
          }

        },

         /**
         * Concat canjs observed array.
         * @param arr
         * @param data
         */
        concatArray: function (arr, data) {
          if(arr != undefined) {
              if(data != undefined) {
                arr.attr(arr.attr().concat(data.attr()));
              }
          }
        },
         /**
         * refresh canjs paginate array.
         * @param arr
         * @param data
         */
        refreshPaginate: function (arr, data) {
            arr.attr('hasPrev',data.hasPrev);
            arr.attr('hasNext',data.hasNext);
            arr.attr('prevPage',data.currentPage-1);
            arr.attr('nextPage',data.currentPage+1);
        },

        /**
         * Refresh canjs observed array.
         * @param arr
         * @param data
         */
        refreshArray: function (arr, data) {
          if(arr == undefined) {
              this.initArray(arr);
          }
          if(arr != undefined) {
              this.logDebug('array','not null');
              this.clearArray(arr);
              if(data != undefined) {
              this.logDebug('array','not null');
                this.setArray(arr, data);
              }
          }
        },
        /**
         * Refresh canjs observed object.
         * @param obj
         * @param data
         */
        refreshObject: function (obj, data) {
          if(obj != undefined) {
//              obj.length = 0;
              if(data != undefined) {
                  this.setObject(obj, data);
              }
          }
        },



        /**
         * Refresh page title to h1#pageTitle's content.
         */
        refreshTitle: function() {
          var $pageTitle = document.getElementById('pageTitle');
          var title = '';
          if($pageTitle != null) title = $pageTitle.innerHTML + ' - ';
          document.title = title + i18n.t('service.name');
        },

         refreshMeta: function() {
          var $pageTitle = document.getElementById('pageTitle'),
              $pageDescription = $('#pageDescription');
          var title = '',
              description = '',
              serviceName = '',
              serviceUrl = '';
          serviceName = i18n.t('service.name');
          serviceUrl = this.getCurrentURL();
          if($pageTitle != null) { title = $pageTitle.innerHTML + ' - '; }
          title = this.truncate70(title + serviceName);
          document.title = title;
          if($pageDescription != null) { description = $pageDescription.text(); }
          description = this.truncate150(description);
           var meta = '<title>'+title+'</title>'+
                      '<meta name="description" content="'+description+'" />'+
                      '<meta property="og:title" content="'+title+'" />'+
                      '<meta property="og:description" content="'+description+'" />'+
                      '<meta property="og:type" content="article">'+
                      '<meta property="og:url" content="'+serviceUrl+'">'+
                      '<meta property="og:site_name" content="'+serviceName+'">'+
                      '<meta name="twitter:card" content="'+description+'">'+
                      '<meta name="twitter:url" content="'+serviceUrl+'">'+
                      '<meta name="twitter:title" content="'+title+'">'+
                      '<meta name="twitter:description" content="'+description+'">'+
                      '<meta itemprop="name" content="'+title+'">'+
                      '<meta itemprop="description" content="'+description+'">';
           var $head = $('head');
           $($head.find('title')).remove();
           $($head.find('meta[name="description"]')).remove();
           $($head.find('meta[property*="og:"]')).remove();
           $($head.find('meta[name*="twitter:"]')).remove();
           $($head.find('meta[itemprop]')).remove();
           $head.prepend(meta);
        },

        /**
         * Set wrapper information of control.
         * @param control
         * @param wrapperName
         */
        setWrapperInfo: function (control, wrapperName) {
          var $wrapper = $('#'+wrapperName);
          control.wrapper = $wrapper.html();
          control.wrapperId = $wrapper.prop('id');
          control.wrapperClass = $wrapper.prop('class');
        },
        /**
         * Replace wrapper information of control.
         * @param control
         * @param wrapperName
         */
        replaceWrapper: function (control, wrapperName) {
          var $wrapper = $('#'+wrapperName);
          $wrapper.replaceWith('<div id="'+control.wrapperId+'" class="'+control.wrapperClass+'">'+control.wrapper+'</div>');
        },

        /**
         * Allocate the view if view is undefined.
         * @param router
         * @param allocate
         */
        allocate: function (router, control) {
            if(control === undefined || control.element === null ) router.allocate();
        },
        /**
         * Enable javascript logger.
         *
         */
        enableLog: function () {
            log.enableAll();
        },
        /**
         * Disable javascript logger.
         *
         */
        disableLog: function () {
            log.disableAll();
        },
        /**
         * Disables all logging below the given level.
         * @param level
         * @private
         */
        setLogLevel: function (level) {
            log.setLevel(level);
        },
        /**
         * Log trace level message into console.
         * It works only after enable logger. {@link module:util/enableLog}
         * @param target
         * @param msg
         */
        logTrace: function (target, msg) {
            log.trace(target + ' : ', msg);
        },
        /**
         * Log debug level message into console.
         * It works only after enable logger. {@link module:util/enableLog}
         * @param target
         * @param msg
         */
        logDebug: function (target, msg) {
            log.debug(target + ' : ', msg);
        },
        /**
         * Log info level message into console.
         * It works only after enable logger. {@link module:util/enableLog}
         * @param target
         * @param msg
         */
        logInfo: function (target, msg) {
            log.info(target + ' : ', msg);
        },
        /**
         * Log warn level message into console.
         * It works only after enable logger. {@link module:util/enableLog}
         * @param target
         * @param msg
         */
        logWarn: function (target, msg) {
            log.warn(target + ' : ', msg);
        },
        /**
         * Log error level message into console.
         * It works only after enable logger. {@link module:util/enableLog}
         * @param target
         * @param msg
         */
        logError: function (target, msg) {
            log.error(target + ' : ', msg);
        },

        /**
         * Log JSON object as debug level message into console.
         * It works only after enable logger. {@link module:util/enableLog}
         * @param target
         * @param obj JSON object
         */
        logJson: function (target, obj) {
            log.debug(target + ' : ', this.stringify(obj));
        },

        /**
         * show message popup.
         * @param msg
         * @param type
         */
        showMessage: function (msg, type) {
            if(msg != undefined && msg.length > 0) {
              var $msgBox = $('#message' + type);
              $msgBox.html(msg);
              $msgBox.fadeIn(200).removeClass('uk-hidden').delay(500).fadeOut(800);
              this.logDebug("showMessage",msg);
            }
        },
        /**
         * show warning message.
         * @param msg
         */
        showWarningMsg: function (msg) {
            this.showMessage(msg, 'AlertWarning');
        },
        /**
         * show error message.
         * @param msg
         */
        showErrorMsg: function (msg) {
            this.showMessage(msg, 'AlertError');
        },
        /**
         * show info message.
         * @param msg
         */
        showInfoMsg: function (msg) {
            this.showMessage(msg, 'AlertInfo');
        },
        /**
         * show success message.
         * @param msg
         */
        showSuccessMsg: function (msg) {
            this.showMessage(msg, 'AlertSuccess');
        },
        /**
         * Put message into message array.
         * @param msg
         */
        putMessage: function (msg) {
            messages[messages.length] = msg;
        },
        /**
         * Show messages that putted before.
         */
        showMessages: function () {
            this.showWarningMsg(messages.join('<br/>'));
            this.clearMessages();
        },
        /**
        * Get messages that was placed before.
        */
        getMessages: function () {
          var msg = "";
          if(messages.length > 0) {
            msg = messages.join('<br/>');
            this.clearMessages();
          }
          return msg;
        },
        /**
         * Clear message array.
         */
        clearMessages: function () {
            messages.length = 0;
        },
        /**
         * DEPRECATED
         * Return Template engine for typeahead
         * @returns {{}}
         */
//        getEngine: function () {
//           var Stache = {};
//           Stache.compile = function (template) {
//                var compile = can.view.stache(template),
//                    render = {
//                        render: function (ctx) {
//                            return compile.render(ctx);
//                        }
//                    };
//                return render;
//           };
//           return Stache;
//        },
        /**
         * Compile Template for typeahead
         * @returns {{}}
         */
        compileTemplate: function (template) {
        var compile = can.view.stache(template);
        return compile;
        },

        /**
         * Refresh app div.
         * @param target
         * @param name
         */
        getFreshApp: function (){
          return this.getFreshDiv('app');
        },
        /**
         * Refresh div.
         * @param name
         * @returns {HTMLElement}
         */
        getFreshDiv: function (name){
          $('#'+name).replaceWith('<div id="'+name+'"></div>');
          return $('#'+name);
        },
        /**
         * handle http status codes.
         * @param xhr
         */
        handleStatus: function (xhr, message, changeRoute) {
            var util = this;
            util.logError('response code', xhr.status);
            switch (xhr.status) {
                case 200:
                    // OK
                    break;
                case 201:
                  // Created
                  break;
                case 304:
                  // Not Modified
                  break;
                case 400:
                    // Bad Request
                    // $(settings.appDivId).html(can.view('views_partial_error_stache', {}));
                    // can.route.attr({'route': ''}, true);
                    break;
                case 401:
                  // Unauthorized//
                  Navbar.load();

                  if(changeRoute) {
                    util.showErrorMsg(message+i18n.t('error.loginPlease'));
                    $(settings.appDivId).html(can.view('views_partial_error_stache', {}));
                    can.route.attr({'route': 'login'}, true);
                  } else {
                    util.showErrorMsg(message);
                  }
                  break;
                case 403:
                  // Forbidden
                  // Navbar.load();
                  util.showMsgAndChangeRoute(message, changeRoute, 'login');
                  break;
                case 404:
                  // Not Found
                  util.showMsgAndChangeRoute(message, changeRoute, '');
                  break;
                // case 405:
                //     // request method not supported by that resource
                //     break;
                // case 409:
                //     // request could not be processed because of conflict
                //     break;
                case 500:
                // Internal Server Error
                  util.showMsgAndChangeRoute(message, changeRoute, '');
                  // $(settings.appDivId).html(can.view('views_partial_error_stache', {}));
                  break;
                default:
                  util.showMsgAndChangeRoute(message, changeRoute, '');

            }
        },

        /**
        * show error message and change route.
        * @param message
        * @param changeRoute
        * @param route
        */
        showMsgAndChangeRoute: function (message, changeRoute, route) {
          this.showErrorMsg(message);
          if(changeRoute) {
            $(settings.appDivId).html(can.view('views_partial_error_stache', {}));
            can.route.attr({'route': ''}, true);
          }
        },
        /**
         * show error message and handle status.
         * @param xhr
         * @param message
         */
        handleStatusWithErrorMsg: function (xhr, message) {
            if(message != undefined && message.length > 0) {
              message = message + "\n<br\>";
            }
            this.handleStatus(xhr, message);
        },

        /**
        * show error message and handle status.
        * @param xhr
        */
        handleError: function (xhr) {
          var message = i18n.t(xhr.responseJSON.messageType);
          this.handleStatusWithErrorMsg(xhr, message);
        },

        /**
         * delete cookie
         * @param name
         * @returns {*}
         */
        deleteCookie: function (name) {
            document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;';
        },

         /**
         * compare two values and return boolean.
         * @param first
         * @param second
         * @returns {boolean}
         */
        compare: function(first, second) {
          return first == second;
        },

        /**
         * compare two values and return true when first value is less than second.
         * @param first
         * @param second
         * @returns {boolean}
         */
        lessThanSecond: function(first, second) {
          return first < second;
        },

        /**
         * compare two values and return true when first value is greater than second.
         * @param first
         * @param second
         * @returns {boolean}
         */
        greaterThanSecond: function(first, second) {
          return first > second;
        },


         /* empty the element.
         * @param element
         */
        empty: function (element) {
          var child;
          while (child = element.lastChild) {
            element.removeChild(child); //clear
          }
        },

         /**
         * Trim the string.
         * @returns trimedStr
         * @param str
         */
        trim: function(str) {
            return str.replace(SPACE_REGEX,'');
        },

         /**
         * Strip non numeric characters of the string.
         * @returns stripped str
         * @param str
         */
        strip: function(str) {
            return str.replace(NON_DIGIT,'');
        },


        /**
         * Truncate the string.
         * @param str
         */
        truncate: function(str, limit) {
//            return str;
            if(str != undefined) {
                if(str.length > limit) {
                    return this.trim(str).substr(0, limit-1)+'...';
                }
            }
            return str;

        },

        /**
         * Truncate the string with default limit 40.
         * @param str
         * @returns {*}
         */
        truncate40: function(str) {
            return this.truncate(str, 40);
        },

        /**
         * Truncate the string with default limit 70.
         * @param str
         * @returns {*}
         */
        truncate70: function(str) {
            return this.truncate(str, 70);
        },

        /**
         * Truncate the string with default limit 150.
         * @param str
         * @returns {*}
         */
        truncate150: function(str) {
            return this.truncate(str, 150);
        },

        /**
         * replace / to blank
         * @param str
         * @returns {*|XML|string|void}
         */
        slashToBlank: function (str) {
            return str.replace(/\//gi,'');
        },
        /**
        * Correct username.
        * @param username
        * @returns {*|XML|string|void}
        */
        correctUserName: function (username) {
          if(username.length > 16) {
            username = username.substring(0,16);
          }
          var correctedName = username.replace(INVALID_USER_NAME_REGEX,'0');
          return correctedName;
        },


         /**
         * sets focus on top of the page
         */
        setFocusOnTop: function() {
            window.scrollTo(0,0);
        }

    };
});
