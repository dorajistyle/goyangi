/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace user
     */

    /**
     * Liked model.
     *  @constructor
     * @type {*}
     * @name user#Liked
     */
    var Liked = Model({
        findAll:  function(params){
          return $.ajax({
            type: 'GET',
            url: API+'/users/'+params.id+'/liked',
            data: params.data
          });
        }
    }, {
    });
    return Liked;
});
