(function () {
    'use strict';

    angular.module('app.dashboard').
    controller('indexController', ['$scope','$location','loginFactory','$http','$resource','dataFactory', 
    			function($scope,$location,loginFactory,$http,$resource,dataFactory){


        firebase.database().ref('Apps').on('value',function (snapshot) {
            $scope.$emit('getApps',snapshot.val())
        })
        firebase.database().ref('Projects').on('value',function (snapshot) {
            $scope.$emit('getProjects',snapshot.val())
        })
        firebase.database().ref('Users').on('value',function (snapshot) {
            $scope.$emit('getUsers',snapshot.val())
        })
       
        $scope.$on('getApps', function (event,data) {
            $scope.applicationsNumb =  Object.keys(data).length
            var keys = Object.keys(data)
            var applications = new Array()
            for (var i = 0; i < keys.length; i++) {
                data[keys[i]].Name = keys[i]
                applications[i] = data[keys[i]]
            }
            dataFactory.setApplications(applications)
            $scope.$apply()
        })
        $scope.$on('getProjects', function (event,data) {
            $scope.projectsNumb =  Object.keys(data).length
            var keys = Object.keys(data)
            var projects = new Array()
            for (var i = 0; i < keys.length; i++) {
                data[keys[i]].Name = keys[i]
                projects[i] = data[keys[i]]
            }
            dataFactory.setProjects(projects)
            $scope.$apply()
        })
        $scope.$on('getUsers', function (event,data) {
            
            $scope.usersNumb =  Object.keys(data).length
            var keys = Object.keys(data)
            var users = new Array()
            for (var i = 0; i < keys.length; i++) {
                data[keys[i]].Name = keys[i]
                users[i] = data[keys[i]]
            }
            dataFactory.setUsers(users)
            $scope.$apply()
        })                
        
    	/*tipo de usuario mockapeado*/
    	switch (loginFactory.userRole()) {
    		case 'admin':
    			$scope.userRole = 1;
    			break;
    		default:
    			// statements_def
    			break;
    	}
    	/*------------------------*/
    	//Esta funcion lo que hace es mostrar el menu del dash dependiendo del nivel de usuario obtenido
    	$scope.userMenu = function(role) {
    		var style = "";
			if (role === $scope.userRole) {
				style = {"display":"all"};
			}else {
				style = {"display":"none"};
			}
			return style;
    	}
    	//Funcion que permite navegar entre views
    	$scope.goTo = function (view) {
    		 $location.path(view);
    	}
        //TODO: sacar esta demoAlert en para produccion
        $scope.demoAlert = function(str){
            alert(str);
        }
        //TODO: re hacer este switch da asco 
        $scope.setActive = function(path){
            return ($location.path().substr(0, path.length) === path) ? 'active' : '';
        }


    }])

})();
