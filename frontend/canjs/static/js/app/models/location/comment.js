/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Location related
     * @author dorajistyle
     * @namespace location
     */

    /**
     * Comment model.
     *  @constructor
     * @type {*}
     * @name location#Comment
     */
    var Comment = Model({
        create: 'POST '+API+'/locations/comments',
        findAll: function(params){
          return $.ajax({
              type: 'GET',
              url: API+'/locations/'+params.id+'/comments',
              data: params.data
          });
        },
        update : function(locationId,commentId){
          return $.ajax({
              type: 'PUT',
              url: API+'/locations/'+locationId+'/comments/'+commentId,
              data: {}
          });
        },
        destroy : function(locationId,commentId){
          return $.ajax({
              type: 'DELETE',
              url: API+'/locations/'+locationId+'/comments/'+commentId,
              data: {}
          });
        }
    }, {
    });
    return Comment;
});
