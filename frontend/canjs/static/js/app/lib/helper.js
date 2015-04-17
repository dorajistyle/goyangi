define(['i18n', 'can', 'app/lib/gravatar', 'util', 'info', 'moment'],
    function (i18n, can, gravatar, util, info, moment) {
    'use strict';

    can.stache.registerHelper('truncate', function (str) {
        return util.truncate40((typeof str === "function" ? str(): str));
    });
    can.stache.registerHelper('strip', function (str) {
        return util.strip((typeof str === "function" ? str(): str));
    });
    can.stache.registerHelper('safeTitle', function (str) {
        return util.slashToBlank(util.truncate40((typeof str === "function" ? str(): str)));
    });

    can.stache.registerHelper('i18n', function (str, options) {
        var i18nOptions = {};
        if(options.hash !== undefined) {
            $.each(options.hash, function(key, option){
                i18nOptions[key] = (typeof option === "function" ? option(): option);
            });
            if(i18nOptions['count'] != undefined) {
                i18nOptions['context'] = (i18nOptions['count'] > 1 ? 'plural' : (i18nOptions['count'] === 0 ? 'zero' : ''));
            }
            if(i18nOptions['str'] != undefined) {
              i18nOptions['sprintf'] = [i18nOptions['str']];
              i18nOptions['postProcess'] = 'sprintf';
            }
        }
        return i18n !== undefined ? i18n.t(str, i18nOptions)
            : str;
    });

    can.stache.registerHelper('timeago', function (time) {
      return moment(typeof time === "function" ? time(): time).fromNow();
    });

    can.stache.registerHelper('compare', function (first, second, options) {
        var result = util.compare((typeof first === "function" ? first(): first),
            (typeof second === "function" ? second(): second));
        return result ? options.fn(this) : options.inverse(this);
    });
    can.stache.registerHelper('lessThan', function (first, second, options) {
        var result = util.lessThanSecond((typeof first === "function" ? first(): first),
            (typeof second === "function" ? second(): second));
        return result ? options.fn(this) : options.inverse(this);
    });
    can.stache.registerHelper('greaterThan', function (first, second, options) {
        var result = util.greaterThanSecond((typeof first === "function" ? first(): first),
            (typeof second === "function" ? second(): second));
        return result ? options.fn(this) : options.inverse(this);
    });

    can.stache.registerHelper('gravatar', function (hash, size) {
      return gravatar.getGravatarURL((typeof hash === "function" ? hash(): hash), (typeof size === "function" ? size(): size), "mm", false);
    });

    can.stache.registerHelper('gravatarSecure', function (hash, size) {
      return gravatar.getGravatarURL((typeof hash === "function" ? hash(): hash), (typeof size === "function" ? size(): size), "mm", true);
    });


    can.stache.registerHelper('articleCategoryName', function (kind, id) {
        return info.getArticleCategoryName((typeof kind === "function" ? kind(): kind), (typeof id === "function" ? id(): id));
    });
    can.stache.registerHelper('articleCategoryRoute', function (id) {
        return info.getArticleCategoryRoute((typeof id === "function" ? id(): id));
    });
});
