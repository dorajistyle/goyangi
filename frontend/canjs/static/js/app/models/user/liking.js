/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace user
     */

    /**
     * Liking model.
     *  @constructor
     * @type {*}
     * @name user#Liking
     */
    var Liking = Model({
        create: 'POST '+API+'/users/likings',
        destroy : function(followingId,followerId){
          return $.ajax({
            type: 'DELETE',
            url: API+'/users/'+followingId+'/likings/'+followerId,
            data: {}
          });
        },
        findAll:  function(params){
          return $.ajax({
            type: 'GET',
            url: API+'/users/'+params.id+'/likings',
            data: params.data
          });
        }
    }, {
    });
    return Liking;
});
