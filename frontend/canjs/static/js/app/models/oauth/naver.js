/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Naver related
  * @author dorajistyle
  * @namespace naver
  */

  /**
  * Naver model.
  * @constructor
  * @type {*}
  * @name naver#Naver
  */
  var Naver = Model({
    findAll: 'GET '+API+'/oauth/naver',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/naver',
        data: {}
      });
    }
  }, {
  });
  return Naver;
});
