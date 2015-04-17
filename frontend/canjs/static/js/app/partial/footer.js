define(['jquery', 'can', 'can/view/stache'],
    function ($, can) {
        var Footer = can.Control.extend({
            init: function () {
            },
            load: function () {
                var footer = this;
                    footer.show();
            },
            show: function () {
                this.element.html(can.view('views_footer_stache', {}));
            }
        });
        var footer = new Footer("#footer");
        return footer;
    });