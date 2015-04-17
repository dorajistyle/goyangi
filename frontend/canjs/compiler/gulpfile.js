var parent_path = '../',
    staticPath = parent_path+'static/',
    scss_path = staticPath+'scss/',
    css_path = staticPath+'css/',
    bower_postfix = 'bower_components/',
    bower_path = staticPath+bower_postfix,
    scss_files = [scss_path +'*.scss', '!'+scss_path+'list.scss'],
    css_files = [css_path+'*.css','!'+css_path+'application.css'];

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



var gulp = require("gulp"),//http://gulpjs.com/
    util = require("gulp-util"),//https://github.com/gulpjs/gulp-util
    sass = require("gulp-sass"),//https://www.npmjs.org/package/gulp-sass
    autoprefixer = require('gulp-autoprefixer'),//https://www.npmjs.org/package/gulp-autoprefixer
    minifyCSS = require('gulp-minify-css'),//https://www.npmjs.org/package/gulp-minify-css
    concat = require('gulp-concat'),
    uglify = require('gulp-uglify'),
    replace = require('gulp-replace'),
    dir = require('yargs').argv.dir,
    uncss = require('gulp-uncss'),
    css_root = "./static/css/",
    js_root = "./static/js/",
    canCompiler = require('can-compile/gulp.js'),
    merge = require('merge-stream'),
    runSequence = require('run-sequence'),
    log = util.log;

var canOptions = {
        //src: [staticPath+'/views/**/*.mustache'],
        src: [staticPath+'/views/**/*.stache'],
        out: staticPath+'/js/views.build.js',
        dest: staticPath+'/js/views.build.js',
        version: '2.1.3',
        wrapper: 'define(["can", "can/view/stache"], function(can) { {{{content}}} });'
        //wrapper: 'define(["can/view/mustache"], function(can) { {{{content}}} });'
    };
//console.log('dir : '+dir);

// Creates a task called 'app-views'.  Must pass in same gulp instance.
canCompiler.task('cancompile', canOptions, gulp);
// Creates a task called 'app-views-watch'. Optional, but convenient.
canCompiler.watch('cancompile', canOptions, gulp);

gulp.task("generateCSS", function() {
   log("Generate CSS files " + (new Date()).toString());
    var scss = gulp.src(scss_files)
        .pipe(sass({ style: 'expanded' }))
                    .pipe(autoprefixer("last 3 version","safari 5", "ie 8", "ie 9"))
        .pipe(gulp.dest(css_path))
    var uikit = gulp.src(bower_path+'uikit/css/uikit.almost-flat.css')
        .pipe(replace('../', parent_path+bower_postfix+'uikit/'))
        .pipe(gulp.dest(css_path));
    var trumbowyg = gulp.src(bower_path+'trumbowyg/dist/ui/trumbowyg.css')
        .pipe(replace('./', parent_path+bower_postfix+'trumbowyg/dist/ui/'))
        .pipe(gulp.dest(css_path));
    return merge(scss, uikit, trumbowyg);
});
gulp.task("minifyCSS", function() {
  return gulp.src(css_files)
        .pipe(concat('application.css'))
        .pipe(minifyCSS({keepBreaks:true}))
        .pipe(gulp.dest(css_path));
});
gulp.task("scss", function(){
  runSequence('generateCSS', 'minifyCSS');
});

gulp.task('scss-watch', function(){
  gulp.watch(scss_path +'*.scss', ['scss']);
});

gulp.task('log-watch', function(){
  log("Watching mustache files and scss files for modifications");
});


gulp.task('default', ['cancompile','scss']);

gulp.task('watch', [
    'cancompile-watch',
    'scss-watch',
    'log-watch'
]);

//gulp.task('default', function() {
//
// gulp.src(dir+'/css/application.css')
//  .pipe(uncss({
//    raw : get_css(),
//    ignore       : ['#error-outer',
//                    '#error-inner',
//                    '.error-container',
//                    '.error-page',
//                    '.btn-large',
//                    '.btn-success',
//                    '.fascinate-font',
//                    '.no-js',
//                    '.chromeframe',
//                    '.container',
//                    '.app-main',
//                    '.row',
//                    '#fb-root',
//                    '#ajaxProgress',
//                    '.hidden',
//                    '.fa',
//                    '.fa-spinner',
//                    '.fa-spin',
//                    '.fa-lg',
//                    '#messageWrapper',
//                    '#messageAlertWarning',
//                    '.alert',
//                    '.alert-warning',
//                    '#messageAlertError',
//                    '.alert-danger',
//                    '#messageAlertInfo',
//                    '.alert-info',
//                    '#messageAlertSuccess',
//                    '.alert-success',
//                    '#navbar',
//                    '#appWrapper',
//                    '#app',
//                    '#popup',
//                    '.fa-sort-asc',
//                    '.fa-sort-desc',
//                    '.order-asc',
//                    '.selected-color',
//                    /\.open/,
//                    /\.mfp.*/],
//      media        : ['                                            (min-width: 1240px)',
//                    '                                (min-width: 750px)',
//                    '                        (min-width: 1240px)',
//                    '              (min-width: 1200px)',
//                    '              (min-width: 992px)',
//                    '            (min-width: 768px)',
//                    '          (min-width: 1240px)',
//                    '        (min-width: 1024px)',
//                    '        (min-width: 768px)',
//                    '      (min-width: 1200px)',
//                    '      (min-width: 1240px)',
//                    '    (min-width: 1200px)',
//                    '    (min-width: 1240px)',
//                    '  (min-width: 1240px)',
//                    '  (min-width: 768px)',
//                    '  (min-width: 992px)',
//                    '(max-width: 1199px)',
//                    '(max-width: 767px)',
//                    '(max-width:767px)',
//                    '(max-width:768px)',
//                    '(min-width: 1024px)',
//                    '(min-width: 1100px) and (max-width: 1240px)',
//                    '(min-width: 480px) and (max-width: 768px)',
//                    '(min-width: 768px) and (max-width: 1024px)',
//                    '(min-width: 768px) and (max-width: 991px)',
//                    '(min-width: 768px) and (max-width: 992px)',
//                    '(min-width: 992px)',
//                    '(min-width: 992px) and (max-width: 1100px)',
//                    '(min-width:1200px)',
//                    '(min-width:768px)',
//                    '(min-width:768px) and (max-width:991px)',
//                    '(min-width:992px)',
//                    '(min-width:992px) and (max-width:1199px)',
//                    'all and (max-width: 900px)',
//                    'screen and (max-width: 480px)',
//                    'screen and (max-width: 750px)',
//                    'screen and (max-width: 767px)',
//                    'screen and (max-width: 768px)',
//                    'screen and (max-width: 800px) and (orientation: landscape), screen and (max-height: 300px)',
//                    'screen and (max-width: 991px)',
//                    'screen and (max-width:400px)',
//                    'screen and (min-width:768px)'
//            ]
//    },
//    files: {
//      '../static/css/application.css': ['../static/views/**/*.mustache']
//    }
//  ))
//  .pipe(gulp.dest(dir+'dest'));
//
//});

//    gulp.task('default', function() {
//
//  //gulp.src(js_root+'*.js')
//  gulp.src([js_root+'jquery-1.11.1.js',js_root+'bootstrap.min.js',js_root+'jquery.easing.min.js', js_root+'jquery.sticky.js',js_root+'jquery.scrollTo.js',js_root+'stellar.js',js_root+'wow.js',js_root+'owl.carousel.js',js_root+'nivo-lightbox.js',js_root+'instafeed.js',js_root+'custom.js'])
//  .pipe(uglify())
//  .pipe(concat('application.js'))
//  .pipe(gulp.dest(js_root))
//
//  //gulp.src(css_root+'*.css')
//  gulp.src([css_root+'bootstrap.css', css_root+'font-awesome.min.css', css_root+'animate.css', css_root+'owl.carousel.css', css_root+'owl.theme.css', css_root+'nivo-lightbox.css', css_root+'responsive.css', css_root+'style.css'])
//  .pipe(minifyCSS({keepBreaks:true}))
//  .pipe(concat('application.css'))
//  .pipe(gulp.dest(css_root))
//
//});
