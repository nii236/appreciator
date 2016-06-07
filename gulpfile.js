var gulp = require('gulp');
var notifier = require('node-notifier');
var sync = require('gulp-sync')(gulp).sync;
var reload = require('gulp-livereload');
var plumber = require ('gulp-plumber');
var sourcemaps = require('gulp-sourcemaps');
var stylelint = require('gulp-stylelint');
var autoprefixer = require('gulp-autoprefixer');
var rename = require('gulp-rename');
var util = require('gulp-util');
var sass = require('gulp-sass');
var cssmin = require('gulp-cssmin');
var child = require('child_process');

// ======
// Locals
// ======

var server = null;

// =========
// Overrides
// =========

/*
 * Override gulp.src() for nicer error handling.
 */
var src = gulp.src;
gulp.src = function() {
	return src.apply(gulp, arguments)
		.pipe(plumber(function(error) {
			util.log(util.colors.red(
	'Error (' + error.plugin + '): ' + error.message
			));
			notifier.notify({
	title: 'Error (' + error.plugin + ')',
	message: error.message.split('\n')[0]
			});
			this.emit('end');
		})
	);
};

// ============
// SERVER TASKS
// ============
// server:build recompiles the Go app
gulp.task('server:build', function() {
	var build = child.spawnSync('go', ['build']);
	if (build.status !== 0) {
		notifier.notify({
			title: 'Go build error',
			message: build.stderr
		});
	}
	return build
});

// server:spawn runs a new instance of the Go app, and connects the STDOUT and
// STDERR to the terminal
gulp.task('server:spawn', function() {
	if (server) server.kill();
	server = child.spawn('./appreciator', ['web']);
	server.stdout.on('data', function(data) {
		var lines = data.toString().split('\n');
		for (var l in lines) {
                        if (lines[l].length) {
                                util.log(lines[l]);
                        }
                }
	});
	server.stderr.on('data', function(data) {
                var lines = data.toString().split('\n');
		for (var l in lines) {
                        if (lines[l].length) {
                                util.log(util.colors.red(lines[l]));
                        }
                }
	});
	server.on('error', function(data) {
		process.stdout.write(data.toString())
		notifier.notify({
			title: 'Go run error',
			message: data.toString()
		});
	})
});

// server:watch will restart the process whenever a view is changed,
// and recompile then run if a *.go file is changed
gulp.task('server:watch', function() {
	gulp.watch(['./views/**/*'], ['server:spawn']);
	gulp.watch(['*/**/*.go'], sync(['server:build', 'server:spawn'], 'server'));
});

// ===========
// ASSET PIPELINE
// ===========
// Watch Files For Changes
gulp.task('assets:watch', function() {
	gulp.watch(['./assets/sass/**/*.scss'], ['assets:stylesheets']);
        gulp.watch(['./assets/js/**/*.js'], ['assets:js']);
});

gulp.task('assets:build', [
	'assets:stylesheets',
	'assets:js'
])
gulp.task('assets:js', function() {

});

gulp.task('assets:stylesheets', function() {
	gulp.src('./assets/sass/**/*.*')
		.pipe(gulp.dest('./assets/scss'));

	return gulp.src('./assets/sass/cutestrap.scss')
		.pipe(stylelint({
			reporters: [{formatter: 'string', console: true}]
		}))
		.pipe(sourcemaps.init())
			.pipe(sass().on('error', sass.logError))
		.pipe(sourcemaps.write())
		// .pipe(autoprefixer())
		.pipe(gulp.dest('./public/css'))
		.pipe(cssmin())
		.pipe(rename({suffix: '.min'}))
		.pipe(gulp.dest('./public/css'));
});

gulp.task('assets:watch', function() {
	gulp.watch(['./assets/sass/**/*.scss'], ['assets:build'])
})

// GULP INTERFACE
gulp.task('build', [
	'assets:build',
	'server:build'
]);

gulp.task('watch', ['assets:build', 'server:build'], function() {
	reload.listen();
	return gulp.start([
                'assets:watch',
		'server:watch',
		'server:spawn'
	]);
});

gulp.task('default', ['build']);
