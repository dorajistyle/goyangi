/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Facebook related
  * @author dorajistyle
  * @namespace facebook
  */

  /**
  * Facebook model.
  * @constructor
  * @type {*}
  * @name facebook#Facebook
  */
  var Facebook = Model({
    findAll: 'GET '+API+'/oauth/facebook',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/facebook',
        data: {}
      });
    }
  }, {
  });
  return Facebook;
});
