define(['can', 'app/models/oauth/oauth', 'app/partial/oauth',
    'util', 'i18n', 'jquery'],
    function (can, OauthModel, Oauth, util, i18n, $) {
    'use strict';


    /**
     * Instance of Setting Contorllers.
     * @private
     */
    var connection;

    /**
     * Control for Setting Profile
     * @private
     * @author dorajistyle
     * @param {string} target
     * @name setting#Connection
     * @constructor
     */
    var Connection = can.Control.extend({
        init: function () {
            util.logInfo('*Setting/Connection', 'Initialized');
        },
        load: function () {
                OauthModel.findAll({}, function (oauth){
                  var templateData = new can.Observe({
                    oauthStatus: oauth.oauthStatus
                  });
                  connection.oauthData = templateData;
                connection.show();
                util.refreshTitle();
            },function (xhr) {
                util.handleStatus(xhr);
            });
        },
        reload: function (oauthData) {
          $.each(connection.oauthData.oauthStatus._data, function( key, value ) {
            connection.oauthData.oauthStatus.attr(key, oauthData[key]);
          });
        },
        /**
         * Show Setting view.
         * @memberof setting#Connection
         */
        show: function () {
            util.logDebug("connection.oauthData.oauthStatus", connection.oauthData.oauthStatus);
            this.element.html(can.view('views_setting_connection_stache', {
              oauthStatus: connection.oauthData.oauthStatus
            }));
            var $oauthWrapper = $(this.element.find('#oauthWrapper'));
            new Oauth($oauthWrapper, {view: connection});
        },
//         '.send-to-facebook click': function (el, ev) {
//             ev.preventDefault();
//             ev.stopPropagation();
//             this.performSendToFacebook();
//             return false;
//         },
//         '.disconnect-to-facebook click': function (el, ev) {
//             ev.preventDefault();
//             ev.stopPropagation();
//             this.performDisconnectToFacebook();
//             return false;
//         },
//         validateSendtoFB: function () {
//             return validation.minLengthTextarea('content', 3, 'validation.contentTooShort')
//         },
//         performDisconnectToFacebook: function () {
//             Facebook.destroy(function () {
//                 connection.load('connection');
//                 window.location.reload(true);
// //            $('#sendToFacebookForm').addClass('uk-hidden');
// //            $('.disconnect-to-facebook-wrapper').addClass('uk-hidden');
// //            $('.connect-to-facebook-wrapper').removeClass('uk-hidden');
//             });
//         },
//         performSendToFacebook: function () {
//             if (this.validateSendtoFB()) {
//                 var form = this.element.find('#sendToFacebookForm');
//                 var values = can.deparam(form.serialize());
//                 var $form = $(form);
//                 if (!$form.data('submitted')) {
//                     $form.data('submitted', true);
//                     var sendBtn = this.element.find('.send-to-facebook');
//                     sendBtn.attr('disabled', 'disabled');
//                     can.when(Facebook.create(values)).then(function (result) {
//                         sendBtn.removeAttr('disabled');
//                         $form.data('submitted', false);
//                         if (result.oauthStatus) {
//                             util.showSuccessMsg(i18n.t('facebook.send.done'));
//                             $('#content').val('');
//                         } else {
//                             util.showErrorMsg(i18n.t('facebook.send.fail'));
//                         }
//                     }, function () {
//                         connection.performDisconnectToFacebook();
//                         util.showErrorMsg(i18n.t('facebook.send.connectionFailed'));
//                     });
//                 }
//             } else {
//                 util.showMessages();
//             }
//         }
    });

    /**
     * Router for setting connection.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name setting#ConnectionRouter
     * @constructor
     */
     var Router = can.Control.extend({
        defaults: {}
        }, {
            init: function () {
                connection = undefined;
                util.logInfo('*setting/Connection/Router', 'Initialized');
            },
            allocate: function () {
                var $app = util.getFreshDiv('setting-connection');
                connection = new Connection($app);
            },
            load: function(page) {
                util.logDebug('*setting/Connection/Router', 'loaded');
                util.allocate(this, connection);
                connection.load(page);
            }
        });

    return Router;
});
