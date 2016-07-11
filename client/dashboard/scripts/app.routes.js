(function () {
    'use strict';
    angular.module('app.dashboard')
    .config(['$routeProvider',function($routeProvider) {
		$routeProvider.
	    when('/adminDashboard',{
			templateUrl: 'views/adminDashboard.html'
		}).
	    when('/adminUsers',{
			templateUrl: 'views/adminUsuarios.html'
		}).
	    when('/adminProjects',{
			templateUrl: 'views/adminProyectos.html'
		}).
		when('/adminProject',{
			templateUrl: 'views/adminProject.html'
		}).
	    when('/adminBalance',{
			templateUrl: 'views/adminBalance.html'
		}).
	    when('/adminApps',{
			templateUrl: 'views/adminAplicaciones.html'
		}).
		when('/projectView',{
			templateUrl: 'views/adminDashboard.html'
		}).
		when('/devDashboard',{
			templateUrl: 'views/clientDashboard.html'
		}).
		when('/devProjects',{
			templateUrl: 'views/clientApps.html'
		}).
		when('/devApps',{
			templateUrl: 'views/adminAplicaciones.html'
		}).		
		when('/devApp',{
			templateUrl: 'views/clientApp.html'
		}).
		when('/devProfile',{
			templateUrl: 'views/clientProfile.html'
		}).
		 when('/devFinance',{
			templateUrl: 'views/adminBalance.html'
		}).
		when('/clientDashboard',{
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
		when('/appPage',{
			templateUrl: 'views/appPage.html'
		}).					
		when('/newApp',{
			templateUrl: 'views/newApp.html'
		});
	}]);
})(); 