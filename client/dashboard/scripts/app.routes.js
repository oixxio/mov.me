(function () {
    'use strict';
    angular.module('app.dashboard')
    .config(['$routeProvider',function($routeProvider) {
		$routeProvider.
	    when('/',{
			templateUrl: 'views/adminDashboard.html'
		}).
	    when('/usuarios',{
			templateUrl: 'views/adminUsuarios.html'
		}).
	    when('/proyectos',{
			templateUrl: 'views/adminProyectos.html'
		}).
	    when('/balance',{
			templateUrl: 'views/adminBalance.html'
		}).
	    when('/aplicaciones',{
			templateUrl: 'views/adminAplicaciones.html'
		}).
		when('/projectView',{
			templateUrl: 'views/adminDashboard.html'
		}).
		when('/client',{
			templateUrl: 'views/clientDashboard.html'
		}).
		when('/clientApps',{
			templateUrl: 'views/clientApps.html'
		}).		
		when('/clientApp',{
			templateUrl: 'views/clientApp.html'
		}).
		when('/clientProfile',{
			templateUrl: 'views/clientProfile.html'
		}).					
		when('/newApp',{
			templateUrl: 'views/newApp.html'
		});
	}]);
})(); 