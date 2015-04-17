/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Kakao related
  * @author dorajistyle
  * @namespace kakao
  */

  /**
  * Kakao model.
  * @constructor
  * @type {*}
  * @name kakao#Kakao
  */
  var Kakao = Model({
    findAll: 'GET '+API+'/oauth/kakao',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/kakao',
        data: {}
      });
    }
  }, {
  });
  return Kakao;
});
