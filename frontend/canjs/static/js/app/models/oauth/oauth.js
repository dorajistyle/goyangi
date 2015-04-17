/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Oauth related
  * @author dorajistyle
  * @namespace oauth
  */

  /**
  * Oauth model.
  * @constructor
  * @type {*}
  * @name oauth#Oauth
  */
  var Oauth = Model({
    findAll: 'GET '+API+'/oauth'
  }, {
  });
  return Oauth;
});
