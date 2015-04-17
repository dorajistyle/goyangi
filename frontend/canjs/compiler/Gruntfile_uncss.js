module.exports = function (grunt) {
    'use strict';
    var staticPath = '../static',
        css_path = staticPath+'/css/application.css';

    function get_css() {
      var css = '.test {}';
      if (!grunt.file.exists(css_path)) {
      // Warn the user if the file doesn't exist.
      grunt.log.warn('CSS file "' + css_path + '" not found.');
    } else {
      // Read the file.
      css = grunt.file.read(css_path);
    }
    return css;
    }




    // Project configuration.
    require('time-grunt')(grunt);
    grunt.initConfig({
//        pkg: grunt.file.readJSON('./package.json'),

    uncss: {
      dist: {
        options: {
          raw : get_css(),
          ignore       : ['#error-outer',
                        '#error-inner',
                        '.error-container',
                        '.error-page',
                        '.btn-large',
                        '.btn-success',
                        '.fascinate-font',
                        '.no-js',
                        '.chromeframe',
                        '.container',
                        '.app-main',
                        '.row',
                        '#fb-root',
                        '#ajaxProgress',
                        '.hidden',
                        '.fa',
                        '.fa-spinner',
                        '.fa-spin',
                        '.fa-lg',
                        '#messageWrapper',
                        '#messageAlertWarning',
                        '.alert',
                        '.alert-warning',
                        '#messageAlertError',
                        '.alert-danger',
                        '#messageAlertInfo',
                        '.alert-info',
                        '#messageAlertSuccess',
                        '.alert-success',
                        '#navbar',
                        '#appWrapper',
                        '#app',
                        '#popup',
                        '.fa-sort-asc',
                        '.fa-sort-desc',
                        '.order-asc',
                        '.selected-color',
                        /\.open/,
                        /\.mfp.*/],
          media        : ['                                            (min-width: 1240px)',
                        '                                (min-width: 750px)',
                        '                        (min-width: 1240px)',
                        '              (min-width: 1200px)',
                        '              (min-width: 992px)',
                        '            (min-width: 768px)',
                        '          (min-width: 1240px)',
                        '        (min-width: 1024px)',
                        '        (min-width: 768px)',
                        '      (min-width: 1200px)',
                        '      (min-width: 1240px)',
                        '    (min-width: 1200px)',
                        '    (min-width: 1240px)',
                        '  (min-width: 1240px)',
                        '  (min-width: 768px)',
                        '  (min-width: 992px)',
                        '(max-width: 1199px)',
                        '(max-width: 767px)',
                        '(max-width:767px)',
                        '(max-width:768px)',
                        '(min-width: 1024px)',
                        '(min-width: 1100px) and (max-width: 1240px)',
                        '(min-width: 480px) and (max-width: 768px)',
                        '(min-width: 768px) and (max-width: 1024px)',
                        '(min-width: 768px) and (max-width: 991px)',
                        '(min-width: 768px) and (max-width: 992px)',
                        '(min-width: 992px)',
                        '(min-width: 992px) and (max-width: 1100px)',
                        '(min-width:1200px)',
                        '(min-width:768px)',
                        '(min-width:768px) and (max-width:991px)',
                        '(min-width:992px)',
                        '(min-width:992px) and (max-width:1199px)',
                        'all and (max-width: 900px)',
                        'screen and (max-width: 480px)',
                        'screen and (max-width: 750px)',
                        'screen and (max-width: 767px)',
                        'screen and (max-width: 768px)',
                        'screen and (max-width: 800px) and (orientation: landscape), screen and (max-height: 300px)',
                        'screen and (max-width: 991px)',
                        'screen and (max-width:400px)',
                        'screen and (min-width:768px)'
                ]
        },
        files: {
          '../static/css/application.css': ['../static/views/**/*.mustache']
        }
      }
    }
  });

  grunt.registerTask('un_css', ['uncss']);
   grunt.registerTask('default', ['uncss']);
  grunt.loadNpmTasks('grunt-uncss');
};