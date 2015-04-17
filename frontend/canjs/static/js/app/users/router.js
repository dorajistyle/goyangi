define(['can', 'app/users/authentications', 'app/users/email-verification', 'app/users/profile', 'app/users/reset-password'],
function (can, Authentications, EmailVerification, Profile, ResetPassword) {
    'use strict';

    /**
     * Router for users.
     * @author dorajistyle
     * @param {string} target
     * @function Router
     * @name users#Routers
     * @constructor
     */

    var Routers = function (target) {
        new Authentications(target);
        new EmailVerification(target);
        new Profile(target);
        new ResetPassword(target);
    };


    return Routers;
});
