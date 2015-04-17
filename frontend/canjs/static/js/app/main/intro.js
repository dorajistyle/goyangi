define(['can', 'util', 'i18n', 'jquery'],
    function (can, util, i18n, $) {
    'use strict';

    var Intro = can.Control.extend({
        init: function () {
            util.logInfo('*Intro', 'Initialized');
        },
        /**
         * Show main view.
         */
        show: function () {
            var mainImage = util.getStaticPath()+'/images/goyangi.jpg';
            this.element.html(can.view('views_intro_greeting_stache', {mainImage: mainImage}));
        }
    });

    return Intro;
});
