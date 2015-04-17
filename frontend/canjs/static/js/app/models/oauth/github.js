/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
  'use strict';

  /**
  * Github related
  * @author dorajistyle
  * @namespace github
  */

  /**
  * Github model.
  * @constructor
  * @type {*}
  * @name github#Github
  */
  var Github = Model({
    findAll: 'GET '+API+'/oauth/github',
    destroy : function(){
      return $.ajax({
        type: 'DELETE',
        url: API+'/oauth/github',
        data: {}
      });
    }
  }, {
  });
  return Github;
});
