requirejs.config({
	baseUrl: "app/js/lib",
	paths: {
		app: "../app"
	},
	shim : {
		jquery : {
			exports : 'jQuery'
		},
		underscore : {
			exports : '_'
		},
		backbone : {
			deps : ['jquery', 'underscore'],
			exports : 'Backbone'
		}
	}
});

requirejs(['app/app']);