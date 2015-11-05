/* global $, wisply */
/**
 * @file Encapsulates the functionality for managing the analysers
 * @author Cristian Sima
 */
/**
 * @namespace AdminEducationModule
 */
var AdminEducationModule = function() {
	'use strict';
	/**
	 * The constructor does nothing
	 * @class List
	 * @memberof AdminAnalysers
	 * @classdesc It manages the operations for deleting the analysers
	 */
	var List = function List() {};
	/**
	 * @memberof Manager
	 */
	List.prototype =
		/** @lends AdminAnalysers.List */
		{
			/**
			 * It activates the listeners for all delete buttons
			 * @fires AnalysersManager#confirmDelete
			 */
			init: function() {
				var instance = this;
				$(".deleteAnalyserButton").click(function(event) {
					event.preventDefault();
					var id = $(this).data("id");
					instance.confirmDelete(id);
				});
			},
			/**
			 * It shows the confirmation dialog
			 * @param  {number} id  The ID of analyser
			 */
			confirmDelete: function(id) {
				var msg = this.getDialogMessage(id);
				wisply.message.dialog(msg);
			},
			/**
			 * It returns the arguements for the confimation dialog
			 * @param  {number} id  The ID of analyser
			 * @return {object}         The arguements of the confimation message
			 */
			getDialogMessage: function(id) {
				/**
				 * It returns the buttons for the dialog
				 * @return {object} The buttons for the dialog
				 */
				function getButtons() {
					var cancelButton,
						mainButton;
					cancelButton = {
						label: "Cancel",
						className: "btn-success",
						callback: function() {
							this.modal('hide');
						}
					};
					mainButton = {
						label: "Delete",
						className: "btn-danger",
						callback: function() {
							instance.delete(id);
						}
					};
					return {
						"cancel": cancelButton,
						"main": mainButton
					};
				}
				var msg,
					instance = this,
					buttons = getButtons();
				msg = {
					title: "We need your confirmation",
					message: "The analyse will be permanently removed.<br /><br /> Are you sure?",
					onEscape: true,
					buttons: buttons
				};
				return msg;
			},
			/**
			 * It delets an analyser
			 * @param  {number} id  The ID of analyser
			 */
			delete: function(id) {
				var request,
					successCallback,
					errorCallback;
				/**
				 * It is called when the deletion has been done. It reloads the page in 2 seconds
				 * @ignore
				 */
				successCallback = function() {
					wisply.message.showSuccess("Done !");
					wisply.reloadPage();
				};
				/**
				 * It is called when the request has failed
				 * @ignore
				 */
				errorCallback = function() {
					wisply.message.showError("There was a problem with your request!");
				};
				request = {
					"url": '/admin/education/analyser/' + id + "/delete",
					"success": successCallback,
					"error": errorCallback
				};
				wisply.message.tellToWait("Removing the analyse...");
				wisply.executePostAjax(request);
			}
		};
	return {
		List: List,
	};
};
$(document).ready(function() {
	"use strict";
	var module = new AdminEducationModule();
	wisply.loadModule("admin-education-home", module);
});
