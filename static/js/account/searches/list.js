/* global $ , wisply*/
/**
 * @file Encapsulates the functionality for managing the searches of the account
 * @author Cristian Sima
 */
/**
 * @namespace AccountSearchesListModule
 */
var AccountSearchesListModule = function() {
	'use strict';
	/**
	 * The constructor does nothing
	 * @class List
	 * @memberof AccountSearchesListModule
	 * @classdesc It encapsulets the functions for the list with the searches
	 */
	var List = function List() {};
	/**
	 * @memberof List
	 */
	List.prototype =
		/** @lends AccountSearchesListModule.List */
		{
			/**
			 * It activates the listener for the delete button
			 * @param [event] event The event which has been fired
			 * @fires List#confirmClear
			 */
			init: function() {
				var instance = this;
				$("#clearSearchHistory").click(function(event) {
					event.preventDefault();
					instance.confirmClear();
				});
			},
			/**
			 * It shows the dialog for confirmation
			 */
			confirmClear: function() {
				var msg = this.getDialogMessage();
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguments for the confirmation dialog
			 * @return {object}         The arguments for the confirmation dialog
			 */
			getDialogMessage: function() {
				/**
				 * It returns the buttons for the dialog
				 * @return {object} The object which contains the buttons
				 */
				function getButtons() {
					var cancelButton,
						mainButton;
					cancelButton = {
						label: "Cancel",
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
					return {
						"cancel": cancelButton,
						"main": mainButton
					};
				}
				var buttons,
					msg,
					instance = this;
				buttons = getButtons();
				msg = {
					title: "We need your confirmation",
					message: "Wisply uses your search history in order to deliver better suggestions and results. You may choose to clear your history, but Wisply will not longer know what is best for you. Once it is deleted you <strong>cannot</strong> recover the history. <br /><br /> Are you sure?",
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
				 * It is called when the r
				 * @ignore
				 */
				errorCallback = function() {
					wisply.message.showError("The request was not successful");
				};
				request = {
					"url": '/account/searches/clear',
					"success": successCallback,
					"error": errorCallback
				};
				wisply.executePostAjax(request);
			}
		};
	return {
		List: List,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new AccountSearchesListModule();
	wisply.loadModule("account-searches-list", module);
});
