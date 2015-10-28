/* global $, wisply, bootbox */
/**
 * @file Encapsulates the functionality for managing the settings page for
 * account
 * @author Cristian Sima
 */
/**
 * @namespace AccountSettingsModule
 */
var AccountSettingsModule = function() {
	'use strict';
	/**
	 * The constructor does nothing
	 * @class Page
	 * @memberof AccountSettingsModule
	 * @classdesc It encapsulets the functions for settings page of the account.
	 */
	var Page = function Page() {};
	/**
	 * @memberof Page
	 */
	Page.prototype =
		/** @lends AccountSettingsModule.Page */
		{
			/**
			 * It activates the listener for delete button
			 * @param [event] event The event which has been fired
			 * @fires AccountSettingsModule#confirmDelete
			 */
			init: function() {
				var instance = this;
				$("#deleteAccountButton").click(function(event) {
					event.preventDefault();
					instance.confirmDelete();
				});
			},
			/**
			 * It shows the promt for typing the password
			 */
			confirmDelete: function() {
				var instance = this,
					cancelButton = {
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
				setTimeout(function() {
					$("#promt-password").focus();
				}, 500);
			},
			/**
			 * It sends a request to the server to delete the account
			 */
			delete: function(password) {
				var request,
					successCallback,
					errorCallback;
				/**
				 * It is called when the account has been deleted. It refreshes the page
				 * @ignore
				 */
				successCallback = function() {
					wisply.reloadPage();
				};
				/**
				 * It is called when the password is not good
				 * @ignore
				 */
				errorCallback = function() {
					wisply.message.showError("The request was not successful");
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
		Page: Page,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new AccountSettingsModule();
	wisply.loadModule("account-settings", module);
});
