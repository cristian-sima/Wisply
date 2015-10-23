/* global jQuery,$, wisply */
/**
 * @file Encapsulates the functionality for managing the Search
 * @author Cristian Sima
 */
/**
 * @namespace Search
 */
var Search = function() {
	'use strict';

	/**
	 * The constructor activates the listeners
	 * @class Manager
	 * @memberof Search
	 * @classdesc It encapsulets the functions for Search.
	 */
	var Manager = function Manager() {};
	/**
	 * @memberof Manager
	 */
	Manager.prototype =
		/** @lends Search.Manager */
		{
			/**
			 * It activates the listener for all delete buttons
			 * @param [event] event The event which has been fired
			 * @fires SearchManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$("#clearSearchHistory").click(function(event){
					event.preventDefault();
					instance.confirmClear();
				});
			},
			/**
			 * It requests the user to confirm
			 */
			confirmClear: function() {
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
					mainButton,
          instance = this;
				cancelButton = {
					label: "No, Thanks",
					className: "btn-primary",
					callback: function() {
						this.modal('hide');
					}
				};
				mainButton = {
					label: "Clear history",
					className: "btn-danger",
					callback: function() {
						instance.clear();
					}
				};
				buttons = {
					"cancel": cancelButton,
					"main": mainButton
				};
				msg = {
					title: "We need your confirmation",
					message: "Wisply uses your search history in order to deliver better suggestions and results. You may opt to clear your history, but Wisply will not longer know what is the best for you. Once it is deleted you can <strong>not</strong> recover the history <br /><br /> Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It clears the history
			 */
			clear: function() {
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
					"url": '/account/search/clear',
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
	var module = new Search();
	wisply.searchManager = new module.Manager();
	wisply.searchManager.init();
});
