require.config({
//    baseUrl: '',
    paths: {
        util: 'app/lib/util',
        info: 'app/lib/info',
        validation: 'app/lib/validation',
        refresh: 'app/lib/refresh',
        loglevel: '../bower_components/loglevel/dist/loglevel.min',
        views: 'views.build',
        text: 'vendor/text',
        facebook: ['//connect.facebook.net/en_US/all',
                     'vendor/facebook.all'
                    ],
        jquery: '../bower_components/jquery/dist/jquery.min',
//        jquery: [
//            '//ajax.googleapis.com/ajax/libs/jquery/1.11.0/jquery.min',
//            '../bower_components/jquery/dist/jquery.min'
//        ],
        googlemaps: '../bower_components/googlemaps-amd/src/googlemaps',
        async: '../bower_components/requirejs-plugins/src/async',
        can: '../bower_components/canjs/amd/can',
        i18n: '../bower_components/i18next/i18next.amd.withJQuery',
        spin: '../bower_components/spin.js/spin',
        placeholders: 'vendor/placeholders.jquery.min',
        uikit: '../bower_components/uikit/js/uikit',
        bloodhound: '../bower_components/typeahead.js/dist/bloodhound.min',
        hammerjs: '../bower_components/hammerjs/hammer.min',
        moment: '../bower_components/moment/min/moment-with-locales.min',
        'jquery.magnific-popup': '../bower_components/magnific-popup/dist/jquery.magnific-popup.min',
        'jquery.typeahead': '../bower_components/typeahead.js/dist/typeahead.jquery.min',
        'jquery.ba-bbq': 'vendor/jquery.ba-bbq.min',
        'jquery.ui.widget': '../bower_components/jquery-file-upload/js/vendor/jquery.ui.widget',
        'jquery.iframe-transport': '../bower_components/jquery-file-upload/js/jquery.iframe-transport',
        'jquery.fileupload': '../bower_components/jquery-file-upload/js/jquery.fileupload',
        'jquery.xdomainrequest': 'vendor/jquery.xdomainrequest.min',
        'jquery.hammer': '../bower_components/jquery-hammerjs/jquery.hammer',
        'jquery.trumbowyg': '../bower_components/trumbowyg/dist/trumbowyg.min'
//        jquery.xml2json: 'vendor/jquery.xml2json'
    },
    googlemaps: {
      params: {
        key: '',
        sensor: 'false',
        libraries: 'places',
        language: 'ko'
      }
    },
    config: {
        uikit: {
            base: '../bower_components/uikit'
        }
    },
    shim: {
          bloodhound: {
              deps: ['jquery'],
              exports: 'Bloodhound'
          },
          'jquery.typeahead': {
              deps: ['jquery'],
              exports: '$.fn.typeahead'
          },
          'jquery.ba-bbq': {
            deps: ['jquery']
          },
          'jquery.trumbowyg': {
            deps: ['jquery'],
            exports: '$.fn.trumbowyg'
          },
          uikit: {
              deps: ['jquery']
          },
          placeholders: {
            deps: ['jquery']
          },
          util: {
              deps: ['jquery']
          },
          facebook: {
          exports: 'FB'
          },
          enforceDefine: true
    }
    // waitSeconds: 14
});
/**
 * @requires jquery
 * @requires util
 * @requires i18n
 * @requires can
 * @requires settings
 * @requires spin
 * @requires app/partial/navbar
 * @requires app/routers
 * @requires can/view/stache
 */
requirejs(['can', 'jquery', 'util', 'i18n', 'moment', 'settings', 'app/partial/navbar',  'app/partial/footer', 'refresh',
          'facebook', 'app/routers','uikit', 'app/lib/helper', 'can/view/stache', 'can/route/pushstate',
           'placeholders', 'jquery.ba-bbq', 'jquery.hammer', 'views', 'app/models/fixtures'],
    function (can, $, util, i18n, moment, settings, Navbar, Footer, Refresh, FB, Routers, UI) {
        'use strict';

        $(document).ready(function () {
            var lang = $('html').attr('lang');
            var i18nOption = settings.i18nOptions;
            lang != undefined
                ? i18nOption.lng = lang
                : i18nOption.lng = 'en';
            settings.useLogger ? util.enableLog() : util.disableLog();
            can.when(i18n.init(i18nOption)).then(function () {
                moment.locale(i18nOption.lng);
                Navbar.load();
                Footer.load();
                new Routers(document.getElementById('appWrapper'));
                FB.init({
                  appId : settings.facebookAppId,
                  status: true,
                  xfbml: true,
                  cookie: false
                });

//                var target = document.getElementById('ajaxProgress');
//                var spinner = new Spin(settings.spinOptions);
//                var $document = $(document);
//                $document.ajaxStart(function () {
//                    spinner.spin(target);
//                    $(target).removeClass('hidden');
//                });
//                $document.ajaxStop(function () {
//                    spinner.stop();
//                    $(target).addClass('hidden');
//                });
                var $document = $(document);
//                $document.ajaxSuccess(function( event, xhr, settings ) {
//                    util.logDebug('success Event', event);
//                    util.logDebug('success Xhr', xhr);
//                });
                $document.ajaxError(function(event, xhr, settings, exception) {
                    util.logDebug('error Event', event);
                    util.logDebug('error Xhr', xhr);
                    Refresh.loadWithException('', xhr);
                });

                $document.on('keypress', 'form', function (e) {
                    var code = e.keyCode || e.which;

                    if (code == 13) {
                        var tagName = e.target.tagName;
                        if((!(tagName == 'DIV') && !(tagName == 'TEXTAREA')) || ($(e.target).hasClass('enter-as-save'))) {
                            var enter = $(this).closest('form').find(':submit').not('.hidden,.cancel');
                             e.preventDefault();
                            if(enter.length > 0) {
                                enter.click();
                                return false;
                            }
                        }
                    }
                });

                UI.ready(function(context) {
                    $('[data-uk-nav]', context).each(function() {
                        var nav = $(this);
                        if (!nav.data('nav')) {
                            var obj = UI.nav(nav, UI.Utils.options(nav.attr('data-uk-nav')));
                        }
                    });
                });
                can.route.ready();
            });
        });
    });
