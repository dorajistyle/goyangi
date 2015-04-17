define(['can', 'app/users/router', 'app/setting/router', 'app/admin/router', 'app/articles/router',
        'app/locations/router','app/upload/router', 'app/pages/router', 'app/main/router', 'app/models/user/user', 'refresh', 'util'],
    function (can, Users, Setting, Admin, Article, Location, Upload, Page, Main, User, Refresh, util) {
    'use strict';

    var Routers = function (target) {
        new Users(target);
        new Setting(target);
        new Admin(target);
        new Article(target);
        new Location(target);
        new Upload(target);
        new Page(target);
        new Main(target);
    };


    return Routers;
});
