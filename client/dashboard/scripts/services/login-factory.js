(function () {
    'use strict';

    angular.module('app.dashboard').
    factory('loginFactory', ['$http', function($http){
    	var login = {};

    	login.userRole = function () {
    		var role = 'admin'; 
    		return  role
    	}
    	return login;
    }]);

})();