/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Article
     * @author dorajiarticle
     * @namespace article
     */

    /**
     * Article model.
     *  @constructor
     * @type {*}
     * @name blog#Article
     */
    var Article = Model({
        create: 'POST '+API+'/articles',
        update: 'PUT '+API+'/articles/{id}',
        destroy: 'DELETE '+API+'/articles/{id}',
        findOne: 'GET '+API+'/articles/{id}',
        findAll:  function(filter){
          return $.ajax({
            type: 'GET',
            url: API+'/articles',
            data: filter
          });
        }
    }, {
    });
    return Article;
});
