/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Twitter related
  * @author dorajistyle
  * @namespace twitter
  */

  /**
  * Twitter model.
  * @constructor
  * @type {*}
  * @name twitter#Twitter
  */
  var Twitter = Model({
    findAll: 'GET '+API+'/oauth/twitter',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/twitter',
        data: {}
      });
    }
  }, {
  });
  return Twitter;
});
