module.exports = function (grunt) {
    var staticPath = '../static';
    var scss_path = staticPath+'/scss';
    var css_path = staticPath+'/css';
//    var css_main_output = css_path+'/main.css';

    // Project configuration.
    require('time-grunt')(grunt);
    grunt.initConfig({
        css_path: css_path,
//        pkg: grunt.file.readJSON('./package.json'),

    shell: {
      scsscompile:{
        command: 'python2 ../css.py '+staticPath,
        options: {
            stdout: true
        }
      }
    },
    cancompile: {
        options: {
            version: '2.1.2',
            wrapper: 'define(["can/view/mustache"], function(can) { {{{content}}} });'
        },
        dist: {
            src: [staticPath+'/views/**/*.mustache'],
            out: staticPath+'/js/views.build.js',
            dest: staticPath+'/js/views.build.js',
            version: '2.1.2',
            wrapper: 'define(["can/view/mustache"], function(can) { {{{content}}} });'
        }
    },
    sass: {                                 // task
        dist: {                             // target
            files: {                        // dictionary of files
                 '<%= css_path %>/main.css' : scss_path +'/main.scss',     // 'destination': 'source'
                 '<%= css_path %>/http-errors.css' : scss_path +'/http-errors.scss'     // 'destination': 'source'
            }
        }
    },
    concat: {
        css: {
            src: [css_path+'/uikit.css', css_path+'/main.css', css_path+'/http-errors.css'],
            dest: css_path+'/application.css'
        }
    },
    watch: {
      run_mustache: {
        files: [staticPath+'/views/**/*.mustache'],
        tasks: ['cancompile']
      },
      run_scss: {
            files: [staticPath+'/scss/**/*.scss'],
            tasks: ['scss', 'concat-css']
      },
        options: { nospawn: true, livereload: true }
    }
  });

  grunt.registerTask('watch', ['watch']);
  grunt.registerTask('mustache', ['cancompile']);
  grunt.registerTask('scss', ['sass']);
  grunt.registerTask('concat-css', ['concat']);
  grunt.registerTask('static', ['cancompile','scss','concat-css']);

  grunt.loadNpmTasks('grunt-sass');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('can-compile');
  grunt.loadNpmTasks('grunt-shell');
  grunt.loadNpmTasks('grunt-contrib-watch');
};
