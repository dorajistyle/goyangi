/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Article related
     * @author dorajistyle
     * @namespace article
     */

    /**
     * Liking model.
     *  @constructor
     * @type {*}
     * @name article#Liking
     */
    var Liking = Model({
        create: 'POST '+API+'/articles/likings',
        destroy : function(articleId,userId){
          return $.ajax({
            type: 'DELETE',
            url: API+'/articles/'+articleId+'/likings/'+userId,
            data: {}
          });
        },
        findAll:  function(params){
          return $.ajax({
            type: 'GET',
            url: API+'/articles/'+params.id+'/likings',
            data: params.data
          });
        }
    }, {
    });
    return Liking;
});
