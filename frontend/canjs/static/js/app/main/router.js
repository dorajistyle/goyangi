define(['can', 'app/main/intro', 'util', 'i18n', 'jquery'],
    function (can, Intro, util, i18n, $) {
    'use strict';

    var intro;

    /**
    * Router for profile.
    * @author dorajistyle
    * @param {string} target
    * @function Router
    * @name users#ProfileRouter
    * @constructor
    */
    var Router = can.Control.extend({
      defaults: {}
    }, {
      init: function () {
        util.logInfo('*Main/Router', 'Initialized')
      },
      'route': function () {
        intro = new Intro(util.getFreshApp());
        intro.show();
      },
      'locales/:locale route': function (data) {
        var lang = $('html').attr('lang');
        if(lang == data.locale) {
          can.route.attr({'route': ''}, true);
        }
        else {
          window.location.replace(data.locale);
        }
      }
    });

    return Router;
});
