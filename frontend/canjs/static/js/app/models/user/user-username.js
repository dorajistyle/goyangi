/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * User related
  * @author dorajistyle
  * @namespace user-username
  */

  /**
  * UserUsername model.
  *  @constructor
  * @type {*}
  * @name user-username#UserUsername
  */
  var UserUsername = Model({
    findOne: 'GET '+API+'/user/username/{username}'
  }, {
  });
  return UserUsername;
});
