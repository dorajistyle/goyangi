/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Linkedin related
  * @author dorajistyle
  * @namespace linkedin
  */

  /**
  * Linkedin model.
  * @constructor
  * @type {*}
  * @name linkedin#Linkedin
  */
  var Linkedin = Model({
    findAll: 'GET '+API+'/oauth/linkedin',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/linkedin',
        data: {}
      });
    }
  }, {
  });
  return Linkedin;
});
