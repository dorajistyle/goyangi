/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Google related
  * @author dorajistyle
  * @namespace google
  */

  /**
  * Google model.
  * @constructor
  * @type {*}
  * @name google#Google
  */
  var Google = Model({
    findAll: 'GET '+API+'/oauth/google',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/google',
        data: {}
      });
    }
  }, {
  });
  return Google;
});
