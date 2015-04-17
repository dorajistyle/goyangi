{
    appDir: '../static',
    baseUrl: 'js',
    mainConfigFile: '../static/js/app.js',
"paths": {
        util: 'app/lib/util',
        info: 'app/lib/info',
        validation: 'app/lib/validation',
        refresh: 'app/lib/refresh',
        loglevel: '../bower_components/loglevel/dist/loglevel.min',
        views: 'views.build',
        text: 'vendor/text',
        facebook: "empty:",
        jquery: '../bower_components/jquery/dist/jquery.min',
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
    },
    fileExclusionRegExp: /^node_modules$|^views$|^scss$|^test$/,
    dir: '../static-build',
    optimizeAllPluginResources: true,
//    inlineText : true,
//   stubModules: ['text'],
    optimize: "uglify2",
    optimizeCss: "standard.keepLines",
    removeCombined: true,
//    uglify2: {
//        output: {
//                beautify: false
//            },
//            compress: {
//                dead_code: true,
//                sequences: true,
//                drop_console: true,
//                global_defs: {
//                    DEBUG: false
//                }
//            },
//            warnings: true,
//            mangle: "toplevel"
//    },
    preserveLicenseComments: false,
    modules: [
        {
            name: 'app'
//            include: ['views'],
//            insertRequire: ["views"]
        }
    ]
}
