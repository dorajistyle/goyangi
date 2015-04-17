/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Article related
     * @author dorajistyle
     * @namespace article
     */

    /**
     * Comment model.
     *  @constructor
     * @type {*}
     * @name article#Comment
     */
    var Comment = Model({
        create: 'POST '+API+'/articles/comments',
        findAll: function(params){
          return $.ajax({
              type: 'GET',
              url: API+'/articles/'+params.id+'/comments',
              data: params.data
          });
        },
        update : function(articleId,commentId){
          return $.ajax({
              type: 'PUT',
              url: API+'/articles/'+articleId+'/comments/'+commentId,
              data: {}
          });
        },
        destroy : function(articleId,commentId){
          return $.ajax({
              type: 'DELETE',
              url: API+'/articles/'+articleId+'/comments/'+commentId,
              data: {}
          });
        }
    }, {
    });
    return Comment;
});
