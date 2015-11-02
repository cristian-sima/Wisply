/* global $, wisply */
/**
 * @file Encapsulates the functionality for managing the processes
 * @author Cristian Sima
 */
/**
 * @namespace ProcessesModule
 */
var ProcessesModule = function() {
	'use strict';

	/**
	 * The constructor activates the listeners
	 * @class Manager
	 * @memberof ProcessesModule
	 * @classdesc It encapsulets the functions for processes
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends ProcessesModule.Manager */
		{
			/**
			 * It activates the listener for all delete buttons
			 */
			init: function() {
				var instance = this;
				$(".deleteProcessButton").click(function(event){
					event.preventDefault();
					var object,
						id;
					object = $(this);
					id = object.data("id");
					instance.confirmDelete(id);
				});
				$(".printPage").click(function(event) {
					event.preventDefault();
					window.print();
				});
			},
			/**
			 * It requests the user to confirm
			 * @param  {number} ID  The ID of the process
			 */
			confirmDelete: function(ID) {
				var msg = this.getDialogMessage(ID);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguements of the confimation message
			 * @param  {number} ID      The ID of the process
			 * @return {object}         The arguements of the confimation message
			 */
			getDialogMessage: function(ID) {
				var buttons,
					cancelButton,
					msg,
					mainButton,
					instance = this;
				cancelButton = {
					label: "No thanks",
					className: "btn-success",
					callback: function() {
						this.modal('hide');
					}
				};
				mainButton = {
					label: "Delete process",
					className: "btn-danger",
					callback: function() {
						instance.delete(ID);
					}
				};
				buttons = {
					"cancel": cancelButton,
					"main": mainButton
				};
				msg = {
					title: "Please confirm!",
					message: "The process will be permanently removed. Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It deletes the process
			 * @param  {number} ID The ID of the process
			 */
			delete: function(ID) {
				var request,
					successCallback,
					errorCallback;
				/**
				 * It is called when the deletion has been done. It reloads the page in 2 seconds
				 * @ignore
				 */
				successCallback = function() {
					wisply.message.showSuccess("The process has been removed! Refreshing page...");
					window.location="/admin/log";
				};
				/**
				 * It is called when there has been problems
				 * @ignore
				 */
				errorCallback = function() {
					wisply.message.showError("There was a problem with your request!");
				};
				request = {
					"url": '/admin/log/harvest/process/' + ID + "/delete",
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
	var module = new ProcessesModule();
	wisply.loadModule("admin-log-harvest-process", module);
});
