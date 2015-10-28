/* global jQuery,$, wisply */
/**
 * @file Encapsulates the functionality for managing the log
 * @author Cristian Sima
 */
/**
 * @namespace LogAdvanceOptionsModule
 */
var LogAdvanceOptionsModule = function() {
	'use strict';
	/**
	 * The constructor activates the listeners
	 * @class Manager
	 * @memberof LogAdvanceOptionsModule
	 * @classdesc It encapsulets the functions for the log
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends LogAdvanceOptionsModule.Manager */
		{
			/**
			 * It activates the listener for delete button
			 */
			init: function() {
				var instance = this;
				$(".deleteLogButton").click(function(event) {
					event.preventDefault();
					instance.confirmDelete();
				});
			},
			/**
			 * It requests the user to confirm
			 */
			confirmDelete: function() {
				var msg = this.getDialogMessage();
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguements of the confimation message
			 * @return {object}         The arguements of the confimation message
			 */
			getDialogMessage: function() {
				var buttons,
					cancelButton,
					msg,
					mainButton;
				cancelButton = {
					label: "Cancel",
					className: "btn-success",
					callback: function() {
						this.modal('hide');
					}
				};
				mainButton = {
					label: "Delete entire log",
					className: "btn-danger",
					callback: function() {
						wisply.logManager.delete();
					}
				};
				buttons = {
					"cancel": cancelButton,
					"main": mainButton
				};
				msg = {
					title: "Please confirm!",
					message: "The entire log will be permanently removed. Do you want to continue?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It deletes the log
			 */
			delete: function() {
				var request,
					successCallback,
					errorCallback;
				/**
				 * It is called when the deletion has been done. It reloads the page in 2 seconds
				 * @ignore
				 */
				successCallback = function() {
					wisply.message.showSuccess("The log has been removed! Refreshing page...");
					window.location = "/admin/log";
				};
				/**
				 * It is called when there has been problems
				 * @ignore
				 */
				errorCallback = function() {
					wisply.message.showError("There was a problem with your request!");
				};
				request = {
					"url": '/admin/log/delete',
					"success": successCallback,
					"error": errorCallback
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
	var module = new LogAdvanceOptionsModule();
	wisply.loadModule("log-advance-options", module);
});
