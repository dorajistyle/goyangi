/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Location related
     * @author dorajistyle
     * @namespace location
     */

    /**
     * Liking model.
     *  @constructor
     * @type {*}
     * @name location#Liking
     */
    var Liking = Model({
        create: 'POST '+API+'/locations/likings',
        destroy : function(locationId,userId){
          return $.ajax({
            type: 'DELETE',
            url: API+'/locations/'+locationId+'/likings/'+userId,
            data: {}
          });
        },
        findAll:  function(params){
          return $.ajax({
            type: 'GET',
            url: API+'/locations/'+params.id+'/likings',
            data: params.data
          });
        }
    }, {
    });
    return Liking;
});
