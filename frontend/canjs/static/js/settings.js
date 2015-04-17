define(['jquery'], function ($) {
    'use strict';
    var exports = {};
    var requireToUrl = require.toUrl('');
    var staticPrefix = 'static/';
    var jsPostfix = '/js/';
    var staticPath = requireToUrl.substring($.inArray( staticPrefix, requireToUrl),requireToUrl.length-jsPostfix.length);
//    if($.inArray( '/',  requireToUrl) == 0) {
//        staticPath = staticPath.substring(1,staticPath.length);
//    }
    exports.canFixture = false;
    exports.host = '';
    exports.staticPath = staticPath;

    exports.i18nOptions = {debug: true,
                            getAsync: true,
                            useLocalStorage: false,
                            localStorageExpirationTime: 86400000,
                            fallbackLng: 'en',
                            resGetPath: staticPath+'/locales/__lng__/__ns__.json'};
    exports.apiPath = exports.host+'/api/v1';
    exports.useLogger = true;
    exports.facebookAppId = '';
    exports.appDivId = '#app';
    exports.savefilesMeta = true;
    exports.saveAllfilesMetaAtOnce = true;
    exports.spinOptions = {
        lines: 7, // The number of lines to draw
        length: 4, // The length of each line
        width: 2, // The line thickness
        radius: 3, // The radius of the inner circle
        corners: 1, // Corner roundness (0..1)
        rotate: 0, // The rotation offset
        direction: 1, // 1: clockwise, -1: counterclockwise
        color: '#000', // #rgb or #rrggbb
        speed: 1.1, // Rounds per second
        trail: 60, // Afterglow percentage
        shadow: false, // Whether to render a shadow
        hwaccel: false, // Whether to use hardware acceleration
        className: 'spinner', // The CSS class to assign to the spinner
        zIndex: 2e9, // The z-index (defaults to 2000000000)
        top: 'auto', // Top position relative to parent in px
        left: 'auto' // Left position relative to parent in px
    };
    return exports;
});
