define(['jquery', 'can', 'app/models/user/user-current', 'can/view/stache'],
    function ($, can, UserCurrent) {
        var Navbar = can.Control.extend({
            init: function () {
            },
            load: function () {
                var navbar = this;
                UserCurrent.findOne({}, function (data) {
                    navbar.data = data;
                    navbar.show();
                }, function(xhr){
                  navbar.show();
                });
            },
            show: function () {
                this.element.html(can.view('views_navbar_stache', {
                    data: this.data
                }));
                console.log('navbar data show :',this.data);
            },
            "#offcanvasNavbar a.internal click": function () {
                $.UIkit.offcanvas.hide(false);
            }
        });
        var navbar = new Navbar("#navbar");
        return navbar;
    });
