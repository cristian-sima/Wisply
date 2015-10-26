/* global jQuery,$, wisply, bootbox */
/**
 * @file Encapsulates the functionality for managing the Settings
 * @author Cristian Sima
 */
/**
 * @namespace Settings
 */
var Settings = function() {
	'use strict';

	/**
	 * The constructor activates the listeners
	 * @class Manager
	 * @memberof Settings
	 * @classdesc It encapsulets the functions for Settings.
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends Settings.Manager */
		{
			/**
			 * It activates the listener for all delete buttons
			 * @param [event] event The event which has been fired
			 * @fires SettingsManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$("#deleteAccountButton").click(function(event){
					event.preventDefault();
					instance.confirmDelete();
				});
			},
			/**
			 * It requests the account to confirm
			 */
			confirmDelete: function() {
				var instance = this,
					cancelButton  = {
						label: "Cancel",
						className: "btn-primary",
						callback: function() {
							this.modal('hide');
						}
					},
					mainButton = {
					label: "Delete Account",
					className: "btn-danger",
						callback: function() {
							var password = $("#promt-password").val();
							instance.delete(password);
						},
					},
					buttons = {
						cancel: cancelButton,
							main: mainButton,
					},
					options = {
						title: "Type your password",
						buttons: buttons,
						onEscape: true,
						message: '<input class="bootbox-input bootbox-input-text form-control" autocomplete="off" type="password" id="promt-password" />',
					};
					bootbox.dialog(options);
					setTimeout(function(){
						$("#promt-password").focus();
					}, 500);
			},
			/**
			 * It deletes the history
			 */
			delete: function(password) {
				var request,
					successCallback,
					errorCallback;
				/**
				 * It is called when request has been performed
				 * @ignore
				 */
				successCallback = function() {
					wisply.reloadPage();
				};
				/**
				 * It is called when there has been problems
				 * @ignore
				 */
				errorCallback = function() {
					wisply.message.showError("There was a problem with your request");
				};
				request = {
					url: '/account/settings/delete',
					success: successCallback,
					error: errorCallback,
					data: {
						password: password,
					},
				};
				wisply.executePostAjax(request);
			}
		};
	return {
		Manager: Manager,
	};
};
jQuery(document).ready(function() {
	"use strict";
	var module = new Settings();
	wisply.SettingsManager = new module.Manager();
	wisply.SettingsManager.init();
});
