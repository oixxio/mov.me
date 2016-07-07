// grab all our packages
var gulp = require('gulp'),
    browserSync = require('browser-sync').create();

// create a task to serve the app
gulp.task('serve', function() {
    // start the proxy for apache
    browserSync.init({
        proxy: "127.0.0.1/client/dashboard"
    });
    gulp.watch("client/**/*.html").on('change',browserSync.reload);
    gulp.watch("client/**/*.js").on('change',browserSync.reload);
    gulp.watch("client/**/*.css").on('change',browserSync.reload);
});