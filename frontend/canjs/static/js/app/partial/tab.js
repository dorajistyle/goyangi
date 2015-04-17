define(['can', 'util', 'jquery'],
    function (can, util, $) {
    'use strict';

    /**
     * Control for Tab
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name share#Tab
     * @constructor
     */
    var Tab = can.Control.extend({
        init: function () {
            this.options.navid = '#'+this.options.route+'Nav';
            util.logInfo('*share/Tab', 'Initialized');
        },
        /**
         * Show
         * @param name
         */
        showTab: function (tabName) {
            $('#'+this.options.route+'Wrapper > div').addClass('uk-hidden');
            $('#'+this.options.route+'-'+tabName).removeClass('uk-hidden');
        },
        activeTab: function (tabName) {
            $('#'+this.options.route+'Nav > ul > li').removeClass('uk-active');
            $('#'+this.options.route+'Nav > ul > li.' + tabName + '-tab').addClass('uk-active');
        },
        '#{route}Nav > ul > li > a click': function (el, ev) {
            ev.preventDefault();
            ev.stopPropagation();
            var target = ev.target;
            util.logDebug('tab route', this.options.route);
            util.logDebug('tab name', $(target).attr('class'));
            can.route.attr({'route': this.options.route+'/:tab', 'tab': $(target).attr('class')});
            return false;
        }
    });
    return Tab;
});
