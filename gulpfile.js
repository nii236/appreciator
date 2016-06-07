var gulp = require('gulp');
var sync = require('gulp-sync')(gulp).sync;
var reload = require('gulp-livereload');
var util = require('gulp-util')
var child = require('child_process')

var server = null;

// SERVER TASKS
// server:build recompiles the Go app
gulp.task('server:build', function() {
  return child.spawnSync('go', ['build']);
});

// server:spawn runs a new instace of the Go app, and connects the STDOUT and
// STDERR to the terminal
gulp.task('server:spawn', function() {
  if (server) server.kill();
  server = child.spawn('./appreciator', ['web']);
  server.stdout.once('data', function() {
    reload.reload('/');
  })
  server.stdout.on('data', function(data) {
    var lines = data.toString().split('\n');
    for (var l in lines)
      if (lines[l].length) {
              util.log(lines[l]);
      }
  });
  server.stderr.on('data', function(data) {
    process.stdout.write(data.toString());
  });
});


// server:watch will restart the process whenever a view is changed,
// and recompile then run if a *.go file is changed
gulp.task('server:watch', function() {
  gulp.watch(['./views/**/*'], ['server:spawn']);
  gulp.watch(['*/**/*.go'], sync(['server:build', 'server:spawn'], 'server'));
});

// GULP INTERFACE
gulp.task('build', [
  'server:build'
]);

gulp.task('watch', ['server:build'], function() {
  reload.listen();
  return gulp.start([
    'server:watch',
    'server:spawn'
  ]);
});

gulp.task('default', ['build']);
