/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Location
     * @author dorajilocation
     * @namespace location
     */

    /**
     * Location model.
     *  @constructor
     * @type {*}
     * @name blog#Location
     */
    var Location = Model({
        create: 'POST '+API+'/locations',
        update: 'PUT '+API+'/locations/{id}',
        destroy: 'DELETE '+API+'/locations/{id}',
        findOne: 'GET '+API+'/locations/{id}',
        findAll:  function(filter){
          return $.ajax({
            type: 'GET',
            url: API+'/locations',
            data: filter
          });
        }
    }, {
    });
    return Location;
});
