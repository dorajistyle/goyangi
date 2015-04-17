/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * Articles
     * @author dorajiupload
     * @namespace upload
     */

    /**
     * Articles model.
     *  @constructor
     * @type {*}
     * @name blog#Articles
     */
    var Articles = Model({
        create : function(params){
          console.log("params.data : "+params.data);
          return $.ajax({
            type: 'POST',
            url: API+'/articles/all',
            contentType: "application/json",
            async: false,
            data: params.data
          });
        }
    }, {
    });
    return Articles;
});
