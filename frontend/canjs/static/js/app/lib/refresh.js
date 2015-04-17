define(['app/partial/navbar', 'util', 'settings', 'jquery'],
    function (Navbar, util) {
    'use strict';

    return {
       /**
         * Refresh navbar and replace hash.
         * @param hash
       */
       load: function(routeMap){
           Navbar.load();
           can.route.attr(routeMap, true);
       },
        /**
         * Refresh navbar and replace hash. If failed show error message on the screen.
         * @param hash
         * @param control
         */
       loadWithException: function(hash, xhr){
           if(!xhr.responseJSON) {
               util.handleStatus(xhr);
           }
       }
    };
});
