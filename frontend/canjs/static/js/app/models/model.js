/*global define*/
define(['can', 'can/model'], function (can, Model) {
    'use strict';

    /**
     * @author dorajistyle
     * @namespace model
     */

    /**
     * Model inherit can/model to return blank data for canjs 1.7.
     * @constructor
     * @type {*}
     * @name model#Model
     */
    var Model = Model({

        models: function(data){
            this.data = data;
            this.data.data = {};
            return can.Model.models.call(this,this.data);}
    }, {
    });
    return Model;
});
