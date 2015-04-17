/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Yahoo related
  * @author dorajistyle
  * @namespace yahoo
  */

  /**
  * Yahoo model.
  * @constructor
  * @type {*}
  * @name yahoo#Yahoo
  */
  var Yahoo = Model({
    findAll: 'GET '+API+'/oauth/yahoo',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/yahoo',
        data: {}
      });
    }
  }, {
  });
  return Yahoo;
});
