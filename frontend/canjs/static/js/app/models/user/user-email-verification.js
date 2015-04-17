/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace user-email-verification
     */

    /**
     * UserEmailVerification model.
     *  @constructor
     * @type {*}
     * @name user-email-verification#UserEmailVerification
     */
    var UserEmailVerification = Model({
      create: 'POST '+API+'/user/send/email/verification/token',
      update: function(data){
        return $.ajax({
          type: 'PUT',
          url: API+'/user/verify/email',
          data: data
        });
      }
    }, {
    });

    return UserEmailVerification;
});
