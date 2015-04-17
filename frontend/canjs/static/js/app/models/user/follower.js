/*global define*/
define(['can', 'app/models/model', 'app/models/api-path'], function (can, Model, API) {
    'use strict';

    /**
     * User related
     * @author dorajistyle
     * @namespace follower
     */

    /**
     * User model.
     *  @constructor
     * @type {*}
     * @name follower#Follower
     */
    var Follower = Model({
        create: 'POST '+API+'/followers',
        destroy: 'DELETE '+API+'/followers/{id}',
        findAll: 'GET '+API+'/followers'
    }, {
    });
    return Follower;
});
