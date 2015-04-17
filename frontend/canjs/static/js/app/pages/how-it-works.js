define(['can', 'util', 'i18n'],
    function (can, util, i18n) {

    'use strict';
    /**
     * Control for new propose
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name pages#HowItWorks
     * @constructor
     */
    var howItWorks;
    var HowItWorks = can.Control.extend({
        init: function () {
            util.logInfo('*Pages/HowItWorks', 'Initialized');

        },
        /**
         * Show Terms and Conditions view.
         * @memberof pages#HowItWorks
         */
        show: function () {
            util.setFocusOnTop();
            var serviceName = i18n.t('service.name');
            this.element.html(can.view('views_page_simple_stache', {
              title: i18n.t('howItWorks.title'),
              content: i18n.t('howItWorks.content', { postProcess: 'sprintf', sprintf: { name: serviceName } })
              }));
        }
    });


    /**
     * Router for HowItWorks.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name pages#HowItWorks-router
     * @constructor
     */
    var Router = can.Control.extend({
        defaults: {}
    }, {
        init: function () {
            util.logInfo('*pages#HowItWorks/Router', 'Initialized');
        },
        allocate: function() {
            var $app = util.getFreshApp();
            howItWorks = new HowItWorks($app);
        },
        'how-it-works route': function () {
            this.allocate();
            howItWorks.show();
        }
    });
    return Router;
});
