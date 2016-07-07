(function () {
    'use strict';

    angular.module('app.dashboard').
    controller('newAppController', ['$scope','$location','$http','$parse', '$window','dataFactory',
    			function($scope,$location,$http,$parse,$window,dataFactory){

    	var storageRef = firebase.storage().ref();
    	var iconsRef = storageRef.child('icons');
    	$scope.newApp = {}
    	$scope.newApp.clientFeatures = new Array({})
    	$scope.newApp.iconPath = 'http://c.dryicons.com/images/icon_sets/blue_velvet/png/128x128/add.png'
    	$scope.newApp.headerSmallPath = 'http://icons.iconarchive.com/icons/dryicons/aesthetica-2/128/image-add-icon.png'
    	$scope.newApp.headerBigPath = 'http://icons.iconarchive.com/icons/dryicons/aesthetica-2/128/image-add-icon.png'
    	$scope.newApp.screenshotPath = 'http://icons.iconarchive.com/icons/dryicons/aesthetica-2/128/image-add-icon.png'
    	$scope.newApp.clientFeatures[0].img = 'http://icons.iconarchive.com/icons/dryicons/aesthetica-2/128/image-add-icon.png'
    	$scope.clientFeatures = [1]
    	
    	

    	$scope.setThumbnail = function(files,path,item,limit) {
    		var file = files[0];
		     //if (file.width <= limit) {
		     	var metadata = {
		        	'contentType': file.type
			      };
			      // Push to child path.
			      var uploadTask = storageRef.child(path+'/'+item+ $scope.newApp.Name).put(file, metadata);
			      // Listen for errors and completion of the upload.
			      // [START oncomplete]
			      uploadTask.on('state_changed', null, function(error) {
			        // [START onfailure]
			        console.error('Upload failed:', error);
			        // [END onfailure]
			      }, function() {
			        console.log('Uploaded',uploadTask.snapshot.totalBytes,'bytes.');
			        console.log(uploadTask.snapshot.metadata);
			        var model = $parse(item)
			        model.assign($scope, uploadTask.snapshot.metadata.downloadURLs[0])			   			        
			        $scope.$apply()
			      });
			      // [END oncomplete]
		     //}else {
		     	//$window.alert('Tamaño de imagen demasiado grande')
		     //}
    	}

    	$scope.upload = function (elm,index) {
    		var element = elm +index
	    	angular.element(element).click()
    	}
    	$scope.guardar = function(files,path,item,limit) {
		     var file = files[0];
		     //if (file.width <= limit) {
		     	var metadata = {
		        	'contentType': file.type
			      };
			      // Push to child path.
			      var uploadTask = storageRef.child(path+'/'+item+ $scope.newApp.Name).put(file, metadata);
			      // Listen for errors and completion of the upload.
			      // [START oncomplete]
			      uploadTask.on('state_changed', null, function(error) {
			        // [START onfailure]
			        console.error('Upload failed:', error);
			        // [END onfailure]
			      }, function() {
			        console.log('Uploaded',uploadTask.snapshot.totalBytes,'bytes.');
			        console.log(uploadTask.snapshot.metadata);
			        switch (item) {
			        	case 'Icon':
			        		$scope.newApp.iconPath = uploadTask.snapshot.metadata.downloadURLs[0];
			        		break;
			        	case 'hs':
			        		$scope.headerSmallPath = uploadTask.snapshot.metadata.downloadURLs[0];
			        		break;
			        	case 'hb':
			        		$scope.headerBigPath = uploadTask.snapshot.metadata.downloadURLs[0];
			        		break;
			        	case 'sh':
			        		$scope.screenshotPath = uploadTask.snapshot.metadata.downloadURLs[0];
			        		break;
			        	default:
			        		// statements_def
			        		break;
			        }			        
			        $scope.$apply()
			      });
			      // [END oncomplete]
		     //}else {
		     	//$window.alert('Tamaño de imagen demasiado grande')
		     //}
		};
		$scope.addClientFeature = function () {
			$scope.clientFeatures.push($scope.clientFeatures.length +1)
		}
		$scope.addFinalClientFeature = function () {
			$scope.clientFeatures.push($scope.clientFeatures.length +1)
		}
    }])

})();